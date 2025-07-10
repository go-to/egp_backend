CREATE INDEX idx_shops_event_id ON shops(event_id);
CREATE INDEX idx_shops_location_shop_id ON shops_location(shop_id);
CREATE INDEX idx_shops_time_composite ON shops_time(shop_id, week_number, day_of_week, is_holiday);
CREATE INDEX idx_stamps_composite ON stamps(shop_id, user_id, deleted_at);
