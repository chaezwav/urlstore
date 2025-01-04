-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS links
(
    id             TEXT    PRIMARY KEY,
    alias          TEXT,
    url            TEXT    NOT NULL,
    flags          TEXT[],
    created_at     DATE    NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS links;
-- +goose StatementEnd
