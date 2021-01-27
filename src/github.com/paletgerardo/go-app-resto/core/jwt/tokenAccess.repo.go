package jwt

import (
	"github.com/paletgerardo/go-app-resto/core/db"
	"github.com/paletgerardo/go-app-resto/core/structs"
	"github.com/pkg/errors"
)

/*func AcaSeBuscaElUsuario(email string) (structs.Usuarios, int, error) {
	var usuario structs.Usuarios

	queryString := `select * from usuarios where email = $1`

	connectionDB := db.GetConnectionToDB()
	defer connectionDB.Close()

	row := connectionDB.QueryRow(queryString, email)
	errorAlParsearDatos := row.Scan(&usuario.Id, &usuario.Nombre, &usuario.Apellido, &usuario.Email, &usuario.Password, &usuario.FechaNacimiento, &usuario.Created_at, &usuario.Updated_at)

	if errorAlParsearDatos != nil && errorAlParsearDatos != sql.ErrNoRows {
		fmt.Println(errorAlParsearDatos.Error())
		return usuario, -1, errorAlParsearDatos
	}

	return usuario, usuario.Id, nil
}*/

func AcaSeBuscaElUsuario(email string) (structs.Usuarios, int, error) {
	var usuario structs.Usuarios

	gormConnection := db.GetGormConnection()

	result := gormConnection.Model(&structs.Usuarios{}).Where("email = ?", email).Find(&usuario)

	if result.Error != nil {
		return usuario, -1, errors.New("Error al buscar usuario login.")
	}

	return usuario, usuario.Id, nil
}
