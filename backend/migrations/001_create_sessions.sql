CREATE TABLE IF NOT EXISTS user_sessions (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    channel_name VARCHAR(100) NOT NULL,
    session_token VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_seen_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_user_sessions_username ON user_sessions (username);

CREATE INDEX IF NOT EXISTS idx_user_sessions_last_seen_at ON user_sessions (last_seen_at);