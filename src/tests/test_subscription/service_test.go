package test_subscription

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tg "gopkg.in/telebot.v3"
	"my-vpn-shop/subscription"
	"testing"
)

//func TestService_GetActualPrice(t *testing.T) {
//	t.Run("zero keys in vpn servers", func(t *testing.T) {
//		const (
//			keysCount     int     = 0
//			totalVPNPrice float64 = 300
//		)
//		_, err := subscription.GetActualPrice(keysCount, totalVPNPrice)
//		require.ErrorIs(t, err, subscription.ErrZeroKeysInServer)
//	})
//	t.Run("zero total price", func(t *testing.T) {
//		const (
//			keysCount     int     = 18
//			totalVPNPrice float64 = 0
//		)
//		expected := tg.Price{
//			Label:  "Актуальная цена за этот месяц",
//			Amount: subscription.MinAmount * 100,
//		}
//		actual, err := subscription.GetActualPrice(keysCount, totalVPNPrice)
//		require.NoError(t, err)
//		assert.Equal(t, expected, actual)
//	})
//	t.Run("negative total price", func(t *testing.T) {
//		const (
//			keysCount     int     = 18
//			totalVPNPrice float64 = -543.01
//		)
//		_, err := subscription.GetActualPrice(keysCount, totalVPNPrice)
//		require.ErrorIs(t, err, subscription.ErrNegativeTotalPrice)
//	})
//	t.Run("get actual price above minimal valid value", func(t *testing.T) {
//		const (
//			keysCount     int     = 3
//			totalVPNPrice float64 = 1000
//		)
//		expected := tg.Price{
//			Label:  "Актуальная цена за этот месяц",
//			Amount: 33333,
//		}
//		actual, err := subscription.GetActualPrice(keysCount, totalVPNPrice)
//		require.NoError(t, err)
//		assert.Equal(t, expected, actual)
//	})
//}

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
}
