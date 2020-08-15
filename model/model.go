package model

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open("mysql",
		"root:15219331409@/stu?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&Student{})
	checkError(err)
}

type Student struct {
	ID       int    `form:"id" binding:"required"`
	Name     string `form:"name"`
	Major    string `form:"major"`
	Sex      int    `form:"sex"`
	Birthday string `form:"birthday"`
	Socre    int    `form:"socre"`
	Note     string
}

type Accout struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}

type StudentSlice struct {
	Students []Student
}

func NewStu(id int, name string, major string,
	sex int, birthday string, socre int, note string) Student {
	return Student{
		ID:       id,
		Name:     name,
		Major:    major,
		Sex:      sex,
		Birthday: birthday,
		Socre:    socre,
		Note:     note,
	}
}

func NewStuSlice() *StudentSlice {
	return &StudentSlice{}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Counts() int {
	var count int
	db.Model(&Student{}).Count(&count)
	return count
}

func (this *StudentSlice) List(i int) []Student {
	db.Offset(i * 10).Limit(10).Find(&this.Students)
	//b, _ := json.Marshal(this.Students)
	return this.Students
}

func (this *StudentSlice) Search(id int) []Student {
	db.Where("ID = ?", id).Find(&this.Students)
	return this.Students
}

func (this *StudentSlice) Add(stu Student) error {
	if err := db.Create(&stu).Error; err != nil {
		return err
	} else {
		return nil
	}
}
func (this *StudentSlice) FindIndex(id int) error {
	for _, val := range this.Students {
		if val.ID == id {
			return nil
		}
	}
	return errors.New("查询的编号不存在")
}

func (this *StudentSlice) Modify(stu Student) {
	db.Model(Student{}).Where(" ID = ?",
		stu.ID).Updates(Student{Name: stu.Name, Major: stu.Major,
		Sex: stu.Sex, Birthday: stu.Birthday, Socre: stu.Socre, Note: stu.Note})
}

func (this *StudentSlice) Delete(id int) error {
	if err := db.Where("ID = ?", id).Delete(Student{}).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (this *StudentSlice) Sort() []Student {
	db.Order("Socre").Find(&this.Students)
	return this.Students
}
