CREATE DATABASE fib;

\c fib

CREATE TABLE IF NOT EXISTS memo (
  n INT NOT NULL UNIQUE, 
  val INT
);

CREATE UNIQUE INDEX n_unique ON memo (n);
CREATE INDEX val_idx ON memo (val);
