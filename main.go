package main

import (
	"fmt"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var (
	UserName string = "root"
	PassWord string = "password"
	IP       string = "127.0.0.1"
	Port     int    = 3306
	DbName   string = "gorm"
	CharSet  string = "utf8mb4"

	db  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", UserName, PassWord, IP, Port, DbName, CharSet)
	db, err = gorm.Open(gmysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalln("连接mysql失败: ", err)
	}
}

func main() {
	Init()
	db = db.Debug()
	//insert()
	//selectTable()
	//update()
	//deleteTable()
}

func deleteTable() {
	// 删除id为24的记录
	//db.Delete(models.User{Id: 24})
	//DELETE FROM `user` WHERE `user`.`id` = 24

	// 删除id为24，爱好为java,邮箱是cd9b@63e2.com.cn的记录
	//db.Where("Hobby = ?", "java").Where("Email = ?","cd9b@63e2.com.cn").Delete(models.User{Id: 24})
	//DELETE FROM `user` WHERE Hobby = 'java' AND Email = 'cd9b@63e2.com.cn' AND `user`.`id` = 24

	//根据主键删除
	// 删除id为1000的记录
	//db.Delete(&models.User{}, 1000)
	// DELETE FROM `user` WHERE `user`.`id` = 1000

	// 删除id为1000，1002，1003的记录
	//var users []models.User
	//db.Delete(&users, []int{1000, 1002, 1003})
	//DELETE FROM `user` WHERE `user`.`id` IN (1000,1002,1003)
}

func update() {
	// 把id大于103的记录 gender改为男
	//db.Model(&models.User{}).Where("id > ?", 103).Update("Hobby", "男")
	// UPDATE `user` SET `gender`='男' WHERE id > 103

	// 把id是106的 hobby 改为 apple
	//db.Model(&models.User{Id: 106}).Update("Hobby", "apple")
	//UPDATE `user` SET `hobby`='apple' WHERE `id` = 106

	//把所有Id为29，Gender为男的Email 改为 66544@souhu.com
	//db.Model(&models.User{Id: 29}).Where("Gender = ?", "男").Update("Email", "66544@souhu.com")
	// UPDATE `user` SET `email`='66544@souhu.com' WHERE Gender = '男' AND `id` = 29

	// 更新多列，把id为107的记录,Gender改为男，Email改为66544@souhu.com
	//db.Model(&models.User{Id: 107}).Updates(&models.User{Gender: "男", Email: "66544@souhu.com"})
	//UPDATE `user` SET `gender`='男',`email`='66544@souhu.com' WHERE `id` = 107

	// 只更新固定字段
	// 更新id为108的记录的Email和Hobby,别的字段会被忽略
	//db.Model(&models.User{Id: 108}).Select("Email", "Hobby").Updates(models.User{
	//	Id:     0,
	//	Name:   "神荼",
	//	Gender: "女",
	//	Hobby:  "java",
	//	Email:  "8787@4ocad.vip",
	//})
	//UPDATE `user` SET `hobby`='java',`email`='8787@4ocad.vip' WHERE `id` = 108

	// 忽略更新字段
	// 更新id为107的所有信息
	//db.Model(&models.User{Id: 107}).Omit("Name").Updates(models.User{
	//	Id:     0,
	//	Name:   "神荼",
	//	Gender: "女",
	//	Hobby:  "云南信息工程大学",
	//	Email:  "8d4a@njhs0.xyz",
	//})
	// UPDATE `user` SET `gender`='女',`hobby`='云南信息工程大学',`email`='8d4a@njhs0.xyz' WHERE `id` = 107
}

func selectTable() {
	//user := models.User{}
	// 获取第一条数据，主键升序
	//db.First(&user)
	//SELECT * FROM `user` ORDER BY `user`.`id` LIMIT 1

	// 获取一条记录，没有指定排序字段
	//db.Take(&user)
	//SELECT * FROM `user` LIMIT 1

	// 获取最后一条记录（主键降序）
	//db.Last(&user)
	//SELECT * FROM `user` ORDER BY `user`.`id` DESC LIMIT 1

	// 查询主键为10的记录
	//db.First(&user, 10)
	//SELECT * FROM `user` WHERE `user`.`id` = 10 ORDER BY `user`.`id` LIMIT 1

	//log.Println(user)

	//var users []models.User

	// find查询所有id为10，20，33的记录
	//db.Find(&users, []int{10, 20, 33})
	//SELECT * FROM `user` WHERE `user`.`id` IN (10,20,33)

	// 获取全部记录
	//db.Find(&users)
	//SELECT * FROM `user`

	// 获取所有hobby为apple的用户
	//db.Where("hobby = ?","apple").Find(&users)
	//SELECT * FROM `user` WHERE hobby = 'apple'

	// find查询所有name为 阿卡丽，思凡，向安琪的记录
	//db.Where("name in ?", []string{"阿卡丽", "思凡", "向安琪"}).Find(&users)
	//SELECT * FROM `user` WHERE name in ('阿卡丽','思凡','向安琪')

	// Map多条件查询
	//db.Where(&models.User{Gender: "女", Hobby: "apple"}).Find(&users)
	//SELECT * FROM `user` WHERE `user`.`gender` = '女' AND `user`.`hobby` = 'apple'

	// 传入指定查询字段,只会查询 Hobby字段相符合的记录，忽视ID
	//db.Where(&models.User{Hobby: "apple", Id: 12}, "Hobby").Find(&users)
	//SELECT * FROM `user` WHERE `user`.`hobby` = 'apple'

	// 内联查询
	//查询所有 Hobby是apple的记录
	//db.Find(&users, "Hobby = ?", "apple")
	//SELECT * FROM `user` WHERE Hobby = 'apple'

	//db.Find(&users, models.User{Hobby: "apple"})
	//SELECT * FROM `user` WHERE `user`.`hobby` = 'apple'

	// Not用法基本同where
	// 查询所有Hobby不是apple的
	//db.Not("Hobby = ?", "apple").Find(&users)
	// SELECT * FROM `user` WHERE NOT Hobby = 'apple'

	// Or条件拼接
	// 寻找 Hobby是apple 或者 Gender是女的记录
	//db.Where("Hobby = ?", "apple").Or("Gender =  ?", "女").Find(&users)
	//SELECT * FROM `user` WHERE Hobby = 'apple' OR Gender =  '女'

	// 查询Name 与 Hobby
	//db.Select("Name", "Hobby").Find(&users)
	//SELECT `name`,`hobby` FROM `user`

	// Order排序
	// 按照姓名倒序
	//db.Order("name desc").Find(&users)
	//SELECT * FROM `user` ORDER BY name desc

	// 按照姓名和id排序，倒序
	//db.Order("name desc").Order("id desc").Find(&users)
	//
	//for _, user := range users {
	//	log.Println(user)
	//}

}

func insert() {
	//user := models.User{Hobby: "唱歌", Email: "26aa@5g7au.cn", Name: "思凡", Gender: "男"}
	// 依照全部字段创建
	//db.Create(&user)
	//INSERT INTO `user` (`name`,`gender`,`hobby`,`email`) VALUES ('思凡','男','唱歌','26aa@5g7au.cn')

	// 选择指定字段创建
	//db.Select("Name", "Email").Create(&user)
	//INSERT INTO `user` (`name`,`email`) VALUES ('思凡','26aa@5g7au.cn')

	// 略传递给略去的字段值
	//db.Omit("Name", "Email").Create(&user)
	// INSERT INTO `user` (`gender`,`hobby`) VALUES ('男','唱歌')

	// 声明一个数组测试批量插入
	//var users = []models.User{{Name: "张三"}, {Name: "李四"}, {Name: "王五"}}
	//db.Create(&users)
	//INSERT INTO `user` (`name`,`gender`,`hobby`,`email`) VALUES ('张三','','',''),('李四','','',''),('王五','','','')

	// 打印出所有的主键值
	//for _, user := range users {
	//	log.Println(user.Id)
	//}

}
