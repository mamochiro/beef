CREATE TABLE IF NOT EXISTS remote_config 
(
    id             SERIAL PRIMARY KEY,
    key            VARCHAR(100),
    value          VARCHAR(100),
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     BIGINT
);
