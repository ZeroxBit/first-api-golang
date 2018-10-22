package commons

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/krlos/firstApi/api/models"
)

var (
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./keys/private.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo private")
	}

	publicBytes, err := ioutil.ReadFile("./keys/public.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo publico")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a privateKey")
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("no se pudo hacer párse a PublicKey")
	}
}

func GenerateJwt(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // para expirar la sesion
			Issuer:    "Krlos",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("No se pude firmar el token")
	}

	return result
}
