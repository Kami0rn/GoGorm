package main

import (
	"context"
	// "crypto/rand"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql,_ := fc()
	fmt.Printf("%v\n====================================\n",sql)
}

var db *gorm.DB

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/tew?parseTime=true"
	dial := mysql.Open(dsn)

	var err error
	db,err = gorm.Open(dial , &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}
	
	// CreateGender("xxxx")
	//db.Migrator().AutoMigrate(Gender{},Test{},Customer{})
	// GetGenders()
	// GetGender(1)
	// GetGenderByName("Female")
	// UpdateGender()
	// UpdateGender2(3,"zzzz")
	// DeleteGender(4)\
	// CreateTest(1,"lol1")
	// CreateTest(2,"lol2")
	// CreateTest(3,"lol3")
	// DeleteTest(2)
	// GetTest()
	// CreateCustomer("Jane" , 2)
	GetCustomers()
	// db.Migrator().CreateTable(Customer{})

}

func CreateCustomer (name string , genderID uint) {
	customer := Customer{
		Name : name ,
		GenderID: genderID,
	}
	tx := db.Create(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(customer)

}

type Customer struct {
	ID uint
	Name string
	Gender Gender
	GenderID uint
}

func GetCustomers() {
	customers := []Customer{}
	tx := db.Preload(clause.Associations).Find(&customers)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	for _, customer := range customers {
		fmt.Printf("%v|%v|%v\n",customer.ID,customer.Name,customer.Gender.Name)
	}
}

func CreateTest (code uint , name string){
	test := Test{Code: code , Name: name}
	db.Create(&test)
}
func GetTest (){
	tests := []Test{}
	db.Find(&tests)
	for _,t := range tests {
		fmt.Println("%v|%v" , t.ID , t.Name)
	}
}
func DeleteTest (id uint){
	db.Delete(&Test{} , id)
}

func PermaDeleteTest (id uint){
	db.Unscoped().Delete(&Test{} , id)
}
func DeleteGender (id uint) {
	tx := db.Delete(Gender{}, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println("Deleted")
	GetGender(id)
}

func UpdateGender(id uint , name string) {
	gender := Gender{}
	tx := db.First(&gender , id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	gender.Name = name 
	db.Save(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}



func UpdateGender2(id uint , name string) {
	gender := Gender{Name: name}
	tx := db.Model(&Gender{}).Where("id=?",id).Updates(gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	} 
	GetGender(id)
}

func GetGenders() {
	genders := []Gender{}
	tx := db.Order("id").Find(&genders)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

func GetGender(id uint) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func GetGenderByName (name string) {
	gender := Gender{}
	tx := db.First(&gender, "name=?" , name)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func CreateGender (name string){
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

type Gender struct {
	ID uint
	Name string `gorm:"unique;size(10)"`
}

type Test struct {
	gorm.Model
	Code uint `gorm:"comment:This is code"`
	Name string `gorm:"column:myname;size:20;unique;default:Hello;not null"`
	
}

func (t Test) TableName() string {
	return "MyTest"
}