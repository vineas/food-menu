package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
		},
		{
			Name:      "Ayam Penyet",
			OrderCode: "ayam_penyet",
			Price:     41250,
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}

func getDrinkMenu(c echo.Context) error {
	drinkMenu := []MenuItem{
		{
			Name:      "Teh Pucuk",
			OrderCode: "teh_pucuk",
			Price:     37500,
		},
		{
			Name:      "Jus Jeruk",
			OrderCode: "jus_jeruk",
			Price:     41250,
		},
		{
			Name:      "Air Putih",
			OrderCode: "air_putih",
			Price:     41250,
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": drinkMenu,
	})
}

func main() {
	e := echo.New()
	// localhost:14045/menu/food
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drinks", getDrinkMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
