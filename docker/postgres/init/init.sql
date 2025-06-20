-- DB作成
CREATE DATABASE IF NOT EXISTS egp;

-- スキーマ作成
CREATE SCHEMA IF NOT EXISTS egp;
COMMENT ON SCHEMA egp IS 'ヱビスビールに合う逸品グランプリ';

-- postgis有効化
CREATE EXTENSION postgis WITH SCHEMA egp;
