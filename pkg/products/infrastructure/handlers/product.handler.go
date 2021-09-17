package handlers

import (
	use_case "api_test_hexagonal/pkg/products/application/use-case"
	"api_test_hexagonal/pkg/products/domain/entities"
	"api_test_hexagonal/pkg/products/infrastructure/services/products"
	"api_test_hexagonal/pkg/utils"
	"encoding/json"
	"net/http"
)

type ProductHandler struct {
	service *products.ProductService
}

func Product() ProductHandler {
	return ProductHandler{}
}

func (product *ProductHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	if utils.VerifyMethodHttp(writer, request.Method, http.MethodGet) { return }

	productUseCase := use_case.ProductUseCase{Repository: product.service}

	productList, err := productUseCase.GetAllProducts()
	if err != nil {
		response := utils.NewResponse(utils.Error, "Error en el get", nil)
		utils.ResponseJSON(writer, http.StatusBadRequest, response)
	}

	response := utils.NewResponse(utils.Message, "Lista de productos", productList)
	utils.ResponseJSON(writer, http.StatusCreated, response)
}

func (product *ProductHandler) Create(writer http.ResponseWriter, request *http.Request)  {
	if utils.VerifyMethodHttp(writer, request.Method, http.MethodPost) { return }

	productUseCase := use_case.ProductUseCase{Repository: product.service}
	var newProduct entities.IProductEntity

	err := json.NewDecoder(request.Body).Decode(&newProduct)
	if err != nil {
		response := utils.NewResponse(utils.Error, "JSON mal formado", nil)
		utils.ResponseJSON(writer, http.StatusBadRequest, response)
		return
	}

	errProductCase := productUseCase.CreateProduct(newProduct)
	if errProductCase != nil {
		response := utils.NewResponse(utils.Error, "Error al crear el producto", nil)
		utils.ResponseJSON(writer, http.StatusBadRequest, response)
		return
	}

	response := utils.NewResponse(utils.Message, "Producto creado con exito", nil)
	utils.ResponseJSON(writer, http.StatusCreated, response)
}

func (product *ProductHandler) Update(writer http.ResponseWriter, request *http.Request)  {
	if utils.VerifyMethodHttp(writer, request.Method, http.MethodPut) { return }

	productUseCase := use_case.ProductUseCase{Repository: product.service}
	var updatedProduct entities.IProductEntity

	err := json.NewDecoder(request.Body).Decode(&updatedProduct)
	if err != nil {
		response := utils.NewResponse(utils.Error, "JSON mal formado", nil)
		utils.ResponseJSON(writer, http.StatusBadRequest, response)
		return
	}

	productID := request.URL.Query().Get("id")

	errProductCase := productUseCase.UpdateProduct(updatedProduct, productID)
	if errProductCase != nil {
		response := utils.NewResponse(utils.Error, "Error al actualizar el producto", nil)
		utils.ResponseJSON(writer, http.StatusBadRequest, response)
		return
	}

	response := utils.NewResponse(utils.Message, "Producto actualizado con exito", nil)
	utils.ResponseJSON(writer, http.StatusOK, response)
}

func (product *ProductHandler) Delete(writer http.ResponseWriter, request *http.Request)  {
	if utils.VerifyMethodHttp(writer, request.Method, http.MethodPut) { return }
	//productUseCase := use_case.ProductUseCase{Repository: product.service}

	//err := productUseCase.DeleteProducts("12345")
}
