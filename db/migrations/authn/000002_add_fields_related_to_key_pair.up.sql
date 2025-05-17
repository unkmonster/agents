ALTER TABLE user_credentials ADD COLUMN alg VARCHAR(10) NULL;
ALTER TABLE user_credentials ADD COLUMN public_key TEXT NULL; 
ALTER TABLE user_credentials ADD COLUMN private_key TEXT NULL;
ALTER TABLE user_credentials ADD COLUMN secret TEXT NULL;