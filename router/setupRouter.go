package router

import (
	stock "stock_tracker/api/stock"

	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(app *fiber.App) {
	apiBackendPrefix := app.Group("/testgo")
	apiRoutes := apiBackendPrefix.Group("/api")
	v1 := apiRoutes.Group("/v1")

	v1.Get("/get-stock", stock.GetStock)
	v1.Get("/get-exchange", stock.GetExchange)

}
