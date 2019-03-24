package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/final-application/src/models"
)

/*
var user = {
    "name": "Mina",
    "loggedin": true,
    "orders": [
        {
            "id": 1,
            "img": "img/img-small/strings.png",
            "imgalt": "string",
            "desc": "A very authentic and beautiful instrument!!",
            "price": 100.0,
            "productname": "Strings",
            "days": 32
        }, {
            "id": 2,
            "img": "img/img-small/redguitar.jpeg",
            "imgalt": "redg",
            "desc": "A really cool red guitar that can produce super cool music!!",
            "price": 299.0,
            "productname": "Red Guitar",
            "days": 99
        }, {
            "id": 3,
            "img": "img/img-small/drums.jpg",
            "imgalt": "drums",
            "desc": "A set of super awesome drums, combined with a guitar, they can product more than amazing music!!",
            "price": 17000.0,
            "productname": "Drums",
            "days": 45
        }
    ]
};

*/

type MockHandler struct {
}

func NewMockHandler() *MockHandler {
	return &MockHandler{}
}

func (mh *MockHandler) GetMainPage(c *gin.Context) {

}
func (mh *MockHandler) GetProducts(c *gin.Context) {

}
func (mh *MockHandler) GetPromos(c *gin.Context) {

}
func (mh *MockHandler) AddUser(c *gin.Context) {
	f, err := os.Open("mockdata.json")
	if err != nil {
		log.Println("Could not add user", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "mock data file not found"})
		return
	}
	var user models.Customer
	err = json.NewDecoder(f).Decode(&user)
	if err != nil {
		log.Println("Could not add user", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "could not decode mock data json file"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (mh *MockHandler) SignIn(c *gin.Context) {
	f, err := os.Open("mockdata.json")
	if err != nil {
		log.Println("Could not add user", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "mock data file not found"})
		return
	}
	var user models.Customer
	err = json.NewDecoder(f).Decode(&user)
	if err != nil {
		log.Println("Could not add user", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "could not decode mock data json file"})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (mh *MockHandler) SignOut(c *gin.Context) {

}
func (mh *MockHandler) GetOrders(c *gin.Context) {

}
func (mh *MockHandler) Charge(c *gin.Context) {

}
