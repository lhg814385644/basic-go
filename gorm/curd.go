package main

import (
	"fmt"
	"gorm.io/driver/mysql"
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
	//db, err := gorm.Open(mysql.Open("root:bb123456@tcp(localhost:3306)/wechatorders"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//// 迁移 schema
	//db.AutoMigrate(&OrdermealIShops{})

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

	// InitO2ODB()
	InitWeChatOrdersDB()
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

// OrdermealIShops represents the table `ordermeal_i_shops`
type OrdermealIShops struct {
	ID              int     `json:"id" comment:"门店id"`
	WLifeShopID     int64   `json:"wlife_shop_id" comment:"微生活门店ID"`
	WLifeCashierID  int64   `json:"wlife_cashier_id" comment:"微生活用来扫付的收银员ID"`
	AlipayStoreID   string  `json:"alipay_store_id" comment:"支付宝支付的时候传的store_id"`
	KoubeiShopID    string  `json:"koubei_shop_id" comment:"口碑店铺ID"`
	BrandID         int     `json:"brand_id" comment:"品牌id"`
	Name            string  `json:"name" comment:"名称"`
	Seq             int     `json:"seq" comment:"排序"`
	Status          int     `json:"status" comment:"状态"`
	CreateTime      string  `json:"create_time" comment:"创建时间"`
	UpdateTime      string  `json:"update_time" comment:"更新时间"`
	ShopKey         string  `json:"shopkey" comment:"门店密码"`
	ShopID          int     `json:"shop_id" comment:"总部对应门店id"`
	Partner         string  `json:"partner" comment:"微信门店id"`
	YazuoShopID     *int64  `json:"yazuo_shop_id" comment:"雅座门店id"`
	YazuoCashierID  *int64  `json:"yazuo_cashier_id" comment:"雅座收银员id"`
	IsOpenWxdc      *int    `json:"is_open_wxdc" comment:"是否开启微信点餐"`
	OldmemShopID    *string `json:"oldmem_shop_id" comment:"旧版微信门店id"`
	OldmemCashierID *string `json:"oldmem_cashier_id" comment:"旧版微信收银员id"`
	IsBigPic        uint16  `json:"is_big_pic" comment:"点餐大小图切换配置1默认大图2小图"`
	UpdateTimeUnix  int64   `json:"update_time_unix" comment:"更新时间Unix时间戳"`
	AreaID          int     `json:"area_id" comment:"区域id"`
	BusinessID      int     `json:"business_id" comment:"商户id"`
	Location        string  `json:"location" comment:"坐标经纬度"`
	Longitude       string  `json:"longitude" comment:"经度"`
	EcoID           *string `json:"eco_id" comment:"生态环境ID"`
	Latitude        string  `json:"latitude" comment:"纬度"`
	BusinessType    uint8   `json:"business_type" comment:"业务类型:1-自营外卖"`
	Bookstatus      *int    `json:"bookstatus" comment:"是否开启预点餐 1 开启"`
	Ordermode       *int    `json:"ordermode" comment:"点餐模式，2先付，3后付"`
	OnlineTime      *string `json:"online_time" comment:"营业时间"`
	OMPConf         []byte  `json:"omp_conf" comment:"omp门店配置"`
}

func InitWeChatOrdersDB() {
	db, err := gorm.Open(mysql.Open("root:bb123456@tcp(localhost:3306)/wechatorders"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&OrdermealIShops{})
	var shops []OrdermealIShops
	if err := db.Model(&OrdermealIShops{}).Select("id,wlife_shop_id,brand_id,partner,business_id").
		// Where("partner <> ?", "").
		Find(&shops).Error; err != nil {
		println(err)
		return
	}
	for _, shop := range shops {
		getOSS(shop.BusinessID, shop.BrandID, shop.ID)
	}
}

func InitO2ODB() {
	db, err := gorm.Open(mysql.Open("root:bb123456@tcp(localhost:3306)/o2o"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Shops{})

	var shops []Shops
	if err := db.Model(&Shops{}).Select("id,sname,mid,bid").
		// Where("partner <> ?", "").
		Find(&shops).Error; err != nil {
		println(err)
		return
	}
	for _, shop := range shops {
		getOSSShops(shop.Mid, shop.Bid, shop.ID)
	}
}

// O2O数据库Shops 表
type Shops struct {
	ID    string `json:"id"`
	SName string `json:"sname"`
	Mid   string `json:"mid"`
	Bid   string `json:"bid"`
}

func getOSS(BusinessID int, BrandID int, ID int) {
	url := fmt.Sprintf("https://menus-omp-nm-tmp.oss-cn-beijing.aliyuncs.com/menus/%v/%v/%v/takeoutsoldout.json?t=%v", BusinessID, BrandID, ID, time.Now().Unix())
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
	fmt.Printf("SUCCESS URL:%s\n", url)
}

func getOSSShops(mid, bid, id string) {
	url := fmt.Sprintf("https://menus-omp-nm-tmp.oss-cn-beijing.aliyuncs.com/menus/%v/%v/%v/takeoutsoldout.json?t=%v", mid, bid, id, time.Now().Unix())
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
	fmt.Printf("SUCCESS URL:%s\n", url)
}
