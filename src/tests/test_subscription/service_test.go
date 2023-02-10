package test_subscription

import (
	"fmt"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"my-vpn-shop/config"
	"my-vpn-shop/db"
	"my-vpn-shop/outline"
	"my-vpn-shop/storage"
	"my-vpn-shop/subscription"
	"testing"
)

func TestService_Connect(t *testing.T) {
	t.Run("connect sub", func(t *testing.T) {
		dbName := fmt.Sprintf("test_store_%s.db", uuid.New())
		sqliteClient, err := db.Connect("sqlite3", dbName)
		require.NoError(t, err)
		sqliteDB := sqliteClient.Client()
		require.NoError(t, sqliteClient.CreateTables())
		sqliteStorage := storage.NewSQlDB(sqliteDB)
		api := outline.NewOutlineClient(config.Get().ApiUrl)
		service := subscription.NewSubscriptionService(sqliteStorage, api)

		var (
			subId             int64 = 123
			subName                 = "test user"
			expectedCountSubs       = 1
		)
		expectedKey, err := service.Connect(subId, subName)
		require.NoError(t, err)
		assert.NotEmpty(t, expectedKey)

		actualSubs, err := sqliteStorage.GetSubscribers()
		require.NoError(t, err)
		assert.Equal(t, expectedCountSubs, len(actualSubs))

		sub, err := sqliteStorage.GetSubByID(subId)
		require.NoError(t, err)
		assert.Equal(t, subName, sub.Name)
		assert.Equal(t, subId, sub.ID)

		actualKey, err := sqliteStorage.GetKeyBySubId(subId)
		require.NoError(t, err)
		assert.Equal(t, expectedKey.Name, actualKey.Name)
		assert.Equal(t, expectedKey.AccessUrl, actualKey.AccessUrl)
		assert.Equal(t, expectedKey.ID, actualKey.ID)
		assert.Equal(t, expectedKey.Subscriber.ID, actualKey.Subscriber.ID)

		require.NoError(t, api.DeleteKey(outline.AccessKey{Id: actualKey.ID}))
	})
}
func TestService_Disconnect(t *testing.T) {
	t.Run("disconnect sub", func(t *testing.T) {
		dbName := fmt.Sprintf("test_store_%s.db", uuid.New())
		sqliteClient, err := db.Connect("sqlite3", dbName)
		require.NoError(t, err)
		sqliteDB := sqliteClient.Client()
		require.NoError(t, sqliteClient.CreateTables())
		sqliteStorage := storage.NewSQlDB(sqliteDB)
		api := outline.NewOutlineClient(config.Get().ApiUrl)
		service := subscription.NewSubscriptionService(sqliteStorage, api)

		var (
			subId             int64 = 123
			subName                 = "test user"
			expectedCountSubs       = 0
		)
		_, err = service.Connect(subId, subName)
		require.NoError(t, err)
		require.NoError(t, service.Disconnect(subId))
		actualSubs, err := sqliteStorage.GetSubscribers()
		require.NoError(t, err)
		assert.Equal(t, expectedCountSubs, len(actualSubs))
	})
	t.Run("disconnect not existing sub", func(t *testing.T) {
		dbName := fmt.Sprintf("test_store_%s.db", uuid.New())
		sqliteClient, err := db.Connect("sqlite3", dbName)
		require.NoError(t, err)
		sqliteDB := sqliteClient.Client()
		require.NoError(t, sqliteClient.CreateTables())
		sqliteStorage := storage.NewSQlDB(sqliteDB)
		api := outline.NewOutlineClient(config.Get().ApiUrl)
		service := subscription.NewSubscriptionService(sqliteStorage, api)

		var (
			subId             int64 = 123
			expectedCountSubs       = 0
		)
		require.Error(t, service.Disconnect(subId))
		actualSubs, err := sqliteStorage.GetSubscribers()
		require.NoError(t, err)
		assert.Equal(t, expectedCountSubs, len(actualSubs))
	})
}

func TestService_FindKey(t *testing.T) {
	t.Run("find key", func(t *testing.T) {
		dbName := fmt.Sprintf("test_store_%s.db", uuid.New())
		sqliteClient, err := db.Connect("sqlite3", dbName)
		require.NoError(t, err)
		sqliteDB := sqliteClient.Client()
		require.NoError(t, sqliteClient.CreateTables())
		sqliteStorage := storage.NewSQlDB(sqliteDB)
		api := outline.NewOutlineClient(config.Get().ApiUrl)
		service := subscription.NewSubscriptionService(sqliteStorage, api)

		var (
			subId   int64 = 123
			subName       = "test user"
		)
		expectedKey, err := service.Connect(subId, subName)
		require.NoError(t, err)

		actualKey, err := service.FindKey(subId)
		require.NoError(t, err)

		assert.Equal(t, expectedKey.Name, actualKey.Name)
		assert.Equal(t, expectedKey.AccessUrl, actualKey.AccessUrl)
		assert.Equal(t, expectedKey.ID, actualKey.ID)
		assert.Equal(t, expectedKey.Subscriber.ID, actualKey.Subscriber.ID)

		require.NoError(t, service.Disconnect(subId))
	})
	t.Run("find not existing key", func(t *testing.T) {
		dbName := fmt.Sprintf("test_store_%s.db", uuid.New())
		sqliteClient, err := db.Connect("sqlite3", dbName)
		if err != nil {
			log.Fatal(err)
			return
		}
		sqliteDB := sqliteClient.Client()
		require.NoError(t, sqliteClient.CreateTables())
		sqliteStorage := storage.NewSQlDB(sqliteDB)
		api := outline.NewOutlineClient(config.Get().ApiUrl)
		service := subscription.NewSubscriptionService(sqliteStorage, api)

		var (
			subId int64 = 123
		)

		actualKey, err := service.FindKey(subId)
		require.Error(t, err)
		assert.Empty(t, actualKey)

	})
}
