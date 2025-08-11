// Code manually maintained by the SDK team.

package ascendsdkgo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/afs-public/ascend-sdk-go/models/components"
)

// ValidateEventPayload validates an incoming webhook request.
// It checks the signature to ensure the request is authentic, verifies the timestamp to prevent replay attacks,
// and unmarshals the payload into an EventMessage.
func ValidateEventPayload(r *http.Request, webhookSecret string, allowedEventAge time.Duration) (*components.EventMessage, error) {
	signatureHeader := r.Header.Get("x-apex-event-signature")
	sendTimeHeader := r.Header.Get("x-apex-event-send-time")
	fmt.Printf("Signature Header: %s\n", signatureHeader)

	if signatureHeader == "" || sendTimeHeader == "" {
		return nil, fmt.Errorf("missing required headers: x-apex-event-signature and/or x-apex-event-send-time")
	}

	var requestBody []byte
	if r.Body != nil {
		defer r.Body.Close()
		var err error
		requestBody, err = io.ReadAll(r.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read request body: %w", err)
		}
	}

	bodyStr := strings.TrimSpace(string(requestBody))

	sendTime, err := time.Parse(time.RFC3339, sendTimeHeader)
	if err != nil {
		return nil, fmt.Errorf("invalid send time format: %w", err)
	}

	if eventOutOfRange(sendTime, allowedEventAge) {
		return nil, fmt.Errorf("event is out of range, it must be sent within the allowed event age")
	}

	err = verifySignature(signatureHeader, sendTimeHeader, bodyStr, webhookSecret)
	if err != nil {
		return nil, fmt.Errorf("signature validation failed: %w", err)
	}

	var eventMessage components.EventMessage
	if err := json.Unmarshal([]byte(bodyStr), &eventMessage); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event message: %w", err)
	}

	return &eventMessage, nil
}

func eventOutOfRange(sendTime time.Time, allowedEventAge time.Duration) bool {
	currentTime := time.Now().UTC()
	// event is too old and was sent before the allowed event age
	isTooOld := sendTime.Before(currentTime.Add(-allowedEventAge))
	// the event is in the future, indicating a potential issue with the send time and should not be processed
	isInTheFuture := sendTime.After(currentTime.Add(allowedEventAge))

	return isTooOld || isInTheFuture
}

func verifySignature(headerSignature, sendTimeHeader, body, secret string) error {
	payload := fmt.Sprintf("%s.%s", body, sendTimeHeader)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(expectedSignature), []byte(headerSignature)) {
		return fmt.Errorf("provided signature does not match calculated signature")
	}

	return nil
}
