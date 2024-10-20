package service

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/model"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/repository"
	"github.com/spf13/viper"
)

type OTPService struct {
	otpRepository  *repository.OTPRepository
	userRepository *repository.UserRepository
}

func NewOTPService(otpRepository *repository.OTPRepository, userRepository *repository.UserRepository) *OTPService {
	return &OTPService{
		otpRepository:  otpRepository,
		userRepository: userRepository,
	}
}

func (s *OTPService) Generate(username string) (*model.OTP, error) {
	user, err := s.userRepository.Find(username)
	if err != nil {
		return nil, err
	}

	otp, err := s.otpRepository.Create(user.Username)
	if err != nil {
		return nil, err
	}

	var otpMessageBuffer bytes.Buffer
	otpMessageTemplate := template.Must(template.New("otp_message").Parse(viper.GetString("otp.message")))
	otpMessageTemplate.Execute(&otpMessageBuffer, map[string]string{
		"Username": otp.Username,
		"Code":     otp.Code,
	})

	otpMessage := strings.ReplaceAll(otpMessageBuffer.String(), "\n", `\n`)

	body := []byte(fmt.Sprintf(`{
		"session": "default",
		"chatId":  "%s",
		"text":    "%s"
	}`, user.Phone[1:]+"@c.us", otpMessage))

	endpoint := viper.GetString("waha.base_url") + "/api/sendText"

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("failed to send otp")
	}

	return otp, nil
}

func (s *OTPService) Verify(username string, code string) (*model.User, error) {
	otp, err := s.otpRepository.Find(username)
	if err != nil {
		return nil, err
	}

	if otp.Code != code {
		return nil, fmt.Errorf("invalid code")
	}

	user, err := s.userRepository.Update(username, &model.User{
		Verified: true,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
