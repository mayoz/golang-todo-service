package todo

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func mockDb() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("mock database can not create: %s", err)
	}

	return db, mock
}

func TestGet(t *testing.T) {
	db, mock := mockDb()
	store := &Store{db}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "text", "completed", "created_at"}).
		AddRow(1, "foo", true, time.Now()).
		AddRow(2, "bar", true, time.Now())

	mock.ExpectQuery("SELECT * FROM todos").WillReturnRows(rows)

	todos, err := store.Get()
	assert.NotEmpty(t, todos)
	assert.NoError(t, err)
	assert.Len(t, todos, 2)
}

func TestGetQueryError(t *testing.T) {
	db, mock := mockDb()
	store := &Store{db}
	defer db.Close()

	mock.ExpectQuery("SELECT * FROM todos").WillReturnError(fmt.Errorf("some error"))

	todos, err := store.Get()
	assert.Nil(t, todos)
	assert.NotEmpty(t, err)
}

func TestStore(t *testing.T) {
	db, mock := mockDb()
	store := &Store{db}
	defer db.Close()

	mockId := int64(rand.Intn(100))
	prep := mock.ExpectPrepare("INSERT INTO todos (text) VALUES (?)")
	prep.ExpectExec().WithArgs("foo").WillReturnResult(sqlmock.NewResult(mockId, 1))

	id, err := store.Store("foo")
	assert.Equal(t, mockId, id)
	assert.NoError(t, err)
}

func TestStoreError(t *testing.T) {
	db, mock := mockDb()
	store := &Store{db}
	defer db.Close()

	prep := mock.ExpectPrepare("INSERT INTO todos (text) VALUES (?)")
	prep.ExpectExec().WithArgs("foo").WillReturnError(fmt.Errorf("some error"))

	id, err := store.Store("foo")
	assert.Equal(t, int64(-1), id)
	assert.Error(t, err)
}

func TestFind(t *testing.T) {
	db, mock := mockDb()
	store := &Store{db}
	defer db.Close()

	mockId := int64(rand.Intn(100))
	rows := sqlmock.NewRows([]string{"id", "text", "completed", "created_at"}).
		AddRow(mockId, "foo", true, time.Now())

	mock.ExpectQuery("SELECT * FROM todos WHERE id = ?").
		WithArgs(mockId).
		WillReturnRows(rows)

	user, err := store.Find(mockId)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestFindError(t *testing.T) {
	db, mock := mockDb()
	store := &Store{db}
	defer db.Close()

	mockId := int64(rand.Intn(100))
	rows := sqlmock.NewRows([]string{"id", "text", "completed", "created_at"})

	mock.ExpectQuery("SELECT * FROM todos WHERE id = ?").
		WithArgs(mockId).
		WillReturnRows(rows)

	user, err := store.Find(mockId)
	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestToggle(t *testing.T) {
	db, mock := mockDb()
	store := &Store{db}
	defer db.Close()

	mockId := int64(rand.Intn(100))
	mock.ExpectExec("UPDATE todos SET completed = !completed WHERE id = ?").
		WithArgs(mockId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := store.Toggle(mockId)
	assert.NoError(t, err)
}

func TestToggleError(t *testing.T) {
	db, mock := mockDb()
	store := &Store{db}
	defer db.Close()

	mockId := int64(rand.Intn(100))
	mock.ExpectExec("UPDATE todos SET completed = !completed WHERE id = ?").
		WithArgs(mockId).
		WillReturnError(fmt.Errorf("some error"))

	err := store.Toggle(mockId)
	assert.Error(t, err)
}

func TestDestroy(t *testing.T) {
	db, mock := mockDb()
	store := &Store{db}
	defer db.Close()

	mockId := int64(rand.Intn(100))
	mock.ExpectExec("DELETE FROM todos WHERE id = ?").
		WithArgs(mockId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := store.Destroy(mockId)
	assert.NoError(t, err)
}

func TestDestroyError(t *testing.T) {
	db, mock := mockDb()
	store := &Store{db}
	defer db.Close()

	mockId := int64(rand.Intn(100))
	mock.ExpectExec("DELETE FROM todos WHERE id = ?").
		WithArgs(mockId).
		WillReturnError(fmt.Errorf("some error"))

	err := store.Destroy(mockId)
	assert.Error(t, err)
}
