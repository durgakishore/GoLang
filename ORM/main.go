package main

import (
	_ "ORM/routers"
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID      int
	Name    string   `orm:"size(100)"`
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
	ID   int
	Age  int16
	User *User `orm:"reverse(one)"` // Reverse relationship
}

const (
	dbAlias       = "company"
	mysqlUser     = "rapiscan"
	mysqlPassword = "Anubis2830"
	mysqlHost     = "127.0.0.1"
	mysqlPort     = 3306
	mysqlDatabase = "SysDb"
	mysqlCharset  = "utf8"
)

var (
	mysqlCon = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		mysqlUser,
		mysqlPassword,
		mysqlHost,
		mysqlPort,
		mysqlDatabase,
		mysqlCharset,
	)
)

func init() {

	var err error

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterModel(new(User), new(Profile))

	// set timezone
	orm.DefaultTimeLoc = time.UTC

	err = orm.RegisterDataBase("company", "mysql", mysqlCon)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
}

func main() {
	beego.Run()
}
