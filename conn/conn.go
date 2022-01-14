package conn

import (
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	createDataBase = "sql/createDataBase.sql"
	createTable = "sql/createTable.sql"
	getShortUrl = "sql/getShortUrl.sql"
	insertUrl = "sql/insertUrl.sql"
)

var (
	CreateBD string
	CreateTable string
	GetShortUrl string
	InsertUrl string
)

var Memory = "postgres"
var DB *sqlx.DB
var Data map[string]string

func createDB() {
	cmd := exec.Command("psql", "-d", "database_name", "-f", "path/to/database.sql")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Panicln(err)
	}

	if err := cmd.Start(); err != nil {
		log.Panicln(err)
	}

	errout, _ := ioutil.ReadAll(stderr)
	if err := cmd.Wait(); err != nil {
		log.Println(errout)
		log.Panicln(err)
	}
}

func init() {
	data, err := ioutil.ReadFile(createDataBase)
	if err != nil {
		log.Fatalln(err)
	}
	CreateBD = string(data);

	data, err = ioutil.ReadFile(createTable)
	if err != nil {
		log.Fatalln(err)
	}
	CreateTable = string(data);

	data, err = ioutil.ReadFile(getShortUrl)
	if err != nil {
		log.Fatalln(err)
	}
	GetShortUrl = string(data);

	data, err = ioutil.ReadFile(insertUrl)
	if err != nil {
		log.Fatalln(err)
	}
	InsertUrl = string(data);

	//createDB()
	DB, err = sqlx.Open("postgres", "user=url_shortened password=url_shortened dbname=url_shortened sslmode=disable");
	if err != nil {
		log.Panicln(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Panicln(err)
	}
	// DB.MustExec(CreateTable)
}