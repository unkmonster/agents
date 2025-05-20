CREATE TABLE IF NOT EXISTS domain_daily_stats(
    domain VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    register_count INT NOT NULL DEFAULT 0,
    recharge_amount INT NOT NULL DEFAULT 0,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE(date, domain),
    INDEX idx_domain (domain)
);