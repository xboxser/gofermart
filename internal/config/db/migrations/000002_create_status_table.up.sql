-- Создание таблицы со статусами заказа
BEGIN;

CREATE TABLE statuses (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

INSERT INTO statuses (name) VALUES ('NEW'), ('PROCESSING'), ('INVALID'), ('PROCESSED');
COMMIT;