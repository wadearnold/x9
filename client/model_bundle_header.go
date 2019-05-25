/*
 * X9 API
 *
 * Moov X9 () implements an HTTP API for creating, parsing and validating X9 (Check21) files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type BundleHeader struct {
	// BundleHeader ID
	Id string `json:"id,omitempty"`
	// A code that identifies the type of bundle. It is the same value as the CollectionTypeIndicator in the CashLetterHeader within which the bundle is contained, unless the CollectionTypeIndicator in the CashLetterHeader is 99.  * `00` - Preliminary Forward Information * `01` - Forward Presentment * `02` - Forward Presentment - Same-Day Settlement * `03` - Return * `04` - Return Notification * `05` - Preliminary Return Notification * `06` - Final Return Notification 
	CollectionTypeIndicator string `json:"collectionTypeIndicator,omitempty"`
	// DestinationRoutingNumber contains the routing and transit number of the institution that receives and processes the cash letter or the bundle.
	DestinationRoutingNumber string `json:"destinationRoutingNumber,omitempty"`
	// ECEInstitutionRoutingNumber contains the routing and transit number of the institution that that creates the bundle header.
	ECEInstitutionRoutingNumber string `json:"eCEInstitutionRoutingNumber,omitempty"`
	// BundleBusinessDate is the business date of the bundle. Format - YYYYMMDD, where - YYYY year, MM month, DD day.
	BundleBusinessDate string `json:"bundleBusinessDate,omitempty"`
	// BundleCreationDate is the date that the bundle is created. Format - YYYYMMDD, where - YYYY year, MM month, DD day.
	BundleCreationDate string `json:"bundleCreationDate,omitempty"`
	// BundleID is number that identifies the bundle, assigned by the institution that creates the bundle.
	BundleID string `json:"bundleID,omitempty"`
	// BundleSequenceNumber is a number assigned by the institution that creates the bundle. Usually denotes the relative position of the bundle within the cash letter.
	BundleSequenceNumber string `json:"bundleSequenceNumber,omitempty"`
	// CycleNumber is a code assigned by the institution that creates the bundle.  Denotes the cycle under which the bundle is created.
	CycleNumber string `json:"cycleNumber,omitempty"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField,omitempty"`
}
