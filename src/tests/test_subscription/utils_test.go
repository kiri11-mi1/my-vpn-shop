package test_subscription

import (
	"github.com/stretchr/testify/assert"
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
}

func TestUtils_GetActualPrice(t *testing.T) {

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
