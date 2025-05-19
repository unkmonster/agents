CREATE TABLE IF NOT EXISTS user_wallets (
    id CHAR(36) NOT NULL PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    wallet_type ENUM('alipay', 'wxpay', 'tron') NOT NULL,
    account VARCHAR(255) NULL,
    qr_code TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, wallet_type),
    INDEX idx_user_id (user_id)
);