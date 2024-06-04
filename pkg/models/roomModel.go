package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/michee/pkg/config"
)

type Room struct {
	RoomID      string    `gorm:"primary_key" json:"roomId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	BedCount    string    `json:"bedCount"`
	Image       string    `json:"image"`
	RoomPrice   string    `json:"roomPrice"`
	RoomService bool      `json:"roomService"`
	OceanView   bool      `json:"oceanView"`
	CityView    bool      `json:"cityView"`
	HotelID     string      `json:"hotelId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
	Hotel       Hotel     `gorm:"foreignKey:HotelID" json:"-"`
	Bookings    []Booking `gorm:"foreignKey:RoomId"`
}

func init() {
	config.Connect()
	DBS = config.GetDB()
	DBS.DropTableIfExists(&Room{})
	DBS.AutoMigrate(&Room{})
}

func (r *Room) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New().String()
	return scope.SetColumn("RoomID", id)
}

func (r *Room) CreateRoom() *Room {
	DBS.NewRecord(r)
	DBS.Create(&r)
	return r
}

func GetAllRoom() []Room {
	var rooms []Room
	DBS.Preload("Bookings").Find(&rooms)
	return rooms
}

func GetRoomById(roomId string) (*Room, *gorm.DB) {
	var getRoom Room
	db := DBS.Preload("Bookings").Where("room_id = ?", roomId).Find(&getRoom)
	return &getRoom, db
}

func DeleteRoomById(roomId string) Room {
	var room Room
	DBS.Where("room_id = ?", roomId).Delete(&room)
	return room
}
