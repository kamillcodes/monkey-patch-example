package person

import (
	"fmt"
	"mpexample/internal/assertutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindById_NotFound(t *testing.T) {
	expectedPerson:= Person{}

	mockRepo := new(MockRepo)
	mockRepo.On("FindById", int64(123)).Return(expectedPerson, false)

	service := NewService(mockRepo)
	person, found :=service.FindById(123)

	assert.False(t, found)
	assert.Equal(t, expectedPerson, person)
}

func Test_FindById_Found(t *testing.T) {
	basePerson := Person {
		Id: 0,
		Name: "Jimmy McJim",
		Age: 123,
		Job: "teacher",
	}
	expectedPerson := Person{
		Id: basePerson.Id,
		Name: fmt.Sprintf("Awesome %s", basePerson.Name),
		Age: basePerson.Age * 2,
		Job: basePerson.Job,
	}

	mockRepo := new(MockRepo)
	mockRepo.On("FindById", int64(123)).Return(basePerson, true)

	service := NewService(mockRepo)
	person, found :=service.FindById(123)

	assert.True(t, found)

	assertutils.AssertEqualIgnoringFields(t, expectedPerson, person, "Name", "Age")
	assert.Equal(t, expectedPerson.Name, person.Name)
	assert.Equal(t, expectedPerson.Age, person.Age)
}