package repositorys

import (
	"goapitemplate/app/model/entity"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	//Login(user dto.LoginDto) dto.GetLoginDto
	IsEmailExists(email string) (tx *gorm.DB)
}

type authRepository struct {
	db *gorm.DB
}

//Create new instance of UserRepository
// Use with New func in service
func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	} //You should implement all methods of Repo int
}

func CryptoPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash password")
	}
	return string(hash)
}

func (r *authRepository) IsEmailExists(email string) (tx *gorm.DB) {
	var user entity.User
	return r.db.Where("email = ?", email).Take(&user)
}

// func (r *authRepository) Login(user dto.LoginDto) dto.GetLoginDto {
// 	getUser := dto.GetLoginDto{}
// 	res := r.db.Where("email = ? AND password = ?", user.Username, user.Password).Take(&getUser)
// 	err := smapping.FillStruct(&getUser, smapping.MapFields(&res))

// 	return res
// }
