package commons

import (
	"ProyectoWeb/models"
	"crypto/rsa"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"time"
)

var (
	privateKey *rsa.PrivateKey
	// Public key se usa para validar el token
	PublicKey *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./keys/private.rsa")
	if err != nil {
		log.Fatal(err)
	}

	publicBytes, err := ioutil.ReadFile("./keys/public.rsa")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal(err)
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal(err)
	}
}

// GenerateJWT Genera el token para el Cliente
func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "Vancode",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)

	if err != nil {
		log.Fatal("No se Pudo firmar el token")
	}

	return result
}
