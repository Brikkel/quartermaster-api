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

// ImplResponse defines an implementation response with error code and the associated body
type ImplResponse struct {
	Code int
	Body interface{}
}