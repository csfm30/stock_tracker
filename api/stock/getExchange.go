package account

import (
	"fmt"
	"log"
	"stock_tracker/methods/notify"
	"stock_tracker/utility"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

// GetExchange
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
func GetExchange(c *fiber.Ctx) error {

	apiKey := "1cb64604af978c9616c12945"
	baseCurrency := "USD"

	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"base_code": baseCurrency,
			"apikey":    apiKey,
		}).
		SetResult(map[string]interface{}{}).
		Get("https://v6.exchangerate-api.com/v6/" + apiKey + "/latest/" + baseCurrency)

	if err != nil {
		log.Fatalf("Error getting exchange rates: %v", err)
	}

	data := resp.Result().(*map[string]interface{})

	conversionRates, ok := (*data)["conversion_rates"].(map[string]interface{})
	if !ok {
		log.Fatal("Missing or invalid 'conversion_rates' field in API response")
	}

	usdToThb, ok := conversionRates["THB"].(float64)
	if !ok {
		log.Fatal("THB conversion rate not found")
	}

	// fmt.Printf("USD to THB exchange rate: %.2f\n", usdToThb)

	if usdToThb > 35 {
		message := fmt.Sprintf("USD to THB exchange rate: %.2f\n", usdToThb)
		go notify.DiscordNotify(message)
	}

	return utility.ResponseSuccess(c, nil)
}
