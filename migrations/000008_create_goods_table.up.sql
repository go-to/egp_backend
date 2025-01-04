CREATE TABLE IF NOT EXISTS goods
(
    id           INT AUTO_INCREMENT COMMENT 'ID'
        PRIMARY KEY,
    event_id     INT                                 NOT NULL COMMENT 'イベントID',
    num_of_stamp INT                                 NOT NULL COMMENT 'スタンプ個数',
    goods_name   VARCHAR(255)                        NOT NULL COMMENT 'グッズ名',
    is_party     TINYINT(1)                          NOT NULL COMMENT 'パーティー参加基準',
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    CONSTRAINT goods_events_id_fk
        FOREIGN KEY (event_id) REFERENCES events (id)
)
    COMMENT 'グッズ';
