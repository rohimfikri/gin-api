package helper

import (
	"fmt"
	"gin-api/core"
	"gin-api/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func GenerateToken(user *model.User) (string, error) {
	claims := JWTUserClaims{
		user.Username,
		user.Email,
		user.FirstName,
		*user.LastName,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(core.ENV.JWT_EXPIRED_HOUR) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    core.ENV.JWT_ISSUER,
			Subject:   user.Username,
			// ID:        user.ID,
			// Audience:  []string{"somebody_else"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(core.ENV.JWT_SECRET_KEY)
}

func ParseToken(tokenString *string) (*JWTUserClaims, error) {
	token, err := jwt.ParseWithClaims(*tokenString, &JWTUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(core.ENV.JWT_SECRET_KEY), nil
	})
	logger := &log.Logger
	if err != nil {
		logger.Fatal().Err(err)
		return nil, err
	} else if claims, ok := token.Claims.(*JWTUserClaims); ok {
		logger.Info().Str("username", claims.Username).Str("issuer", claims.RegisteredClaims.Issuer)
		return claims, nil
	} else {
		logger.Fatal().Msg("unknown claims type, cannot proceed")
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}

}
