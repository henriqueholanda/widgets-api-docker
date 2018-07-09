package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/henriqueholanda/widgets-api-docker/services/database"
	"github.com/henriqueholanda/widgets-api-docker/entities"
	"gopkg.in/mgo.v2/bson"
	"regexp"
)

type UsersHandler struct {
	http.Handler
}

var mongo = database.GetDatabaseConnection()
var collection = mongo.DB("widgets_api").C("users")

func (UsersHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	hasUserId, _ := regexp.MatchString("/users/.+", request.URL.Path)
	if hasUserId {
		id := bson.ObjectId(request.URL.Query().Get("id"))
		data, err := fetchOne(id)
		if err != nil {
			fmt.Println("Failed to get single user: " + err.Error())
		}
		successResponse(response, data)
	}
	data, err := fetchAll()
	if err != nil {
		fmt.Println("Failed to get all users: " + err.Error())
	}
	successResponse(response, data)
}

func successResponse(response http.ResponseWriter, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	jsonResponse, _ := json.Marshal(data)
	response.Write([]byte(jsonResponse))
}

func fetchAll() (interface{}, error) {
	users := entities.Users{}
	err := collection.Find(nil).All(&users)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}

func fetchOne(id bson.ObjectId) (interface{}, error) {
	user := entities.User{}
	err := collection.Find("{_id: " + id + "}").One(&user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}
