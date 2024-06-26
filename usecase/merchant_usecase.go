package usecase

import (
	"fmt"

	"github.com/josevitorrodriguess/productsAPI/model"
	"github.com/josevitorrodriguess/productsAPI/repository"
	"github.com/josevitorrodriguess/productsAPI/services"
)

type MerchantUsecase struct {
	repository repository.MerchantRepository
}

func NewMerchantUsecase(repo repository.MerchantRepository) MerchantUsecase {
	return MerchantUsecase{
		repository: repo,
	}
}

func (mu *MerchantUsecase) GetMerchants() ([]model.Merchant, error) {
	return mu.repository.GetMerchants()
}

func (mu *MerchantUsecase) GetMerchantById(id_merchant int) (*model.Merchant, error) {
	merchant, err := mu.repository.GetMerchantByID(id_merchant)
	if err != nil {
		return nil, err
	}

	return merchant, nil
}

func (mu *MerchantUsecase) CreateMerchant(merchant model.Merchant) (model.Merchant, error) {

	isValidEmail, err := services.ValidateEmail(merchant.Email)
	if err != nil {
		return model.Merchant{}, fmt.Errorf("failed to validate email: %v", err)
	}
	if !isValidEmail {
		return model.Merchant{}, fmt.Errorf("invalid email format: %s", merchant.Email)
	}

	isValidCNPJ, err := services.ValidateCNPJ(merchant.CNPJ)
	if err != nil {
		return model.Merchant{}, fmt.Errorf("failed to validate CNPJ %v", merchant.CNPJ)
	}
	if !isValidCNPJ {
		return model.Merchant{}, fmt.Errorf("invalid cnpj format: %s", merchant.CNPJ)
	}
	

	merchantId, err := mu.repository.CreateMerchant(merchant)
	if err != nil {
		return model.Merchant{}, err
	}

	merchant.Password = services.SHA256Encoder(merchant.Password)
	merchant.ID = merchantId

	email := model.Email{
		Subject: "Welcome! Your account was created",
		Body:    "Your account sucessfull created.",
	}

	services.SendMail(merchant.Email, email)

	return merchant, nil
}

func (mu *MerchantUsecase) DeleteMerchant(id_merchant int) error {
	return mu.repository.DeleteMerchant(id_merchant)
}
