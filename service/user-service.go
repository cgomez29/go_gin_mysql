package service

import (
	"log"

	"github.com/cgomez29/api-gin/dto"
	"github.com/cgomez29/api-gin/models"
	"github.com/cgomez29/api-gin/repository"
	"github.com/mashingan/smapping"
)

//UserService is a contract.....
type UserService interface {
	Update(user dto.UserUpdateDTO) models.User
	Profile(userID string) models.User
}

type userService struct {
	userRepository repository.UserRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) models.User {
	userToUpdate := models.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) models.User {
	return service.userRepository.ProfileUser(userID)
}
