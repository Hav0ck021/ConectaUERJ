package repository

import (
	"errors"
	"log/slog"
	"time"

	"gorm.io/gorm"

	"github.com/OVillas/user-api/config/database"
	"github.com/OVillas/user-api/model"
)

type userRepository struct{}

func NewUserRepository() model.UserRepository {
	return userRepository{}
}

func (ur userRepository) Create(user model.User) error {
	log := slog.With(
		slog.String("func", "Create"),
		slog.String("repository", "user"))

	db, err := database.NewMysqlConnection()
	if err != nil {
		log.Error("Error connecting to the database", err)
		return err
	}

	now := time.Now()

	user.CreatedAt = now
	user.LastModified = now

	result := db.Create(&user)

	if result.Error != nil {
		log.Error("Error to create user in database", result.Error)
		return result.Error
	}

	log.Info("create repository executed successfully")
	return nil
}

func (ur userRepository) GetAll() ([]model.User, error) {
	log := slog.With(
		slog.String("func", "GetAll"),
		slog.String("repository", "user"))

	db, err := database.NewMysqlConnection()
	if err != nil {
		log.Error("Error connecting to the database", err)
		return nil, err
	}
	var users []model.User

	result := db.Find(&users)

	err = result.Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Error: ", err)
		return nil, err
	}

	log.Info("get all repository executed successfully")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return users, nil
}

func (ur userRepository) GetById(id string) (*model.User, error) {
	log := slog.With(
		slog.String("func", "GetById"),
		slog.String("repository", "user"))

	db, err := database.NewMysqlConnection()
	if err != nil {
		log.Error("Error connecting to the database", err)
		return nil, err
	}

	var user model.User
	err = db.Where("id = ?", id).First(&user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	log.Info("get by id service executed successfully")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, nil
}

func (ur userRepository) GetByName(name string) ([]model.User, error) {
	log := slog.With(
		slog.String("func", "GetByName"),
		slog.String("repository", "user"))

	db, err := database.NewMysqlConnection()
	if err != nil {
		log.Error("Error connecting to the database", err)
		return nil, err
	}

	var users []model.User

	err = db.Where("name LIKE ?", "%"+name+"%").Find(&users).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Error: ", err)
		return nil, err
	}

	log.Info("get by name repository executed successfully")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return users, nil
}

func (ur userRepository) GetByEmail(email string) (*model.User, error) {
	log := slog.With(
		slog.String("func", "GetByEmail"),
		slog.String("repository", "user"))

	db, err := database.NewMysqlConnection()
	if err != nil {
		log.Error("Error connecting to the database", err)
		return nil, err
	}

	var user model.User
	err = db.Where("email = ?", email).First(&user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Error: ", err)
		return nil, err
	}

	log.Info("get by email repository executed successfully")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, nil
}

func (ur userRepository) Update(id string, user model.User) error {
	log := slog.With(
		slog.String("func", "Create"),
		slog.String("repository", "user"))

	db, err := database.NewMysqlConnection()
	if err != nil {
		log.Error("Error connecting to the database", err)
		return err
	}

	err = db.Model(&model.User{}).Where("id = ?", id).Updates(model.User{Name: user.Name, Email: user.Email}).Error
	if err != nil {
		log.Error("Error: ", err)
		return err
	}

	log.Info("update repository executed successfully")
	return nil
}

func (ur userRepository) Delete(id string) error {
	log := slog.With(
		slog.String("func", "Delete"),
		slog.String("repository", "user"))

	db, err := database.NewMysqlConnection()
	if err != nil {
		log.Error("Error: ", err)
		return err
	}
	err = db.Delete(&model.User{}, "id = ?", id).Error
	if err != nil {
		log.Error("Error: ", err)
		return err
	}

	log.Info("delete repository executed successfully")
	return nil
}

func (ur userRepository) UpdatePassword(id string, password string) error {
	log := slog.With(
		slog.String("func", "updatePassword"),
		slog.String("repository", "user"))

	db, err := database.NewMysqlConnection()
	if err != nil {
		log.Error("Error connecting to the database", err)
		return err
	}

	err = db.Model(&model.User{}).Where("id = ?", id).Updates(model.User{Password: password}).Error
	if err != nil {
		log.Error("Error: ", err)
		return err
	}

	log.Info("update password repository executed successfully")
	return nil
}

func (ur userRepository) UpdateConfirmedEmail(id string) error {
	log := slog.With(
		slog.String("func", "UpdateConfirmedEmail"),
		slog.String("repository", "user"))

	db, err := database.NewMysqlConnection()
	if err != nil {
		log.Error("Error connecting to the database", err)
		return err
	}

	err = db.Model(&model.User{}).Where("id = ?", id).Updates(model.User{IsEmailConfirmed: true}).Error
	if err != nil {
		log.Error("Error: ", err)
		return err
	}

	log.Info("update confirmed email repository executed successfully")
	return nil
}
