//just a class for exercise to test oob
//需要暴露给外部包的成员变量或方法，需要大写字母开头
package lib
import "fmt"
type Student struct{
	name string
	Score string
}

func (s *Student) GetName()string {
	return s.name
}


func (s *Student) SetName(name string) {
	s.name = name
}

func (s *Student) Say() {
	fmt.Println("I'm a student")
}

type GirlStudent struct{
	Student
	height int32
}

func (s *GirlStudent) Say() {
	fmt.Println("I'm a girl student")
}

type Human interface{
	Say()
}

type Teacher struct{
	name string
}
func (t *Teacher) Say() {
	fmt.Println("I'm a Teacher")
}

//
func HumanFactory(name string) Human {
    switch name {
    case "student":
        return &Student{name: "luodao", Score:"30"}
    case "teacher":
        return &Teacher{name: "test"}
    default:
        panic("No such human")
    }
}