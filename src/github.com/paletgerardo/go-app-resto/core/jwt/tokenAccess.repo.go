package jwt

import (
	"database/sql"
	"fmt"
	"github.com/paletgerardo/go-app-resto/core/db"
	"github.com/paletgerardo/go-app-resto/core/structs"
)

func AcaSeBuscaElUsuario(email string) (structs.UserLogin, bool, int) {
	var usuario structs.UserLogin

	queryString := `select * from usuarios where email = $1`

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	row := connectionDB.QueryRow(queryString, email)
	errorAlParsearDatos := row.Scan(&usuario.Id, &usuario.Nombre, &usuario.Apellido, &usuario.Email, &usuario.Password, &usuario.FechaNacimiento, &usuario.Created, &usuario.Updated)

	if errorAlParsearDatos != nil && errorAlParsearDatos != sql.ErrNoRows {
		fmt.Println(errorAlParsearDatos.Error())
		return usuario, false, -1
	}

	return usuario, true, usuario.Id
}
