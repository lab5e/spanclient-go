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
// OutputType the model 'OutputType'
type OutputType string

// List of OutputType
const (
	UNDEFINED OutputType = "undefined"
	WEBHOOK OutputType = "webhook"
	UDP OutputType = "udp"
	MQTT OutputType = "mqtt"
	IFTTT OutputType = "ifttt"
)