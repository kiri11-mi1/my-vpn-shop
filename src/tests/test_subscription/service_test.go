package test_subscription

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tg "gopkg.in/telebot.v3"
	"my-vpn-shop/subscription"
	"testing"
)

func TestService_GetInvoice(t *testing.T) {
	t.Run("get invoice", func(t *testing.T) {
		var (
			providerToken = "test token"
			totalVPNPrice = 600.0
			keysCount     = 6
			image         = tg.File{FileURL: subscription.InvoiceImage}
			price         = tg.Price{Label: subscription.Label, Amount: 100 * 100}
			expected      = tg.Invoice{
				Title:       subscription.InvoiceTitle,
				Description: subscription.InvoiceDescription,
				Payload:     subscription.InvoicePayload,
				Currency:    subscription.InvoiceCurrency,
				Token:       providerToken,
				Prices:      []tg.Price{price},
				Photo:       &tg.Photo{File: image},
			}
		)
		actual, err := subscription.GetInvoice(keysCount, totalVPNPrice, providerToken)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("get invoice with minimal price value", func(t *testing.T) {
		var (
			providerToken = "test token"
			totalVPNPrice = 100.0
			keysCount     = 6
			image         = tg.File{FileURL: subscription.InvoiceImage}
			price         = tg.Price{Label: subscription.Label, Amount: subscription.MinAmount * 100}
			expected      = tg.Invoice{
				Title:       subscription.InvoiceTitle,
				Description: subscription.InvoiceDescription,
				Payload:     subscription.InvoicePayload,
				Currency:    subscription.InvoiceCurrency,
				Token:       providerToken,
				Prices:      []tg.Price{price},
				Photo:       &tg.Photo{File: image},
			}
		)
		actual, err := subscription.GetInvoice(keysCount, totalVPNPrice, providerToken)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("zero count keys", func(t *testing.T) {
		var (
			providerToken = "test token"
			totalVPNPrice = 100.0
			keysCount     = 0
		)
		actual, err := subscription.GetInvoice(keysCount, totalVPNPrice, providerToken)
		require.ErrorIs(t, err, subscription.ErrZeroKeysInServer)
		assert.Empty(t, actual)
	})
	t.Run("negative total price", func(t *testing.T) {
		var (
			providerToken = "test token"
			keysCount     = 18
			totalVPNPrice = -543.01
		)
		actual, err := subscription.GetInvoice(keysCount, totalVPNPrice, providerToken)
		require.ErrorIs(t, err, subscription.ErrNegativeTotalPrice)
		assert.Empty(t, actual)
	})
	t.Run("zero total price", func(t *testing.T) {
		var (
			providerToken = "test token"
			keysCount     = 18
			totalVPNPrice = 0.0
			image         = tg.File{FileURL: subscription.InvoiceImage}
			price         = tg.Price{Label: subscription.Label, Amount: subscription.MinAmount * 100}
			expected      = tg.Invoice{
				Title:       subscription.InvoiceTitle,
				Description: subscription.InvoiceDescription,
				Payload:     subscription.InvoicePayload,
				Currency:    subscription.InvoiceCurrency,
				Token:       providerToken,
				Prices:      []tg.Price{price},
				Photo:       &tg.Photo{File: image},
			}
		)
		actual, err := subscription.GetInvoice(keysCount, totalVPNPrice, providerToken)
		require.NoError(t, err, subscription.ErrNegativeTotalPrice)
		assert.Equal(t, expected, actual)
	})
}

func TestService_GetName(t *testing.T) {
	t.Run("get name sub", func(t *testing.T) {
		var (
			username       = "test"
			id       int64 = 245
			expected       = "sub_test_245"
		)
		actual := subscription.GetName(username, id)
		assert.Equal(t, expected, actual)
	})
	t.Run("get name with empty username", func(t *testing.T) {
		var (
			username       = ""
			id       int64 = 245
			expected       = "sub__245"
		)
		actual := subscription.GetName(username, id)
		assert.Equal(t, expected, actual)
	})
}
