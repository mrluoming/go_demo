//封装文件读写
package lib
import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
//读文件，返回数组，数组的元素为一行
func ReadFromFile(fileName string) string {
	var result string = ""
	fmt.Println("fileName:", fileName)
	if contents,err := ioutil.ReadFile(fileName);err == nil {
        //因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
        result = string(contents)
        fmt.Println(result)
    }
	return result
}

//写文件，追加
func WriteToFile(fileName string,content string) bool {
	/*
	//使用ioutil.WriteFile会每次都清空，没法追加
	data :=  []byte(content)
    if ioutil.WriteFile(fileName,data,0644) == nil {
		fmt.Println("写入文件成功:",content)
		return true	
	}else{
		return false
	}
	*/

	fileObj,err := os.OpenFile(fileName,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
    if err != nil {
        fmt.Println("Failed to open the file",err.Error())
		os.Exit(2)
		return false
    }
    if  _,err := io.WriteString(fileObj,content);err == nil {
        fmt.Println("Successful appending the data:",content)
    }
	return true
}