package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NikhilSharmaWe/scribblifly/pkg/storage"
	"github.com/joho/godotenv"
)

type application struct {
	infoLog      *log.Logger
	errLog       *log.Logger
	accountModel storage.AccountStorage
	scriptModel  storage.ScriptStorage
}

func main() {
	err := godotenv.Load("vars.env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := storage.NewPostgresStore()
	if err != nil {
		errLog.Fatal(err)
	}

	err = storage.Init(db)
	if err != nil {
		errLog.Fatal(err)
	}

	accountStorage := storage.AccountModel{DB: db}
	scriptStorage := storage.ScriptModel{DB: db}
	app := application{
		infoLog:      infoLog,
		errLog:       errLog,
		accountModel: &accountStorage,
		scriptModel:  &scriptStorage,
	}

	acc := storage.Account{
		Username:  "nikhilsharmawe",
		FirstName: "Nikhil",
		LastName:  "Sharma",
	}

	// try inserting some data in the databases
	if err := app.accountModel.Create(acc); err != nil {
		errLog.Fatal(err)
	}

	fmt.Println("Finished")
}
