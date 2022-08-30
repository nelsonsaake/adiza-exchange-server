package archetypes

import "fmt"

var (
	// source: https://pravatar.cc/
	AvatarPlaceholder = "https://i.pravatar.cc/200?img=70"
)

func NewAvatarPlaceholder(size, id int) string {

	var _size interface{}

	if size == 0 {
		_size = ""
	} else {
		_size = size
	}

	return fmt.Sprintf("https://i.pravatar.cc/%v?img=%d", _size, id)
}
