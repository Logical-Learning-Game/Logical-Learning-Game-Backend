CREATE TABLE IF NOT EXISTS player
(
  player_id VARCHAR(255) PRIMARY KEY,
  email     VARCHAR(255) DEFAULT NULL,
  name      VARCHAR(255) DEFAULT NULL
);