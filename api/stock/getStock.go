package account

import (
	"fmt"
	"stock_tracker/logs"
	"stock_tracker/utility"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type reqGetStock struct {
	Name string `json:"name"`
}

type responseGetStock struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Time  string  `json:"time"`
}

// GetAuthAgora
// @Summary ดึงข้อมูล appid cert
// @Description ดึงข้อมูล appid cert
// @Security ApiKeyAuth
// @Tags App Store
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseSuccess{data=resGetAuthAgora}
// @Failure 401 {object} models.ResponseError401
// @Failure 500 {object} models.ResponseError
// @Router /v1/get-auth-agora [get]
func GetStock(c *fiber.Ctx) error {

	apiToken := viper.GetString("token.stock")

	// db := database.DBConn
	reqGetStock := new(reqGetStock)

	//CheckInput
	if err := c.BodyParser(reqGetStock); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	stockName := strings.TrimSpace(reqGetStock.Name)

	if reqGetStock.Name == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}

	fmt.Println("Key", apiToken)

	// client := resty.New()
	// resp, err := client.R().
	// 	SetQueryParams(map[string]string{
	// 		"symbol": stockName,
	// 		"token":  apiToken,
	// 	}).
	// 	SetResult(map[string]interface{}{}).
	// 	Get("https://finnhub.io/api/v1/quote")

	// if err != nil {
	// 	log.Fatalf("API request failed: %v", err)
	// }

	// data := resp.Result().(*map[string]interface{})

	// price := (*data)["c"]

	// priceFloat := price.(float64)

	// if priceFloat == 0 {
	// 	return utility.ResponseError(c, fiber.StatusBadRequest, "stock_not_found")
	// }

	// timestamp := int64((*data)["t"].(float64))

	// if priceFloat < 400 {
	// 	message := fmt.Sprintf("Stock %s is %.2f , Time: %s", stockName, priceFloat, time.Unix(timestamp, 0).Format(time.RFC1123))
	// 	go notify.DiscordNotify(message)
	// }
	// responseGetStock := responseGetStock{
	// 	Name:  stockName,
	// 	Price: priceFloat,
	// 	Time:  time.Unix(timestamp, 0).Format(time.RFC1123),
	// }

	return utility.ResponseSuccess(c, stockName)
}
