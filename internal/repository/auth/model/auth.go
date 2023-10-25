package model

import (
	"database/sql"
	"time"
)

type UserU struct {
	ID    int64  `db:"id"`
	Name  string `db:"username"`
	Email string `db:"email"`
	Role  int    `db:"role"`
}

type UserC struct {
	UU       UserU  `db:""`
	Password string `db:"password"`
}

type User struct {
	UC        UserC        `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
