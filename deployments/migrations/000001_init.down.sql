BEGIN;

DROP TABLE IF EXISTS launchpad_reservations;
DROP TABLE IF EXISTS trips_schedule;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS launchpads;
DROP TABLE IF EXISTS destinations;

DROP TYPE IF EXISTS gender;
DROP TYPE IF EXISTS status;
DROP EXTENSION IF EXISTS "uuid-ossp";

COMMIT;