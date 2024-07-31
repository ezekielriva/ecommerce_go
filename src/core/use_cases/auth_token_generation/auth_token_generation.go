package authtokengeneration

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/ezekielriva/ecommerce_go/src/core/entities"
)

func GenerateAuthToken(cred *entities.UserCredentials) error {
	randomToken := make([]byte, 32)
	_, err := rand.Read(randomToken)

	if err != nil {
		return err
	}

	cred.AuthToken = base64.URLEncoding.EncodeToString(randomToken)
	cred.AuthTokenExp = time.Now().Add(time.Minute * 60)

	return nil
}
