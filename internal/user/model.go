package user

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}