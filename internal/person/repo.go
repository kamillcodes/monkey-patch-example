package person

type IRepo interface {
	FindById(id int64) (Person, bool)
}

type Repository struct {
	records map[int64]*Person
}

func NewRepository() *Repository {
	return &Repository{}
}

func (repo *Repository) FindById(id int64) (Person, bool) {
	person, ok := repo.records[id]

	return *person, ok
}
