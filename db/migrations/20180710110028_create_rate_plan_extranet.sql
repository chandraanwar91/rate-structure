
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE rate_plan_extranet (
  id int(10) NOT NULL AUTO_INCREMENT,
  room_id bigint(20) unsigned NOT NULL,
  room_type varchar(255) COLLATE utf8_unicode_ci,
  hotel_id bigint(20) unsigned NOT NULL,
  travel_start date NOT NULL,
  travel_end date NOT NULL,
  travel_days set('1','2','3','4','5','6','7') COLLATE utf8_unicode_ci DEFAULT NULL,
  booking_start date NOT NULL,
  booking_end date NOT NULL,
  booking_days set('1','2','3','4','5','6','7') COLLATE utf8_unicode_ci DEFAULT NULL,
  currency char(3) COLLATE utf8_unicode_ci NOT NULL DEFAULT 'IDR',
  discount_type enum('percentage','amount','none') COLLATE utf8_unicode_ci NOT NULL DEFAULT 'none',
  nominal decimal(12,2) DEFAULT '0.00',
  commission_added decimal(12,2) DEFAULT '0.00',
  minimum_stay tinyint(2) unsigned NOT NULL DEFAULT '1',
  cancellation_policy enum('non_refundable','follow_hotel_policy') COLLATE utf8_unicode_ci NOT NULL DEFAULT 'non_refundable',
  applicable_to enum('apps','member') COLLATE utf8_unicode_ci NOT NULL,
  status enum('active','disabled','removed') COLLATE utf8_unicode_ci NOT NULL DEFAULT 'active',
  created_by mediumint(6) unsigned NOT NULL DEFAULT '0',
  modified_by mediumint(6) unsigned NOT NULL DEFAULT '0',
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (id),
  KEY hotel_id (hotel_id),
  KEY room_id (room_id),
  KEY travel_startdate (travel_start),
  KEY travel_enddate (travel_end),
  KEY booking_startdate (booking_start),
  KEY booking_enddate (booking_end),
  KEY travel_days (travel_days),
  KEY booking_days (booking_days)
);



-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE rate_plan_extranet;