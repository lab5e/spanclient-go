/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 3.0.0
 * Contact: dev@lab5e.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package spanclient
// DataDumpResponse struct for DataDumpResponse
type DataDumpResponse struct {
	Collections []DumpedCollection `json:"collections,omitempty"`
	Profile UserProfile `json:"profile,omitempty"`
	Teams []Team `json:"teams,omitempty"`
	Tokens []Token `json:"tokens,omitempty"`
}