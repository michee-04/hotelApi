package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/michee/pkg/config"
)


type Booking struct {
	ID             string    `gorm:"primary_key"`
	UserName       string    `json:"userName"`
	UserEmail      string    `json:"userEmail"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	Currency       string    `json:"currency"`
	TotalPrice     string    `json:"totalPrice"`
	HotelIdB       string    `json:"hotelIdB"`
	HotelOwnerId   string    `json:"hotelOwnerId"`
	RoomId         string    `json:"roomId"`
	PaymentStatus  bool      `json:"paymentStatus"`
	PaymentId      string    `json:"paymentId"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	Hotel          Hotel     `gorm:"foreignKey:HotelIdB" json:"-"`
	User           User      `gorm:"foreignKey:HotelOwnerId" json:"-"`
	Rooms          Room      `gorm:"foreignKey:RoomId" json:"-"`
}

func init() {
	config.Connect()
	DBS = config.GetDB()
	DBS.DropTableIfExists(&Booking{})
	DBS.AutoMigrate(&Booking{})
}

func (r *Booking) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New().String()
	return scope.SetColumn("ID", id) 
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
	db := DBS.Where("id = ?", bookingId).Find(&getBooking)
	return &getBooking, db
}

func DeleteById(bookingId string) Booking {
	var booking Booking
	DBS.Where("id = ?", bookingId).Delete(&booking)
	return booking
}
