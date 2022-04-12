package ots_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/ots"
)

var (
	userName = "userName"
	apiKey   = "apiKey"
)

func TestClient_Status(t *testing.T) {
	ctrl := gomock.NewController(t)

	httpClientMock := NewMockHTTPClient(ctrl)

	t.Run("Ok", func(t *testing.T) {
		ctx := context.Background()
		expected := ots.Nominal

		res := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"status": "nominal"}`)),
		}

		httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)

		actual, err := ots.New(userName, apiKey, ots.WithHTTPClient(httpClientMock)).
			Status(ctx)

		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		ctx := context.Background()

		res := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       io.NopCloser(strings.NewReader(`{}`)),
		}

		httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)

		actual, err := ots.New(userName, apiKey, ots.WithHTTPClient(httpClientMock)).
			Status(ctx)

		assert.Empty(t, actual)
		assert.ErrorContains(t, err, "api error")
	})
}

func TestClient_Share(t *testing.T) {
	ctrl := gomock.NewController(t)

	httpClientMock := NewMockHTTPClient(ctrl)

	secret := "secret"

	t.Run("Ok", func(t *testing.T) {
		ctx := context.Background()

		expected := ots.Metadata{
			CustID:      "custID",
			MetadataKey: "metadataKey",
			SecretKey:   "secretKey",
			TTL:         ots.Duration(time.Hour),
			MetadataTTL: ots.Duration(2 * time.Hour),
			SecretTTL:   ots.Duration(3 * time.Hour),
			Recipient: []string{
				"foo@bar.baz",
			},
			PassphraseRequired: true,
			Created:            ots.Time(time.Now().Add(-2 * time.Hour).Round(time.Second)),
			Updated:            ots.Time(time.Now().Add(-1 * time.Hour).Round(time.Second)),
		}

		body, err := json.Marshal(&expected)
		assert.NoError(t, err)

		res := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(body)),
		}

		httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)

		actual, err := ots.New(userName, apiKey, ots.WithHTTPClient(httpClientMock)).
			Share(ctx, secret)

		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		ctx := context.Background()

		res := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       io.NopCloser(strings.NewReader(`{}`)),
		}

		httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)

		actual, err := ots.New(userName, apiKey, ots.WithHTTPClient(httpClientMock)).
			Share(ctx, secret)

		assert.Empty(t, actual)
		assert.ErrorContains(t, err, "api error")
	})
}

func TestClient_Generate(t *testing.T) {
	ctrl := gomock.NewController(t)

	httpClientMock := NewMockHTTPClient(ctrl)

	t.Run("Ok", func(t *testing.T) {
		ctx := context.Background()

		value := "value"

		expected := ots.Metadata{
			CustID:      "custID",
			Value:       &value,
			MetadataKey: "metadataKey",
			SecretKey:   "secretKey",
			TTL:         ots.Duration(time.Hour),
			MetadataTTL: ots.Duration(2 * time.Hour),
			SecretTTL:   ots.Duration(3 * time.Hour),
			Recipient: []string{
				"foo@bar.baz",
			},
			PassphraseRequired: true,
			Created:            ots.Time(time.Now().Add(-2 * time.Hour).Round(time.Second)),
			Updated:            ots.Time(time.Now().Add(-1 * time.Hour).Round(time.Second)),
		}

		body, err := json.Marshal(&expected)
		assert.NoError(t, err)

		res := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(body)),
		}

		httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)

		actual, err := ots.New(userName, apiKey, ots.WithHTTPClient(httpClientMock)).
			Generate(ctx)

		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		ctx := context.Background()

		res := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       io.NopCloser(strings.NewReader(`{}`)),
		}

		httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)

		actual, err := ots.New(userName, apiKey, ots.WithHTTPClient(httpClientMock)).
			Generate(ctx)

		assert.Empty(t, actual)
		assert.ErrorContains(t, err, "api error")
	})
}

func TestClient_Secret(t *testing.T) {
	ctrl := gomock.NewController(t)

	httpClientMock := NewMockHTTPClient(ctrl)

	secretKey := "secretKey"

	t.Run("Ok", func(t *testing.T) {
		ctx := context.Background()

		expected := ots.Secret{
			SecretKey: secretKey,
			Value:     "value",
		}

		body, err := json.Marshal(&expected)
		assert.NoError(t, err)

		res := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(body)),
		}

		httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)

		actual, err := ots.New(userName, apiKey, ots.WithHTTPClient(httpClientMock)).
			Secret(ctx, secretKey)

		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		ctx := context.Background()

		res := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       io.NopCloser(strings.NewReader(`{}`)),
		}

		httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)

		actual, err := ots.New(userName, apiKey, ots.WithHTTPClient(httpClientMock)).
			Secret(ctx, secretKey)

		assert.Empty(t, actual)
		assert.ErrorContains(t, err, "api error")
	})
}

func TestClient_Metadata(t *testing.T) {
	ctrl := gomock.NewController(t)

	httpClientMock := NewMockHTTPClient(ctrl)

	metadataKey := "metadataKey"

	t.Run("Ok", func(t *testing.T) {
		ctx := context.Background()

		value := "value"

		expected := ots.Metadata{
			CustID:      "custID",
			Value:       &value,
			MetadataKey: "metadataKey",
			SecretKey:   "secretKey",
			TTL:         ots.Duration(time.Hour),
			MetadataTTL: ots.Duration(2 * time.Hour),
			SecretTTL:   ots.Duration(3 * time.Hour),
			Recipient: []string{
				"foo@bar.baz",
			},
			PassphraseRequired: true,
			Created:            ots.Time(time.Now().Add(-2 * time.Hour).Round(time.Second)),
			Updated:            ots.Time(time.Now().Add(-1 * time.Hour).Round(time.Second)),
		}

		body, err := json.Marshal(&expected)
		assert.NoError(t, err)

		res := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(body)),
		}

		httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)

		actual, err := ots.New(userName, apiKey, ots.WithHTTPClient(httpClientMock)).
			Metadata(ctx, metadataKey)

		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		ctx := context.Background()

		res := &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       io.NopCloser(strings.NewReader(`{}`)),
		}

		httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)

		actual, err := ots.New(userName, apiKey, ots.WithHTTPClient(httpClientMock)).
			Generate(ctx)

		assert.Empty(t, actual)
		assert.ErrorContains(t, err, "api error")
	})
}

func TestClient_Burn(t *testing.T) {
	ctrl := gomock.NewController(t)

	httpClientMock := NewMockHTTPClient(ctrl)

	metadataKey := "metadataKey"

	t.Run("Ok", func(t *testing.T) {
		ctx := context.Background()

		value := "value"

		expected := ots.Metadata{
			CustID:      "custID",
			Value:       &value,
			MetadataKey: "metadataKey",
			SecretKey:   "secretKey",
			TTL:         ots.Duration(time.Hour),
			MetadataTTL: ots.Duration(2 * time.Hour),
			SecretTTL:   ots.Duration(3 * time.Hour),
			Recipient: []string{
				"foo@bar.baz",
			},
			PassphraseRequired: true,
			Created:            ots.Time(time.Now().Add(-2 * time.Hour).Round(time.Second)),
			Updated:            ots.Time(time.Now().Add(-1 * time.Hour).Round(time.Second)),
		}

		body, err := json.Marshal(&expected)
		assert.NoError(t, err)

		res := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(body)),
		}

		httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)

		err = ots.New(userName, apiKey, ots.WithHTTPClient(httpClientMock)).
			Burn(ctx, metadataKey)

		assert.NoError(t, err)
	})

	//	t.Run("Error", func(t *testing.T) {
	//		ctx := context.Background()
	//
	//		res := &http.Response{
	//			StatusCode: http.StatusInternalServerError,
	//			Body:       io.NopCloser(strings.NewReader(`{}`)),
	//		}
	//
	//		httpClientMock.EXPECT().Do(gomock.Any()).Return(res, nil)
	//
	//		actual, err := ots.New(userName, apiKey, ots.WithHTTPClient(httpClientMock)).
	//			Generate(ctx)
	//
	//		assert.Empty(t, actual)
	//		assert.ErrorContains(t, err, "api error")
	//	})
}
