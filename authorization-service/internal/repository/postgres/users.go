package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/entity"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/repository"
)

const (
	Client = "client"
)

type UsersRepo struct {
	*sql.DB
}

func NewUsersRepo(pg *sql.DB) repository.Users {
	return &UsersRepo{pg}
}

func (ur *UsersRepo) Store(ctx context.Context, user entity.User) (int32, error) {
	tx, err := ur.DB.Begin()
	if err != nil {
		return 0, fmt.Errorf("UsersRepo - Store - Begin: %w", err)
	}
	defer tx.Rollback()

	res, err := tx.ExecContext(ctx, `
		INSERT INTO users(name, phone, email, password, role)
			VALUES($1, $2, $3, $4, $5)
	`, user.Name, user.Phone, user.Email, user.Password, Client)
	if err != nil {
		return 0, fmt.Errorf("UsersRepo - Store - ExecContext #1: %w", err)
	}

	rows, err := res.RowsAffected()
	if rows != 1 || err != nil {
		return 0, fmt.Errorf("UsersRepo - Store - RowsAffected #1: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("UsersRepo - Store - LastInsertId #1: %w", err)
	}

	res, err = tx.ExecContext(ctx, `
		INSERT INTO sessions (user_id)
			VALUES ($1)
	`, int32(id))
	if err != nil {
		return 0, fmt.Errorf("UsersRepo - Store - ExecContext #2: %w", err)
	}

	rows, err = res.RowsAffected()
	if rows != 1 || err != nil {
		return 0, fmt.Errorf("UsersRepo - Store - RowsAffected #2: %w", err)
	}

	res, err = tx.ExecContext(ctx, `
		INSERT INTO verifications (user_id, verification_token, verification_ttl, verified)
			VALUES ($1, $2, $3, $4)
	`, int32(id), user.VerificationToken, user.VerificationTTL, false)
	if err != nil {
		return 0, fmt.Errorf("UsersRepo - Store - ExecContext #3: %w", err)
	}

	rows, err = res.RowsAffected()
	if rows != 1 || err != nil {
		return 0, fmt.Errorf("UsersRepo - Store - RowsAffected #3: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("UsersRepo - Store - Commit: %w", err)
	}

	return int32(id), nil
}

func (ur *UsersRepo) Delete(ctx context.Context, userId int32) error {
	res, err := ur.DB.ExecContext(ctx, `
		DELETE FROM users
		WHERE id = $1
	`, userId)

	if err != nil {
		return fmt.Errorf("UsersRepo - Delete - ExecContext: %w", err)
	}

	rows, err := res.RowsAffected()
	if rows != 1 || err != nil {
		return fmt.Errorf("UsersRepo - Delete - RowsAffected: %w", err)
	}

	return nil
}

func (ur *UsersRepo) GetByPhone(ctx context.Context, phone string) (entity.User, error) {
	user := entity.User{}

	row := ur.DB.QueryRowContext(ctx, `
		SELECT id, name, phone, email, password, role, session_token, session_ttl,
		verification_token, verification_ttl, verified
		FROM users
		INNER JOIN sessions ON users.id = sessions.user_id
		INNER JOIN verifications ON users.id = verifications.user_id
		WHERE phone = $1
	`, phone)

	var email sql.NullString
	var sessionToken sql.NullString
	var sessionTTL sql.NullTime
	var VerificationToken sql.NullString
	var VerificationTTL sql.NullTime

	err := row.Scan(&user.Id, &user.Name, &user.Phone, &email, &user.Password,
		&user.Role, &sessionToken, &sessionTTL, &VerificationToken,
		&VerificationTTL, &user.Verified)

	if errors.Is(err, sql.ErrNoRows) {
		return user, entity.ErrUserDoesNotExist
	} else if err != nil {
		return user, fmt.Errorf("UsersRepo - GetByPhone - Scan: %w", err)
	}

	user.Email = email.String
	user.SessionToken = sessionToken.String
	user.SessionTTL = sessionTTL.Time
	user.VerificationToken = VerificationToken.String
	user.VerificationTTL = VerificationTTL.Time

	return user, nil
}

func (ur *UsersRepo) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	user := entity.User{}

	row := ur.DB.QueryRowContext(ctx, `
		SELECT id, name, phone, email, password, role, session_token, session_ttl,
		verification_token, verification_ttl, verified
		FROM users
		INNER JOIN sessions ON users.id = sessions.user_id
		INNER JOIN verifications ON users.id = verifications.user_id
		WHERE email = $1
	`, email)

	var sessionToken sql.NullString
	var sessionTTL sql.NullTime
	var VerificationToken sql.NullString
	var VerificationTTL sql.NullTime

	err := row.Scan(&user.Id, &user.Name, &user.Phone, &email, &user.Password,
		&user.Role, &sessionToken, &sessionTTL, &VerificationToken,
		&VerificationTTL, &user.Verified)

	if errors.Is(err, sql.ErrNoRows) {
		return user, entity.ErrUserDoesNotExist
	} else if err != nil {
		return user, fmt.Errorf("UsersRepo - GetByPhone - Scan: %w", err)
	}

	user.SessionToken = sessionToken.String
	user.SessionTTL = sessionTTL.Time
	user.VerificationToken = VerificationToken.String
	user.VerificationTTL = VerificationTTL.Time

	return user, nil
}

func (ur *UsersRepo) GetById(ctx context.Context, id int32) (entity.User, error) {
	user := entity.User{}

	row := ur.DB.QueryRowContext(ctx, `
		SELECT id, name, phone, email, password, role, session_token, session_ttl,
		verification_token, verification_ttl, verified
		FROM users
		INNER JOIN sessions ON users.id = sessions.user_id
		INNER JOIN verifications ON users.id = verifications.user_id
		WHERE id = $1
	`, id)

	var email sql.NullString
	var sessionToken sql.NullString
	var sessionTTL sql.NullTime
	var VerificationToken sql.NullString
	var VerificationTTL sql.NullTime

	err := row.Scan(&user.Id, &user.Name, &user.Phone, &email, &user.Password,
		&user.Role, &sessionToken, &sessionTTL, &VerificationToken,
		&VerificationTTL, &user.Verified)

	if errors.Is(err, sql.ErrNoRows) {
		return user, entity.ErrUserDoesNotExist
	} else if err != nil {
		return user, fmt.Errorf("UsersRepo - GetById - Scan: %w", err)
	}

	user.Email = email.String
	user.SessionToken = sessionToken.String
	user.SessionTTL = sessionTTL.Time
	user.VerificationToken = VerificationToken.String
	user.VerificationTTL = VerificationTTL.Time

	return user, nil
}

func (ur *UsersRepo) GetByToken(ctx context.Context, token string) (entity.User, error) {
	user := entity.User{}

	row := ur.DB.QueryRowContext(ctx, `
		SELECT id, name, phone, email, password, role, session_token, session_ttl,
		verification_token, verification_ttl, verified
		FROM users
		INNER JOIN sessions ON users.id = sessions.user_id
		INNER JOIN verifications ON users.id = verifications.user_id
		WHERE session_token = $1
	`, token)

	var email sql.NullString
	var sessionToken sql.NullString
	var sessionTTL sql.NullTime
	var VerificationToken sql.NullString
	var VerificationTTL sql.NullTime

	err := row.Scan(&user.Id, &user.Name, &user.Phone, &email, &user.Password,
		&user.Role, &sessionToken, &sessionTTL, &VerificationToken,
		&VerificationTTL, &user.Verified)

	if errors.Is(err, sql.ErrNoRows) {
		return user, entity.ErrUserDoesNotExist
	} else if err != nil {
		return user, fmt.Errorf("UsersRepo - GetByToken - Scan: %w", err)
	}

	user.Email = email.String
	user.SessionToken = sessionToken.String
	user.SessionTTL = sessionTTL.Time
	user.VerificationToken = VerificationToken.String
	user.VerificationTTL = VerificationTTL.Time

	return user, nil
}

func (ur *UsersRepo) UpdateSession(ctx context.Context, user entity.User) error {
	res, err := ur.DB.ExecContext(ctx, `
		UPDATE sessions
		SET session_token = $1, session_ttl = $2
		WHERE user_id = $3
	`, user.SessionToken, user.SessionTTL, user.Id)

	if err != nil {
		return fmt.Errorf("UsersRepo - UpdateSession - ExecContext: %w", err)
	}

	rows, err := res.RowsAffected()
	if rows != 1 || err != nil {
		return fmt.Errorf("UsersRepo - UpdateSession - RowsAffected: %w", err)
	}

	return nil
}

func (ur *UsersRepo) UpdateVerification(ctx context.Context, user entity.User) error {
	res, err := ur.DB.ExecContext(ctx, `
		UPDATE verifications
		SET verified = $1
		WHERE user_id = $2
	`, true, user.Id)

	if err != nil {
		return fmt.Errorf("UsersRepo - UpdateVerification - ExecContext: %w", err)
	}

	rows, err := res.RowsAffected()
	if rows != 1 || err != nil {
		return fmt.Errorf("UsersRepo - UpdateVerification - RowsAffected: %w", err)
	}

	return nil
}
