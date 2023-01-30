CREATE TABLE IF NOT EXISTS login_log
(
    player_id VARCHAR(255) NOT NULL REFERENCES player(player_id),
    login_at  TIMESTAMPTZ NOT NULL DEFAULT (now())
);