package auth

import (
	"fmt"
	"nom/structs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	Issuer        string
	Audience      string
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

func (j *Auth) GenerateTokenpar(user *structs.JwtUser) (structs.TokenPairs, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	claims["sub"] = fmt.Sprint(user.ID)
	claims["aud"] = j.Audience
	claims["iss"] = j.Issuer
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"

	claims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()

	signAccessToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return structs.TokenPairs{}, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = fmt.Sprint(user.ID)
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()
	// Set expiry for the refresh token
	refreshTokenClaims["exp"] = time.Now().UTC().Add(j.RefreshExpiry).Unix()

	signedRefreshToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return structs.TokenPairs{}, err
	}

	var tokenPairs = structs.TokenPairs{
		Token:        signAccessToken,
		RefreshToken: signedRefreshToken,
	}

	return tokenPairs, nil

}
