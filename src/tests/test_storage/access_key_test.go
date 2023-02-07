package test_storage

import (
	"fmt"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"log"
	"my-vpn-shop/db"
	"my-vpn-shop/storage"
	"testing"
)

func TestAccessKey_InsertAccessKey(t *testing.T) {
	t.Run("insert access key in db", func(t *testing.T) {
		dbName := fmt.Sprintf("test_store_%s.db", uuid.New())
		sqliteClient, err := db.Connect("sqlite3", dbName)
		if err != nil {
			log.Fatal(err)
			return
		}
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
		require.NoError(t, sqliteStorage.DeleteSubscriber(testSub.ID))
	})
}
