package models

import (
	"time"

	"microservices/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Name     string             `bson:"name"`
	Created  time.Time          `bson:"created"`
	Updated  time.Time          `bson:"updated"`
}

func (u *User) ToProtoBuffer() *pb.User {
	return &pb.User{
		Id:       u.Id.Hex(),
		Email:    u.Email,
		Password: u.Password,
		Name:     u.Name,
		Created:  u.Created.Unix(),
		Updated:  u.Updated.Unix(),
	}
}

func (u *User) FromProtoBuffer(user *pb.User) {
	u.Id, _ = primitive.ObjectIDFromHex(user.Id)
	u.Email = user.Email
	u.Password = user.Password
	u.Name = user.Name
	u.Created = time.Unix(user.Created, 0)
	u.Updated = time.Unix(user.Updated, 0)
}
