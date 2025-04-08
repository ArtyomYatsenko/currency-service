CREATE TABLE currencies (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE currency_rates (
    created_date TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_date TIMESTAMP NOT NULL DEFAULT NOW(),
    basic_currency INT NOT NULL REFERENCES currencies(id),
    other_currency INT NOT NULL REFERENCES currencies(id)
);