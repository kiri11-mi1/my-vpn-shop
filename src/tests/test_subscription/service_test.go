package test_subscription

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tg "gopkg.in/telebot.v3"
	"my-vpn-shop/subscription"
	"testing"
)

func TestService_GetActualPrice(t *testing.T) {
	t.Run("get actual price", func(t *testing.T) {
		const (
			keysCount     int     = 6
			totalVPNPrice float64 = 300
		)
		expected := tg.Price{
			Label:  "Актуальная цена за этот месяц",
			Amount: 8741,
		}
		actual, err := subscription.GetActualPrice(keysCount, totalVPNPrice)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("zero keys in vpn servers", func(t *testing.T) {
		const (
			keysCount     int     = 0
			totalVPNPrice float64 = 300
		)
		_, err := subscription.GetActualPrice(keysCount, totalVPNPrice)
		require.ErrorIs(t, err, subscription.ErrZeroKeysInServer)
	})
	t.Run("zero total price", func(t *testing.T) {
		const (
			keysCount     int     = 18
			totalVPNPrice float64 = 0
		)
		expected := tg.Price{
			Label:  "Актуальная цена за этот месяц",
			Amount: 8741,
		}
		actual, err := subscription.GetActualPrice(keysCount, totalVPNPrice)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("negative total price", func(t *testing.T) {
		const (
			keysCount     int     = 18
			totalVPNPrice float64 = -543.01
		)
		_, err := subscription.GetActualPrice(keysCount, totalVPNPrice)
		require.ErrorIs(t, err, subscription.ErrNegativeTotalPrice)
	})
}
