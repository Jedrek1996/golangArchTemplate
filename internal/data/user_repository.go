package data

import (
	"database/sql"

	"template/internal/interfaces"
	"template/internal/model"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) interfaces.UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) CreateUser(user *model.User) (*model.User, error) {
	stmt, err := r.db.Prepare("INSERT INTO users(name, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(id)
	return user, nil
}

func (r *UserRepositoryImpl) GetUser(id int) (*model.User, error) {
	row := r.db.QueryRow("SELECT id, name, email, password FROM users WHERE id = ?", id)
	user := &model.User{}
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
