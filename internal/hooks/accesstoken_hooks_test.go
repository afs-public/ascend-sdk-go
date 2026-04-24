package hooks

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/afs-public/ascend-sdk-go/internal/config"
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type hookUser struct {
	request         *http.Request
	hookCtx         BeforeRequestContext
	accessTokenHook *AccessTokenHook
}

func generateTestPrivateKey(t *testing.T) string {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate RSA key: %v", err)
	}
	der, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatalf("failed to marshal private key: %v", err)
	}
	return base64.StdEncoding.EncodeToString(der)
}

func newTokenServer(t *testing.T) *httptest.Server {
	t.Helper()
	tokenCounter := 0
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCounter++
		t.Logf("Test server received: %s %s", r.Method, r.URL.Path)
		t.Logf("Headers: x-api-key=%s", r.Header.Get("x-api-key"))

		if r.URL.Path != "/iam/v1/serviceAccounts:generateAccessToken" {
			t.Errorf("unexpected path: %s", r.URL.Path)
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var body map[string]string
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("failed to decode request body: %v", err)
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		t.Logf("JWS received: %s...", body["jws"][:20])

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": fmt.Sprintf("test-access-token-%d", tokenCounter),
			"expires_in":   float64(7200),
		})
	}))
}

func setUpHookUser(t *testing.T, serverURL string) hookUser {
	t.Helper()

	parsed, err := url.Parse(serverURL)
	if err != nil {
		t.Fatalf("failed to parse server URL: %v", err)
	}

	apiKey := "test-api-key"
	privateKey := generateTestPrivateKey(t)
	security := components.Security{
		APIKey: &apiKey,
		ServiceAccountCreds: &components.ServiceAccountCreds{
			PrivateKey:   privateKey,
			Name:         "Test Name",
			Organization: "Test Organization",
			Type:         "Test Type",
		},
	}

	return hookUser{
		request: &http.Request{
			Method: "GET",
			Header: http.Header{},
			URL: &url.URL{
				Scheme: parsed.Scheme,
				Host:   parsed.Host,
				Path:   "/v1/accounts",
			},
		},
		hookCtx: BeforeRequestContext{
			HookContext: HookContext{
				SecuritySource: func(ctx context.Context) (interface{}, error) {
					return security, nil
				},
				Context: context.Background(),
				SDKConfiguration: config.SDKConfiguration{
					Client: http.DefaultClient,
				},
			},
		},
		accessTokenHook: &AccessTokenHook{},
	}
}

func newSlowTokenServer(t *testing.T, delay time.Duration) *httptest.Server {
	t.Helper()
	var tokenCounter int64
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count := atomic.AddInt64(&tokenCounter, 1)
		t.Logf("[req %d] Server received request, sleeping %v to widen race window", count, delay)

		time.Sleep(delay)

		if r.URL.Path != "/iam/v1/serviceAccounts:generateAccessToken" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": fmt.Sprintf("test-access-token-%d", count),
			"expires_in":   float64(7200),
		})
	}))
}

func TestAccessTokenHook_BeforeRequest(t *testing.T) {
	// Reset global token state so each test starts fresh
	accessToken = ""
	accessTokenExpiration = time.Time{}

	server := newTokenServer(t)
	defer server.Close()

	hu := setUpHookUser(t, server.URL)

	result, err := hu.accessTokenHook.BeforeRequest(hu.hookCtx, hu.request)
	if err != nil {
		t.Fatalf("BeforeRequest failed: %v", err)
	}

	if got := result.Header.Get("x-api-key"); got != "test-api-key" {
		t.Errorf("x-api-key header = %q, want %q", got, "test-api-key")
	}
	if got := result.Header.Get("Authorization"); got != fmt.Sprintf("Bearer %s", "test-access-token-1") {
		t.Errorf("Authorization header = %q, want %q", got, "Bearer test-access-token-1")
	}

	t.Logf("Request after BeforeRequest:")
	t.Logf("  x-api-key: %s", result.Header.Get("x-api-key"))
	t.Logf("  Authorization: %s", result.Header.Get("Authorization"))
}

func TestAccessTokenHook_BeforeRequest_RaceCondition(t *testing.T) {
	accessToken = ""
	accessTokenExpiration = time.Time{}

	server := newSlowTokenServer(t, 100*time.Millisecond)
	defer server.Close()

	goroutines := 2
	var barrier sync.WaitGroup
	barrier.Add(goroutines)

	type result struct {
		req *http.Request
		err error
	}
	results := make([]result, goroutines)

	var done sync.WaitGroup
	done.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			defer done.Done()

			hu := setUpHookUser(t, server.URL)

			// Both goroutines wait here until all are ready,
			// then release simultaneously so they both see
			// accessTokenExpired() == true before either finishes.
			barrier.Done()
			barrier.Wait()

			req, err := hu.accessTokenHook.BeforeRequest(hu.hookCtx, hu.request)
			results[idx] = result{req: req, err: err}
		}(i)
	}

	done.Wait()

	for i, r := range results {
		if r.err != nil {
			t.Errorf("goroutine %d: BeforeRequest failed: %v", i, r.err)
			continue
		}
		token := r.req.Header.Get("Authorization")
		t.Logf("goroutine %d: Authorization = %s", i, token)
	}

	// Both goroutines should get the same token if there's proper synchronization.
	// With the current code they may get different tokens (test-access-token-1 vs
	// test-access-token-2), proving the race condition exists.
	token0 := results[0].req.Header.Get("Authorization")
	token1 := results[1].req.Header.Get("Authorization")
	if token0 != token1 {
		t.Errorf("expected both goroutines to get the same token, got %q and %q", token0, token1)
	}
}
