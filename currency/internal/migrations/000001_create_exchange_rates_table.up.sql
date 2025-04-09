CREATE TABLE currencies (
    id SERIAL PRIMARY KEY,
    created_date TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_date TIMESTAMP ,
    basic_currency VARCHAR NOT NULL,
    other_currency VARCHAR NOT NULL,
    meaning numeric NOT NULL
);
