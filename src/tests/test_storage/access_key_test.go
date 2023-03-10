package test_storage

import (
	"fmt"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"my-vpn-shop/db"
	"my-vpn-shop/storage"
	"testing"
)

func TestAccessKey_InsertAccessKey(t *testing.T) {
	t.Run("insert access key in db", func(t *testing.T) {
		dbName := fmt.Sprintf("test_store_%s.db", uuid.New())
		sqliteClient, err := db.Connect("sqlite3", dbName)
		require.NoError(t, err)
		sqliteDB := sqliteClient.Client()
		require.NoError(t, sqliteClient.CreateTables())
		sqliteStorage := storage.NewSQlDB(sqliteDB)
		var (
			keyID              = "test_id"
			keyAccessUrl       = "https://example.com"
			keyName            = "test key name"
			subID        int64 = 123
			subName            = "test user"
		)

		testSub, err := sqliteStorage.InsertSubscriber(subID, subName)
		require.NoError(t, err)

		expected := db.AccessKey{
			ID:         keyID,
			AccessUrl:  keyAccessUrl,
			Name:       keyName,
			Subscriber: testSub,
		}
		actual, err := sqliteStorage.InsertAccessKey(keyID, keyName, keyAccessUrl, testSub)

		require.NoError(t, err)
		require.NotEmpty(t, actual)
		require.Equal(t, expected, actual)
	})
}

func TestAccessKey_GetKeyBySubId(t *testing.T) {
	t.Run("get key by sub id", func(t *testing.T) {
		dbName := fmt.Sprintf("test_store_%s.db", uuid.New())
		sqliteClient, err := db.Connect("sqlite3", dbName)
		require.NoError(t, err)
		sqliteDB := sqliteClient.Client()
		require.NoError(t, sqliteClient.CreateTables())
		sqliteStorage := storage.NewSQlDB(sqliteDB)
		var (
			keyID              = "test_id"
			keyAccessUrl       = "https://example.com"
			keyName            = "test key name"
			subID        int64 = 123
			subName            = "test user"
		)

		testSub, err := sqliteStorage.InsertSubscriber(subID, subName)
		require.NoError(t, err)

		expected := db.AccessKey{
			ID:         keyID,
			AccessUrl:  keyAccessUrl,
			Name:       keyName,
			Subscriber: db.Subscriber{ID: testSub.ID},
		}
		_, err = sqliteStorage.InsertAccessKey(keyID, keyName, keyAccessUrl, testSub)
		require.NoError(t, err)

		actual, err := sqliteStorage.GetKeyBySubId(testSub.ID)
		require.NotEmpty(t, actual)
		require.Equal(t, expected, actual)
	})
	t.Run("get key by not existing sub id", func(t *testing.T) {
		dbName := fmt.Sprintf("test_store_%s.db", uuid.New())
		sqliteClient, err := db.Connect("sqlite3", dbName)
		require.NoError(t, err)
		sqliteDB := sqliteClient.Client()
		require.NoError(t, sqliteClient.CreateTables())
		sqliteStorage := storage.NewSQlDB(sqliteDB)
		var subID int64 = 123

		actual, err := sqliteStorage.GetKeyBySubId(subID)
		require.Error(t, err)
		require.Empty(t, actual)
	})
}
