package crud

//ObtenerConexion crea una conexión a la bbdd
func ObtenerConexion() (db *sql.DB, er error) {
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