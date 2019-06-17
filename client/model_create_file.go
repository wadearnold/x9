/*
 * ImageCashLetter API
 *
 * Moov Image Cash Letter (ICL) implements an HTTP API for creating, parsing and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CreateFile struct {
	// File ID
	Id          string       `json:"id,omitempty"`
	FileHeader  FileHeader   `json:"fileHeader"`
	CashLetters []CashLetter `json:"cashLetters,omitempty"`
	Bundles     []Bundle     `json:"bundles,omitempty"`
	FileControl FileControl  `json:"fileControl,omitempty"`
}
