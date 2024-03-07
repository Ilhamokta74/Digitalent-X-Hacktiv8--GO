package main

import (
	"Gorm/database"
	"Gorm/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	createUser("johndoe@gmail.com")
	getUserById(1)
	updateUserById(1, "johnjohn@gmail.com")
	createProduct(1, "YLO", "YYYY")
	getUserWithProducts()
	deleteProductById(1)
}

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data : ", err)
		return
	}

	fmt.Println("New User Data : ", User)
}

func getUserById(id uint) {
	db := database.GetDB()

	User := models.User{}

	err := db.First(&User, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error Finding User : ", err)
	}

	fmt.Printf("User Data : %+v \n", User)
}

func updateUserById(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error Updated User Data : ", err)
		return
	}

	fmt.Printf("Update user's Email : %+v \n", user.Email)
}

func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	product := models.Product{
		UserID: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&product).Error

	if err != nil {
		fmt.Println("Error Create Product Data : ", err.Error())
		return
	}

	fmt.Println("New Product Data : ", product)
}

func getUserWithProducts() {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error Getting User Datas With Products : ", err.Error())
		return
	}

	fmt.Println("User Datas With Products")
	fmt.Printf("%+v", users)
}

func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}
	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error Deleting Products : ", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted", id)
}
