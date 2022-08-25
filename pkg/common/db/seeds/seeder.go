package seeds

import (
	"errors"
	"fmt"

	"github.com/borisa99/gin-starter/pkg/common/models"
	"gorm.io/gorm"
)

type Model interface {
	models.Role | models.User | models.UserRole
}

func Plant(db *gorm.DB) {
	roles := []models.Role{
		{Name: "admin"},
		{Name: "base"},
	}

	values, err := Seeder(db, models.Role{}, roles)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(values)
		users := []models.User{
			{FirstName: "admin",
				LastName:     "admin",
				Email:        "admin@mail.com",
				PasswordHash: "$2a$10$npPMoi7q5PWy3RBZtI.zteisiP28MYP0cttG35Hqsqj9gB74uv54S",
				Roles: []models.UserRole{
					{RoleId: values[0].ID},
				},
			},
		}
		Seeder(db, models.User{}, users)

	}

}

func Seeder[M Model](db *gorm.DB, model M, values []M) ([]M, error) {
	ok := db.Migrator().HasTable(&model)
	if !ok {
		return nil, errors.New("table does not exist")
	}

	var exists bool
	db.Model(&model).Select("count(*) > 0").Find(&exists)

	if !exists {
		db.Create(&values)
		return values, nil
	}

	return nil, errors.New("rows already exist")
}
