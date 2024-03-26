package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPaassword(password string)(string,error){
	hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	
	if err!=nil{
		return "", err
	}

	return string(hashedPassword), nil
}

func VerifyPassword(plainTextPassword,hashedPassword string)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(plainTextPassword))
	return err==nil
}