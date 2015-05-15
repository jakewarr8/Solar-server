package main

import "crypto/sha256"
import "code.google.com/p/go.crypto/bcrypt"
import "code.google.com/p/go.crypto/pbkdf2"

type User struct {
	UserName	string  
	Password	string
	Hash		[]byte
	Salt		[]byte
	ID		int
}

func Clear(b []byte) {
	for i:=0; i<len(b); i++ {
		b[i] = 0;
	}
}

func Crypt(password []byte) ([]byte, error) {
	defer Clear(password)
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func HashPassword(password, salt []byte) []byte {
	defer Clear(password)
	return pbkdf2.Key(password, salt, 4096, sha256.Size, sha256.New)
}

