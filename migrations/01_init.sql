CREATE TABLE racks (
    id SERIAL PRIMARY KEY,
    "name" CHAR NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL
);

CREATE UNIQUE INDEX products_idx ON products("name");

CREATE TABLE products_racks (
    id SERIAL PRIMARY KEY,
    product_id SERIAL NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    rack_id SERIAL NOT NULL REFERENCES racks(id) ON DELETE CASCADE,
    is_basic BOOLEAN NOT NULL
);

CREATE UNIQUE INDEX products_racks_idx ON products_racks(product_id, rack_id);

CREATE TABLE products_orders (
    id SERIAL PRIMARY KEY,
    product_id SERIAL NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    order_id SERIAL NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    quantity INT NOT NULL DEFAULT 1
);

CREATE UNIQUE INDEX products_orders_idx ON products_orders(product_id, order_id);

INSERT INTO racks(id, "name") VALUES (1, 'А');
INSERT INTO racks(id, "name") VALUES (2, 'Б');
INSERT INTO racks(id, "name") VALUES (3, 'В');
INSERT INTO racks(id, "name") VALUES (4, 'Ж');
INSERT INTO racks(id, "name") VALUES (5, 'З');

INSERT INTO products(id, "name") VALUES (1, 'Ноутбук');
INSERT INTO products(id, "name") VALUES (2, 'Телевизор');
INSERT INTO products(id, "name") VALUES (3, 'Телефон');
INSERT INTO products(id, "name") VALUES (4, 'Системный блок');
INSERT INTO products(id, "name") VALUES (5, 'Часы');
INSERT INTO products(id, "name") VALUES (6, 'Микрофон');

INSERT INTO products_racks(id, rack_id, product_id, is_basic) VALUES (1, 1, 1, true); --ноутбук
INSERT INTO products_racks(id, rack_id, product_id, is_basic) VALUES (2, 1, 2, true); --телевизор
INSERT INTO products_racks(id, rack_id, product_id, is_basic) VALUES (3, 2, 3, true); --телефон
INSERT INTO products_racks(id, rack_id, product_id, is_basic) VALUES (5, 4, 4, true); --системный блок
INSERT INTO products_racks(id, rack_id, product_id, is_basic) VALUES (4, 4, 5, true); --часы
INSERT INTO products_racks(id, rack_id, product_id, is_basic) VALUES (6, 4, 6, true); --микрофон
INSERT INTO products_racks(id, rack_id, product_id, is_basic) VALUES (7, 5, 3, false); --телефон доп стеллаж З
INSERT INTO products_racks(id, rack_id, product_id, is_basic) VALUES (8, 3, 3, false); --телефон доп стеллаж В
INSERT INTO products_racks(id, rack_id, product_id, is_basic) VALUES (9, 1, 5, false); --часы доп стеллаж А

INSERT INTO orders(id) VALUES (10);
INSERT INTO orders(id) VALUES (11);
INSERT INTO orders(id) VALUES (14);
INSERT INTO orders(id) VALUES (15);

INSERT INTO products_orders(id, product_id, order_id, quantity) VALUES (1, 1, 10, 2);
INSERT INTO products_orders(id, product_id, order_id, quantity) VALUES (2, 2, 11, 3);
INSERT INTO products_orders(id, product_id, order_id, quantity) VALUES (3, 1, 14, 3);
INSERT INTO products_orders(id, product_id, order_id, quantity) VALUES (4, 3, 10, 1);
INSERT INTO products_orders(id, product_id, order_id, quantity) VALUES (5, 4, 14, 4);
INSERT INTO products_orders(id, product_id, order_id, quantity) VALUES (6, 5, 15, 1);
INSERT INTO products_orders(id, product_id, order_id, quantity) VALUES (7, 6, 10, 1);
