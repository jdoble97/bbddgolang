package main

import (
	"fmt"

	"./crud"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := crud.ObtenerConexion()
	if err != nil {
		fmt.Println("Error en la base de datos:", err)
	}
	defer db.Close() //Con esto cerramos la conexion cuando se termine la función

	err = db.Ping() //Comprobamos si hay conexión
	if err != nil {
		fmt.Println("Error conectando a la bbdd:", err)
	}
	fmt.Println("Conexión exitosa...")
	contacto := crud.Contacto{
		Nombre:    "Jorge",
		Direccion: "Calle Jaccobinia",
		Correo:    "jorgegonzalezw97@gmail.com",
	}

}
