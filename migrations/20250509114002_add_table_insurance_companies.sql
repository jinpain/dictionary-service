-- +goose Up
-- +goose StatementBegin
CREATE TABLE insurance_companies (
    id SERIAL PRIMARY KEY,
    type INT4 NOT NULL,
    name VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    address TEXT,
    phone CHAR(11),
    email VARCHAR(32),
    website VARCHAR(255),
    ogrn CHAR(13),
    okpo VARCHAR(10),
    okato VARCHAR(11),
    created_ts TIMESTAMP DEFAULT NULL,
    created_by VARCHAR(32) DEFAULT NULL,
    updated_ts TIMESTAMP DEFAULT NULL,
    updated_by VARCHAR(32) DEFAULT NULL,
    deleted_ts TIMESTAMP DEFAULT NULL,
    deleted_by VARCHAR(32) DEFAULT NULL,
    version INT DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd