-- +goose Up
create table products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text not null,
    value int not null
);

-- +goose Down
drop table products;
