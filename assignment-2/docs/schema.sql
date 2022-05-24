DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS items;

CREATE TABLE orders(
   order_id INT GENERATED ALWAYS AS IDENTITY,
   customer_name VARCHAR(255) NOT NULL,
   ordered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY(order_id)
);

CREATE TABLE items(
   item_id INT GENERATED ALWAYS AS IDENTITY,
   item_code INT,
   description TEXT,
   quantity INT,
   order_id INT NOT NULL,
   PRIMARY KEY(item_id),
   CONSTRAINT fk_customer
      FOREIGN KEY(order_id)
	  REFERENCES orders(order_id)
);