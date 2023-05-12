package crypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"pwdkeeper/internal/app/initconfig"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
)

func HashAndSalt(pwd []byte) string {    
    // Use GenerateFromPassword to hash & salt pwd.
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost. 
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Error().Err(err)
    }    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return string(hash)
}

func GenAuthToken(password string) (authToken string) {
	//sign user with HMAC, using SHA256
	h := hmac.New(sha256.New, initconfig.Key)
	h.Write([]byte(password))
	dst := h.Sum(nil)
	authToken = string(dst) + string(password)
	return hex.EncodeToString([]byte(authToken))
}
/*
At registration time, a uniform random symmetric key (key1) is generated from a CSPRNG using client-side scripting running in the user's web browser.
The user chooses a good password ( preferably dicewire or Bip39), 
and another symmetric key (key2) is derived from this password using a password based key derivation function (such as PBKDF2, Argon2, etc). 
Then, key1 is encrypted using key2 in-browser, then encrypted key1 is stored on the server.

When the user logs in, encrypted key1 is downloaded to the user's web browser. 
The user provides the password, key2 is derived from the password, and encrypted key1 is decrypted in-browser. 
key1 is used to encrypt all the user's secrets in-browser, so that only the encrypted secrets are uploaded to the server.
*/
func generateRandom(size int) ([]byte, error) {
    b := make([]byte, size)
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }

    return b, nil
}

func Key2build(password string) []byte{
	dk := pbkdf2.Key([]byte(password), initconfig.Salt, 4096, 32, sha1.New)
	return dk
}

func Key1build() []byte{
    // будем использовать AES256, создав ключ длиной 32 байта
    key, err := generateRandom(2 * aes.BlockSize) // ключ шифрования
    if err != nil {
        fmt.Printf("error: %v\n", err)
        return nil
    }
	return key
}

func EncryptKey1(key1, key2 []byte) (key1enc []byte){
	aesblock, err := aes.NewCipher(key2)
    if err != nil {
        fmt.Printf("error: %v\n", err)
        return 
    }

    aesgcm, err := cipher.NewGCM(aesblock)
    if err != nil {
        fmt.Printf("error: %v\n", err)
        return 
    }

    // создаём вектор инициализации
    nonce, err := generateRandom(aesgcm.NonceSize())
	fmt.Printf("aesgcm.NonceSize() %v\n", aesgcm.NonceSize())
	fmt.Printf("nonce %v\n", nonce)
    if err != nil {
        fmt.Printf("error: %v\n", err)
        return 
    }

    dst := aesgcm.Seal(nil, nonce, key1, nil) // зашифровываем
    fmt.Printf("encrypted: %x\n", dst)
	fmt.Printf("Seal %v\n", dst)
	return dst
}

func DecryptKey1(key1enc, key2 []byte) (key1decrypted []byte){
	aesblock, err := aes.NewCipher(key2)
    if err != nil {
        fmt.Printf("error NewCipher: %v\n", err)
        return nil
    }
	aesgcm, err := cipher.NewGCM(aesblock)
    if err != nil {
        fmt.Printf("error NewGCM: %v\n", err)
        return nil
    }
	    // создаём вектор инициализации
		fmt.Printf("aesgcm.NonceSize() %v\n", aesgcm.NonceSize())
	nonce := key2[len(key2)-aesgcm.NonceSize():]
	fmt.Printf("nonce %v\n", nonce)
    key1, err := aesgcm.Open(nil, nonce, key1enc, nil) // расшифровываем
    if err != nil {
        fmt.Printf("error Open: %v\n", err)
        return nil
    }
    fmt.Printf("decrypted: %s\n", key1)
	return key1
}



