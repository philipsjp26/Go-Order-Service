-- +goose Up
-- +goose StatementBegin
INSERT INTO menus (name) VALUES('tengkleng'), ('tongseng');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
