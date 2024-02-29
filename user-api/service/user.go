package service

import (
	"log/slog"

	"github.com/OVillas/user-api/model"
)

type userService struct {
	userRepository model.UserRepository
}

func NewUserService(userRepository model.UserRepository) model.UserService {
	return userService{
		userRepository: userRepository,
	}
}

func (us userService) Create(userPayLoad model.UserPayLoad) error {
	log := slog.With(
		slog.String("service", "user"),
		slog.String("func", "Create"))

	userResponse, err := us.userRepository.GetByEmail(userPayLoad.Email)
	if err != nil {
		log.Error("Error trying to get user from repository")
		return model.ErrGetUser
	}

	if userResponse != nil {
		log.Warn("There is already a registered user with this email: " + userPayLoad.Email)
		return model.ErrUserAlreadyRegistered
	}

	hashedPassword, err := Hash(userPayLoad.Password)
	if err != nil {
		log.Error("Error trying to hashed password")
		return model.ErrHashPassword
	}

	user, err := userPayLoad.ToUser(string(hashedPassword))
	if err != nil {
		log.Error("Error trying to convert userPayload to User")
		return model.ErrConvertUserPayLoadToUser
	}

	if err := us.userRepository.Create(*user); err != nil {
		log.Error("Error: ", err)
		return model.ErrCreateUser
	}

	log.Info("success to create user")
	return nil
}

func (us userService) GetAll() ([]model.UserResponse, error) {
	log := slog.With(
		slog.String("service", "user"),
		slog.String("func", "GetAll"))

	users, err := us.userRepository.GetAll()
	if err != nil {
		log.Error("Error: ", err)
		return nil, model.ErrGetUser
	}

	log.Info("get all service executed successfully")
	if users == nil {
		return nil, nil
	}

	var usersResponse []model.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, *user.ToUserResponse())
	}

	return usersResponse, nil
}

func (us userService) GetById(id string) (*model.UserResponse, error) {
	log := slog.With(
		slog.String("service", "user"),
		slog.String("func", "GetById"))

	user, err := us.userRepository.GetById(id)
	if err != nil {
		log.Error("Error: ", err)
		return nil, model.ErrGetUser
	}

	log.Info("get all service executed successfully")
	if user == nil {
		return nil, nil
	}

	userResponse := user.ToUserResponse()

	return userResponse, err
}

func (us userService) GetByName(name string) ([]model.UserResponse, error) {
	log := slog.With(
		slog.String("service", "user"),
		slog.String("func", "GetAll"))

	users, err := us.userRepository.GetByName(name)
	if err != nil {
		log.Error("Error: ", err)
		return nil, model.ErrGetUser
	}

	log.Info("get all service executed successfully")
	if users == nil {
		return nil, nil
	}

	var usersResponse []model.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, *user.ToUserResponse())
	}

	return usersResponse, err
}

func (us userService) GetByEmail(email string) (*model.UserResponse, error) {
	log := slog.With(
		slog.String("service", "user"),
		slog.String("func", "GetByEmail"))

	user, err := us.userRepository.GetByEmail(email)
	if err != nil {
		log.Error("Error: ", err)
		return nil, model.ErrGetUser
	}

	log.Info("get by email service executed successfully")
	if user == nil {
		return nil, nil
	}

	userResponse := user.ToUserResponse()

	return userResponse, nil
}

func (us userService) Update(id string, userUpdate model.UserUpdatePayLoad) error {
	log := slog.With(
		slog.String("service", "user"),
		slog.String("func", "update"))

	user, err := us.userRepository.GetById(id)
	if err != nil {
		log.Error("Error: ", err)
		return model.ErrGetUser
	}

	if user == nil {
		log.Warn("User not found to update")
		return model.ErrUserNotFound
	}

	if user.Email == userUpdate.Email {
		log.Warn("Email same as above")
		return model.ErrSameEmail
	}

	if userUpdate.Email != "" {
		user.Email = userUpdate.Email
	}

	if userUpdate.Name != "" {
		user.Name = userUpdate.Name
	}

	if err := us.userRepository.Update(id, *user); err != nil {
		log.Error("Error: ", err)
		return model.ErrCreateUser
	}

	return nil
}

func (us userService) Delete(id string) error {
	log := slog.With(
		slog.String("service", "user"),
		slog.String("func", "delete"))

	user, err := us.userRepository.GetById(id)
	if err != nil {
		log.Error("Error trying to get user from repository")
		return model.ErrGetUser
	}

	if user == nil {
		log.Warn("User not found to delete")
		return model.ErrUserNotFound
	}

	if err := us.userRepository.Delete(id); err != nil {
		log.Error("Error: ", err)
		return model.ErrDeleteUser
	}

	return nil
}
