package main

import (
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	//db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13306)/webook"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}

	// 迁移 schema
	//db.AutoMigrate(&User{})

	// Create
	//db.Create(&Product{Code: "D42", Price: 100})

	// Read
	//var product Product
	//db.First(&product, 1)                 // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	//// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	//// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - 删除 product
	//db.Delete(&product, 1)

	// TODO: query
	//var user User
	//if err := db.First(&user, 1).Error; err != nil {
	//	println(err)
	//	return
	//}
	//println(user.Email)

	getOSS("94043c3b-e059-4f59-a207-ca0526b9af01", "94043c3b-e14d-448b-a17d-c7de875a3758", "9407e414-a0f4-4eb3-a1df-052a4bdfe87e")
}

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	NickName     string `json:"nick_name"`    // 昵称
	Birthday     string `json:"birthday"`     // 生日
	Introduction string `json:"introduction"` // 简介
}

// todo: model

func getOSS(bid, brid, sid string) {
	url := fmt.Sprintf("https://menus-omp-nm-tmp.oss-cn-beijing.aliyuncs.com/menus/%v/%v/%v/takeoutsoldout.json?t=%v", bid, brid, sid, time.Now().Unix())
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.StatusCode)
		return
	}

	// Defer closing the response body
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}
