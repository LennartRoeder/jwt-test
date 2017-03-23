package routes

import (
	"encoding/json"
	"fmt"
	"jwt-test/models"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TODO: Improve error messages

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := new(models.User)
	json.NewDecoder(r.Body).Decode(&user)

	err := models.CreateUser(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "A user \"%s\" already exists", user.Name)
		return
	}

	log.Printf("User created: %s (%s)\n", user.Name, user.Id)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, user.Id)
}

func GetUsers(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	result, err := models.GetUsers()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	log.Println("users:", result)

	res, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", res)
}

func GetUser(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	result, err := models.GetUserById(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	log.Println("user:", result)

	res, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := new(models.User)
	json.NewDecoder(r.Body).Decode(&user)

	user.Id = ps.ByName("id")

	err := models.UpdateUser(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "%s updated!", user.Name)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	err := models.DeleteUser(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "%s updated!", id)
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := new(models.User)
	json.NewDecoder(r.Body).Decode(&user)

	token, err := models.Login(*user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	log.Printf("User %s logged in", user.Name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", token)
}

func Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
