CREATE EXTENSION IF NOT EXISTS citext;

CREATE DOMAIN user_name AS citext
CHECK (
    char_length(VALUE) >= 2 AND
    char_length(VALUE) <= 16 AND
    VALUE ~ '^[a-zA-Z0-9_]+$'
);

CREATE TABLE IF NOT EXISTS users {
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    name user_name UNIQUE NOT NULL,
    display_name VARCHAR(32) NOT NULL,
    bio VARCHAR(256),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
}
