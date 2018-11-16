package main

import (
    "fmt"
    "lib"
    "os"
    "path/filepath"
)
func main() {
    fmt.Println("Hello, World!")
    var a int = 12
    var b int = 13
    var c = &b
    ret := lib.Add(a,b)
    b = ret
    fmt.Println("b:",b,"c:",c)

    if (ret < 100 && b > 10 ) {
        fmt.Println("ret:", ret)
    }
    //定义数组
    var arr = [...]int{10,100,200}
    //循环打印数组元素
    for i:=0; i< len(arr); i++ {
        fmt.Println("arr[i]:",arr[i])
    }

    for i, num := range arr {
        fmt.Println("i:",i,"num:",num)
    }

    //传值和指针使用
    var x int = 10
    var y int = 20
    lib.Swap(&x, &y)
    fmt.Println("x:",x,"y:",y)

    var ptr *int
    if(ptr == nil){
        fmt.Println("is nil,ptr", ptr)
    }else{
        fmt.Println("not nil,ptr:", ptr)
    }
    

    //切片
    sliceArr := make([]string,3,5)
    fmt.Println("sliceArr:", sliceArr)
    fmt.Println("sliceArr,len:", len(sliceArr),"cap:",cap(sliceArr))

    //map
    nameMap := map[string]int32{"name":12,"age":1}
    for k,v := range nameMap {
        fmt.Println("k:",k,"v:",v)
    }
    var path string = getCurrentPath()
    fmt.Println("path", path)
    lib.ReadFromFile(path + "/lm.txt")

    lib.WriteToFile(path + "/lm.txt", "ok!\n")

    //mysql
    
    db, err := lib.ConnectMySQL("user", "password", "127.0.0.1", 3306, "testdb")

	if err != nil || db == nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return
	}
	defer db.Close()
 
	//查询数据，指定字段名，返回sql.Rows结果集
    rows, queryErr := db.Query("select id,name from tb_admin")
    if queryErr != nil {
		fmt.Println("Query时出错:", queryErr.Error())
		return
    }
    
	id := 0
	name := ""
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, name)
    }
    defer rows.Close()
    

    //oob encapsulation
    studentObj := lib.Student{ Score:"31"}
    fmt.Println("studentObj", studentObj)
    fmt.Println("studentObj.Score", studentObj.Score)
    fmt.Println("studentObj.name:", studentObj.GetName())
    studentObj.SetName("test")
    fmt.Println("studentObj.name:", studentObj.GetName())
    
    //oob inherit
    girlStudent := lib.GirlStudent{ }
    girlStudent.SetName("girl")
    girlStudent.Score = "abc"
    fmt.Println("girlStudent", girlStudent)
    girlStudent.Say()

    //oob Polymorphisn
    human := lib.HumanFactory("teacher")
    human.Say()
    human = lib.HumanFactory("student")
    human.Say()
}

func getCurrentPath() string {
    fmt.Println("os.Args[0]:", os.Args[0])
    s, err := filepath.Abs(filepath.Dir(os.Args[0]))
    fmt.Println("s:", s)
    checkErr(err)
    return s
}
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}