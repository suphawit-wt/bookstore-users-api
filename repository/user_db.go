package repository

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type userRepositoryDB struct {
	db *sqlx.DB
}

func NewUserRepositoryDB(db *sqlx.DB) UserRepository {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) GetAll() ([]User, error) {
	query := "SELECT * FROM users;"

	users := []User{}
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r userRepositoryDB) GetById(id int64) (*User, error) {
	query := "SELECT * FROM users WHERE id = @p1;"

	user := User{}
	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r userRepositoryDB) GetByStatus(status string) ([]User, error) {
	query := "SELECT * FROM users WHERE status = @p1;"

	users := []User{}
	err := r.db.Select(&users, query, status)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r userRepositoryDB) Create(req User) error {
	query := "INSERT INTO users (first_name, last_name, email, password) VALUES (@p1, @p2, @p3, @p4);"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(req.FirstName, req.LastName, req.Email, hashPass)
	if err != nil {
		return err
	}

	return nil
}

func (r userRepositoryDB) Update(id int64, req User) error {
	query := "UPDATE users SET first_name = @p1, last_name = @p2, email = @p3, password = @p4 WHERE id = @p5;"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(req.FirstName, req.LastName, req.Email, hashPass, id)
	if err != nil {
		return err
	}

	return nil
}

func (r userRepositoryDB) Delete(id int64) error {
	query := "DELETE FROM users WHERE id = @p1;"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
