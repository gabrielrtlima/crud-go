package model

import "golang.org/x/crypto/bcrypt"

func (ud *userDomain) EncryptPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	ud.password = string(hashedPassword)
}
