-- Создание таблицы с запросами на списание средств
BEGIN;

CREATE TABLE withdrawal (
  id SERIAL PRIMARY KEY,
  sum NUMERIC(15,2) NOT NULL,
  order_id BIGINT NOT NULL,
  user_id INTEGER NOT NULL,
  update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  
  CONSTRAINT fk_balance_users 
    FOREIGN KEY (user_id) 
    REFERENCES users(id)
    ON UPDATE CASCADE
);

COMMIT; 