package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"dissys/internal/database"
)

// repository to handle postgres operations
type Repository struct {
	db *pgxpool.Pool
}

// NewRepository creates a new repository
func NewRepository(datastores *database.DataStores) *Repository {
	return &Repository{
		db: datastores.Postgres,
	}
}

// CreateTx creates a new transaction
func (r *Repository) CreateTx(ctx context.Context) (pgx.Tx, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to create transaction")
		return nil, err
	}

	return tx, nil
}

// RollbackTx rolls back a transaction
func (r *Repository) CommitTx(ctx context.Context, tx pgx.Tx) error {
	// commit transaction or rollback if error
	if err := tx.Commit(ctx); err != nil {
		log.Error().Err(err).Msg("failed to commit transaction")
		return err
	}
	defer tx.Rollback(ctx)

	return nil
}

// retrieve user credentials from database
func (r *Repository) FetchCredentials(ctx context.Context, email string) (*UserCreds, error) {
	creds := &UserCreds{}

	err := r.db.QueryRow(ctx, QueryGetUserCreds, email).Scan(
		&creds.UserID, &creds.Password, &creds.Status, &creds.Role,
	)
	if err != nil {
		// check if user not found
		if errors.Is(err, pgx.ErrNoRows) {
			errMsg := fmt.Sprintf("user with email %s not found", email)
			return nil, echo.NewHTTPError(http.StatusNotFound, map[string]string{"error": errMsg})
		}

		log.Error().Err(err).Msg("failed to fetch user credentials")
		return nil, err
	}

	return creds, nil
}

// insert user into database
func (r *Repository) InsertUser(ctx context.Context, tx pgx.Tx, user *User) error {
	_, err := tx.Exec(ctx, QueryRegisterUser,
		&user.UserID, &user.Email, &user.Password, &user.Created, &user.Status, &user.Role,
	)
	if err != nil {
		// check if user already exists
		pgErr := &pgconn.PgError{}
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			errMsg := fmt.Sprintf("user with email %s already exists", user.Email)
			return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": errMsg})
		}

		log.Error().Err(err).Msg("failed to insert user")
		return err
	}

	return nil
}
