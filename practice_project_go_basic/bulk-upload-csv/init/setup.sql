-- 1. Bảng Categories (Master Data)
CREATE TABLE IF NOT EXISTS categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL
) ENGINE=InnoDB;

-- 2. Bảng Products (Master Data)
CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    sku VARCHAR(100) NOT NULL UNIQUE,
    category_id INT,
    CONSTRAINT fk_product_category 
        FOREIGN KEY (category_id) REFERENCES categories(id)
        ON DELETE SET NULL
) ENGINE=InnoDB;

-- 3. Bảng Warehouses (Master Data)
CREATE TABLE IF NOT EXISTS warehouses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    code VARCHAR(50) NOT NULL UNIQUE
) ENGINE=InnoDB;

-- 4. Bảng Inventory Transactions (Bảng nghiệp vụ)
CREATE TABLE IF NOT EXISTS inventory_transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    category_id INT NOT NULL,
    warehouse_id INT NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    transaction_type ENUM('IN', 'OUT', 'ADJUST') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_trans_product FOREIGN KEY (product_id) REFERENCES products(id),
    CONSTRAINT fk_trans_category FOREIGN KEY (category_id) REFERENCES categories(id),
    CONSTRAINT fk_trans_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
) ENGINE=InnoDB;

-- Chèn Master Data mẫu
INSERT INTO categories (code, name) VALUES ('ELEC', 'Electronics'), ('FURN', 'Furniture');

INSERT INTO products (sku, category_id) VALUES ('LAP-DELL-01', 1), ('CHAIR-OFF-02', 2);

INSERT INTO warehouses (code) VALUES ('WH-HCM-01'), ('WH-HN-02');

-- Chèn thử 1 giao dịch nhập kho
INSERT INTO inventory_transactions (product_id, category_id, warehouse_id, quantity, transaction_type) 
VALUES (1, 1, 1, 10, 'IN');