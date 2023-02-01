package test_outline

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"my-vpn-shop/outline"
	"os"
	"testing"
)

func TestApi_GetKeys(t *testing.T) {
	t.Run("get keys from outline api", func(t *testing.T) {
		var (
			client = outline.NewOutlineClient(os.Getenv("VPN_URL_API"))
		)
		keys, err := client.GetKeys()
		require.NoError(t, err)
		assert.NotEmpty(t, keys)
	})

	t.Run("not valid url", func(t *testing.T) {
		var (
			client = outline.NewOutlineClient("https://example.com")
		)
		keys, err := client.GetKeys()
		require.Error(t, err)
		assert.Empty(t, keys)
	})
}
