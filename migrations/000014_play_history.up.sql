CREATE TABLE IF NOT EXISTS play_history
(
    id                BIGSERIAL PRIMARY KEY,
    game_session_id   BIGINT      NOT NULL REFERENCES game_session (id),
    action_step       INTEGER     NOT NULL,
    number_of_command INTEGER     NOT NULL,
    is_finited        BOOLEAN     NOT NULL,
    is_completed      BOOLEAN     NOT NULL,
    submit_datetime   TIMESTAMPTZ NOT NULL DEFAULT (now())
);