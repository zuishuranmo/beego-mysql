package db_mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "go-sql-driver/mysql"
	"sgwe/models"
)

var Db *sql.DB

func init()  {
	fmt.Println("连接mysql数据库")
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIP := config.String("db_ip")
	dbName := config.String("db_name")
	dbCharset := config.String("db_charset")

	connUrl := dbUser+":"+dbPassword+"@tcp"+ dbIP +dbName+dbCharset
	fmt.Println(connUrl)
	DB,err :=sql.Open(dbDriver,connUrl)
	if  err != nil{
		panic("数据库连接失败，请查看配置")
	}
	fmt.Println("数据库连接成功")
	Db = DB
}

func InsertUser(user models.User)(int64,error)  {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(user.Password))
	bytes := hashMd5.Sum(nil)
	user.Password = hex.EncodeToString(bytes)
	fmt.Println("将要保存的用户名:",user.Nick,"密码:",user.Password,"名字：",user.Name,"地址：",user.Address,"生日：",user.Birthday)
	result,err := Db.Exec("insert into user(nick, password, name, address, birthday) values(?,?,?,?,?)",user.Nick,user.Password,user.Name,user.Address,user.Birthday)
	if err != nil {//保存数据时遇到错误
		return -1,err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return -1,err
	}
	return id,nil
}