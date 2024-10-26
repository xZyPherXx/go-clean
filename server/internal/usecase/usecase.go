package usecase

import (
	"pwa/internal/entity"
	"pwa/internal/repository"
)

type ProductUsecase struct {
	repo repository.ProductRepository
}

func NewUsecase(repo repository.ProductRepository) ProductUsecase {

	return ProductUsecase{repo: repo}
}

func (uc ProductUsecase) CreateProduct(product entity.Product) (entity.Product, error) {

	// newProduct, err := uc.repo.Create(product)
	// if err != nil {

	// 	return newProduct, err
	// }

	return uc.repo.Create(product)
}

func (uc ProductUsecase) GetProducts() ([]entity.Product, error) {

	// products, err := uc.repo.Get()
	// if err != nil {

	// 	return products, err
	// }

	return uc.repo.Get()
}

func (uc ProductUsecase) GetProductById(id string) (entity.Product, error) {

	return uc.repo.GetById(id)
}

func (uc ProductUsecase) DeleteProductById(id string) error {

	return uc.repo.DeleteById(id)
}

func (uc ProductUsecase) UpdateProductById(productToUpdate entity.Product, id string) (entity.Product, error) {

	productById, err := uc.repo.GetById(id)
	if err != nil {

		return productById, err
	}

	productToUpdate.ID = productById.ID
	return uc.repo.UpdateById(productToUpdate)
}
