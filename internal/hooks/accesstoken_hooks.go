package hooks

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/dgrijalva/jwt-go"
)

type AccessTokenHook struct{}

var (
	_                     beforeRequestHook = (*AccessTokenHook)(nil)
	accessToken           string
	accessTokenExpiration time.Time
)

func (i *AccessTokenHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	// modify the request object before it is sent, such as adding headers or query parameters, or return an error to stop the request from being sent
	secSource := hookCtx.SecuritySource
	// converty secSource to components.Security
	sec, err := secSource(hookCtx.Context)
	if err != nil {
		return nil, err
	}
	customSec := sec.(components.Security)

	if customSec.APIKey == nil {
		return nil, errors.New("apikey is empty")
	} else {
		req.Header.Set("x-api-key", *customSec.APIKey)
	}

	if customSec.ServiceAccountCreds == nil {
		return nil, errors.New("service account creds is empty")
	} else {

		url := req.URL.Scheme + "://" + req.URL.Host
		accessToken, err := getAccessToken(url, *customSec.APIKey, *customSec.ServiceAccountCreds)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error getting access token: %s", err))
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}

	return req, nil
}

func getAccessToken(serverUrl string, apiKey string, creds components.ServiceAccountCreds) (string, error) {
	if accessTokenExpired() {
		// get new access token
		jws, err := getJws(creds)
		if err != nil {
			return "", err
		}

		accessToken, err = generateJwt(serverUrl, apiKey, jws)
		if err != nil {
			return "", err
		}

		return accessToken, nil

	}
	return accessToken, nil

}

func generateJwt(serverUrl string, apiKey string, jws string) (string, error) {
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

	req, err := http.NewRequest("POST", opURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", errors.New("failed to create http request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("failed to send http request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("failed to get access token")
	} else {

		// parse response body to get access token
		// return access token
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", errors.New("failed to read response body")
		}

		var result map[string]interface{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return "", errors.New("failed to unmarshal response body")
		}

		token := result["access_token"].(string)
		expiresIn := result["expires_in"].(float64)

		accessToken = token
		accessTokenExpiration = time.Now().Add(time.Duration(expiresIn) * time.Second)
		return accessToken, nil
	}

}

func getJws(creds components.ServiceAccountCreds) (string, error) {
	privateKeyContent := creds.PrivateKey
	privateKeyContent = strings.ReplaceAll(privateKeyContent, "\\n", "")
	privateKeyContent = strings.ReplaceAll(privateKeyContent, "\r", "")
	privateKeyContent = strings.ReplaceAll(privateKeyContent, "-----BEGIN PRIVATE KEY-----", "")
	privateKeyContent = strings.ReplaceAll(privateKeyContent, "-----END PRIVATE KEY-----", "")

	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyContent)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to decode private key %v", err))
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
