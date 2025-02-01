CREATE TABLE IF NOT EXISTS config
(
    id         INT AUTO_INCREMENT COMMENT 'ID'
        PRIMARY KEY,
    conf_name  VARCHAR(255)                        NOT NULL COMMENT '設定項目名',
    conf_value VARCHAR(255)                        NOT NULL COMMENT '設定値',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL
) comment '設定値';
