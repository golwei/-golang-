package main

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User `orm:"rel(fk)"` //设置一对多关系
	//	Tags  []*Tag `orm:"rel(m2m)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Profile), new(Post))
}

//===================
package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:654654@tcp(127.0.0.1:3306)/test?charset=utf8")
}

func main() {
	// orm.RegisterModel...
	// orm.RegisterDataBase...

	orm.RunCommand()
}




