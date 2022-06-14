// Entry point for application's monitor which actualizes  launches of SpaceX
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"github.com/vbetsun/space-trouble/configs"
	"github.com/vbetsun/space-trouble/internal/storage/psql"
	"go.uber.org/zap"
)

type Response struct {
	LaunchpadID string `json:"launchpad"`
	DateUTC     string `json:"date_utc"`
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	if err := configs.LoadConfig("configs"); err != nil {
		logger.Fatal(fmt.Sprintf("can't read config: %v", err))
	}
	dbHost := viper.GetString("POSTGRES_HOST")
	if dbHost == "" {
		dbHost = viper.GetString("db.host")
	}
	apiURL := viper.GetString("api.url")
	db, err := psql.NewDB(psql.Config{
		Host:     dbHost,
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
		Logger:   logger,
	})
	if err != nil {
		logger.Fatal(fmt.Sprintf("can't connect to the DB %v", err))
	}
	store := psql.NewStorage(db)

	for {
		launches, err := fetchLaunches(apiURL)
		errCount := 0
		if err != nil {
			logger.Fatal(err.Error())
		}
		for _, launch := range launches {
			err := store.Launchpad.AddReservation(launch.LaunchpadID, launch.DateUTC)
			if err != nil {
				errCount += 1
				logger.Error(err.Error())
			}
		}
		logger.Info(fmt.Sprintf("launches synchronized. Total: %d, errors: %d", len(launches), errCount))
		time.Sleep(20 * time.Second) // regarding the documentation, cache renews every 20 sec
	}
}

// fetchLaunches it is a function for getting information from spaceX api about launces
func fetchLaunches(apiURL string) ([]Response, error) {
	var data []Response
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodGet, apiURL+"/launches", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error request response with status code %d", resp.StatusCode)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
