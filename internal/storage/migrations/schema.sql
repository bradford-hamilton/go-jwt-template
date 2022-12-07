-- Create table for users --
CREATE TABLE IF NOT EXISTS users (
  id SERIAL NOT NULL PRIMARY KEY,
  first_name VARCHAR (50) NOT NULL,
  last_name VARCHAR (50) NOT NULL,
  email VARCHAR (50) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create function for keeping `updated_at` updated with current timestamp
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS trigger AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$
    LANGUAGE 'plpgsql';

-- Create trigger to call the function we created to keep updated_at updated
DROP trigger IF EXISTS users_updated_at ON users;
CREATE trigger users_updated_at
    BEFORE UPDATE ON users FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

-- Create a unique index across the appropriate columns
CREATE UNIQUE INDEX unique_users_idx
    ON users (email);
