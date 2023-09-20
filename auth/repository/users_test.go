package repository

import (
	"fmt"
	"microservices/auth/models"
	"microservices/db"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
}

func TestUserRepositorySave(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := primitive.NewObjectID()
	user := &models.User{
		Id:       id,
		Name:     "test",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "test",
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	foundUser, err := r.GetById(id)
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)

}

func TestUsersRepositoryGetById(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := primitive.NewObjectID()
	user := &models.User{
		Id:       id,
		Name:     "test",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "test",
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	foundUser, err := r.GetById(id)
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user.Id, foundUser.Id)
	assert.Equal(t, user.Name, foundUser.Name)
	assert.Equal(t, user.Email, foundUser.Email)
	assert.Equal(t, user.Password, foundUser.Password)
}
