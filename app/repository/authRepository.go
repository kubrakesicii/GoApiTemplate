package repository

import (
	"goapitemplate/app/model/entity"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	//Login(user dto.LoginDto) dto.GetLoginDto
	//baseRepo generic.BaseRepository
	IsEmailExists(email string) (tx *gorm.DB)
	//GetById(id int) entity.User
}

type AuthRepository struct {
	DB *gorm.DB
}

func CryptoPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash password")
	}
	return string(hash)
}

func (r *AuthRepository) IsEmailExists(email string) (tx *gorm.DB) {
	var user entity.User
	return r.DB.Where("email = ?", email).Take(&user)
}

// func (r *AuthRepository) GetById(id int) entity.User {
// 	var user entity.User

// 	//err := r.baseRepo.FindById(user, uint(id))
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return user
// }

// func (r *authRepository) Login(user dto.LoginDto) (*dto.GetLoginDto, error) {
// 	getUser := dto.GetLoginDto{}
// 	res := r.DB.Where("email = ? AND password = ?", user.Username, user.Password).Take(&getUser)
// 	//err := smapping.FillStruct(&getUser, smapping.MapFields(&res))

// 	// return nil,err
// 	// return &getUser,nil
// }
