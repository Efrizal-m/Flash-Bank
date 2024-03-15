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
