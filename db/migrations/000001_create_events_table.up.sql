CREATE TABLE IF NOT EXISTS events
(
    id         INT AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    name       VARCHAR(255) NOT NULL COMMENT 'イベント名',
    year       INT          NOT NULL COMMENT '開催年',
    start_date DATE         NOT NULL COMMENT '開始日',
    end_date   DATE         NOT NULL COMMENT '終了日',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
    COMMENT 'イベント';
