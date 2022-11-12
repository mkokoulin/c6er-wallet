package jwt

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mkokoulin/c6er-wallet.git/internal/config"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}

func CreateToken(userID string, cfg *config.Config) (*TokenDetails, error) {
	td := &TokenDetails{
		AtExpires: time.Now().Add(time.Minute * time.Duration(cfg.AccessTokenLiveTimeMinutes)).Unix(),
		RtExpires: time.Now().Add(time.Minute * time.Duration(cfg.RefreshTokenLiveTimeHours)).Unix(),
	}

	atClaims := jwt.MapClaims{
		"exp":     td.AtExpires,
		"user_id": userID,
	}

	rtClaims := jwt.MapClaims{
		"exp":     td.RtExpires,
		"user_id": userID,
	}

	atWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	rtWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	at, err := atWithClaims.SignedString([]byte(cfg.AccessTokenSecret))
	if err != nil {
		return nil, err
	}

	rt, err := rtWithClaims.SignedString([]byte(cfg.RefreshTokenSecret))
	if err != nil {
		return nil, err
	}

	td.AccessToken = at
	td.RefreshToken = rt

	log.Println("token has been generated")

	return td, nil
}

func RefreshToken(r *http.Request, cfg *config.Config) (*TokenDetails, error) {
	_, refreshToken := ExtractToken(r);

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.RefreshTokenSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims);

	if ok && token.Valid {
		userID := claims["user_id"].(string);

		log.Println("refreshing token has started");

		fmt.Println("userID", userID);

		td, err := CreateToken(userID, cfg);
		if err != nil {
			return nil, err
		}

		return td, nil
	} else {
		return nil, errors.New("refresh token expired")
	}
}

func ValidateToken(r *http.Request, cfg *config.Config) (*jwt.Token, string, error) {
	var userID string
	at, _ := ExtractToken(r)

	accessToken, err := jwt.Parse(at, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.AccessTokenSecret), nil
	})
	if err != nil {
		return nil, userID, err
	}

	userID = accessToken.Claims.(jwt.MapClaims)["user_id"].(string)

	if _, ok := accessToken.Claims.(jwt.MapClaims); !ok && !accessToken.Valid {
		return nil, "", errors.New("expired token")
	}

	return accessToken, userID, nil
}

func ExtractToken(r *http.Request) (string, string) {
	var at string
	var rt string
	for _, c := range r.Cookies() {
		if c.Name == "access_token" {
			at = c.Value;
		}
		if c.Name == "refresh_token" {
			rt = c.Value;
		}
	}

	return at, rt
}