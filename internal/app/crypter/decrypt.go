package crypter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"pwdkeeper/internal/app/initconfig"

	"github.com/rs/zerolog/log"
)

func CheckUserAuthandData(msg string) (validSign bool, val string) {
	var (
		data []byte // декодированное сообщение с подписью
		usefuldata   string // полезное содержимое
		err  error
		sign []byte // HMAC-подпись от полезного содержимого
	)
	validSign = false
	data, err = hex.DecodeString(msg)
	if err != nil {
		panic(err)
	}
	log.Debug().Str("data", string(data)).Msgf("hex.DecodeString(msg)")
	usefuldata = string(data[sha256.Size:])
	val = usefuldata
	h := hmac.New(sha256.New, initconfig.Key)
	h.Write(data[sha256.Size:])
	sign = h.Sum(nil)
	if hmac.Equal(sign, data[:sha256.Size]) {
		log.Info().Msgf("Подпись подлинная. содержимое:%v", usefuldata)
		validSign = true
	} else {
		log.Warn().Msgf("Подпись неверна. Где-то ошибка! содержимое:%v", usefuldata)
	}
	return validSign, val
}