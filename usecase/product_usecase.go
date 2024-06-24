package usecase

import (
	"github.com/josevitorrodriguess/productsAPI/model"
	"github.com/josevitorrodriguess/productsAPI/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error){
	products, err := pu.repository.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return  nil, err
	}

	return product, nil
}

func ( pu *ProductUsecase) DeleteProduct(id_product int) error {
		return pu.repository.DeleteProduct(id_product)
}
