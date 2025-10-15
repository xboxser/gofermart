-- Создаем индексы на поля основные
BEGIN;

CREATE INDEX idx_users_email ON users (login);
CREATE INDEX idx_number_email ON orders (number);

COMMIT; 