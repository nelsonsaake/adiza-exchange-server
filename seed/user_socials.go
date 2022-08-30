package seed

import (
	"projects/adiza-exchange-server/gen"
	"projects/adiza-exchange-server/models"
	"projects/adiza-exchange-server/repo"
	"projects/adiza-exchange-server/smpids"
	"projects/adiza-exchange-server/userids"

	"github.com/bxcodec/faker"
)

func UserSocials(count int) {

	userIDs, err := userids.GetIDs()
	if err != nil {
		panic(err)
	}

	smpIDs, err := smpids.GetIDs()
	if err != nil {
		panic(err)
	}

	defer repo.Logger.LogMode(repo.Logger.Mode())
	repo.Logger.WarnMode()

	for i := 0; i < count; i++ {

		x := gen.Int(len(userIDs))
		y := gen.Int(len(smpIDs))
		handle := faker.NAME + gen.Code(5)

		repo.DB.Save(
			&models.UserSocial{
				UserID:                userIDs[x],
				SocialMediaPlatformID: smpIDs[y],
				Handle:                handle,
			},
		)
	}
}
