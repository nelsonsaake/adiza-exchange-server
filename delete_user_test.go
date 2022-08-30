package main

import (
	"fmt"
	"projects/adiza-exchange-server/models"
	"projects/adiza-exchange-server/repo"
	"testing"
)

func usersStuff(userID uint, t *testing.T) ([]models.UserSocial, []models.Scan) {

	var (
		socials = []models.UserSocial{}
		links   = []models.Scan{}
		err     error
	)

	err = repo.DB.Find(&socials, map[string]interface{}{"user_id": userID}).Error
	if err != nil {
		t.Error(err)
	}

	err = repo.DB.Find(&links, map[string]interface{}{"user_id": userID, "scanned_user_id": userID}).Error
	if err != nil {
		t.Error(err)
	}

	fmt.Println("no of user socials: ", len(socials))

	fmt.Println("no of user links: ", len(links))

	fmt.Println()

	return socials, links
}

func TestDeleteUser(t *testing.T) {

	var (
		user    = &models.User{}
		socials = []models.UserSocial{}
		links   = []models.Scan{}
	)

	err := repo.DB.First(user).Error
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(user)

	fmt.Println("BEFORE DELETE")

	usersStuff(user.ID, t)

	err = repo.DB.Unscoped().Delete(user).Error
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("AFTER DELETE")

	socials, links = usersStuff(user.ID, t)

	if len(socials) != 0 {
		t.Error("user socials still exists after user deleted")
	}

	if len(links) != 0 {
		t.Error("user links still exists after user deletes")
	}
}
