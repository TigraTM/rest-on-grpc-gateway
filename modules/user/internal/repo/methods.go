package repo

import (
	"context"
	"fmt"

	"rest-on-grpc-gateway/modules/user/internal/domain"
)

// CreateUser create user in db.
func (r *Repo) CreateUser(ctx context.Context, newUser *domain.User) (*domain.User, error) {
	const query = `INSERT INTO
						"user".users
							(name,
							email,
							password)
					VALUES
							(:name,
							:email,
							:password)
					RETURNING
							id,
							name,
							email,
							password
					`

	row, err := r.DB.NamedQueryContext(ctx, query, toRepo(newUser))
	if err != nil {
		return nil, fmt.Errorf("r.DB.NamedQueryContext: %w", convertErr(err))
	}
	defer row.Close()

	var user = &User{}
	for row.Next() {
		if err := row.StructScan(user); err != nil {
			return nil, fmt.Errorf("row.StructScan: %w", err)
		}
	}

	return toDomain(user), nil
}

// GetUserByID search user in db.
func (r *Repo) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	const query = ` SELECT
						id,
						name,
						email,
						password
					FROM
						"user".users
					WHERE
						id = $1`

	var user = &User{}
	if err := r.DB.GetContext(ctx, user, query, id); err != nil {
		return nil, fmt.Errorf("r.DB.GetContext: %w", convertErr(err))
	}

	return toDomain(user), nil
}

// UpdateUserByID update user in db by id.
func (r *Repo) UpdateUserByID(ctx context.Context, id int, name, email string) (*domain.User, error) {
	const query = `UPDATE 
						"user".users
					SET
						name = $1, 
						email = $2
					WHERE
						id = $3
					RETURNING 
    					id, 
						name, 
						email`

	row := r.DB.QueryRowxContext(ctx, query, name, email, id)

	var user = &User{}
	if err := row.StructScan(user); err != nil {
		return nil, fmt.Errorf("row.StructScan: %w", err)
	}

	return toDomain(user), nil
}

// UpdateUserPasswordByID update user password in db by id.
func (r *Repo) UpdateUserPasswordByID(ctx context.Context, id int, password string) error {
	const query = `UPDATE 
						"user".users
					SET
						password = $1 
					WHERE
						id = $2
					RETURNING 
    					id, 
						name, 
						email`

	_, err := r.DB.ExecContext(ctx, query, password, id)
	if err != nil {
		return fmt.Errorf("r.DB.ExecContext: %w", err)
	}

	return nil
}

// DeleteUserByID delete user in db by id.
func (r *Repo) DeleteUserByID(ctx context.Context, id int) error {
	const query = `	DELETE
					FROM
						"user".users
					WHERE
						id = $1 `

	_, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("r.DB.ExecContext: %w", err)
	}

	return nil
}
