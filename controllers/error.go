package controllers

type ErrorController struct {
	BaseController
}

func (o *ErrorController) Error404() {
	o.ServeError(404)
}

func (o *ErrorController) Error500() {
	o.ServeError(500)
}
