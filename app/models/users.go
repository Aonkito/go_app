package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	Todos     []Todo
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5)`

	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where id = $1`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = $1, email = $2 where id = $3`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = $1`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	// user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where email = $1`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	cmdCreate := `insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4)`
	_, err = Db.Exec(cmdCreate, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	cmdGet := `select id, uuid, email, user_id, created_at from sessions where user_id = $1 and email =$2 order by created_at desc limit 1`
	err = Db.QueryRow(cmdGet, u.ID, u.Email).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	return
}

func (session *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at from sessions where uuid =$1`
	err = Db.QueryRow(cmd, session.UUID).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)

	return err == nil && session.ID != 0, err
}

func (session *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = $1`
	_, err = Db.Exec(cmd, session.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (session *Session) GetUserBySession() (user User, err error) {
	cmd := `select id, uuid, name, email, created_at from users where id =$1`
	err = Db.QueryRow(cmd, session.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt)
	return user, err
}
