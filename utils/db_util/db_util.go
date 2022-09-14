package db_util

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

//定义全局的db对象，我们执行数据库操作主要通过他实现。
var _db *gorm.DB

//包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
func init() {
	//配置MySQL连接参数
	username := "root"  //账号
	password := "dchat.db" //密码
	host := "192.168.1.231" //数据库地址，可以是Ip或者域名
	port := 7002 //数据库端口
	Dbname := "otc" //数据库名
	timeout := "30s" //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Printf("testConnPoll db=%v, err=%v\n", _db, err)
	if err != nil {
		fmt.Println("数据库连接失败！")
	}

	//_db.Set("gorm:table_options", "ENGINE=InnoDB")
	sqlDB, _ := _db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Hour)
	//defer sqlDB.Close()
}

//获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
//不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetDB() *gorm.DB {
	return _db
}
