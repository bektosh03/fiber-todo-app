DROP INDEX IF EXISTS workspace_id_index;
DROP INDEX IF EXISTS users_id_index;
DROP INDEX IF EXISTS todos_id_index;

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS workspaces CASCADE;
DROP TABLE IF EXISTS users_auth CASCADE;