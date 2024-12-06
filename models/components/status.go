// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// Status - The status message serves as the general-purpose service error message. Each status message includes a gRPC error code, error message, and error details.
type Status struct {
	// The code field contains an enum value of google.rpc.Code.
	Code *int `json:"code,omitempty"`
	// The details field contains one or more technical error details.
	Details []GoogleProtobufAny `json:"details,omitempty"`
	// The message field contains human-friendly messages about the error.
	Message *string `json:"message,omitempty"`
}

func (o *Status) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *Status) GetDetails() []GoogleProtobufAny {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *Status) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}