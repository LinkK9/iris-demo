package repo

import (
	"math/rand"
	"testing"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestAddNewUser(t *testing.T) {
	assert := assert.New(t)
	pass := fake.Password(true, true, true, true, false, 10)
	user := MyUser{
		ID: ksuid.New().String(),
		User: fake.Name(),
		Pass: pass,
		Email: fake.Email(),
		BirthDay: fake.Date().String(),
		Job: fake.JobTitle(),
	}
	want := len(users) + 1
	AddNewUser(user)
	assert.Equal(len(users), want)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)
	want := len(users) - 1
	index := rand.Intn(want + 1)
	DeleteUser(users[index].ID)
	assert.Equal(len(users), want)
}

func TestModifyUser(t *testing.T) {
	assert := assert.New(t)

	// lấy random 1 ID sẵn có
	index := rand.Intn(len(users))
	id := users[index].ID
	user := MyUser{
		ID: id,
		User: fake.Name(),
		Email: fake.Email(),
		BirthDay: fake.Date().String(),
		Job: fake.JobTitle(),
	}
	ModifyUser(user)
	assert.Equal(users[index].User, user.User)
	assert.Equal(users[index].Email, user.Email)
	assert.Equal(users[index].BirthDay, user.BirthDay)
	assert.Equal(users[index].Job, user.Job)
}