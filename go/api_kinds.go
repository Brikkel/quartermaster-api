/*
 * Swagger K8S Quartermaster - OpenAPI 3.0
 *
 * This is the API design of the K8S quartermaster tool, the goal is to effectively obtain K8S resources, which can then be used for YAML configurations or other pruposes.  Some useful links: - [Swagger](https://swagger.io/) - [Kubernetes](https://kubernetes.io/docs/concepts/configuration/)
 *
 * API version: 0.0.1
 * Contact: rik.van.brakel@sue.nl
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	// "encoding/json"
	"net/http"
	"strings"
	// "github.com/go-chi/chi/v5"
)

// KindsAPIController binds http requests to an api service and writes the service results to the http response
type KindsAPIController struct {
	service      KindsAPIServicer
	errorHandler ErrorHandler
}

// KindsAPIOption for how the controller is set up.
type KindsAPIOption func(*KindsAPIController)

// WithKindsAPIErrorHandler inject ErrorHandler into controller
func WithKindsAPIErrorHandler(h ErrorHandler) KindsAPIOption {
	return func(c *KindsAPIController) {
		c.errorHandler = h
	}
}

// NewKindsAPIController creates a default api controller
func NewKindsAPIController(s KindsAPIServicer, opts ...KindsAPIOption) Router {
	controller := &KindsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the KindsAPIController
func (c *KindsAPIController) Routes() Routes {
	return Routes{
		"GetKinds": Route{
			strings.ToUpper("Get"),
			"/v3/kinds",
			c.GetKinds,
		},
	}
}

// GetKinds - Returns pet inventories by status
func (c *KindsAPIController) GetKinds(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetKinds(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
