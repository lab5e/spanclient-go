/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.3 factual-kahlil
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanclient
// SystemInfoResponse struct for SystemInfoResponse
type SystemInfoResponse struct {
	Version string `json:"version,omitempty"`
	BuildDate string `json:"buildDate,omitempty"`
	ReleaseName string `json:"releaseName,omitempty"`
	DefaultFieldMask FieldMask `json:"defaultFieldMask,omitempty"`
	ForcedFieldMask FieldMask `json:"forcedFieldMask,omitempty"`
}
