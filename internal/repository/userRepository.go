package repository

import (
	"context"
	"fmt"
	"job-portal/internal/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Conn) CreateU(ctx context.Context, nu models.NewUser) (models.User, error) {

	// We hash the user's password for storage in the database.
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("generating password hash: %w", err)
	}
	u := models.User{
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: string(hashedPass),
		Dob:          nu.Dob,
	}
	err = s.db.Create(&u).Error
	if err != nil {
		return models.User{}, err
	}

	// Successfully created the record, return the user.
	return u, nil
}

func (s *Conn) AuthenticateUser(ctx context.Context, email, password string) (jwt.RegisteredClaims,
	error) {
	var u models.User
	tx := s.db.Where("email = ?", email).First(&u)
	if tx.Error != nil {
		return jwt.RegisteredClaims{}, tx.Error
	}

	// We check if the provided password matches the hashed password in the database.
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}

	// Successful authentication! Generate JWT claims.
	c := jwt.RegisteredClaims{
		Issuer:    "job-portal-api",
		Subject:   strconv.FormatUint(uint64(u.ID), 10),
		Audience:  jwt.ClaimStrings{"students"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	return c, nil
}

func (s *Conn) CheckUserData(ctx context.Context, email string, dob string) bool {
	var data models.User
	// tx := s.db.Debug().Where("email = ? AND dob = ?", email, dob).Find(&data)
	tx := s.db.Where("email = ?", email).Where("dob = ?", dob).First(&data)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return false
	}
	return true
}

func (s *Conn) SavePassword(ctx context.Context, otp models.CheckOtp) bool {

	if otp.NewPassword != otp.ConfirmPassword {
		log.Error().Msg("password not matched")
		return false
	}
	Pass := otp.NewPassword
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(Pass), bcrypt.DefaultCost)

	if err != nil {
		log.Error().Msg("hashed password error")
		return false
	}

	tx := s.db.Model(&models.User{}).Where("email = ? ", otp.Email).Update("PasswordHash", hashedPass)

	if tx.Error != nil || tx.RowsAffected == 0 {
		return false
	}

	return true
}
