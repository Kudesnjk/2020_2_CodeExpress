package test

import (
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2020_2_CodeExpress/models"

	"github.com/go-park-mail-ru/2020_2_CodeExpress/repositories"
)

func TestCheckUserCreatesSuccess(t *testing.T) {
	newUser := &models.User{
		Name:     "Daniil",
		Email:    "daniil@mail.ru",
		Password: "123456dD",
	}

	expectedUser := &models.User{
		ID:       0,
		Name:     "Daniil",
		Email:    "daniil@mail.ru",
		Password: "123456dD",
	}

	sImpl := repositories.NewUserRepImpl()

	err := sImpl.CreateUser(newUser)
	if err != nil {
		t.Fatalf("CheckUserCreates failed, error:  %s", err)
	}

	if !reflect.DeepEqual(expectedUser, newUser) {
		t.Fatalf("CheckUserCreates failed, expected: %v, result: %v", expectedUser, newUser)
	}
}

func TestCheckUserExistsSuccess(t *testing.T) {
	sImpl := repositories.NewUserRepImpl()

	createUser := func(s repositories.UserRep) error {
		newUser := &models.User{
			ID: 0,
			Name:     "Daniil",
			Email:    "daniil@mail.ru",
			Password: "123456dD",
		}
		return s.CheckUserExists(newUser)
	}

	err := createUser(sImpl)
	if err != nil {
		t.Fatalf("TestCheckUserExists failed, error: %s", err)
	}
}

func TestCheckUserExistsEmailFailed(t *testing.T) {
	sImpl := repositories.NewUserRepImpl()

	createUser := func(s repositories.UserRep) error {
		newUser1 := &models.User{
			ID: 0,
			Name:     "Daniil21",
			Email:    "daniil@mail.ru",
			Password: "123456dD",
		}

		newUser2 := &models.User{
			ID: 1,
			Name:     "Daniil",
			Email:    "daniil@mail.ru",
			Password: "123456dD",
		}
		s.CreateUser(newUser1)
		return s.CheckUserExists(newUser2)
	}

	err := createUser(sImpl)
	if err == nil {
		t.Fatalf("TestCheckUserExists not failed")
	}
}

func TestCheckUserExistsUsernameFailed(t *testing.T) {
	sImpl := repositories.NewUserRepImpl()

	createUser := func(s repositories.UserRep) error {
		newUser1 := &models.User{
			ID: 0,
			Name:     "Daniil",
			Email:    "daniil2@mail.ru",
			Password: "123456dD",
		}

		newUser2 := &models.User{
			ID: 1,
			Name:     "Daniil",
			Email:    "daniil@mail.ru",
			Password: "123456dD",
		}
		s.CreateUser(newUser1)
		return s.CheckUserExists(newUser2)
	}

	err := createUser(sImpl)
	if err == nil {
		t.Fatalf("TestCheckUserExists not failed")
	}
}