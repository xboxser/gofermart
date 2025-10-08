-- Удаляем индексы основных полей
BEGIN;

DROP INDEX idx_users_email;
DROP INDEX idx_number_email;

COMMIT; 