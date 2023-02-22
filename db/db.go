package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gopkg.in/ini.v1"
)

func establishConnection() *sql.DB {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf(cfg.Section("cockroach").Key("URL").String(), "fundd"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successful Connection")

	return db
}

func ReadFromTable(table string) {
	db := establishConnection()
	sql := fmt.Sprintf("SELECT * FROM %s;", table)
	db.Exec(sql)

}

func writeToTable(newUser user, table string) string {
	db := establishConnection()
	defer db.Close()

	sql := fmt.Sprintf("UPSERT into %s (name, password, email) values (\"%s\",\"%x\",\"%s\");", table, newUser.Name, newUser.Password, newUser.Email)

	if _, err := db.Exec(sql); err != nil {
		log.Fatalf("This failed because of the following %v ", err)
	}
	return "User Created Successfully"
}
