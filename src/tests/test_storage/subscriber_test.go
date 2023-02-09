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
	"time"
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
		assert.IsType(t, time.Now(), actual.PayedAt)
	})
	t.Run("insert subscriber in db with empty name", func(t *testing.T) {
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
			subName       = ""
		)

		actual, err := sqliteStorage.InsertSubscriber(subID, subName)
		assert.NotEmpty(t, actual)
		require.NoError(t, err)
		assert.Equal(t, subID, actual.ID)
		assert.Equal(t, subName, actual.Name)
		assert.IsType(t, time.Now(), actual.PayedAt)
	})
	t.Run("insert existing subscriber", func(t *testing.T) {
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
			subName       = ""
		)

		_, err = sqliteStorage.InsertSubscriber(subID, subName)
		require.NoError(t, err)

		_, err = sqliteStorage.InsertSubscriber(subID, subName)
		require.Error(t, err)
	})
}

func TestSubscriber_DeleteSubscriber(t *testing.T) {
	t.Run("delete subscriber from db", func(t *testing.T) {
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

		_, err = sqliteStorage.InsertSubscriber(subID, subName)
		require.NoError(t, err)
		require.NoError(t, sqliteStorage.DeleteSubscriber(subID))
		subs, err := sqliteStorage.GetSubscribers()
		require.NoError(t, err)
		assert.Zero(t, len(subs))
	})
	t.Run("delete not existing subscriber from db", func(t *testing.T) {
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
			notExistingID int64 = 123
			lenSubArray         = 10
		)
		for i := 1; i <= lenSubArray; i++ {
			_, err = sqliteStorage.InsertSubscriber(int64(i), fmt.Sprintf("test sub %d", i))
			require.NoError(t, err)
		}
		require.NoError(t, sqliteStorage.DeleteSubscriber(notExistingID))
		subs, err := sqliteStorage.GetSubscribers()
		require.NoError(t, err)
		assert.Equal(t, lenSubArray, len(subs))
	})
}

func TestSubscriber_GetSubscribers(t *testing.T) {
	t.Run("get all subs", func(t *testing.T) {
		dbName := fmt.Sprintf("test_store_%s.db", uuid.New())
		sqliteClient, err := db.Connect("sqlite3", dbName)
		if err != nil {
			log.Fatal(err)
			return
		}
		sqliteDB := sqliteClient.Client()
		require.NoError(t, sqliteClient.CreateTables())
		sqliteStorage := storage.NewSQlDB(sqliteDB)
		var lenSubArray = 10
		for i := 1; i <= lenSubArray; i++ {
			_, err = sqliteStorage.InsertSubscriber(int64(i), fmt.Sprintf("test sub %d", i))
			require.NoError(t, err)
		}
		subs, err := sqliteStorage.GetSubscribers()
		require.NoError(t, err)
		assert.Equal(t, lenSubArray, len(subs))
	})
}

func TestSubscriber_GetSubByID(t *testing.T) {
	t.Run("get sub from db", func(t *testing.T) {
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
			lenSubArray       = 10
			subID       int64 = 12
			subName           = fmt.Sprintf("test sub %d", subID)
		)
		for i := 1; i <= lenSubArray; i++ {
			_, err = sqliteStorage.InsertSubscriber(int64(i), fmt.Sprintf("test sub %d", i))
			require.NoError(t, err)
		}
		_, err = sqliteStorage.InsertSubscriber(subID, subName)
		require.NoError(t, err)

		actual, err := sqliteStorage.GetSubByID(subID)
		require.NoError(t, err)
		require.NotEmpty(t, actual)
		require.Equal(t, subID, actual.ID)
		require.Equal(t, subName, actual.Name)
	})
	t.Run("get not existing sub", func(t *testing.T) {
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
			subID int64 = 12
		)

		actual, err := sqliteStorage.GetSubByID(subID)
		require.Error(t, storage.ErrSubNotFound)
		require.Empty(t, actual)
	})
	t.Run("check valid date", func(t *testing.T) {
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
			subID    int64 = 123
			subName        = "test sub"
			expected       = time.Now().Format("2006-01-02")
		)

		_, err = sqliteStorage.InsertSubscriber(subID, subName)
		require.NoError(t, err)

		sub, err := sqliteStorage.GetSubByID(subID)
		require.NoError(t, err)

		assert.Equal(t, expected, sub.PayedAt.Format("2006-01-02"))
	})
}

func TestSubscriber_UpdateSubscriberPayedAt(t *testing.T) {
	t.Run("update sub pay date", func(t *testing.T) {
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
			subID        int64 = 12
			subName            = fmt.Sprintf("test sub %d", subID)
			expectedDate       = time.Now()
		)
		_, err = sqliteStorage.InsertSubscriber(subID, subName)
		require.NoError(t, err)

		require.NoError(t, sqliteStorage.UpdateSubscriberPayedAt(subID))
		sub, err := sqliteStorage.GetSubByID(subID)
		require.NoError(t, err)
		actualDate := sub.PayedAt

		require.Equal(t, expectedDate.Day(), actualDate.Day())
		require.Equal(t, expectedDate.Year(), actualDate.Year())
		require.Equal(t, expectedDate.Month(), actualDate.Month())
	})
	t.Run("update not existing sub pay date", func(t *testing.T) {
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
			subID int64 = 12
		)
		require.ErrorIs(t, sqliteStorage.UpdateSubscriberPayedAt(subID), storage.ErrSubNotFound)
	})

}
