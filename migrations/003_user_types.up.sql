CREATE TABLE IF NOT EXISTS user_types
(
    user_types VARCHAR(100)
);
INSERT INTO user_types
VALUES ('user'),
       ('owner'),
       ('admin');

ALTER TABLE IF EXISTS workspaces
    ADD COLUMN owner UUID NOT NULL REFERENCES users (user_id);

CREATE TABLE IF NOT EXISTS companies
(
    id    UUID         NOT NULL PRIMARY KEY,
    name  VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    owner UUID         NOT NULL REFERENCES users (user_id)
);
CREATE INDEX IF NOT EXISTS company_id_index ON companies (id);

ALTER TABLE IF EXISTS users
    ADD COLUMN company_id UUID[] DEFAULT '{}';