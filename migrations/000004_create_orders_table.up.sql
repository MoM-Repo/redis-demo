-- Создание таблицы заказов для демонстрации сложных запросов
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    product_name VARCHAR(255) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);
CREATE INDEX IF NOT EXISTS idx_orders_created_at ON orders(created_at);

DO $$
DECLARE
    i INTEGER;
    j INTEGER;
    product_names TEXT[] := ARRAY['iPhone 15', 'MacBook Pro', 'iPad Air', 'AirPods', 'Apple Watch', 'Samsung Galaxy', 'Dell Laptop', 'Sony Headphones'];
    statuses TEXT[] := ARRAY['pending', 'completed', 'shipped', 'cancelled'];
BEGIN
    FOR i IN 1..1000 LOOP
        FOR j IN 1..(3 + (i % 5)) LOOP
            INSERT INTO orders (user_id, product_name, amount, status, created_at, updated_at)
            VALUES (
                i,
                product_names[1 + (j % array_length(product_names, 1))],
                (50 + (j * 25) + (i % 1000))::DECIMAL(10,2),
                statuses[1 + (j % array_length(statuses, 1))],
                NOW() - INTERVAL '1 day' * (j % 30),
                NOW() - INTERVAL '1 day' * (j % 30)
            );
        END LOOP;
    END LOOP;
END $$;
