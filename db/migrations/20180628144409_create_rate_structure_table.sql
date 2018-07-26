
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE rate_backoffice (
  id int(10) NOT NULL AUTO_INCREMENT,
  booking_start date NOT NULL,
  booking_end date NOT NULL,
  staying_start date NOT NULL,
  staying_end date NOT NULL,
  discount_type enum('amount','percentage') COLLATE utf8_unicode_ci NOT NULL,
  currency enum('IDR','USD') COLLATE utf8_unicode_ci NOT NULL,
  nominal decimal(12,2) unsigned NOT NULL DEFAULT '0.00',
  available_to set('direct_contract','expedia','hotelbeds') COLLATE utf8_unicode_ci NOT NULL,
  applicable_to enum('apps','member') COLLATE utf8_unicode_ci NOT NULL,
  country_id text COLLATE utf8_unicode_ci,
  city_id text COLLATE utf8_unicode_ci,
  hotel_group_id text COLLATE utf8_unicode_ci,
  hotel_id text COLLATE utf8_unicode_ci,
  star_rating text COLLATE utf8_unicode_ci,
  hotel_tag_id text COLLATE utf8_unicode_ci,
  description text COLLATE utf8_unicode_ci NOT NULL,
  status enum('enable','disable','removed') COLLATE utf8_unicode_ci NOT NULL,
  created_by mediumint(6) unsigned NOT NULL DEFAULT 0,
  modified_by mediumint(6) unsigned NOT NULL DEFAULT 0,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE rate_backoffice;