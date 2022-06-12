BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE gender AS ENUM ('male', 'female');

CREATE TYPE status AS ENUM ('active', 'cancelled');

CREATE TABLE users (
	id uuid DEFAULT uuid_generate_v4(),
	first_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
    gender gender NOT NULL,
    birthday DATE NOT NULL,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,
	PRIMARY KEY (id),
    UNIQUE (first_name, last_name)
);

CREATE TABLE launchpads (
	id VARCHAR(24) NOT NULL,
    name VARCHAR(20) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE destinations (
	id uuid DEFAULT uuid_generate_v4(),
    name VARCHAR(60) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE orders (
	id uuid DEFAULT uuid_generate_v4(),
	launchpad_id VARCHAR(24) NOT NULL,
	destination_id uuid NOT NULL,
    status status NOT NULL,
    launch_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
	PRIMARY KEY (id),
    FOREIGN KEY (launchpad_id) REFERENCES launchpads(id) ON DELETE RESTRICT,
    FOREIGN KEY (destination_id) REFERENCES destinations(id) ON DELETE RESTRICT
);

CREATE TABLE orders_users (
	order_id uuid NOT NULL,
	user_id uuid NOT NULL,
	PRIMARY KEY (order_id, user_id),
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE RESTRICT,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT
);

CREATE TABLE trips_schedule (
    destination_id uuid NOT NULL,
    date DATE NOT NULL,
	PRIMARY KEY (destination_id, date),
    FOREIGN KEY (destination_id) REFERENCES destinations(id) ON DELETE RESTRICT
);

CREATE TABLE launchpad_reservations (
    launchpad_id VARCHAR(24) NOT NULL,
    date DATE NOT NULL,
    reserved BOOLEAN NOT NULL DEFAULT false,
	PRIMARY KEY (launchpad_id, date),
    FOREIGN KEY (launchpad_id) REFERENCES launchpads(id) ON DELETE RESTRICT
);
COMMIT;