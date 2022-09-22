CREATE TABLE player
(
    player_id VARCHAR(255) NOT NULL,
    email     VARCHAR(255) DEFAULT NULL,
    PRIMARY KEY (player_id)
);

CREATE TABLE login_log
(
    player_id  VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    FOREIGN KEY (player_id) REFERENCES player (player_id)
);
