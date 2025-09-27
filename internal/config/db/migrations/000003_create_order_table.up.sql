-- Создание таблицы с заказами
BEGIN;

CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  number BIGINT NOT NULL,
  accrual NUMERIC(15,2) NOT NULL,
  status_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  uploaded_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  
  CONSTRAINT fk_orders_status 
    FOREIGN KEY (status_id) 
    REFERENCES statuses(id)
    ON UPDATE CASCADE,
  CONSTRAINT fk_orders_users 
    FOREIGN KEY (user_id) 
    REFERENCES users(id)
    ON UPDATE CASCADE
);

COMMIT; 