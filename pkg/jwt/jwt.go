package jwt

import (
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

func ExtractClaims(tokenString string, signingKey []byte) (jwtgo.MapClaims, error) {
	claims := jwtgo.MapClaims{}
	if tokenString == "" {
		claims["role"] = "unauthorized"
		return claims, nil
	}
	token, err := jwtgo.ParseWithClaims(tokenString, claims, func(token *jwtgo.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !(ok && token.Valid) {
		err = fmt.Errorf("invalid jwt token")
		return nil, err
	}

	return claims, nil
}

func GenerateAccessJWT(id string, isAdmin, isOwner bool, signingKey []byte) (string, error) {
	claims := jwtgo.MapClaims{}
	claims["sub"] = id
	if isAdmin {
		claims["role"] = "admin"
	} else if isOwner {
		claims["role"] = "owner"
	} else {
		claims["role"] = "user"
	}
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshJWT(id string, isAdmin, isOwner bool, signingKey []byte) (string, error) {
	claims := jwtgo.MapClaims{}
	claims["sub"] = id
	if isAdmin {
		claims["role"] = "admin"
	} else if isOwner {
		claims["role"] = "owner"
	} else {
		claims["role"] = "user"
	}
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
