package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/module/helper"
	modelProduct "github.com/nurulafifah149/golang/module/model/product"
	_ "github.com/nurulafifah149/golang/module/model/token"
	svcProduct "github.com/nurulafifah149/golang/module/service/product"
	"github.com/nurulafifah149/golang/pkg/logger"
	responseTemplate "github.com/nurulafifah149/golang/pkg/response"
)

type ProductHandlerImpl struct {
	ProductSVC svcProduct.ProductService
}

func NewProductHandler(ProductSVC svcProduct.ProductService) ProductHandler {
	return &ProductHandlerImpl{
		ProductSVC: ProductSVC,
	}
}

// @Summary Get All Product example
// @Schemes
// @Security Bearer
// @Description how to get all Product
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} responseTemplate.WebResponseSuccess{data=[]modelProduct.Product}
// @Router /product/getall [get]
func (p *ProductHandlerImpl) GetAll(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "Product Handler Invoked", "ProductHandler - Getall", nil)

	logger.LogMyApp("i", "Hit Product Service", "ProductHandler - Getall", nil)
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}

	data, _ := p.ProductSVC.GetAll(ctx, accessClaim.AccessClaims.UserId, accessClaim.AccessClaims.Role)
	if len(data) < 1 {
		logger.LogMyApp("i", "Render Response", "ProductHandler - Getall", nil)
		ctx.JSON(http.StatusNotFound, responseTemplate.WebResponseFailed{
			Message: "Data Is Not Found",
			Error:   "Data Is Not Found",
		})
		return
	}
	logger.LogMyApp("i", "Render Response", "ProductHandler - Getall", nil)
	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success GET all Product",
		Data:    data,
	})

	return
}

// @Summary Get product example
// @Schemes
// @Security Bearer
// @Description how to get product
// @Tags products
// @Accept json
// @Produce json
// @Param        id    path     int  false  "id product"
// @Success 200 {object} responseTemplate.WebResponseSuccess{data=modelProduct.Product}
// @Router /product/getone/{id} [get]
func (p *ProductHandlerImpl) GetById(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "product Handler Invoked", "ProductHandler - GetById", nil)
	idProduct, err := helper.GetIdAndConvertToInt(ctx)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit GetIdAndConvertInt helper", "ProductHandler - GetById", err)
		return
	}

	//hit service
	data, err := p.ProductSVC.GetById(ctx, idProduct)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit product Service", "ProductHandler - GetById", err)
		return
	}

	//render
	logger.LogMyApp("i", "Render Response", "ProductHandler - GetById", nil)
	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success GET product",
		Data:    data,
	})

}

// @Summary update product example
// @Schemes
// @Security Bearer
// @Description how to update product
// @Tags products
// @Accept json
// @Produce json
// @Param        id    path     int  false  "id product"
// @Param	request	body	modelProduct.ProductCreateRequest	true	"Input Data product"
// @Success 200 {object} responseTemplate.WebResponseSuccess{data=modelProduct.Product}
// @Router /product/updateproduct/{id} [put]
func (p *ProductHandlerImpl) Update(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "product Handler Invoked", "ProductHandler - Update", nil)

	var data modelProduct.ProductCreateRequest
	logger.LogMyApp("i", "GET ID product FROM PARAMS", "ProductHandler - Update", nil)
	idProduct, err := helper.GetIdAndConvertToInt(ctx)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit GetIdAndConvertInt helper", "ProductHandler - Update", err)
		return
	}

	logger.LogMyApp("i", "GET product DATA FROM JSON", "ProductHandler - Update", nil)
	err = ctx.BindJSON(&data)
	if err != nil {
		logger.LogMyApp("e", "Error When Get Params data", "ProductHandler - Update", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	logger.LogMyApp("i", "GET USERDATA FROM CTX", "ProductHandler - Update", nil)
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}

	logger.LogMyApp("i", "HIT product SERVICE", "ProductHandler - Update", nil)
	dataResp, err := p.ProductSVC.Update(ctx, modelProduct.Product{
		Id:          idProduct,
		Title:       data.Title,
		Description: data.Description,
	}, accessClaim.AccessClaims.UserId)

	if err != nil {
		logger.LogMyApp("e", "Error WHEN HIT product SERVICE", "ProductHandler - Update", err)
		return
	}

	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success Update product",
		Data:    dataResp,
	})
}

// @Summary create product example
// @Schemes
// @Security Bearer
// @Description how to Create product
// @Tags products
// @Accept json
// @Produce json
// @Param	request	body	modelProduct.ProductCreateRequest	true	"Input Data product"
// @Success 201 {object} responseTemplate.WebResponseSuccess{data=modelProduct.Product}
// @Router /product/createproduct [post]
func (p *ProductHandlerImpl) Create(ctx *gin.Context) {
	//panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "product Handler Invoked", "ProductHandler - Create", nil)

	var data modelProduct.ProductCreateRequest
	logger.LogMyApp("i", "GET product DATA FROM JSON", "ProductHandler - Create", nil)
	err := ctx.BindJSON(&data)
	if err != nil {
		logger.LogMyApp("e", "Error When Get Params data", "ProductHandler - Create", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	logger.LogMyApp("i", "GET USERDATA FROM CTX", "ProductHandler - Create", nil)
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}

	logger.LogMyApp("i", "HIT product SERVICE", "ProductHandler - Create", nil)
	dataResp, err := p.ProductSVC.Create(ctx, modelProduct.Product{
		UserId:      accessClaim.AccessClaims.UserId,
		Title:       data.Title,
		Description: data.Description,
	})
	if err != nil {
		logger.LogMyApp("e", "Error WHEN HIT product SERVICE", "ProductHandler - Create", err)
		return
	}

	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success Create product",
		Data:    dataResp,
	})

}

// @Summary delete product example
// @Schemes
// @Security Bearer
// @Description how to Delete product
// @Tags products
// @Accept json
// @Produce json
// @Param        id    path     int  false  "id product"
// @Success 200 {object} responseTemplate.WebResponseSuccess{}
// @Router /product/deleteproduct/{id} [delete]
func (p *ProductHandlerImpl) Delete(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "product Handler Invoked", "ProductHandler - Delete", nil)

	logger.LogMyApp("i", "GET ID product FROM PARAMS", "ProductHandler - Delete", nil)
	idProduct, err := helper.GetIdAndConvertToInt(ctx)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit GetIdAndConvertInt helper", "ProductHandler - Delete", err)
		return
	}

	logger.LogMyApp("i", "GET USERDATA FROM CTX", "ProductHandler - Delete", nil)
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}

	logger.LogMyApp("i", "HIT product SERVICE", "ProductHandler - Delete", nil)
	err = p.ProductSVC.Delete(ctx, idProduct, accessClaim.AccessClaims.UserId)
	if err != nil {
		logger.LogMyApp("e", "Error WHEN HIT product SERVICE", "ProductHandler - Delete", err)
		return
	}

	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success Delete product",
		Data:    nil,
	})
}
