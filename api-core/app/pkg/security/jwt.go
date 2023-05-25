package security

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dnguyenngoc/doc-extractor/pkg/dto"
	"github.com/dnguyenngoc/doc-extractor/pkg/setting"
)

func CreateBearerToken(infoClient dto.InfoAccount) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["client_id"] = infoClient.ClientID
	claims["grant_type"] = infoClient.GrantType
	claims["exp"] = time.Now().Add(time.Second * time.Duration(setting.Config.GetDuration("ApiTokenExpire"))).Unix()

	tokenString, err := token.SignedString([]byte(setting.Config.GetString("ApiTokenSecret")))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
