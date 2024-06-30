package person

import "github.com/stretchr/testify/mock"

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) FindById(id int64) (Person, bool) {
	args := m.Called(id)

	return args.Get(0).(Person), args.Bool(1)
}