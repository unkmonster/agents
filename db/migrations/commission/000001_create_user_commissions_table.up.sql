CREATE TABLE IF NOT EXISTS user_commissions (
    id CHAR(36) UNIQUE PRIMARY KEY NOT NULL,
    user_id CHAR(36) UNIQUE NOT NULL,
    today_commission INT NOT NULL DEFAULT 0,
    total_commission INT NOT NULL DEFAULT 0,
    settled_commission INT NOT NULL DEFAULT 0
);