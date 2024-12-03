// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// Token - Represents a JSON Web Token for access to the system; See: https://datatracker.ietf.org/doc/html/rfc6749#section-5.1
type Token struct {
	// The access token issued by the authorization server
	AccessToken *string `json:"access_token,omitempty"`
	// The lifetime in seconds of the access token
	ExpiresIn *int64 `json:"expires_in,omitempty"`
	// The token type for this access token
	TokenType *string `json:"token_type,omitempty"`
}

func (o *Token) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}

func (o *Token) GetExpiresIn() *int64 {
	if o == nil {
		return nil
	}
	return o.ExpiresIn
}

func (o *Token) GetTokenType() *string {
	if o == nil {
		return nil
	}
	return o.TokenType
}
