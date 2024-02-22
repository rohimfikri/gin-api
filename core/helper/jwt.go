package core_helper

import (
	"fmt"
	core_config "gin-api/core/config"
	core_model "gin-api/core/model"
	core_type "gin-api/core/type"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func GenerateToken(user *core_model.User) (string, error) {
	claims := core_type.JWTUserClaims{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  *user.LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(core_config.ENV.JWT_EXPIRED_HOUR) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    core_config.ENV.JWT_ISSUER,
			Subject:   user.Username,
			// ID:        user.ID,
			// Audience:  []string{"somebody_else"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(core_config.ENV.JWT_SECRET_KEY)
}

func ParseToken(tokenString *string) (*core_type.JWTUserClaims, error) {
	token, err := jwt.ParseWithClaims(*tokenString, &core_type.JWTUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(core_config.ENV.JWT_SECRET_KEY), nil
	})
	logger := &log.Logger
	if err != nil {
		logger.Fatal().Err(err)
		return nil, err
	} else if claims, ok := token.Claims.(*core_type.JWTUserClaims); ok {
		logger.Info().Str("username", claims.Username).Str("issuer", claims.RegisteredClaims.Issuer)
		return claims, nil
	} else {
		logger.Fatal().Msg("unknown claims type, cannot proceed")
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}

}
