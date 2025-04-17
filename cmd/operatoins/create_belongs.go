package operatoins

import (
	"fmt"

	"github.com/Akashkarmokar/gorm_setup/cmd/models"
	"gorm.io/gorm"
)

func OperationForBelongToModel(db *gorm.DB) {

	var doesExistUser models.User
	result := db.Where("name = ?", "Helsenki").Find(&doesExistUser)
	if result.RowsAffected == 0 {
		createdUser := models.User{
			Name: "Helsenki",
		}

		db.Create(&createdUser)
		fmt.Println("Profile Code: ", createdUser.ProfileCode)

	}

	var doesExistProfile models.Profile
	result = db.Where("first_name = ?", "Ex. First Name").Find(&doesExistProfile)
	if result.RowsAffected == 0 {
		createdProfile := models.Profile{
			FirstName: "Ex. First Name",
			LastName:  "Ex. Last Name",
		}

		db.Create(&createdProfile)

		var users models.User
		db.Where("name = ?", "Helsenki").Find(&users).Limit(1)

		// Update User to store profile Id
		db.Model(models.User{}).Where("name = ?", "Helsenki").Update("profile_code",  createdProfile.ID)

	}

}
