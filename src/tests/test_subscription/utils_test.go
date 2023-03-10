package test_subscription

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tg "gopkg.in/telebot.v3"
	"my-vpn-shop/subscription"
	"testing"
	"time"
)

func TestUtils_IsTimeOutPay(t *testing.T) {
	t.Run("test time out", func(t *testing.T) {
		var (
			inputLastPayDate = time.Date(
				2023,
				01,
				12,
				0,
				0,
				0,
				0,
				time.Now().Location(),
			)
			inputCurrentTime = time.Date(
				2023,
				02,
				15,
				0,
				0,
				0,
				0,
				time.Now().Location(),
			)
		)
		actual := subscription.IsTimeOutPay(inputLastPayDate, inputCurrentTime)
		assert.True(t, actual)
	})
	t.Run("not time out by day", func(t *testing.T) {
		var (
			inputLastPayDate = time.Date(
				2023,
				01,
				12,
				0,
				0,
				0,
				0,
				time.Now().Location(),
			)
			inputCurrentTime = time.Date(
				2023,
				02,
				11,
				0,
				0,
				0,
				0,
				time.Now().Location(),
			)
		)
		actual := subscription.IsTimeOutPay(inputLastPayDate, inputCurrentTime)
		assert.False(t, actual)
	})
	t.Run("not time out", func(t *testing.T) {
		var (
			inputLastPayDate = time.Date(
				2023,
				01,
				12,
				0,
				0,
				0,
				0,
				time.Now().Location(),
			)
			inputCurrentTime = time.Date(
				2023,
				02,
				12,
				4,
				34,
				1,
				13,
				time.Now().Location(),
			)
		)
		actual := subscription.IsTimeOutPay(inputLastPayDate, inputCurrentTime)
		assert.False(t, actual)
	})
	t.Run("equal day, month, year - not time out", func(t *testing.T) {
		var (
			inputLastPayDate = time.Date(
				2023,
				01,
				12,
				0,
				0,
				0,
				0,
				time.Now().Location(),
			)
			inputCurrentTime = time.Date(
				2023,
				01,
				12,
				4,
				34,
				1,
				13,
				time.Now().Location(),
			)
		)
		actual := subscription.IsTimeOutPay(inputLastPayDate, inputCurrentTime)
		assert.False(t, actual)
	})
	t.Run("next year time out", func(t *testing.T) {
		var (
			inputLastPayDate = time.Date(
				2022,
				12,
				12,
				0,
				0,
				0,
				0,
				time.Now().Location(),
			)
			inputCurrentTime = time.Date(
				2023,
				01,
				13,
				5,
				45,
				2,
				0,
				time.Now().Location(),
			)
		)
		actual := subscription.IsTimeOutPay(inputLastPayDate, inputCurrentTime)
		assert.True(t, actual)
	})
}

func TestUtils_IsPayDay(t *testing.T) {
	t.Run("is a pay day", func(t *testing.T) {
		var (
			inputLastPayDate = time.Date(
				2023,
				01,
				12,
				0,
				0,
				0,
				0,
				time.Now().Location(),
			)
			inputCurrentTime = time.Date(
				2023,
				02,
				12,
				5,
				45,
				2,
				0,
				time.Now().Location(),
			)
		)
		actual := subscription.IsPayDay(inputLastPayDate, inputCurrentTime)
		assert.True(t, actual)
	})
	t.Run("date before pay day", func(t *testing.T) {
		var (
			inputLastPayDate = time.Date(
				2023,
				01,
				12,
				0,
				0,
				0,
				0,
				time.Now().Location(),
			)
			inputCurrentTime = time.Date(
				2023,
				02,
				11,
				5,
				45,
				2,
				0,
				time.Now().Location(),
			)
		)
		actual := subscription.IsPayDay(inputLastPayDate, inputCurrentTime)
		assert.False(t, actual)
	})
	t.Run("date after pay day", func(t *testing.T) {
		var (
			inputLastPayDate = time.Date(
				2023,
				01,
				12,
				0,
				0,
				0,
				0,
				time.Now().Location(),
			)
			inputCurrentTime = time.Date(
				2023,
				02,
				13,
				5,
				45,
				2,
				0,
				time.Now().Location(),
			)
		)
		actual := subscription.IsPayDay(inputLastPayDate, inputCurrentTime)
		assert.False(t, actual)
	})
	t.Run("next year pay date", func(t *testing.T) {
		var (
			inputLastPayDate = time.Date(
				2022,
				12,
				12,
				0,
				0,
				0,
				0,
				time.Now().Location(),
			)
			inputCurrentTime = time.Date(
				2023,
				01,
				12,
				5,
				45,
				2,
				0,
				time.Now().Location(),
			)
		)
		actual := subscription.IsPayDay(inputLastPayDate, inputCurrentTime)
		assert.True(t, actual)
	})
}

func TestUtils_GetActualPrice(t *testing.T) {
	t.Run("get valid price", func(t *testing.T) {
		var (
			subCount      = 3
			totalVPNPrice = 350.0
			expected      = tg.Price{Label: subscription.Label, Amount: 116.66 * 100}
		)
		actual, err := subscription.GetActualPrice(subCount, totalVPNPrice)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)

	})
	t.Run("get minimal price", func(t *testing.T) {
		var (
			subCount      = 16
			totalVPNPrice = 350.0
			expected      = tg.Price{Label: subscription.Label, Amount: subscription.MinAmount * 100}
		)
		actual, err := subscription.GetActualPrice(subCount, totalVPNPrice)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("error zero keys", func(t *testing.T) {
		var (
			subCount      = 0
			totalVPNPrice = 350.0
		)
		_, err := subscription.GetActualPrice(subCount, totalVPNPrice)
		require.ErrorIs(t, err, subscription.ErrZeroKeysInServer)

	})
	t.Run("zero total price", func(t *testing.T) {
		var (
			subCount      = 14
			totalVPNPrice = 0.0
			expected      = tg.Price{Label: subscription.Label, Amount: subscription.MinAmount * 100}
		)
		actual, err := subscription.GetActualPrice(subCount, totalVPNPrice)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("negative total price", func(t *testing.T) {
		var (
			subCount      = 14
			totalVPNPrice = -12.0
		)
		_, err := subscription.GetActualPrice(subCount, totalVPNPrice)
		require.ErrorIs(t, err, subscription.ErrNegativeTotalPrice)
	})
}

func TestService_GetInvoice(t *testing.T) {
	t.Run("get invoice", func(t *testing.T) {
		var (
			subCount      = 14
			totalVPNPrice = 350.45
			providerToken = "token test"
		)
		price, err := subscription.GetActualPrice(subCount, totalVPNPrice)
		require.NoError(t, err)
		expected := tg.Invoice{
			Title:       subscription.InvoiceTitle,
			Description: subscription.InvoiceDescription,
			Payload:     subscription.InvoicePayload,
			Currency:    subscription.InvoiceCurrency,
			Token:       providerToken,
			Prices:      []tg.Price{price},
			Photo:       &tg.Photo{File: tg.File{FileURL: subscription.InvoiceImage}},
		}
		actual, err := subscription.GetInvoice(subCount, providerToken, totalVPNPrice)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("zero keys", func(t *testing.T) {
		var (
			subCount      = 0
			totalVPNPrice = 350.45
			providerToken = "token test"
		)
		_, err := subscription.GetInvoice(subCount, providerToken, totalVPNPrice)
		require.ErrorIs(t, err, subscription.ErrZeroKeysInServer)
	})
	t.Run("negative total price", func(t *testing.T) {
		var (
			subCount      = 35
			totalVPNPrice = -350.45
			providerToken = "token test"
		)
		_, err := subscription.GetInvoice(subCount, providerToken, totalVPNPrice)
		require.ErrorIs(t, err, subscription.ErrNegativeTotalPrice)
	})
}

func TestUtils_GetName(t *testing.T) {
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
