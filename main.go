package main

import (
	"database/sql"
	"fmt"
)

func obtenerConexion() (db *sql.DB, er error) {
	usuario := "root"
	pass := ""
	host := "tcp(127.0.0.1:3306)"
	nombrebbdd := "ejemplo"
	//usuario:contraseña@host/nombreBaseDeDatos
	db, er = sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombrebbdd))
	if er != nil {
		return nil, er
	}
	return db, nil
}
func main() {
	db, err := obtenerConexion()
	if err != nil {
		fmt.Println("Error en la base de datos:", err)
	}
	defer db.Close() //Con esto cerramos la conexion cuando se termine la función

	err = db.Ping() //Comprobamos si hay conexión
	if err != nil {

	}

}
