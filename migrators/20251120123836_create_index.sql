-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_user_id ON subscriptions(user_id);
CREATE INDEX idx_servece_name ON subscriptions(service_name);
CREATE INDEX idx_start_date ON subscriptions(start_date);
CREATE INDEX idx_end_date ON subscriptions(end_date);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_user_id;
DROP INDEX IF EXISTS idx_servece_name;
DROP INDEX IF EXISTS idx_start_date;
DROP INDEX IF EXISTS idx_end_date;
-- +goose StatementEnd
