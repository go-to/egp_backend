CREATE TABLE IF NOT EXISTS egp.goods
(
    id           SERIAL PRIMARY KEY,
    event_id     INT                                 NOT NULL,
    num_of_stamp INT                                 NOT NULL,
    goods_name   VARCHAR(255)                        NOT NULL,
    is_party     BOOLEAN                             NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    CONSTRAINT goods_events_id_fk
        FOREIGN KEY (event_id) REFERENCES egp.events (id)
);
COMMENT ON TABLE egp.goods IS 'グッズ';
COMMENT ON COLUMN egp.goods.id IS 'ID';
COMMENT ON COLUMN egp.goods.event_id IS 'イベントID';
COMMENT ON COLUMN egp.goods.num_of_stamp IS 'スタンプ個数';
COMMENT ON COLUMN egp.goods.goods_name IS 'グッズ名';
COMMENT ON COLUMN egp.goods.is_party IS 'パーティー参加基準';
