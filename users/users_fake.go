package users

import (
	"projects/adiza-exchange-server/archetypes"
	"projects/adiza-exchange-server/gen"
	"projects/adiza-exchange-server/models"
	"projects/adiza-exchange-server/passwords"

	"github.com/icrowley/fake"
)

func Fake() (user models.User) {
	user = models.User{
		FirstName:   fake.FirstName(),
		LastName:    fake.LastName(),
		JobTitle:    fake.JobTitle(),
		Email:       fake.EmailAddress(),
		PhoneNumber: fake.Phone(),
		Password:    passwords.New(),
		Photo:       archetypes.NewAvatarPlaceholder(200, gen.Int(70)),
	}
	return
}
