package main

import (
	"bufio"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
)

type Monst666 struct {
	Name   string
	Age    int
	Skills string
}

func main003() {
	con, _ := redis.Dial("tcp", "api.ikanned.com:16379")
	defer con.Close()
	reader := bufio.NewReader(os.Stdin)
	var m []string
	var line string
	var j int = 0
	for {
		println("请输入姓名-年龄-技能")
		line, _ = reader.ReadString('\n')
		//line = strings.Trim(line, "\n")
		if line == "exit\n" {
			break
		}
		m = make([]string, 3)
		m[j] = line

	}

	con.Do("HMSet", "monster", "name", m[0], "age", m[1], "skill", m[2])
	res, _ := redis.Strings(con.Do("HMGet", "monster", "name", "age", "skill"))
	fmt.Println("monster:=", res)
}

func main() {
	conn, _ := redis.Dial("tcp", "ikanned.com:16379")

	var monsSlice = make([]Monst666, 0)
	for i := 0; i < 3; i++ {
		var name, skills string
		var age int
		fmt.Print("Name >")
		fmt.Scanf("%s\n", &name)
		fmt.Print("Age >")
		fmt.Scanf("%d\n", &age)
		fmt.Print("Skills >")
		fmt.Scanf("%s\n", &skills)
		monsSlice = append(monsSlice, Monst666{
			Name:   name,
			Age:    age,
			Skills: skills,
		})
	}
	for _, i := range monsSlice {
		insertToRedis(conn, i.Name, i.Age, i.Skills)
		fmt.Println(i)
	}
}

func insertToRedis(conn redis.Conn, name string, age int, skills string) {
	conn.Do("hSet", "Monster_"+name, "Name", name, "Age", age, "Skills", skills)
}
