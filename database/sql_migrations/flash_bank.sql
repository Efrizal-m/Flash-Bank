-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE customer (
    customer_id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    idcard_number VARCHAR(20),
    address VARCHAR(512),
    cif VARCHAR(20),
    created_at TIMESTAMP without time zone,
    updated_at TIMESTAMP without time zone
);

CREATE TABLE saldo (
    saldo_id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customer(customer_id),
    saldo DECIMAL(10, 2),
    transaction_date DATE
);

CREATE TABLE transaction (
    transaction_id SERIAL PRIMARY KEY,
    saldo_id INT REFERENCES saldo(saldo_id),
    volume DECIMAL(10, 2),
    transaction_type VARCHAR(20),
    transaction_date DATE
);

CREATE TABLE report (
    report_id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES transaction(transaction_id),
    volume_in DECIMAL(10, 2),
    volume_out DECIMAL(10, 2),
    transaction_date DATE
);

CREATE TYPE enum_type AS ENUM ('admin', 'user');

CREATE TABLE user_admin (
    id SERIAL PRIMARY KEY,
    username VARCHAR(256),
    password VARCHAR(256), -- or hashed_password if you store hashed passwords
    role enum_type,
    created_at TIMESTAMP without time zone,
    updated_at TIMESTAMP without time zone
);

-- +migrate StatementEnd
