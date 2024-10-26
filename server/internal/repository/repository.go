package repository

import (
	"errors"
	"pwa/internal/entity"

	"gorm.io/gorm"
)

const base_img_url = "/images/"

type ProductModel struct {
	gorm.Model
	ID          uint
	Name        string
	Description string
	Price       float64
	ImageURL    string
}

func (ProductModel) TableName() string {
	return "product"
}

type ProductRepository interface {
	Create(product entity.Product) (entity.Product, error)
	Get() ([]entity.Product, error)
	GetById(id string) (entity.Product, error)
	DeleteById(id string) error
	UpdateById(product entity.Product) (entity.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) ProductRepository {

	return repository{db: db}
}

func (repo repository) Create(product entity.Product) (entity.Product, error) {

	model := &ProductModel{

		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		ImageURL:    base_img_url + product.ImageURL,
	}
	query := repo.db.Create(model)

	if query.Error != nil {

		return product, query.Error
	}

	return entity.Product{

		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		ImageURL:    model.ImageURL,
	}, nil

}

func (repo repository) Get() ([]entity.Product, error) {

	models := []ProductModel{}
	query := repo.db.Find(&models)
	if query.Error != nil {

		return []entity.Product{}, query.Error
	}

	products := []entity.Product{}
	for _, model := range models {

		products = append(products, entity.Product{

			ID:          model.ID,
			Name:        model.Name,
			Description: model.Description,
			Price:       model.Price,
			ImageURL:    model.ImageURL,
		})

	}

	return products, nil
}

func (repo repository) GetById(id string) (entity.Product, error) {

	model := ProductModel{}
	query := repo.db.Where("id = ?", id).Find(&model)
	if query.Error != nil {

		return entity.Product{}, query.Error
	}

	if query.RowsAffected == 0 {

		return entity.Product{}, errors.New("product not found")
	}

	return entity.Product{

		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		ImageURL:    model.ImageURL,
	}, nil

}

func (repo repository) DeleteById(id string) error {

	model := ProductModel{}
	query := repo.db.Where("id = ?", id).Find(&model)
	if query.Error != nil {

		return query.Error
	}

	if query.RowsAffected == 0 {

		return errors.New("product not found")
	}

	if query := repo.db.Delete(&model).Error; query != nil {

		return query
	}

	return nil
}

func (repo repository) UpdateById(product entity.Product) (entity.Product, error) {

	model := ProductModel{}
	if query := repo.db.First(&model, product.ID).Error; query != nil {

		return entity.Product{}, query
	}

	model.Name = product.Name
	model.Description = product.Description
	model.Price = product.Price
	model.ImageURL = product.ImageURL
	if query := repo.db.Save(&model).Error; query != nil {

		return entity.Product{}, query
	}

	return entity.Product{

		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		ImageURL:    model.ImageURL,
	}, nil

}
