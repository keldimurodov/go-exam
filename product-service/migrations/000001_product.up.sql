CREATE TABLE product(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    product_name VARCHAR(100),
    product_price INT,
    product_about VARCHAR(300),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updeted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);