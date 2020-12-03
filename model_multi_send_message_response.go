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
// MultiSendMessageResponse Broadcast message result. The errors array contains the list of errors ocurred when sending a message.
type MultiSendMessageResponse struct {
	Errors []MessageSendResult `json:"errors,omitempty"`
	Sent int32 `json:"sent,omitempty"`
	Failed int32 `json:"failed,omitempty"`
}
