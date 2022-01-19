package conn

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	createTable = "sql/createTable.sql"
	getShortUrl = "sql/getShortUrl.sql"
	getOriginalUrl = "sql/getOriginalUrl.sql"
	insertUrl = "sql/insertUrl.sql"
	checkShortUrl = "sql/checkShortUrl.sql"
)

var (
	CreateTable string
	GetShortUrl string
	GetOriginalUrl string
	InsertUrl string
	CheckShortUrl string

	ServAddr string
	Memory string
)

var DB *sqlx.DB
var Data map[string]string

func init() {
	Data = make(map[string]string)

	data, err := ioutil.ReadFile(createTable)
	if err != nil {
		log.Fatalln(err)
	}
	CreateTable = string(data);

	data, err = ioutil.ReadFile(getShortUrl)
	if err != nil {
		log.Fatalln(err)
	}
	GetShortUrl = string(data);

	data, err = ioutil.ReadFile(getOriginalUrl)
	if err != nil {
		log.Fatalln(err)
	}
	GetOriginalUrl = string(data);

	data, err = ioutil.ReadFile(insertUrl)
	if err != nil {
		log.Fatalln(err)
	}
	InsertUrl = string(data);

	data, err = ioutil.ReadFile(checkShortUrl)
	if err != nil {
		log.Fatalln(err)
	}
	CheckShortUrl = string(data);

	err = godotenv.Load("arg.env")

	if err != nil {
		log.Println(err)
	}

	ServAddr = os.Getenv("SERV_ADDR")
	if ServAddr == "" {
		ServAddr = "127.0.0.1:5000"
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("DB_USERNAME")
	if user == "" {
		user = "url_shortened"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "url_shortened"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "url_shortened"
	}

	Memory = os.Getenv("MEMORY")
	if Memory == "" {
		Memory = "in-memory"
	} else if Memory != "in-memory" && Memory != "postgres" {
		log.Panicln("service does not support " + Memory + " storage")
	}

	if Memory == "postgres" {
		DB, err = sqlx.Open("postgres", "host="+host+" port="+port+" user="+user+" password="+password+" dbname="+dbName+" sslmode=disable");
		if err != nil {
			log.Panicln(err)
		}
		err = DB.Ping()
		if err != nil {
			log.Panicln(err)
		}
		DB.MustExec(CreateTable)
	}
}