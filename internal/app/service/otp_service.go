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
	otpRepository *repository.OTPRepository
}

func NewOTPService(otpRepository *repository.OTPRepository) *OTPService {
	return &OTPService{
		otpRepository: otpRepository,
	}
}

func (s *OTPService) Generate(user *model.User) (*model.OTP, error) {
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

func (s *OTPService) Verify(username string, code string) error {
	otp, err := s.otpRepository.Find(username)
	if err != nil {
		return fmt.Errorf("invalid user")
	}

	if otp.Code != code {
		return fmt.Errorf("invalid code")
	}

	return nil
}
