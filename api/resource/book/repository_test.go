package book_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/tnaucoin/cloudnativego/api/resource/book"
	mockDB "github.com/tnaucoin/cloudnativego/mock/db"
	testUtil "github.com/tnaucoin/cloudnativego/util/test"
	"testing"
	"time"
)

func TestRepository_List(t *testing.T) {
	t.Parallel()
	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)
	repo := book.NewRepository(db)
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

	repo := book.NewRepository(db)
	id := uuid.New()
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO \"books\" ").
		WithArgs(id, "Title", "Author", mockDB.AnyTime{}, "", "", mockDB.AnyTime{}, mockDB.AnyTime{}, mockDB.AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	b := &book.Book{ID: id, Title: "Title", Author: "Author", PublishedDate: time.Now()}
	_, err = repo.Create(b)
	testUtil.NoError(t, err)
}
