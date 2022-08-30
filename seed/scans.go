package seed

import (
	"projects/adiza-exchange-server/gen"
	"projects/adiza-exchange-server/models"
	"projects/adiza-exchange-server/repo"
	"projects/adiza-exchange-server/userids"
)

func Scans(count int) {
	userIDs, err := userids.GetIDs()
	if err != nil {
		panic(err)
	}

	defer repo.Logger.LogMode(repo.Logger.Mode())
	repo.Logger.WarnMode()

	for i := 0; i < count; i++ {

		x := gen.Int(len(userIDs))
		y := gen.Int(len(userIDs))

		repo.DB.Save(
			&models.Scan{
				UserID:        userIDs[x],
				ScannedUserID: userIDs[y],
			},
		)
	}
}
