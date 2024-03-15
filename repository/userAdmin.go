package repository

import (
	"database/sql"
	"flashbank/structs"
)

func GetAllUserAdmin(db *sql.DB) (results []structs.UserAdmin, err error) {
	sql := "SELECT * FROM user_admin"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = structs.UserAdmin{}
		err = rows.Scan(&user.ID, &user.UserName, &user.Password, &user.Role)
		if err != nil {
			panic(err)
		}
		results = append(results, user)
	}
	return
}

func Login(db *sql.DB, userAdmin structs.UserAdmin) (userAdmin_id int, err error) {
	sql := `
	SELECT id FROM user_admin WHERE
	username = $1 AND password = $2`
	err = db.QueryRow(sql, userAdmin.UserName, userAdmin.Password).Scan(&userAdmin.ID)
	userAdmin_id = userAdmin.ID
	return
}

func Register(db *sql.DB, userAdmin structs.UserAdmin) (err error) {
	sql := "INSERT INTO user_admin (username, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	errs := db.QueryRow(sql, userAdmin.UserName, userAdmin.Password, userAdmin.Role, userAdmin.CreatedAt, userAdmin.UpdatedAt)
	return errs.Err()
}
