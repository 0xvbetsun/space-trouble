// package dto is using for storing objects for sending data between layers in application
package dto

import "time"

// User represents data-transfer object for user entity
type User struct {
	FirstName string
	LastName  string
	Gender    string
	Birthday  time.Time
}
