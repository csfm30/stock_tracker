package middleware

import (
	"os"
	"stock_tracker/logs"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

func CreateAuthToken(env string, oneId string, userId string) (accessToken string, refreshToken string, err error) {
	// Create AccessToken
	token := jwt.New(jwt.SigningMethodHS256)
	refId := uuid.NewV4()
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["type"] = "access"
	claims["env"] = env
	claims["one_id"] = oneId
	claims["user_id"] = userId
	claims["ref_id"] = refId

	sessionTimeAccess := 1440 // Access อายุ 1 วัน
	if os.Getenv("ENV") == "dev" {
		if oneId == "3045279722" || oneId == "10743797158438861" || oneId == "804228540928" {
			// sessionTimeAccess = 1440
			sessionTimeAccess = 1
		} else {
			sessionTimeAccess = 1440
		}
	}
	//sessionTimeAccess := 1
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(sessionTimeAccess)).Unix()
	//if sessionTimeAccess > 0 {
	//	claims["exp"] = time.Now().Add(time.Minute * time.Duration(sessionTimeAccess)).Unix()
	//}

	// Generate encoded token and send it as response.
	accessToken, err = token.SignedString([]byte(viper.GetString("auth.access")))
	if err != nil {
		logs.Error(err)
		return "create_token_fail", "", err
	}

	// Create RefreshToken
	tokenRefresh := jwt.New(jwt.SigningMethodHS256)
	claimsRefresh := tokenRefresh.Claims.(jwt.MapClaims)
	claimsRefresh["type"] = "refresh"
	claimsRefresh["env"] = env
	claimsRefresh["one_id"] = oneId
	claimsRefresh["user_id"] = userId
	claimsRefresh["ref_id"] = refId

	sessionTimeRefresh := 129600 // Refresh อายุ 3 เดือน
	claimsRefresh["exp"] = time.Now().Add(time.Minute * time.Duration(sessionTimeRefresh)).Unix()
	//if sessionTimeRefresh > 0 {
	//	claimsRefresh["exp"] = time.Now().Add(time.Minute * time.Duration(sessionTimeRefresh)).Unix()
	//}

	refreshToken, err = tokenRefresh.SignedString([]byte(viper.GetString("auth.refresh")))
	if err != nil {
		logs.Error(err)
		return "create_token_fail", "", err
	}

	return accessToken, refreshToken, err
}

func AuthJwt() fiber.Handler {
	return jwtWare.New(jwtWare.Config{
		SigningKey:   []byte(viper.GetString("auth.access")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {

	if err.Error() == "Missing or malformed JWT" {
		// return c.Status(fiber.StatusBadRequest).SendString("Missing or malformed JWT")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "invalid_or_expired_jwt",
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  fiber.StatusUnauthorized,
		"message": "invalid_or_expired_jwt",
	})
}
