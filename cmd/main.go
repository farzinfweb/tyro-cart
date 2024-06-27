package main

import (
	"cart/api/handler"
	"cart/impl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

func main() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find .env file: ", err)
	}

	dbUri := viper.GetString("DB_URI")
	dbName := viper.GetString("DB_NAME")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	db := client.Database(dbName)
	rpo := impl.NewCartRepo(db)
	cartService := impl.NewCartService(rpo)

	port := viper.GetString("PORT")

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/cart", func(c echo.Context) error {
		return handler.AddCartItem(c, cartService)
	})
	e.Logger.Fatal(e.Start(":" + port))

}
