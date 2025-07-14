package tests

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testSecret     = "my-super-secret-webhook-key"
	testAllowedAge = 5 * time.Minute

	testMessageID    = "msg_2WpGqg8fH9jZ7kY6aB3dX2eF1c"
	testEventType    = "transaction.created"
	testAccountID    = "acc_1a2b3c4d5e6f7g8h9i"
	testClientID     = "cli_j0k1l2m3n4p5q6r7s8"
	testPartitionKey = "user_xyz"
)

func createTestBody(publishTime time.Time) string {
	// mimic the EventMessage structure
	eventPayload := map[string]interface{}{
		"message_id":    testMessageID,
		"event_type":    testEventType,
		"account_id":    testAccountID,
		"client_id":     testClientID,
		"name":          fmt.Sprintf("messages/%s", testMessageID),
		"partition_key": testPartitionKey,
		"publish_time":  publishTime.Format(time.RFC3339Nano),
		"data": map[string]interface{}{
			"transactionId": "txn_abcdef123456",
			"amount":        100.50,
			"currency":      "USD",
			"status":        "completed",
		},
	}

	bodyBytes, err := json.Marshal(eventPayload)

	if err != nil {
		panic(fmt.Sprintf("Failed to marshal test body: %v", err))
	}

	return string(bodyBytes)
}

func generateSignature(sendTimeStr, body, secret string) string {
	payload := fmt.Sprintf("%s.%s", body, sendTimeStr)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}

func TestEvents(t *testing.T) {
	fmt.Println("Starting test for Events.ValidatePayload")

	validSendTime := time.Now().UTC().Truncate(time.Second)
	validBody := createTestBody(validSendTime)

	testCases := []struct {
		name            string
		sendTime        time.Time
		body            string
		signature       string
		headers         map[string]string
		expectError     bool
		errorContains   string
		validateMessage func(t *testing.T, msg *components.EventMessage)
	}{
		{
			name:        "Success - Valid Payload",
			sendTime:    validSendTime,
			body:        validBody,
			headers:     map[string]string{},
			expectError: false,
			validateMessage: func(t *testing.T, msg *components.EventMessage) {
				if msg == nil {
					t.Fatal("expected EventMessage to be non-nil")
				}
				if msg.MessageID == nil || *msg.MessageID != testMessageID {
					t.Errorf("expected message ID %q, got %v", testMessageID, msg.MessageID)
				}
				if msg.EventType == nil || *msg.EventType != testEventType {
					t.Errorf("expected event type %q, got %v", testEventType, msg.EventType)
				}
				if msg.AccountID == nil || *msg.AccountID != testAccountID {
					t.Errorf("expected account ID %q, got %v", testAccountID, msg.AccountID)
				}
				if msg.ClientID == nil || *msg.ClientID != testClientID {
					t.Errorf("expected client ID %q, got %v", testClientID, msg.ClientID)
				}
				expectedName := fmt.Sprintf("messages/%s", testMessageID)
				if msg.Name == nil || *msg.Name != expectedName {
					t.Errorf("expected name %q, got %v", expectedName, msg.Name)
				}
				if msg.PublishTime == nil || !msg.PublishTime.Equal(validSendTime) {
					t.Errorf("expected publish time %v, got %v", validSendTime, msg.PublishTime)
				}
				if msg.Data == nil {
					t.Fatal("expected data payload to be non-nil")
				}
				if status, ok := msg.Data["status"].(string); !ok || status != "completed" {
					t.Errorf("expected data.status to be 'completed', got %v", msg.Data["status"])
				}
			},
		},
		{
			name:          "Error - Invalid Signature",
			sendTime:      validSendTime,
			body:          validBody,
			signature:     "this-is-not-the-correct-signature",
			headers:       map[string]string{},
			expectError:   true,
			errorContains: "signature validation failed",
		},
		{
			name:          "Error - Expired Event",
			sendTime:      time.Now().UTC().Add(-(testAllowedAge + time.Minute)),
			body:          validBody,
			headers:       map[string]string{},
			expectError:   true,
			errorContains: "event has expired",
		},
		{
			name:     "Error - Missing Signature Header",
			sendTime: time.Now().UTC(),
			body:     validBody,
			headers: map[string]string{
				"x-apex-event-signature": "", // Remove the signature header
			},
			expectError:   true,
			errorContains: "missing required headers",
		},
		{
			name:     "Error - Missing Send Time Header",
			sendTime: time.Now().UTC(),
			body:     validBody,
			headers: map[string]string{
				"x-apex-event-send-time": "", // Remove the time header
			},
			expectError:   true,
			errorContains: "missing required headers",
		},
		{
			name:     "Error - Invalid Send Time Format",
			sendTime: time.Now().UTC(),
			body:     validBody,
			headers: map[string]string{
				"x-apex-event-send-time": "not-a-valid-time",
			},
			expectError:   true,
			errorContains: "invalid send time format",
		},
		{
			name:          "Error - Invalid JSON Body",
			sendTime:      time.Now().UTC(),
			body:          `{"id":"evt_123", "type": "user.created" "data":{}}`, // Malformed JSON
			headers:       map[string]string{},
			expectError:   true,
			errorContains: "failed to unmarshal event message",
		},
	}

	for _, tc := range testCases {
		fmt.Printf("Running test case: %s\n", tc.name)

		t.Run(tc.name, func(t *testing.T) {

			// generate the request to the webhook endpoint
			req := httptest.NewRequest("POST", "/webhook", strings.NewReader(tc.body))

			// set headers of the request
			sendTimeStr := tc.sendTime.Format(time.RFC3339)
			signature := tc.signature

			if signature == "" {
				signature = generateSignature(sendTimeStr, tc.body, testSecret)
			}
			req.Header.Set("x-apex-event-signature", signature)
			req.Header.Set("x-apex-event-send-time", sendTimeStr)

			for key, val := range tc.headers {
				if val == "" {
					// header was intended to be removed for test
					req.Header.Del(key)
				} else {
					req.Header.Set(key, val)
				}
			}

			// run validation using the test target
			eventMessage, err := ascendsdkgo.ValidateEventPayload(req, testSecret, testAllowedAge)

			// assert results
			if tc.expectError {
				if err == nil {
					t.Fatal("expected an error, but got nil")
				}
				if !strings.Contains(err.Error(), tc.errorContains) {
					t.Errorf("expected error to contain %q, but got: %v", tc.errorContains, err)
				}
			} else {
				if err != nil {
					t.Fatalf("did not expect an error, but got: %v", err)
				}
				if tc.validateMessage != nil {
					tc.validateMessage(t, eventMessage)
				}
			}
		})
	}
}

func TestIntegration_GetEventMessageAndValidate(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	events, err := sdk.Reader.ListEventMessages(ctx, nil, nil, nil)
	assert.Equal(t, 200, events.HTTPMeta.Response.StatusCode)

	eventMessageId := events.ListEventMessagesResponse.EventMessages[0].MessageID

	if eventMessageId == nil {
		t.Fatalf("expected a valid messageID, but got nil")
	}

	event, err := sdk.Reader.GetEventMessage(ctx, *eventMessageId)
	assert.Equal(t, 200, event.HTTPMeta.Response.StatusCode)

	// simulate webhook request
	simulatedBodyBytes, err := json.Marshal(event.GetEventMessage())
	require.NoError(t, err)
	simulatedBody := string(simulatedBodyBytes)

	fmt.Printf("Simulated body: %s\n", simulatedBody)

	secretKey := "super-secret-key"
	sendTime := time.Now().UTC()
	sendTimeStr := sendTime.Format(time.RFC3339)

	signature := generateSignature(sendTimeStr, simulatedBody, secretKey)

	simulatedRequest := httptest.NewRequest(http.MethodPost, "/webhook", strings.NewReader(simulatedBody))
	simulatedRequest.Header.Set("x-apex-event-signature", signature)
	simulatedRequest.Header.Set("x-apex-event-send-time", sendTimeStr)
	simulatedRequest.Header.Set("Content-Type", "application/json")

	// pass to validation function
	validatedEvent, err := ascendsdkgo.ValidateEventPayload(simulatedRequest, secretKey, 5*time.Minute)
	require.NoError(t, err)
	assert.NotNil(t, validatedEvent, "expected validated event to be non-nil")
}
