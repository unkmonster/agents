ALTER TABLE user_credentials
  DROP COLUMN alg,
  DROP COLUMN public_key,
  DROP COLUMN private_key,
  DROP COLUMN secret,
  DROP COLUMN token_key;
