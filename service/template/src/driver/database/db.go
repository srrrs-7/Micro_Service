package database

type DatabaseInterface interface {
	dbConnect()
}

type Database struct {
}

func NewDatabase() DatabaseInterface {
	return &Database{}
}

func (d *Database) dbConnect() {

}
