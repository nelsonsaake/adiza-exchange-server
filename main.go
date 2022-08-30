package main

import (
	"fmt"
	"projects/adiza-exchange-server/models"
	"projects/adiza-exchange-server/repo"
	"projects/adiza-exchange-server/seed"
)

func PrintRepoUser(id uint) {
	user := models.User{}
	user.ID = id

	err := repo.DB.First(&user).Error
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}

func SeedRepo() {
	seed.Users(5)
	seed.UserSocials(30)
	seed.Scans(30)
}

func main() {

	SeedRepo()

	users := []models.User{}
	scans := []models.Scan{}

	repo.DB.Find(&users)
	repo.DB.Find(&scans)

	fmt.Println("users.len: ", len(users))
	fmt.Println("scans.len: ", len(scans))

	PrintRepoUser(1)

	repo.DB.Delete(&models.User{}, 1)

	PrintRepoUser(1)

	repo.DB.Find(&users)
	repo.DB.Find(&scans)

	fmt.Println("users.len: ", len(users))
	fmt.Println("scans.len: ", len(scans))

}
