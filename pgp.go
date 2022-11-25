package pgp

import (
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/pgp", new(PGP))
}

type PGP struct{}

func EncryptMessageArmored(key, plaintext string) (string, error) {
	return helper.EncryptMessageArmored(key, plaintext)
}

func DecryptMessageArmored(privateKey string, passphrase []byte, ciphertext string) (string, error) {
	return helper.DecryptMessageArmored(privateKey, passphrase, ciphertext)
}
