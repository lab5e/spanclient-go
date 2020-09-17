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
// Output struct for Output
type Output struct {
	OutputId string `json:"outputId,omitempty"`
	CollectionId string `json:"collectionId,omitempty"`
	Type OutputType `json:"type,omitempty"`
	Config OutputConfig `json:"config,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
}