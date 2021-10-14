package database

/**
init function will be ACCESS FIRST AND ALWAYS ACCESS if package access
usually used for comunicate with the database, to open a connection to the database
 */
var connection string

func init() { // first call
	connection = "MySQL"
}

func GetDatabaseConnection() string{
	return connection
}