package user

import "fmt"

var (
	FindById = func(id int64) (User, bool) {
		user, ok := findById(id)
		if !ok {
			return user, ok
		}

		user.Age = user.Age * 2
		user.Name = fmt.Sprintf("Awesome %s", user.Name)

		return user, ok
	}
)