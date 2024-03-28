package sales

import (
	"hacktiv8-techrawih-go-product-sale-concurrency/internal/pkg/http/request/sales"
	"hacktiv8-techrawih-go-product-sale-concurrency/internal/pkg/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPController interface {
	Add(c *gin.Context)
}

// NewHTTPController ...
func NewHTTPController(salesService Service) HTTPController {
	return &httpController{
		salesService: salesService,
	}
}

type httpController struct {
	salesService Service
}

func (controller *httpController) Add(ctx *gin.Context) {
	var request sales.Import
	err := ctx.ShouldBind(&request)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	src, err := request.File.Open()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	defer src.Close()

	request.FilePath = "internal/pkg/shared/import-files/" + request.File.Filename
	err = ctx.SaveUploadedFile(&request.File, request.FilePath)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	_, err = controller.salesService.SaveAll(request)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	response.SuccessCustomMessage(ctx, http.StatusOK, "Import success")
}
