package seed

import (
	"projects/adiza-exchange-server/repo"
	"projects/adiza-exchange-server/users"
)

func Users(count int) {
	for i := 0; i < count; i++ {
		user := users.Fake()
		repo.DB.Save(&user)
	}
}
