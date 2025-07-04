-- スキーマ作成
CREATE SCHEMA IF NOT EXISTS egp;
COMMENT ON SCHEMA egp IS 'ヱビスビールに合う逸品グランプリ';

-- postgis有効化
CREATE EXTENSION postgis WITH SCHEMA egp;
