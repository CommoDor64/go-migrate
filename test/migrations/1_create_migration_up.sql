CREATE TABLE IF NOT EXISTS migration
(
    id      SERIAL PRIMARY KEY,
    uuid    UUID        NOT NULL DEFAULT uuid_generate_v4(),
    created timestamp   NOT NULL DEFAULT NOW(),
    name    varchar(50) NOT NULL,
    hash    text        UNIQUE NOT NULL
);