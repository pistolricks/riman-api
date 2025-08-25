CREATE INDEX IF NOT EXISTS orders_title_idx ON orders USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS orders_genres_idx ON orders USING GIN (genres);
