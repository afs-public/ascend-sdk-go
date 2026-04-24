package hooks

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/afs-public/ascend-sdk-go/internal/config"
	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/golang-jwt/jwt/v4"
)

type AccessTokenHook struct{}

var (
	_                     beforeRequestHook = (*AccessTokenHook)(nil)
	accessTokenMu         sync.Mutex
	accessToken           string
	accessTokenExpiration time.Time
)

func (i *AccessTokenHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	// modify the request object before it is sent, such as adding headers or query parameters, or return an error to stop the request from being sent
	// convert secSource to components.Security
	sec, err := hookCtx.SecuritySource(hookCtx.Context)
	if err != nil {
		return nil, err
	}
	customSec, ok := sec.(components.Security)
	if !ok {
		return nil, errors.New("security source did not return components.Security")
	}

	if customSec.APIKey == nil {
		return nil, errors.New("apikey is empty")
	}
	req.Header.Set("x-api-key", *customSec.APIKey)

	if customSec.ServiceAccountCreds == nil {
		return nil, errors.New("service account creds is empty")
	}
	serverURL := req.URL.Scheme + "://" + req.URL.Host
	accessToken, err := getAccessToken(hookCtx.Context, serverURL, *customSec.APIKey, *customSec.ServiceAccountCreds, hookCtx.SDKConfiguration.Client)
	if err != nil {
		return nil, fmt.Errorf("error getting access token: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	return req, nil
}

func getAccessToken(ctx context.Context, serverUrl string, apiKey string, creds components.ServiceAccountCreds, client config.HTTPClient) (string, error) {
	accessTokenMu.Lock()
	defer accessTokenMu.Unlock()

	if accessTokenExpired() {
		jws, err := getJws(creds)
		if err != nil {
			return "", err
		}

		accessToken, err = generateJwt(ctx, serverUrl, apiKey, jws, client)
		if err != nil {
			return "", err
		}

		return accessToken, nil
	}
	return accessToken, nil
}

func generateJwt(ctx context.Context, serverUrl string, apiKey string, jws string, client config.HTTPClient) (string, error) {
	opURL, err := url.JoinPath(serverUrl, "/iam/v1/serviceAccounts:generateAccessToken")

	if err != nil {
		return "", err
	}

	data := map[string]string{
		"jws": jws,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", errors.New("failed to marshal json data")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", errors.New("failed to create http request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("failed to send http request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Error generating service account token [url: %s, status: %d, error_text: %s]", opURL, resp.StatusCode, string(body))
	}

	// parse response body to get access token
	// return access token
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("failed to read response body")
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", errors.New("failed to unmarshal response body")
	}

	token, ok := result["access_token"].(string)
	if !ok {
		return "", errors.New("response missing or invalid 'access_token' field")
	}
	expiresIn, ok := result["expires_in"].(float64)
	if !ok {
		return "", errors.New("response missing or invalid 'expires_in' field")
	}

	accessToken = token
	// Add 1 hour safety buffer to refresh tokens before they actually expire
	accessTokenExpiration = time.Now().Add(time.Duration(math.Max(expiresIn-3600, 60)) * time.Second)
	return accessToken, nil

}

func getJws(creds components.ServiceAccountCreds) (string, error) {
	privateKeyContent := creds.PrivateKey
	privateKeyContent = strings.ReplaceAll(privateKeyContent, "\n", "")
	privateKeyContent = strings.ReplaceAll(privateKeyContent, "\\n", "")
	privateKeyContent = strings.ReplaceAll(privateKeyContent, "\r", "")
	privateKeyContent = strings.ReplaceAll(privateKeyContent, "\\r", "")
	privateKeyContent = strings.ReplaceAll(privateKeyContent, "-----BEGIN PRIVATE KEY-----", "")
	privateKeyContent = strings.ReplaceAll(privateKeyContent, "-----END PRIVATE KEY-----", "")

	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyContent)
	if err != nil {
		return "", fmt.Errorf("failed to decode private key: %w", err)
	}

	// Parse the private key
	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBytes)
	if err != nil {
		return "", err
	}

	// Prepare claims for the JWT
	nowIsoDateTime := time.Now().UTC().Format(time.RFC3339)
	claims := jwt.MapClaims{
		"iss":          "issuer",
		"sub":          "subject",
		"name":         creds.Name,
		"organization": creds.Organization,
		"datetime":     nowIsoDateTime,
		"exp":          time.Now().Add(5 * time.Minute).Unix(),
	}

	// Create a JWS and sign it with the private key using RS256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func accessTokenExpired() bool {
	return time.Now().After(accessTokenExpiration)
}
