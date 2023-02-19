package postgres_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/entity"
	p "github.com/mrsubudei/chat-bot-backend/authorization-service/internal/repository/postgres"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	query1 := regexp.QuoteMeta(`
	INSERT INTO users(name, phone, email, password, role)
			VALUES($1, $2, $3, $4, $5)
	`)

	query2 := regexp.QuoteMeta(`
	INSERT INTO sessions (user_id)
			VALUES ($1)
	`)

	query3 := regexp.QuoteMeta(`
	INSERT INTO verifications (user_id)
			VALUES ($1)
	`)

	user := entity.User{
		Id:       1,
		Name:     "Alice",
		Phone:    "87776542154",
		Email:    "alice@gmail.com",
		Password: "abcd",
		Role:     p.Client,
	}
	mock.ExpectBegin()
	mock.ExpectExec(query1).WithArgs(user.Name, user.Phone, user.Email,
		user.Password, user.Role).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(query2).WithArgs(user.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(query3).WithArgs(user.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	ctx := context.Background()
	repo := p.NewUsersRepo(db)
	if id, err := repo.Store(ctx, user); err != nil {
		t.Fatalf("error was not expected: %s", err)
	} else if id != 1 {
		t.Fatalf("error was not expected id: %d, got: %d", 1, id)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestStoreShouldRollback(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query1 := regexp.QuoteMeta(`
	INSERT INTO users(name, phone, email, password, role)
			VALUES($1, $2, $3, $4, $5)
	`)
	user := entity.User{
		Id:       1,
		Name:     "Alice",
		Phone:    "87776542154",
		Email:    "alice@gmail.com",
		Password: "abcd",
		Role:     p.Client,
	}
	mock.ExpectBegin()
	mock.ExpectExec(query1).WithArgs(user.Name, user.Phone, user.Email,
		user.Password, user.Role).WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	ctx := context.Background()
	repo := p.NewUsersRepo(db)
	if _, err := repo.Store(ctx, user); err == nil {
		t.Fatal("expected error")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetByPhone(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	phone := "87776542154"

	rows := sqlmock.NewRows([]string{"id", "name", "phone", "email", "password",
		"role", "session_token", "session_ttl", "verification_token", "verified"}).
		AddRow(1, "Vasya", phone, "vasye@gmail.com", "pass", p.Client,
			"token", time.Now(), "056841", false)

	query := regexp.QuoteMeta(`
	SELECT id, name, phone, email, password, role, session_token, session_ttl,
		verification_token, verified
		FROM users
		INNER JOIN sessions ON users.id = sessions.user_id
		INNER JOIN verifications ON users.id = verifications.user_id
		WHERE phone = $1
	`)

	mock.ExpectQuery(query).WithArgs(phone).
		WillReturnRows(rows)

	ctx := context.Background()
	repo := p.NewUsersRepo(db)
	user, err := repo.GetByPhone(ctx, phone)
	assert.NoError(t, err)
	assert.NotNil(t, user)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	id := int32(2)

	rows := sqlmock.NewRows([]string{"id", "name", "phone", "email", "password",
		"role", "session_token", "session_ttl", "verification_token", "verified"}).
		AddRow(id, "Vasya", "87776548754", "vasye@gmail.com", "pass", p.Client,
			"token", time.Now(), "056841", false)

	query := regexp.QuoteMeta(`
	SELECT id, name, phone, email, password, role, session_token, session_ttl,
		verification_token, verified
		FROM users
		INNER JOIN sessions ON users.id = sessions.user_id
		INNER JOIN verifications ON users.id = verifications.user_id
		WHERE id = $1
	`)

	mock.ExpectQuery(query).WithArgs(id).
		WillReturnRows(rows)

	ctx := context.Background()
	repo := p.NewUsersRepo(db)
	user, err := repo.GetById(ctx, id)
	assert.NoError(t, err)
	assert.NotNil(t, user)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetByToken(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	token := "token"

	rows := sqlmock.NewRows([]string{"id", "name", "phone", "email", "password",
		"role", "session_token", "session_ttl", "verification_token", "verified"}).
		AddRow(4, "Vasya", "87776548754", "vasye@gmail.com", "pass", p.Client,
			token, time.Now(), "056841", false)

	query := regexp.QuoteMeta(`
	SELECT id, name, phone, email, password, role, session_token, session_ttl,
		verification_token, verified
		FROM users
		INNER JOIN sessions ON users.id = sessions.user_id
		INNER JOIN verifications ON users.id = verifications.user_id
		WHERE session_token = $1
	`)

	mock.ExpectQuery(query).WithArgs(token).
		WillReturnRows(rows)

	ctx := context.Background()
	repo := p.NewUsersRepo(db)
	user, err := repo.GetByToken(ctx, token)
	assert.NoError(t, err)
	assert.NotNil(t, user)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateSession(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	query := regexp.QuoteMeta(`
	UPDATE sessions
		SET session_token = $1, session_ttl = $2
		WHERE user_id = $3
	`)

	user := entity.User{
		Id:           5,
		SessionToken: "token",
		SessionTTL:   time.Now(),
	}

	mock.ExpectExec(query).WithArgs(user.SessionToken, user.SessionTTL, user.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	repo := p.NewUsersRepo(db)
	err = repo.UpdateSession(ctx, user)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateVerification(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	query := regexp.QuoteMeta(`
	UPDATE verifications
		SET verification_token = $1, verified = $2
		WHERE user_id = $3
	`)

	user := entity.User{
		Id:                5,
		VerificationToken: "05498",
		Verified:          false,
	}

	mock.ExpectExec(query).WithArgs(user.VerificationToken, user.Verified, user.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	repo := p.NewUsersRepo(db)
	err = repo.UpdateVerification(ctx, user)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}
