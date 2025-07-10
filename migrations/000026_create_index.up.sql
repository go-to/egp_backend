CREATE INDEX idx_shops_event_year ON events(year);
CREATE INDEX idx_shops_no_name_menu ON shops(no, shop_name, menu_name);
CREATE INDEX idx_shops_location_spatial ON shops_location USING GIST(location);
CREATE INDEX idx_shops_time_temporal ON shops_time(start_time, end_time);
