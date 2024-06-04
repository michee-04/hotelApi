package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/michee/pkg/config"
)

type Booking struct{
	ID        string `gorm:"primary_key"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	Country      string `json:"country"`
	State        string `json:"state"`
	City         string `json:"city"`
	Localisation string `json:"localisation"`
	Restaurant   bool   `json:"restaurant"`
	HotelIdB string `json:"hotelIdB"`
	RoomId string `json:"roomId"`
	CreatedAt time.Time 
	UpdatedAt time.Time  
	DeletedAt *time.Time
	Hotel       Hotel     `gorm:"foreignKey:HotelIDB" json:"-"`
	Rooms       Room     `gorm:"foreignKey:RoomId" json:"-"`
}


func init() {
	config.Connect()
	DBS = config.GetDB()
	DBS.DropTableIfExists(&Room{})
	DBS.AutoMigrate(&Room{})
}

func (r *Booking) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New().String()
	return scope.SetColumn("RoomID", id)
}

func (b *Booking) CreateBooking() *Booking {
	DBS.NewRecord(b)
	DBS.Create(&b)
	return b
}

func GetAllBooking() []Booking {
	var bookings []Booking
	DBS.Find(&bookings)
	return bookings
}

func GetBookingById(bookingId string) (*Booking, *gorm.DB) {
	var getBooking Booking
	db := DBS.Where("booking_id = ?", bookingId).Find(&getBooking)
	return &getBooking, db
}

func DeleteById(bookingId string) Room {
	var booking Room
	DBS.Where("booking_id = ?", bookingId).Delete(&booking)
	return booking
}
