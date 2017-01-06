package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/howeyc/gopass"
	"github.com/kless/osutil/user/crypt/sha512_crypt"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
)

func secureRandomAlphaString(length int) string {
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = secureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(letterBytes) {
			result[i] = letterBytes[idx]
			i++
		}
	}

	return string(result)
}

func secureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}

func passwordPrompt() (string, bool) {
	var p1, p2 []byte
	var err error

	fmt.Print("Password: ")
	if p1, err = gopass.GetPasswdMasked(); err != nil {
		return "", false
	}

	fmt.Print("Confirm:  ")
	if p2, err = gopass.GetPasswdMasked(); err != nil {
		return "", false
	}

	if bytes.Equal(p1, p2) {
		return string(p1), true
	}

	return "", false
}

func main() {
	var passwordString string
	var saltString string

	fs := flag.NewFlagSet("mkpasswd", flag.ContinueOnError)
	fs.StringVar(&passwordString, "password", "",
		"Optional password argument")
	fs.StringVar(&saltString, "salt", "",
		"Optional salt argument without prefix")

	if err := fs.Parse(os.Args[1:]); err != nil {
		os.Exit(1)
	}

	c := sha512_crypt.New()
	saltPrefix := sha512_crypt.MagicPrefix
	saltMaxLen := sha512_crypt.SaltLenMax

	if saltString != "" {
		if len(saltString) > saltMaxLen {
			fmt.Printf("Warning specified salt greater than max length (%d).  "+
				"Salt will be truncated.\n", saltMaxLen)
		}
		saltString = fmt.Sprintf("%s%s", saltPrefix, saltString)
	} else {
		saltString = fmt.Sprintf("%s%s", saltPrefix, secureRandomAlphaString(saltMaxLen))
	}

	if passwordString == "" {
		var ok bool
		for {
			if passwordString, ok = passwordPrompt(); ok {
				break
			} else {
				fmt.Print("Password mismatch or error.  Please try again.\n")
			}
		}
	}

	shadowHash, err := c.Generate([]byte(passwordString), []byte(saltString))
	if err != nil {
		fmt.Printf("Failed to generate shadow hash: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s\n", shadowHash)
}
