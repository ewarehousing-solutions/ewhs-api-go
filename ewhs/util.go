package ewhs

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"
)

// VerifyWebhookRequest Verifies a webhook http request sent by eWarehousing.
// The body of the request is still readable after invoking the method.
func VerifyWebhookRequest(httpRequest *http.Request, secret string) bool {
	shopifySha256 := httpRequest.Header.Get("X-Hmac-Sha256")
	actualMac := []byte(shopifySha256)

	mac := hmac.New(sha256.New, []byte(secret))
	requestBody, _ := io.ReadAll(httpRequest.Body)
	httpRequest.Body = io.NopCloser(bytes.NewBuffer(requestBody))
	mac.Write(requestBody)
	macSum := mac.Sum(nil)
	expectedMac := []byte(base64.StdEncoding.EncodeToString(macSum))

	return hmac.Equal(actualMac, expectedMac)
}
