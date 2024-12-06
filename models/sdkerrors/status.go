// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"

	"github.com/afs-public/ascend-sdk-go/models/components"
)

// Status - The status message serves as the general-purpose service error message. Each status message includes a gRPC error code, error message, and error details.
type Status struct {
	// The code field contains an enum value of google.rpc.Code.
	Code *int `json:"code,omitempty"`
	// The details field contains one or more technical error details.
	Details []components.GoogleProtobufAny `json:"details,omitempty"`
	// The message field contains human-friendly messages about the error.
	Message *string `json:"message,omitempty"`
}

var _ error = &Status{}

func (e *Status) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
