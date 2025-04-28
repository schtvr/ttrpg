package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/schtvr/ttrpg/backend/internal/store"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username  string    `json:"username"`
	Password  string    `json:"hashed_password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserRegistration(user *User) error {
	users := store.GetMongoClient().Database("ttrpg").Collection("users")
	// Check if the user already exists
	var existingUser User
	if err := users.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser); err == nil {
		return fmt.Errorf("user already exists")
	}
	_, err := users.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}

func UserLogin(user *User) error {
	users := store.GetMongoClient().Database("ttrpg").Collection("users")

	var record User
	if err := users.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&record); err != nil {
		return fmt.Errorf("error unmarshalling ")
	}

	if err := CheckPasswordHash(record.Password, user.Password); err != nil {
		return fmt.Errorf("password does not match record for %s", user.Username)
	}
	return nil
}
