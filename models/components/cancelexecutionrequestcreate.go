// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CancelExecutionRequestCreate - A request for canceling a single execution.
type CancelExecutionRequestCreate struct {
	// The name of the execution to cancel.
	Name string `json:"name"`
}

func (o *CancelExecutionRequestCreate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
