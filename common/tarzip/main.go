package main

import (
	"encoding/json"
	"fmt"
)

type shell struct {
	S []string
}
type UserReq struct {
	Name       string           `json:"name"`
	Account    string           `json:"account"`
	Password   string           `json:"password"`
	Roles      map[string][]int `json:"roles"`       // 角色列表
	CreatorId  int64            `json:"creator_id"`  // 操作人的Id（谁添加的）
	Comment    string           `json:"comment"`     // 备注，添加信息
	Status     string           `json:"status"`      // 当前状态，是否审批通过
	VerifierId int64            `json:"verifier_id"` // 审核人员ID
}

type CalloutAgentConfig struct {
	AgentID int64  `json:"agent_id"`
	Sip     string `json:"sip"`
	Phone   string `json:"phone"`
}

type CalloutEntConfig struct {
	CalloutSwitch string               `json:"callout_switch"`
	Configs       []CalloutAgentConfig `json:"configs"`
}

type ChangePasswordReq struct {
	Account      string `json:"account"`
	ID           int64  `json:"id"`
	PrePassword  string `json:"pre_password"`
	NewPassword1 string `json:"new_password1"`
	NewPassword2 string `json:"new_password2"`
}

type UserSearchReq struct {
	Name    string `json:"name"`
	Account string `json:"account"`
	Product string `json:"product"` // 产品线，来鼓还是美洽
	Roles   []int  `json:"roles"`   // 角色列表
	Status  string `json:"status"`  // 当前状态，是否审批通过
}

func main() {
	//u := UserReq{
	//	Name:     "hello",
	//	Account:  "lql@163.com",
	//	Password: "123454",
	//	Roles:    map[string][]int{"lagui": []int{1, 2, 3}, "meiqia": []int{2, 3, 4}},
	//	Comment:  "测试",
	//}
	//u:=ChangePasswordReq{
	//	Account:      "lql@163.com",
	//	PrePassword:  "12345",
	//	NewPassword1: "123456",
	//	NewPassword2: "123456",
	//}

	//u := UserSearchReq{
	//	Name:    "hello",
	//	Account: "ldfd@",
	//	Product: "laigu",
	//	Roles:   []int{1, 2, 3},
	//	Status:  "noactive",
	//}

	b, _ := json.Marshal([]string{"fdkfd"})
	fmt.Println(string(b))

}
