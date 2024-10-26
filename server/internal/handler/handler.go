package handler

import (
	"fmt"
	"pwa/internal/entity"
	"pwa/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	CreateProduct(c *fiber.Ctx) error
	GetProducts(c *fiber.Ctx) error
	GetProductById(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
	UpdateProductById(c *fiber.Ctx) error
}

type handler struct {
	uc usecase.ProductUsecase
}

type payload struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
}

type response struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
}

func NewHandler(uc usecase.ProductUsecase) ProductHandler {

	return &handler{uc: uc}
}

func (hdr handler) CreateProduct(c *fiber.Ctx) error {

	payload := payload{}
	if err := getPayload(c, &payload); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"payload error": err.Error()})
	}

	newProduct, err := hdr.uc.CreateProduct(entity.Product{
		ID:          payload.ID,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		ImageURL:    payload.ImageURL,
	})
	if err != nil {

		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"create error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{

		"product created": response{
			ID:          newProduct.ID,
			Name:        newProduct.Name,
			Description: newProduct.Description,
			Price:       newProduct.Price,
			ImageURL:    newProduct.ImageURL,
		},
	})
}

func (hdr handler) GetProducts(c *fiber.Ctx) error {

	products, err := hdr.uc.GetProducts()
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"get error": err.Error()})
	}

	responses := []response{}
	for _, product := range products {

		responses = append(responses, response{

			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			ImageURL:    product.ImageURL,
		})

	}

	return c.Status(fiber.StatusOK).JSON(responses)
}

func (hdr handler) GetProductById(c *fiber.Ctx) error {

	product, err := hdr.uc.GetProductById(c.Params("id"))
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"product by id error": err.Error()})
	}

	response := response{

		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		ImageURL:    product.ImageURL,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"product by id": response})
}

func (hdr handler) DeleteProduct(c *fiber.Ctx) error {

	err := hdr.uc.DeleteProductById(c.Params("id"))
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"delete error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("product id '%s' has been deleted", c.Params("id"))})
}

func (hdr handler) UpdateProductById(c *fiber.Ctx) error {

	payload := payload{}
	if err := getPayload(c, &payload); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"payload error": err.Error()})
	}

	product, err := hdr.uc.UpdateProductById(entity.Product{

		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		ImageURL:    payload.ImageURL,
	}, c.Params("id"))
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"update product error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"product has been updated": response{

		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		ImageURL:    product.ImageURL,
	}})

}

func getPayload(c *fiber.Ctx, payload *payload) error {

	if err := c.BodyParser(payload); err != nil {

		return err
	}

	return nil
}
