ALTER TABLE egp.shops
    ALTER COLUMN no TYPE varchar(255) USING no::varchar(255);
