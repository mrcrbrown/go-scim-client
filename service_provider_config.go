/*
 * SCIM API
 *
 * SCIM V2 API implemented by RingCentral
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package scim

type ServiceProviderConfig struct {
	AuthenticationSchemes *AuthenticationSchemes `json:"authenticationSchemes,omitempty"`

	Bulk *BulkSupported `json:"bulk,omitempty"`

	ChangePassword *Supported `json:"changePassword,omitempty"`

	Etag *Supported `json:"etag,omitempty"`

	Filter *FilterSupported `json:"filter,omitempty"`

	Patch *Supported `json:"patch,omitempty"`

	Schemas []string `json:"schemas,omitempty"`

	Sort *Supported `json:"sort,omitempty"`

	XmlDataFormat *Supported `json:"xmlDataFormat,omitempty"`
}
