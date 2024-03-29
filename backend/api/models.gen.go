// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package api

// AbbreviationResponse defines model for AbbreviationResponse.
type AbbreviationResponse struct {
	// Category The category of the abbreviation
	Category *string `json:"category,omitempty"`

	// Meaning The meaning of the abbreviation
	Meaning *string `json:"meaning,omitempty"`

	// Name The name of the abbreviation
	Name *string `json:"name,omitempty"`
}
