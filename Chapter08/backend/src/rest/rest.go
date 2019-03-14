package rest

import (
	"fmt" //"github.com/gin-gonic/autotls"

	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}

func RunAPIWithHandler(address string, h HandlerInterface) error {
	//Get gin's default engine
	r := gin.Default()
	//r.Use(MyCustomLogger())
	//load homepage
	r.GET("/", h.GetMainPage)
	//get products
	r.GET("/products", h.GetProducts)
	//get promos
	r.GET("/promos", h.GetPromos)
	/*
		//post user sign in
		r.POST("/user/signin", h.SignIn)
		//post user sign out
		r.POST("/user/:id/signout", h.SignOut)
		//get user orders
		r.GET("/user/:id/orders", h.GetOrders)
		//post purchase charge
		r.POST("/user/charge", h.Charge)
	*/

	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}

	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("", h.AddUser)
	}

	r.Static("/img", "../public/img")
	//return autotls.Run(r, address)
	return r.Run(address)
	//return r.RunTLS(address, "cert.pem", "key.pem")
}

func MyCustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("************************************")
		c.Next()
		fmt.Println("************************************")
	}
}
