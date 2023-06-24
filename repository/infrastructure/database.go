package infrastructre

import (
	"fmt"
	"zog/delivery/models"
	"zog/domain/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dsn string
var KEY5 = "host=localhost user=edwin dbname=edwin password=acid port=5432 sslmode=disable"

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	dsn = os.Getenv("KEY5")

// }
func ConnectToDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(KEY5), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	DB = db
	DB.AutoMigrate(&entity.TicketDetails{}, &entity.OtpKey{}, &models.Signup{}, &entity.Admin{}, &entity.User{}, &entity.Ticket{}, &entity.Apparel{}, &entity.CartItem{}, &entity.Cart{}, &entity.Wishlist{}, &entity.Order{}, &entity.OrderItem{}, &entity.Address{}, &entity.Inventory{}, &entity.Invoice{}, &entity.Return{}, &entity.Coupon{}, &entity.UsedCoupon{}, &entity.Offer{})
	return db, nil
}
