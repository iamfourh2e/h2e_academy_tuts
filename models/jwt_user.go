package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// identify current user when fully success login or created we provide the token key
type UserClaimModel struct {
	FullName string `json:"fullName"`
	ID       string `json:"id"`
	Role     string `json:"role"`
	*jwt.StandardClaims
}

func (u *UserClaimModel) GenerateToken(userModel *UserModel) (string, error) {
	u.FullName = userModel.FullName
	u.ID = userModel.ID
	u.Role = userModel.Role
	u.StandardClaims = &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	return token.SignedString([]byte("verysecret"))
}

// client -> token -> server
// check whether the token is valid
func ClaimToken(token string) (*UserClaimModel, error) {
	t, err := jwt.ParseWithClaims(token, &UserClaimModel{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("verysecret"), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := t.Claims.(*UserClaimModel)
	if ok && t.Valid {
		return claims, nil
	}
	return nil, nil
}
