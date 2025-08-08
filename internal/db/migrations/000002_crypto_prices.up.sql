CREATE TABLE prices (
    crypto_name TEXT NOT NULL REFERENCES cryptocurrencies(crypto_name) ON DELETE CASCADE,
    price NUMERIC NOT NULL,
    recorded_at BIGINT NOT NULL
);
