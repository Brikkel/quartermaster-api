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




type Kinds struct {

	ApiVersionPath string `json:"apiVersionPath"`

	Kinds []Kind `json:"kinds,omitempty"`
}

// AssertKindsRequired checks if the required fields are not zero-ed
func AssertKindsRequired(obj Kinds) error {
	elements := map[string]interface{}{
		"apiVersionPath": obj.ApiVersionPath,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Kinds {
		if err := AssertKindRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertKindsConstraints checks if the values respects the defined constraints
func AssertKindsConstraints(obj Kinds) error {
	return nil
}