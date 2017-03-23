package models

import (
	"fmt"
)

const mgoSessionCollName = "session"

func Login(user User) (string, error) {
	//c := util.GetMgoCollection(mgoSessionCollName)

	// check credentials
	fmt.Println(user.Name)
	fmt.Println(user.Password)

	// create token
	//auth.CreateToken()

	// return token
	return "", nil
}
