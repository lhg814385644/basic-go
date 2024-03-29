package domain

// 领域对象

// User 是业务概念。他不一定和数据库中的表字段一一对应，而dao.User则是直接映射到数据库表中的字段。
type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	NickName     string `json:"nick_name"`    // 昵称
	Birthday     string `json:"birthday"`     // 生日
	Introduction string `json:"introduction"` // 简介
}
