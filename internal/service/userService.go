package service

import (
	"context"
	"errors"
	"fmt"
	"job-portal/internal/models"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (r NewService) CreateUser(ctx context.Context, nu models.NewUser) (models.User, error) {
	userDetails, err := r.rp.CreateU(ctx, nu)
	if err != nil {
		return models.User{}, err
	}
	return userDetails, err
}
func (r NewService) Authenticate(ctx context.Context, email string, password string) (jwt.RegisteredClaims, error) {
	userData, err := r.rp.AuthenticateUser(ctx, email, password)
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}
	return userData, err
}

func (r *NewService) CheckEmail(ctx context.Context, passwordRequest models.UserRequest) (string, error) {

	email := passwordRequest.Email
	dob := passwordRequest.Dob
	//dob := passwordRequest.Dob
	err := r.rp.CheckUserData(ctx, email, dob)

	if !err {
		return "Email not found", errors.New("")
	}

	// fmt.Println("email", email)

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(90000) + 10000
	otp := strconv.Itoa(randomNumber)

	r.rdb.AddOtpToCache(email, otp)

	// Sender's email address and password
	from := "suryatejamulabagal@gmail.com"
	password := "rejz mrjt ypkw lyfc"

	//email := passwordRequest.Email
	// Recipient's email address
	to := email

	// SMTP server details
	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	// Message content
	message := []byte("Subject: otp details\n\n  " + "your otp for changing password " + otp)

	// Authentication information
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// SMTP connection
	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	ok := smtp.SendMail(smtpAddr, auth, from, []string{to}, message)
	if ok != nil {
		fmt.Println("Error sending email:", err)
		return "otp not sent", errors.New("")
	}

	fmt.Println("email sent successfully!")

	return "otp sent succesfully", nil

}

func (r NewService) CheckOtpResponse(ctx context.Context, otpVerification models.CheckOtp) (string, error) {

	email := otpVerification.Email
	otp := otpVerification.Otp
	otpData, err := r.rdb.CheckOtpRequest(email)
	if err != nil {
		log.Error().Msg("email details not found")
		return "No such email exists in database", errors.New("")
	}

	if otp != otpData {
		log.Error().Msg("Otp Not Matched")
		return "Failed to verify otp", errors.New("")
	}

	if otpVerification.NewPassword != otpVerification.ConfirmPassword {
		log.Error().Msg("Password not matched")
		return "password not matched", errors.New("")

	}
	savePasswordToDatabase := r.rp.SavePassword(ctx, otpVerification)

	if !savePasswordToDatabase {
		log.Error().Msg("Failed to store password in db")
		return "Failed to change password", errors.New("")
	}

	err = r.rdb.DeleteCacheData(email)
	if err != nil {
		return "Password not matched", errors.New("")
	}

	return "password changed successfully", nil

}
