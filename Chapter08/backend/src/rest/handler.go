package rest

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/minaandrawos/Hands-On-Full-Stack-Development-with-Go/7-Testing-and-benchmarking/backend/src/dblayer"
	"github.com/minaandrawos/Hands-On-Full-Stack-Development-with-Go/7-Testing-and-benchmarking/backend/src/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
)

type HandlerInterface interface {
	GetMainPage(c *gin.Context)
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetOrders(c *gin.Context)
	Charge(c *gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler() (HandlerInterface, error) {
	return NewHandlerWithParams("mysql", "root:root@/gomusic")
}

func NewHandlerWithParams(dbtype, conn string) (HandlerInterface, error) {
	db, err := dblayer.NewORM(dbtype, conn)
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}

func NewHandlerWithDB(db dblayer.DBLayer) HandlerInterface {
	return &Handler{db: db}
}

func (h *Handler) GetMainPage(c *gin.Context) {
	log.Println("Main page....")
	c.String(http.StatusOK, "Main page for secure API!!")
	//fmt.Fprintf(c.Writer, "Main page for secure API!!")
}

func (h *Handler) GetProducts(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	products, err := h.db.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(products))
	c.JSON(http.StatusOK, products)
}

func (h *Handler) GetPromos(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	promos, err := h.db.GetPromos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, promos)
}

func (h *Handler) AddUser(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err = h.db.AddUser(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *Handler) SignIn(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err = h.db.SignInUser(customer.Email, customer.Pass)
	if err != nil {
		if err == dblayer.ErrINVALIDPASSWORD {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *Handler) SignOut(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.db.SignOutUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *Handler) GetOrders(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orders, err := h.db.GetCustomerOrdersByID(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *Handler) Charge(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	request := struct {
		models.Order
		Remember    bool   `json:"rememberCard"`
		UseExisting bool   `json:"useExisting"`
		Token       string `json:"token"`
	}{}

	err := c.ShouldBindJSON(&request)
	log.Printf("request: %+v \n", request)
	if err != nil {
		c.JSON(http.StatusBadRequest, request)
		return
	}
	// Set your secret key: remember to change this to your live secret key in production
	// Keys can be obtained from: https://dashboard.stripe.com/account/apikeys
	// They key below is just for testing
	stripe.Key = "sk_test_4eC39HqLyjWDarjtT1zdp7dc"
	//test cards available at:	https://stripe.com/docs/testing#cards
	//setting charge parameters
	chargeP := &stripe.ChargeParams{
		Amount:      stripe.Int64(int64(request.Price)),
		Currency:    stripe.String("usd"),
		Description: stripe.String("GoMusic charge..."),
	}
	stripeCustomerID := ""
	//Either remembercard or use exeisting should be enabled but not both
	if request.UseExisting {
		//use existing
		log.Println("Getting credit card id...")
		stripeCustomerID, err = h.db.GetCreditCardCID(request.CustomerID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		cp := &stripe.CustomerParams{}
		cp.SetSource(request.Token)
		customer, err := customer.New(cp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		stripeCustomerID = customer.ID
		if request.Remember {
			//save card!!
			err = h.db.SaveCreditCardForCustomer(request.CustomerID, stripeCustomerID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}
	//we should check if the customer already ordered the same item or not but for simplicity, let's assume it's a new order
	chargeP.Customer = stripe.String(stripeCustomerID)
	_, err = charge.New(chargeP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.db.AddOrder(request.Order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
