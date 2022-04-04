package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"silocorp/golang-restful-api/helper"
	"silocorp/golang-restful-api/model/web"
	"silocorp/golang-restful-api/service"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	inputStruct := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &inputStruct)

	categoryResponse := controller.CategoryService.Create(request.Context(), inputStruct)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	inputStruct := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &inputStruct)

	categoryID, err := strconv.Atoi(params.ByName("categoryID"))
	helper.PanicIfError(err)

	inputStruct.ID = int64(categoryID)

	categoryResponse := controller.CategoryService.Update(request.Context(), inputStruct)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryID, err := strconv.Atoi(params.ByName("categoryID"))
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), int64(categoryID))
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CategoryControllerImpl) FindByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryID, err := strconv.Atoi(params.ByName("categoryID"))
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindByID(request.Context(), int64(categoryID))
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
