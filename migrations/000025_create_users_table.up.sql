CREATE TABLE IF NOT EXISTS egp.users
(
    id                SERIAL
        PRIMARY KEY,
    user_id           VARCHAR(255) NOT NULL,
    anonymous_user_id VARCHAR(255) NOT NULL,
    display_name      VARCHAR(255) NOT NULL,
    status            INTEGER      NOT NULL,
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at        TIMESTAMP
);
COMMENT ON TABLE egp.users IS 'ユーザー';
COMMENT ON COLUMN egp.users.id IS 'ID';
COMMENT ON COLUMN egp.users.user_id IS 'ユーザーID';
COMMENT ON COLUMN egp.users.anonymous_user_id IS '匿名ユーザーID';
COMMENT ON COLUMN egp.users.display_name IS 'ユーザー名';
COMMENT ON COLUMN egp.users.status IS 'ステータス';
