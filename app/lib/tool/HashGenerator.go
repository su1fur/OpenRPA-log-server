package tool

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func GenerateHash(str string, keyLength int) string {
	secret := time.Now().String()

	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(str))

	hashedStr := hex.EncodeToString(hash.Sum(nil))

	rand.Seed(time.Now().UnixNano())
	cutOutPoint := rand.Intn(len(hashedStr) - keyLength)

	fmt.Print(hashedStr, " ~> ")

	return string([]rune(hashedStr)[cutOutPoint : cutOutPoint+keyLength])
}

func GenerateSha256(str string) string {
	converted := sha256.Sum256([]byte(str))
	return hex.EncodeToString(converted[:])
}
