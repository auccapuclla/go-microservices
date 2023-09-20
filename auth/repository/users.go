package repository

import (
	"context"
	"microservices/auth/models"
	"microservices/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const UsersCollection = "users"

type UsersRepository interface {
	Save(user *models.User) error
	GetById(id primitive.ObjectID) (user *models.User, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetAll() (users []*models.User, err error)
	Update(user *models.User) error
	Delete(id string) error
	DeleteAll() error
}

type usersRepository struct {
	c *mongo.Collection
}

func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{c: conn.DB().Collection(UsersCollection)}
}

func (r *usersRepository) Save(user *models.User) error {
	_, err := r.c.InsertOne(context.TODO(), user)
	return err
}

func (r *usersRepository) GetById(id primitive.ObjectID) (user *models.User, err error) {
	opts := options.FindOne()
	err = r.c.FindOne(
		context.TODO(),
		bson.D{{"_id", id}},
		opts,
	).Decode(&user)
	// println(user.Password)
	return user, err
}

func (r *usersRepository) GetByEmail(email string) (user *models.User, err error) {
	err = r.c.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return user, err
}

func (r *usersRepository) GetAll() (users []*models.User, err error) {
	cur, err := r.c.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var user *models.User
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *usersRepository) Update(user *models.User) error {
	_, err := r.c.UpdateOne(context.TODO(), bson.M{"_id": user.Id}, bson.M{"$set": user})
	return err
}

func (r *usersRepository) Delete(id string) error {
	_, err := r.c.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (r *usersRepository) DeleteAll() error {
	return r.c.Drop(context.Background())
}
