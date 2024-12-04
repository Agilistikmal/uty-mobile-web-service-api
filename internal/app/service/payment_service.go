package service

import (
	"context"
	"fmt"

	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/model"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/repository"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/pkg"
	"github.com/go-playground/validator/v10"
	"github.com/xendit/xendit-go/v6"
	"github.com/xendit/xendit-go/v6/payment_request"
)

type PaymentService struct {
	xenditClient      *xendit.APIClient
	paymentRepository *repository.PaymentRepository
	userRepository    *repository.UserRepository
	validator         *validator.Validate
}

func NewPaymentService(xenditClient *xendit.APIClient, paymentRepository *repository.PaymentRepository, userRepository *repository.UserRepository, validator *validator.Validate) *PaymentService {
	return &PaymentService{
		xenditClient:      xenditClient,
		paymentRepository: paymentRepository,
		userRepository:    userRepository,

		validator: validator,
	}
}

// func (s *PaymentService) Create(payment *model.Payment) (*model.Payment, error) {
// 	err := s.validator.Struct(payment)
// 	if err != nil {
// 		return nil, err
// 	}

// 	_, err = s.userRepository.Find(payment.Username)
// 	if err != nil {
// 		return nil, err
// 	}

// 	referenceID := "AGL-" + pkg.RandomString(8)

// 	createInvoiceRequest := *invoice.NewCreateInvoiceRequest(referenceID, float64(1000))

// 	invoice, _, xdtErr := s.xenditClient.InvoiceApi.CreateInvoice(context.Background()).
// 		CreateInvoiceRequest(createInvoiceRequest).
// 		Execute()
// 	if xdtErr != nil {
// 		return nil, xdtErr
// 	}

// 	payment = &model.Payment{
// 		ID:          *invoice.Id,
// 		ReferenceID: invoice.ExternalId,
// 		Url:         invoice.InvoiceUrl,
// 		Username:    payment.Username,
// 		Amount:      int(invoice.Amount),
// 		Status:      "PENDING",
// 	}

// 	payment, err = s.paymentRepository.Create(payment)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return payment, nil
// }

func (s *PaymentService) Create(payment *model.Payment) (*model.Payment, error) {
	err := s.validator.Struct(payment)
	if err != nil {
		return nil, err
	}

	_, err = s.userRepository.Find(payment.Username)
	if err != nil {
		return nil, err
	}

	referenceID := "AGL-" + pkg.RandomString(8)
	amount := float64(1000)

	payload := payment_request.PaymentRequestParameters{
		ReferenceId: &referenceID,
		Amount:      &amount,
		Currency:    payment_request.PAYMENTREQUESTCURRENCY_IDR,
		PaymentMethod: payment_request.NewPaymentMethodParameters(
			payment_request.PAYMENTMETHODTYPE_QR_CODE,
			payment_request.PAYMENTMETHODREUSABILITY_ONE_TIME_USE,
		),
	}

	response, _, _ := s.xenditClient.PaymentRequestApi.
		CreatePaymentRequest(context.Background()).
		PaymentRequestParameters(payload).
		Execute()

	if response.Id == "" {
		return nil, fmt.Errorf("Failed to create payment")
	}

	payment = &model.Payment{
		ID:          response.Id,
		ReferenceID: response.ReferenceId,
		Username:    payment.Username,
		Amount:      int(*response.Amount),
		Status:      "PENDING",
		QrString:    *response.PaymentMethod.QrCode.Get().ChannelProperties.QrString,
	}

	payment, err = s.paymentRepository.Create(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *PaymentService) FindByID(id string) (*model.Payment, error) {
	payment, err := s.paymentRepository.FindByID(id)
	return payment, err
}

// func (s *PaymentService) FindByReferenceID(referenceID string) (*model.Payment, error) {
// 	payment, err := s.paymentRepository.FindByReferenceID(referenceID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if payment.Status == "PENDING" {
// 		invoice, _, _ := s.xenditClient.InvoiceApi.GetInvoiceById(context.Background(), payment.ID).Execute()
// 		payment.Status = string(invoice.Status)
// 		payment, err = s.paymentRepository.Update(payment.ID, payment)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	if payment.Status == "PAID" {
// 		user, err := s.userRepository.Find(payment.Username)
// 		if err != nil {
// 			return nil, err
// 		}

// 		if user.Verified == false {
// 			user.Verified = true

// 			_, err = s.userRepository.Update(user.Username, user)
// 			if err != nil {
// 				return nil, err
// 			}
// 		}
// 	}

// 	return payment, nil
// }

func (s *PaymentService) FindByReferenceID(referenceID string) (*model.Payment, error) {
	payment, err := s.paymentRepository.FindByReferenceID(referenceID)
	if err != nil {
		return nil, err
	}

	if payment.Status == "PENDING" {
		response, _, _ := s.xenditClient.PaymentRequestApi.GetPaymentRequestByID(context.Background(), payment.ID).Execute()
		payment.Status = string(response.Status)
		payment, err = s.paymentRepository.Update(payment.ID, payment)
		if err != nil {
			return nil, err
		}
	}

	if payment.Status == "SUCCEEDED" {
		user, err := s.userRepository.Find(payment.Username)
		if err != nil {
			return nil, err
		}

		if user.Verified == false {
			user.Verified = true

			_, err = s.userRepository.Update(user.Username, user)
			if err != nil {
				return nil, err
			}
		}
	}

	return payment, nil
}
