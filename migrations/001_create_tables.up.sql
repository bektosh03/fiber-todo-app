CREATE TABLE IF NOT EXISTS users_auth
(
    id            UUID         NOT NULL PRIMARY KEY,
    refresh_token TEXT,
    email         VARCHAR(255) NOT NULL UNIQUE,
    password      TEXT         NOT NULL
);
CREATE TABLE IF NOT EXISTS workspaces
(
    id         UUID        NOT NULL PRIMARY KEY,
    name       VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL
);
CREATE TYPE gender AS ENUM ('male' ,'female');
CREATE TABLE IF NOT EXISTS users
(
    user_id       UUID UNIQUE  NOT NULL REFERENCES users_auth (id),
    workspace_id  UUID[]       NOT NULL DEFAULT '{}',
    email         VARCHAR(255) NOT NULL UNIQUE,
    first_name    VARCHAR(255) NOT NULL,
    last_name     VARCHAR(255),
    username      VARCHAR(255) NOT NULL UNIQUE,
    age           int8,
    gender        gender,
    profile_photo TEXT,
    created_at    TIMESTAMPTZ  NOT NULL
);
CREATE TABLE IF NOT EXISTS todos
(
    id          UUID         NOT NULL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    due_date    TIMESTAMPTZ,
    assignee    UUID,
    issuer      UUID         NOT NULL REFERENCES users (user_id),
    created_at  TIMESTAMPTZ  NOT NULL,
    completed   BOOLEAN DEFAULT FALSE
);
CREATE INDEX IF NOT EXISTS workspace_id_index ON workspaces (id);
CREATE INDEX IF NOT EXISTS users_id_index ON users (user_id);
CREATE INDEX IF NOT EXISTS todos_id_index ON todos (id);