package user

var (
	records map[int64]*User

	findById = func(id int64) (User, bool) {
		user, ok := records[id]

		return *user, ok
	}
)
