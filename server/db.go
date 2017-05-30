package server

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"time"
)

type chat struct {
	Id      int    `pk:"auto"`
	User    string `orm:"size(32)"`
	Content string `orm:"size(1024)"`
	Time    time.Time
}

func InitDb(conf *viper.Viper) {
	conn := conf.GetString("user") + ":" + conf.GetString("pwd") + "@/" + conf.GetString("dbname") + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterModel(new(chat))
	orm.RegisterDataBase("default", "mysql", conn)
	orm.RunSyncdb("default", false, false)
	dbconn := orm.NewOrm()
	_, err := dbconn.Raw("alter table chat  convert to character set utf8").Exec()
	if err != nil {
		log.Println("转编码失败:", err)
	}

}

func save2db(user, content string) {
	c := new(chat)
	c.Content = content
	c.User = user
	c.Time = time.Now()
	dbconn := orm.NewOrm()
	_, err := dbconn.Insert(c)
	if err != nil {
		log.Println("插入数据出错:", err)
	}
}
