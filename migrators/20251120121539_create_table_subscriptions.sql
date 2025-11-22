-- +goose Up
-- +goose StatementBegin
CREATE TABLE subscriptions(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL, CHECK (price>=0),
    user_id UUID NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE, 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subscriptions
-- +goose StatementEnd
