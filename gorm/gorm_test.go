package gorm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestInitialDBClient(t *testing.T) {
	err := InitialDBClient("sillyhat:sillyhat@tcp(127.0.0.1:3308)/sillyhat", true)
	assert.Nil(t, err)
}

type Userinfo struct {
	Id               int64
	Name             string
	Age              int
	Birthday         time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Description      string
	IsDelete         int       `gorm:"column:is_delete"`
	CreatedDate      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	LastModifiedDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (u Userinfo) String() string {
	return fmt.Sprintf(`{ Id : %d, Name : %s, Age : %d, Birthday : %s, Description : %s, IsDelete : %d, CreatedDate : %s, LastModifiedDate : %s}`, u.Id, u.Name, u.Age, u.Birthday, u.Description, u.IsDelete, u.CreatedDate, u.LastModifiedDate)
}

func (Userinfo) TableName() string {
	return "userinfo"
}

func TestFindAllUserInfo(t *testing.T) {
	err := InitialDBClient("sillyhat:sillyhat@tcp(127.0.0.1:3308)/sillyhat", true)
	assert.Nil(t, err)
	var users []Userinfo
	db, err := Client.GetDBClient()
	assert.Nil(t, err)
	defer db.Close()
	db.Find(&users)
	//Client.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	assert.EqualValues(t, len(users), 981)
}

func TestInsert(t *testing.T) {
	err := InitialDBClient("sillyhat:sillyhat@tcp(127.0.0.1:3308)/sillyhat", true)
	assert.Nil(t, err)
	user := Userinfo{Name: "haha"}
	db, err := Client.GetDBClient()
	assert.Nil(t, err)
	defer db.Close()
	db.Create(&user)
	//Client.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
}

func TestFindFirst(t *testing.T) {
	err := InitialDBClient("sillyhat:sillyhat@tcp(127.0.0.1:3308)/sillyhat?parseTime=True", true)
	assert.Nil(t, err)
	var user Userinfo
	db, err := Client.GetDBClient()
	assert.Nil(t, err)
	defer db.Close()
	db.Select([]string{"id", "name", "age", "TIMESTAMP(birthday) birthday", "description", "(is_delete = b'1')  is_delete", "created_date", "last_modified_date"}).First(&user)
	log.Println(user)
	//Client.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	assert.EqualValues(t, user.Id, 1)
	assert.EqualValues(t, user.Name, "test name")
	assert.EqualValues(t, user.Age, 21)
	assert.EqualValues(t, user.Description, "This is description")
	assert.EqualValues(t, user.IsDelete, 1)
}
