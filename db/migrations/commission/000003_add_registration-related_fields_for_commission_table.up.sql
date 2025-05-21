ALTER TABLE user_commissions 
    ADD COLUMN today_registration_count INT NOT NULL DEFAULT 0,
    ADD COLUMN total_registration_count INT NOT NULL DEFAULT 0;