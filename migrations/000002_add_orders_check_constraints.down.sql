ALTER TABLE orders DROP CONSTRAINT IF EXISTS orders_runtime_check;

ALTER TABLE orders DROP CONSTRAINT IF EXISTS orders_year_check;

ALTER TABLE orders DROP CONSTRAINT IF EXISTS genres_length_check;
