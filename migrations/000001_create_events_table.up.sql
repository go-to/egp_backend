CREATE TABLE IF NOT EXISTS egp.events
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    year       INT          NOT NULL,
    start_date DATE         NOT NULL,
    end_date   DATE         NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE egp.events IS 'イベント';
COMMENT ON COLUMN egp.events.id IS 'ID';
COMMENT ON COLUMN egp.events.name IS 'イベント名';
COMMENT ON COLUMN egp.events.year IS '開催年';
COMMENT ON COLUMN egp.events.start_date IS '開始日';
COMMENT ON COLUMN egp.events.end_date IS '終了日';
