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

	quartermaster "github.com/brikkel/k8s-api-spec/openapi-v3"
)

// KindsAPIService is a service that implements the logic for the KindsAPIServicer
// This service should implement the business logic for every endpoint for the KindsAPI API.
// Include any external packages or services that will be required by this service.
type KindsAPIService struct {
}

// NewKindsAPIService creates a default api service
func NewKindsAPIService() KindsAPIServicer {
	return &KindsAPIService{}
}

// GetKinds - Returns pet inventories by status
func (s *KindsAPIService) GetKinds(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetKinds with the required logic for this service method.
	// Add api_kinds_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	myKinds, err := quartermaster.GetAllKinds("http://localhost:8001")
	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}
	if myKinds == nil {
		return Response(http.StatusNotFound, nil), errors.New("No kinds found")
	}

	return Response(200, myKinds), nil
}
