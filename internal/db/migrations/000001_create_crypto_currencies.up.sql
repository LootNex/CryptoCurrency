CREATE TABLE cryptocurrencies (
    id SERIAL PRIMARY KEY,
    crypto_name TEXT NOT NULL UNIQUE
);