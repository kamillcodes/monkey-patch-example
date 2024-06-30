package user

import (
	"fmt"
	"mpexample/internal/assertutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindById_NotFound(t *testing.T) {
	originalFindById := findById
	defer func() {
		findById = originalFindById
	}()

	expectedUser := User{}
	findById = func(id int64) (User, bool) {
		return expectedUser, false
	}

	user, found := FindById(123)

	assert.False(t, found)
	assert.Equal(t, expectedUser, user)
}

func Test_FindById_Found(t *testing.T) {
	originalFindById := findById
	defer func() {
		findById = originalFindById
	}()

	baseUser := User {
		Id: 0,
		Name: "Jimmy McJim",
		Age: 123,
		Job: "teacher",
	}
	expectedUser := User{
		Id: baseUser.Id,
		Name: fmt.Sprintf("Awesome %s", baseUser.Name),
		Age: baseUser.Age * 2,
		Job: baseUser.Job,
	}

	findById = func(id int64) (User, bool) {
		return baseUser, true
	}

	user, found := FindById(123)

	assert.True(t, found)

	assertutils.AssertEqualIgnoringFields(t, expectedUser, user, "Name", "Age")
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, expectedUser.Age, user.Age)
}