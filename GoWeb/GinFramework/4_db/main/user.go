package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goPro/GoWeb/GinFramework/4_db/utily"
)

type User struct {
	Id   int
	Name string
	Addr string
}

// 添加user方法一
func (user *User) AddUser() error {
	sqlStr := "insert into user1(Id,Name,Addr)values(?,?,?) "
	inStmt, err := utily.Db.Prepare(sqlStr) //inStmt有执行和查询工作，是准备好的状态
	if err != nil {
		fmt.Println("预编译出现异常:", err)
		return err
	}
	_, err = inStmt.Exec("1", "admin", "nanjin")
	if err != nil {
		fmt.Println("执行出现异常:", err)
		return err
	}
	return nil
}

// 添加user方法二
func (user *User) AddUser2() error {
	sqlStr := "insert into user1(Id,Name,Addr)values(?,?,?) "
	_, err := utily.Db.Exec(sqlStr, "2", "admin2", "nanjin")
	if err != nil {
		fmt.Println("执行出现异常:", err)
		return err
	}
	return nil
}

func (user *User) GetUserById() (*User, error) {
	sqlStr := "select Id,Name,Addr from user1 where id=?"
	row := utily.Db.QueryRow(sqlStr, user.Id)
	var id int
	var name string
	var addr string
	err := row.Scan(&id, &name, &addr)
	if err != nil {
		return nil, err
	}
	u := &User{
		Id:   id,
		Name: name,
		Addr: addr,
	}
	return u, nil
}

// 查询数据库所有记录
func (user *User) GetUsers() ([]*User, error) {
	sqlStr := "select Id,Name,Addr from user1"
	rows, err := utily.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var users []*User
	for rows.Next() {
		var id int
		var name string
		var addr string
		err := rows.Scan(&id, &name, &addr)
		if err != nil {
			return nil, err
		}
		u := &User{
			Id:   id,
			Name: name,
			Addr: addr,
		}
		users = append(users, u)
	}
	return users, nil
}

func main() {
	user := User{}
	user.AddUser()
	user.AddUser2()
}
