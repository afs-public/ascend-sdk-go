// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// TransferCreate - An account transfer which contains the receiving and delivering party information, assets being transferred, NSCC status information, etc.
type TransferCreate struct {
	// The assets being transferred (Cash, Equities, etc.)
	Assets []AssetCreate `json:"assets,omitempty"`
	// User supplied comment
	Comment *string `json:"comment,omitempty"`
	// The delivering/receiving party information
	Deliverer TransferAccountCreate `json:"deliverer"`
	// An associated NSCC transfer identifier, if applicable
	OriginalControlNumber *string `json:"original_control_number,omitempty"`
}

func (o *TransferCreate) GetAssets() []AssetCreate {
	if o == nil {
		return nil
	}
	return o.Assets
}

func (o *TransferCreate) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *TransferCreate) GetDeliverer() TransferAccountCreate {
	if o == nil {
		return TransferAccountCreate{}
	}
	return o.Deliverer
}

func (o *TransferCreate) GetOriginalControlNumber() *string {
	if o == nil {
		return nil
	}
	return o.OriginalControlNumber
}
