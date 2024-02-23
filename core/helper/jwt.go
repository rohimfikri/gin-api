package core_helper

import (
	"fmt"
	core_model "gin-api/core/model"
	core_type "gin-api/core/type"
	"slices"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lithammer/shortuuid"
	"github.com/rs/zerolog/log"
)

func GenerateToken(user *core_model.User, secretKey *string, issuer *string, expiredHour *uint8) (string, string, error) {
	uuid := shortuuid.New()
	claims := core_type.JWTUserClaims{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  *user.LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(*expiredHour) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    *issuer,
			Subject:   user.Username,
			ID:        uuid,
			// Audience:  []string{"somebody_else"},
		},
	}

	refreshClaims := core_type.JWTUserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(*expiredHour*2) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    *issuer,
			Subject:   user.Username,
			ID:        uuid,
			// Audience:  []string{"somebody_else"},
		},
	}

	var token string
	var refreshToken string
	var err error
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(*secretKey)
	if err != nil {
		logger := &log.Logger
		logger.Fatal().Err(err).Str("logtype", "GenerateToken")
	} else {
		refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(*secretKey)
	}

	return token, refreshToken, err
}

func ValidateToken(tokenString *string, secretKey *string) (*core_type.JWTUserClaims, error) {
	token, err := jwt.ParseWithClaims(*tokenString, &core_type.JWTUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(*secretKey), nil
	})
	logger := &log.Logger
	if err != nil {
		logger.Fatal().Err(err).Str("logtype", "ValidateToken")
		return nil, err
	} else if claims, ok := token.Claims.(*core_type.JWTUserClaims); ok {
		if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
			logger.Info().Str("logtype", "ValidateToken").Msg("token is expired")
			return nil, fmt.Errorf("token is expired")
		} else {
			return claims, nil
		}
	} else {
		logger.Fatal().Str("logtype", "ValidateToken").Msg("unknown claims type, cannot proceed")
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}

}

func CheckPublicRouter(path *string, publicRoute *[]string) bool {
	ret := false

	if slices.Contains(*publicRoute, *path) {
		ret = true
	}

	return ret
}
