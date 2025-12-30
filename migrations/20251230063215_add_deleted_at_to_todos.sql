-- +goose Up
ALTER TABLE todos
ADD COLUMN deleted_at TIMESTAMP NULL;

CREATE INDEX IF NOT EXISTS idx_todos_deleted_at
ON todos (deleted_at);

-- +goose Down
DROP INDEX IF EXISTS idx_todos_deleted_at;
ALTER TABLE todos
DROP COLUMN deleted_at;