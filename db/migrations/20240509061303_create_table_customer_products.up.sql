CREATE TABLE customer_products (
  id VARCHAR(255) PRIMARY KEY,
  customerId VARCHAR(255) NOT NULL,
  productId VARCHAR(255) NOT NULL,
  quantity INT NOT NULL,
  paid INT NOT NULL,
  change INt NOT NULL,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (productId) REFERENCES products(id),
  FOREIGN KEY (customerId) REFERENCES customer(id)
);