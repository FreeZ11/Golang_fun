package user

import (
	"database/sql"
	"fun_project/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * from users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	user := new(types.User)
	for rows.Next() {
		user, err = scanRowToUser(rows)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (s *Store) GetUserByLogin(login string) (*types.GetUserPayload, error) {
	rows, err := s.db.Query("SELECT firstName, lastName, email, login, createdAt from users WHERE login = ?", login)
	if err != nil {
		return nil, err
	}

	user := new(types.GetUserPayload)
	for rows.Next() {
		user, err = scanRowToGetUserPayload(rows)
		if err != nil {
			return nil, err
		}
	}

	return user, nil

}

func (s *Store) CreateUser(user types.User) error {
	return nil
}

func scanRowToUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Login,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func scanRowToGetUserPayload(rows *sql.Rows) (*types.GetUserPayload, error) {
	user := new(types.GetUserPayload)
	err := rows.Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Login,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
