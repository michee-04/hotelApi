package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/michee/pkg/config"
)

type Hotel struct {
	ID        string `gorm:"primary_key"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	Country      string `json:"country"`
	State        string `json:"state"`
	City         string `json:"city"`
	Localisation string `json:"localisation"`
	Restaurant   bool   `json:"restaurant"`
	UserId string `json:"userId"`
	CreatedAt time.Time 
	UpdatedAt time.Time  
	DeletedAt *time.Time
	Rooms        []Room `gorm:"foreignKey:HotelId"`
	Bookings []Booking `gorm:"foreignKey:HotelIdB"`
	User       User     `gorm:"foreignKey:UserId" json:"-"`
}

func init() {
	config.Connect()
	DBS = config.GetDB()
	DBS.DropTableIfExists(&Hotel{})
	DBS.AutoMigrate(&Hotel{})
}

func (user *Hotel) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New().String()
	return scope.SetColumn("ID", id)
}

func (h *Hotel) CreateHotel() *Hotel {
	DBS.NewRecord(h)
	DBS.Create(&h)
	return h
}

func GetAllHotel() []Hotel {
	var Hotels []Hotel
	DBS.Preload("Rooms").Preload("Bookings").Find(&Hotels)
	return Hotels
}

func GetHotelById(Id string) (*Hotel, *gorm.DB) {
	var getHotel Hotel
	db := DBS.Preload("Rooms").Preload("Rooms").Where("ID = ?", Id).Find(&getHotel)
	return &getHotel, db
}

func DeleteHotelId(Id string) Hotel {
	var hotel Hotel
	DBS.Where("ID = ?", Id).Delete(&hotel)
	return hotel
}
