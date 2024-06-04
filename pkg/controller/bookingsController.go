package controller

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/michee/pkg/models"
	"github.com/michee/pkg/utils"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
)

const STRIPE_KEY = "sk_test_51OvUhw07wkPDDzgNQxBvybs2qXEVRqRdXiLHXJt85kwMPKtT7YiVTvTpdwXPZJyQRJbQFPVT23aGCY2XLkCk8bla00BR6eOqgq"

var NewBooking models.Booking

func init() {
	stripe.Key = STRIPE_KEY
}

func generatePaymentID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}


func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	startDate, err := time.Parse("2006-01-02", booking.StartDate)
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		return
	}
	endDate, err := time.Parse("2006-01-02", booking.EndDate)
	if err != nil {
		http.Error(w, "Invalid end date format", http.StatusBadRequest)
		return
	}

	duration := endDate.Sub(startDate)
	numDays := int(duration.Hours() / 24) 

	roomId := booking.RoomId

	room, _ := models.GetRoomById(roomId)
	
	roomPrice := room.RoomPrice


	roomPriceInt, err := strconv.Atoi(roomPrice)
	if err != nil {
			http.Error(w, "Failed to convert room price to integer", http.StatusInternalServerError)
			return
	}
	
	totalPrice := roomPriceInt * numDays
	
	booking.TotalPrice = strconv.Itoa(totalPrice)

	paymentID, err := generatePaymentID()
	if err != nil {
		http.Error(w, "Failed to generate payment ID", http.StatusInternalServerError)
		return
	}

	booking.PaymentId = paymentID
	booking.PaymentStatus = true

	booking.CreateBooking()

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(booking.Currency),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Booking for " + booking.UserName),
					},
					UnitAmount: stripe.Int64(int64(totalPrice)), 
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String("payment"),
		SuccessURL: stripe.String("http://localhost:5000/booking/success?session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String("http://localhost:5000/booking/cancel"),
	}

	s, err := session.New(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(s)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetBooking(w http.ResponseWriter, r *http.Request) {
	bookings := models.GetAllBooking()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

func GetBookingById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookingId := params["bookingId"]
	booking, _ := models.GetBookingById(bookingId)
	res, _ := json.Marshal(booking)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	bookingUpdate := models.Booking{}
	utils.ParseBody(r, &bookingUpdate)
	booking := mux.Vars(r)
	bookingId := booking["bookingId"]
	bookingDetail, db := models.GetBookingById(bookingId)

	if bookingUpdate.Currency != ""{
		bookingDetail.Currency = bookingUpdate.Currency
	}
	if bookingUpdate.StartDate != ""{
		bookingDetail.StartDate = bookingUpdate.StartDate
	}
	if bookingUpdate.EndDate != ""{
		bookingDetail.EndDate = bookingUpdate.EndDate
	}
	if bookingUpdate.UserName != ""{
		bookingDetail.UserName = bookingUpdate.UserName
	}
	if bookingUpdate.UserEmail != ""{
		bookingDetail.UserEmail = bookingUpdate.UserEmail
	}

	db.Save(&bookingDetail)
	res, _ := json.Marshal(bookingDetail)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookingId := params["id"]
	books := models.DeleteById(bookingId)
	res, _ := json.Marshal(books)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	w.WriteHeader(http.StatusNoContent)
}
