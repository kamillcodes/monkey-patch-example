package person

import "fmt"

type IService interface {
	Init() error
	FindById(id int64) (*Person, bool)
}

type Service struct {
	repo IRepo
}

func NewService(repo IRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindById(id int64) (Person, bool) {
	person, ok := s.repo.FindById(id)
	if !ok {
		return person, ok
	}

	person.Age = person.Age * 2
	person.Name = fmt.Sprintf("Awesome %s", person.Name)

	return person, ok
}