-- Дополнительные индексы для таблицы users
CREATE INDEX IF NOT EXISTS idx_users_name ON users(name);
CREATE INDEX IF NOT EXISTS idx_users_age ON users(age);
CREATE INDEX IF NOT EXISTS idx_users_email_deleted_at ON users(email, deleted_at);
