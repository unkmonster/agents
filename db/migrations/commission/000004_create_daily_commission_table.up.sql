CREATE TABLE IF NOT EXISTS daily_user_commissions (
    date DATE NOT NULL,
    user_id CHAR(36) NOT NULL,
    indirect_recharge_amount INT NOT NULL DEFAULT 0,
    indirect_registration_count INT NOT NULL DEFAULT 0,
    direct_recharge_amount INT NOT NULL DEFAULT 0,
    direct_registration_count INT NOT NULL DEFAULT 0,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_user_id (user_id),
    UNIQUE(user_id, date)
);