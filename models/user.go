package models

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// the purpose here is to represent user
type UserModel struct {
	//bson using with mongodb
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	FullName string `json:"fullName" bson:"full_name"`
	Tel      string `json:"tel,omitempty" bson:"tel,omitempty"`
}

// we need to set id before insert
func (u *UserModel) BeforeInsert() {
	u.ID = primitive.NewObjectID().Hex()
}

// we need to hash password before insert
func (u *UserModel) HashPassword() {
	macData := hmac.New(sha256.New, []byte(u.Password))
	hash := macData.Sum(nil)
	u.Password = base64.StdEncoding.EncodeToString(hash)
}

// the purpose to use for create , read , update ,delete
type UserService struct {
	Client  *mongo.Client
	DBName  string
	UserCol *mongo.Collection
}

func NewUserService(client *mongo.Client, dbName string) *UserService {
	userCol := client.Database(dbName).Collection("users")
	return &UserService{
		UserCol: userCol,
	}
}

func (s *UserService) CreateUser(user *UserModel) (*mongo.InsertOneResult, error) {
	res, err := s.UserCol.InsertOne(context.Background(), user)
	return res, err
}
