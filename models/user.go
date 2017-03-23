package models

import (
	"jwt-test/util"

	"crypto/sha256"

	"encoding/hex"

	"github.com/google/uuid"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

const mgoUserCollName = "user"

func CreateUser(user *User) error {
	c := util.GetMgoCollection("user")

	// set uuid
	user.Id = uuid.New().String()

	// hash password
	hash := sha256.New()
	hash.Write([]byte(user.Password))

	user.Password = hex.EncodeToString(hash.Sum(nil))

	return c.Insert(user)
}

func GetUsers() ([]User, error) {
	c := util.GetMgoCollection(mgoUserCollName)

	result := []User{}
	err := c.Find(nil).All(&result)

	return result, err
}

func GetUserById(id string) (User, error) {
	c := util.GetMgoCollection(mgoUserCollName)

	result := User{}
	err := c.Find(bson.M{"id": id}).One(&result)

	return result, err
}

func UpdateUser(user *User) error {
	c := util.GetMgoCollection(mgoUserCollName)

	return c.Update(bson.M{"id": user.Id}, user)
}

func DeleteUser(id string) error {
	c := util.GetMgoCollection(mgoUserCollName)

	return c.Remove(bson.M{"id": id})
}
