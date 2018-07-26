package models

import(
	"time"
)
type (
	Rate struct {
		CheckIn     time.Time   `json:"check_in"`
		CheckOut	time.Time	`json:"check_out"`
		BookingDate	time.Time	`json:"booking_date"`
		RoomIds		string  `json:"room_ids"`
		Platform	string	`json:"platform"`
		IsUserLogin	int	`json:"is_user_login"`
	}
)

