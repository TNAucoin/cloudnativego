package book

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	mockDB "github.com/tnaucoin/cloudnativego/mock/db"
	testUtil "github.com/tnaucoin/cloudnativego/util/test"
	"testing"
)

func TestRepository_List(t *testing.T) {
	t.Parallel()
	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)
	repo := NewRepository(db)
	mockRows := sqlmock.NewRows([]string{"id", "title", "author"}).
		AddRow(uuid.New(), "Book1", "Author1").
		AddRow(uuid.New(), "Book2", "Author2")
	mock.ExpectQuery("^SELECT (.+) FROM \"books\"").WillReturnRows(mockRows)
	books, err := repo.List()
	testUtil.NoError(t, err)
	testUtil.Equal(t, len(books), 2)
}

func TestRepository_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)
	repo := NewRepository(db)
	id := uuid.New()
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO \"books\" ").
		WithArgs(id, "Title", "Author", mockDB.AnyTime{}, "", "", mockDB.AnyTime{}, mockDB.AnyTime{}, nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	book := &book
	_, err = repo.Create(book)
	testUtil.NoError(t, err)
}
