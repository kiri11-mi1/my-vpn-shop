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

	t.Run("empty url", func(t *testing.T) {
		var (
			client = outline.NewOutlineClient("")
		)
		keys, err := client.GetKeys()
		require.Error(t, err)
		assert.Empty(t, keys)
	})
}

func TestApi_CreateKey_DeleteKey(t *testing.T) {
	t.Run("create key via outline api and delete it", func(t *testing.T) {
		var (
			client = outline.NewOutlineClient(os.Getenv("VPN_URL_API"))
		)
		key, err := client.CreateKey()
		require.NoError(t, err)
		assert.NotEmpty(t, key)
		require.NoError(t, client.DeleteKey(key))
	})
	t.Run("delete not existing key", func(t *testing.T) {
		var (
			client = outline.NewOutlineClient(os.Getenv("VPN_URL_API"))
		)
		key, err := client.CreateKey()
		require.NoError(t, err)
		assert.NotEmpty(t, key)
		require.NoError(t, client.DeleteKey(key))

		require.ErrorIs(t, client.DeleteKey(key), outline.ErrInApi)

	})
}

func TestApi_ChangeKeyName(t *testing.T) {
	t.Run("change key name", func(t *testing.T) {
		var (
			client = outline.NewOutlineClient(os.Getenv("VPN_URL_API"))
		)
		key, err := client.CreateKey()
		require.NoError(t, err)
		assert.NotEmpty(t, key)

		require.NoError(t, client.ChangeKeyName("test_key", key))

		require.NoError(t, client.DeleteKey(key))
	})
	t.Run("empty key name", func(t *testing.T) {
		var (
			client = outline.NewOutlineClient(os.Getenv("VPN_URL_API"))
		)
		key, err := client.CreateKey()
		require.NoError(t, err)
		assert.NotEmpty(t, key)

		require.ErrorIs(t, client.ChangeKeyName("", key), outline.ErrInApi)

		require.NoError(t, client.DeleteKey(key))
	})
	t.Run("not existing key", func(t *testing.T) {
		var (
			client = outline.NewOutlineClient(os.Getenv("VPN_URL_API"))
		)
		key, err := client.CreateKey()
		require.NoError(t, err)
		assert.NotEmpty(t, key)
		require.NoError(t, client.DeleteKey(key))

		require.ErrorIs(t, client.ChangeKeyName("test key", key), outline.ErrInApi)
	})
}
