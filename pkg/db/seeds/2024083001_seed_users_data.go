package seeds

import (
	"time"

	"github.com/deveshmishra34/groot/pkg/db/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {

	var s = &gormigrate.Migration{}
	s.ID = "2024083001_seed_users_data"

	s.Migrate = func(db *gorm.DB) error {

		var err error
		var users []*models.User = []*models.User{}

		users = append(users, &models.User{
			FirstName: "John",
			LastName:  "Doe",
			Phone:     "8800000001",
			Username:  "8800000001",
			Email:     "johndoe@gmail.com",
			Status:    "active",
			Password:  "",
			Otp:       "123456",
			OtpExpiry: time.Now().Add(15 * time.Minute),
		})

		users = append(users, &models.User{
			FirstName: "Rakesh",
			LastName:  "Goyal",
			Phone:     "8800000002",
			Username:  "8800000002",
			Email:     "rakeshgoyal@gmail.com",
			Status:    "locked",
			Password:  "",
			Otp:       "123456",
			OtpExpiry: time.Now().Add(15 * time.Minute),
		})

		for _, user := range users {
			if user.Password != "" {
				// password := []byte(user.Password)
				// hashedPassword, err := bcrypt.GenerateFromPassword(
				// 	password,
				// 	rand.Intn(bcrypt.MaxCost-bcrypt.MinCost)+bcrypt.MinCost,
				// )
				// user.Password = string(hashedPassword)
				if err != nil {
					logFail(s.ID, err)
					return err
				}

			}
			err = user.Save()
			if err != nil {
				logFail(s.ID, err)
				return err
			}
		}

		logSuccess(s.ID)
		return nil
	}

	AddSeed(s)
}
