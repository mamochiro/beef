CREATE TABLE IF NOT EXISTS splash_page
(
    id                  SERIAL PRIMARY KEY,
    splash_image        TEXT,
    splash_type         VARCHAR(50),
    splash_data         TEXT,
    is_international    BOOLEAN,
    is_show             BOOLEAN,
    created_at          TIMESTAMP,
    updated_at          TIMESTAMP
);