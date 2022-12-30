-- +goose Up
-- +goose StatementBegin
CREATE TABLE Offers (
                        id uuid primary key DEFAULT gen_random_uuid(),
                        title VARCHAR NOT NULL,
                        description TEXT NOT NULL,
                        photo_url varchar NOT NULL,
                        price float4,
                        created_at timestamp NOT NULL DEFAULT (now() AT TIME ZONE ('utc'))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS Offers;
-- +goose StatementEnd
