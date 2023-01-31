package test_subscription

import (
	"github.com/stretchr/testify/assert"
	"my-vpn-shop/subscription"
	"testing"
)

func TestService_GetActualPrice(t *testing.T) {
	t.Run("get actual price", func(t *testing.T) {
		const (
			keysCount     = 6
			totalVPNPrice = 300
			expected      = 50.0
		)
		assert.Equal(t, expected, subscription.GetActualPrice(keysCount, totalVPNPrice))
	})
}
