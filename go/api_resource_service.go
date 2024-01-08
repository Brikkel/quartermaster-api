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
	"context"
	"errors"
	"net/http"
	"net/url"

	quartermaster "github.com/brikkel/k8s-api-spec/openapi-v3"
)

// ResourceAPIService is a service that implements the logic for the ResourceAPIServicer
// This service should implement the business logic for every endpoint for the ResourceAPI API.
// Include any external packages or services that will be required by this service.
type ResourceAPIService struct {
}

// NewResourceAPIService creates a default api service
func NewResourceAPIService() ResourceAPIServicer {
	return &ResourceAPIService{}
}

// GetParameters - Gives parameters of a resource
func (s *ResourceAPIService) GetParameters(ctx context.Context, fullResourceName string, versionPath string) (ImplResponse, error) {
	// TODO - update GetParameters with the required logic for this service method.
	// Add api_resource_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// The path is url encoded, so needs to be decoded
	rawpath, err := url.PathUnescape(versionPath)
	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	myResource, err := quartermaster.GetResource("http://localhost:8001", rawpath, fullResourceName)
	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}
	if myResource == nil {
		return Response(http.StatusNotFound, nil), errors.New("No Resource found")
	}

	return Response(200, myResource), nil
}