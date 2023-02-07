package test_storage

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"my-vpn-shop/db"
	"my-vpn-shop/storage"
	"testing"
)

func TestSubscriber_InsertSubscriber(t *testing.T) {
	t.Run("insert subscriber in db", func(t *testing.T) {
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
			subID   int64 = 123
			subName       = "test user"
		)

		actual, err := sqliteStorage.InsertSubscriber(subID, subName)
		assert.NotEmpty(t, actual)
		require.NoError(t, err)
		assert.Equal(t, subID, actual.ID)
		assert.Equal(t, subName, actual.Name)
	})
}
