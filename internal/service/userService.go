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

func (r *NewService) CheckEmail(ctx context.Context, passwordRequest models.UserRequest) (models.Response, error) {

	email := passwordRequest.Email
	dob := passwordRequest.Dob
	//dob := passwordRequest.Dob
	userData, err := r.rp.CheckUserData(ctx, email)
	if err != nil {
		log.Error().Err(err).Msg("error from user db")
		return models.Response{Msg: "Email data not found"}, errors.New("")
	}

	if userData.Dob != dob {
		log.Info().Msg("dod not matched")
		return models.Response{Msg: "Dob not macthed.. Enter valid dob"}, errors.New("")
	}

	if userData.Email != email {
		log.Info().Msg("email not matched")
		return models.Response{Msg: "Email is not matched.. Enter valid email address"}, errors.New("")
	}

	// fmt.Println("email", email)

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(90000) + 10000
	otp := strconv.Itoa(randomNumber)

	response := r.rdb.AddOtpToCache(email, otp)

	if !response {
		return models.Response{Msg: "failed to add cache"}, errors.New("")
	}

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
		return models.Response{Msg: "otp not send"}, errors.New("")
	}

	//fmt.Println("email sent successfully!")

	return models.Response{Msg: "Otp has sent to your email " + email + " successfully"}, nil

}

func (r NewService) CheckOtpResponse(ctx context.Context, otpVerification models.CheckOtp) (models.Response, error) {

	email := otpVerification.Email
	otp := otpVerification.Otp

	otpData, err := r.rdb.CheckOtpRequest(email)
	if err != nil {
		log.Error().Err(err).Msg("email details not found")
		return models.Response{Msg: "Enter valid email id"}, err
	}

	if otp != otpData {

		return models.Response{Msg: "Otp not Matched .. Please enter valid otp"}, errors.New("")
	}

	if otpVerification.NewPassword != otpVerification.ConfirmPassword {
		log.Error().Msg("Password not matched")
		return models.Response{Msg: "New Password and Confirm Password not matched"}, errors.New("")

	}
	savePasswordToDatabase := r.rp.SavePassword(ctx, otpVerification)

	if !savePasswordToDatabase {
		log.Error().Msg("Failed to store password in db")
		return models.Response{Msg: "Failed to store password in database"}, errors.New("")
	}

	err = r.rdb.DeleteCacheData(email)
	if err != nil {
		log.Error().Msg("otp not found in cache")
		return models.Response{Msg: ""}, errors.New("")
	}

	return models.Response{Msg: "You have changed your account password linked to  " + email + "  successfully"}, nil

}
