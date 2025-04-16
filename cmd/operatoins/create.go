package operatoins

import (
	"github.com/Akashkarmokar/gorm_setup/cmd/models"
	"gorm.io/gorm"
)

func InsertToTheDataBase(db *gorm.DB) {

	// Without Nullable Column Value
	createdUser := models.User{
		Name: "Jhon",
	}

	// insert To the DB
	// result := db.Create(&createdUser)

	// fmt.Println(createdUser.ID)
	// fmt.Println(createdUser.Name)
	// fmt.Println(result.Error)
	// fmt.Println(result.RowsAffected)
	db.Create(&createdUser)

	// To insert Multiple at a time
	makeUser := func(name string, email string) *models.User {
		user := models.User{}
		if name != "" {
			user.Name = name
		}
		if email != "" {
			user.Email = &email
		}
		return &user

	}
	users := []*models.User{
		makeUser("Jhon", ""),
		makeUser("Example User", ""),
		makeUser("User with email", "email address"),
	}
	db.Create(users)
}
