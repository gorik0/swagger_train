package handler

import "swagger/data"


//swagger:response productResponse
type productResponseWrapper struct {

	//	in:body
	Body []data.Product
}


//swagger:response noContentResponse
type noContentResponseWrapper struct {}




//swagger:response errorResponse
type errorResponseWrapper struct {

	//	in:body
	Body GenericError

}

//swagger:response errorValidate
type errorValidationWrapper struct {

	//	in:body
	Body ValidationError
}

//swagger:parameters createProduct
type createProductParamWrapper struct {

	//	in:body
	Body data.Product
}

//swagger:parameters updateProduct
type idWrapper struct {

	//	in:path
	ID int `json:"id"`
}







