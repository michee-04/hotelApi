package models

type Room struct{
	Title string `json:"title"`
	Description string `json:"description"`
	BedCount string `json:"bedCount"`
	Image string `json:"image"`
	RoomPrice string `json:"roomPrice"`
	RoomService bool `json:"roomService"`
	OceanView bool `json:"oceanView"`
	CityView bool `json:"cityView"`

	HotelId string
	Hotels Hotel `gorm:"foreignKey:HotelId;references:Id"`

}