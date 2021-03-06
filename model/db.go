package model

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "goBlog/modules/log"
    "time"
)

var db *gorm.DB

type dbConnection struct {
	Name     string `toml:"name"`
	UserName string `toml:"user_name"`
	Pwd      string `toml:"pwd"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
}

var conn = dbConnection {
    "goBlog",
    "root",
    "my7536308",
    "127.0.0.1",
    "3306",
}

// 存储表结构的数组，用于创建表
var dbTables = []interface {} {
    &User{},
    &Post{},
    &Comment{},
    &Tag{},
}

type DBModel struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func DB() *gorm.DB {

    if db == nil {
        log.Debugf("create new db")

        newDb, err := newDB(conn)

        
        if err != nil {
            log.Debugf("create error")
            panic(err)
        }
        
        if newDb == nil {
            panic("newDb is nil")
        }
        log.Debugf("db is created")
        
        
        
        // TODO: gorm.Logger不确定是否可用
        // newDb.SetLogger(gorm.Logger{})
        newDb.LogMode(true)
        newDb.DB().SetMaxIdleConns(10)
        newDb.DB().SetMaxOpenConns(100)

        
        db = newDb

        _db := &mydb{db}
        _db.createTables(dbTables)
    }
    
    // 创建
    return db
}

func newDB(conn dbConnection) (*gorm.DB, error) {
    var err error
    sqlConnection := conn.UserName + ":" + conn.Pwd + "@tcp(" + conn.Host + ":" + conn.Port + ")/" + conn.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
    
    db, err := gorm.Open("mysql", sqlConnection)
    db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")

    // for _, table := range dbTables {
    //     db.AutoMigrate(table)
    // }

    if err != nil {
        if err != nil {
            log.Debugf("newDB error")
        }
        return nil, err
    }

    return db, nil
}


// 主要是给gorm.DB添加新的方法
type mydb struct {
    *gorm.DB
}

func (db *mydb) createTables(tables []interface{}) {

    if db == nil {
        return
    }

    for _, table := range tables {
        if db.HasTable(table) == false {
            log.Debugf("table %T dose not exists, create it \n", table)
            // 设置默认编码格式为utf8
            db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(table)
            } else {
                log.Debugf("%T table has created \n", table)
        }
    }
}