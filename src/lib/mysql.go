// connect mysql and curd
package lib
import (
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strconv"
)

type MyDb struct{
	*sql.DB
}

//connect funciton
func ConnectMySQL(user string, pwd string, ip string, port int, dbName string) (*MyDb, error){
	portString := strconv.Itoa(port)
	connString := user + ":" +pwd + "@tcp(" +ip+":"+portString+")/"+dbName+"?charset=utf8"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Println("connect failed for connString:", connString)
		fmt.Println(err)
		return nil, err
	}
	mydb := &MyDb{DB: db}
	fmt.Println("connect successfully for connString:", connString)
	return mydb, nil
}

//query
func (this *MyDb) Query(query string, args ...interface{}) (*sql.Rows, error) {
	r, err := this.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return r, nil
}

//excute (insert,update,delete)
func (this *MyDb) Excute(strSql string, args ...interface{}) (int64, error) {

	result, err := this.Exec(strSql, args...)
	if err != nil {
		return 0, err
	}
	var affNum int64
	affNum, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affNum, nil
}

