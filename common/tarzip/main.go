package main

import (
	"encoding/json"
	"fmt"
	"time"
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

type EntAuditAddReq struct {
	EntIds     []int64 `json:"ent_ids"`     // 企业ID
	Status     string  `json:"status"`      // 状态，黑，灰等,枚举类型
	Reason     []int64 `json:"reason"`      // 加入的原因，使用枚举类型
	Source     int64   `json:"source"`      // 封禁来源
	Comment    string  `json:"comment"`     // 备注
	AddExtra   bool    `json:"add_extra"`   // 是否把附加资源也加入
	AddComment bool    `json:"add_comment"` // 是否把附加资源也加入
	ForceAdd   bool    `json:"force_add"`   // 是否是强制添加
}
type EntAuditSearchReq struct {
	EntIds  []int64   `json:"ent_id"`  // 企业ID
	Status  string    `json:"status"`  // 状态，黑，灰等,枚举类型
	Reason  int64     `json:"reason"`  // 加入的原因，使用枚举类型
	Source  int64     `json:"source"`  // 封禁来源
	Product string    `json:"product"` // 产品线，来鼓还是美洽
	Comment string    `json:"comment"` // 备注
	Token   string    `json:"token"`   // 企业token
	Name    string    `json:"name"`    // 企业名字
	Phone   string    `json:"phone"`   // 电话
	Email   string    `json:"email"`
	StartAt time.Time `json:"start_at"` // 搜索开始时间
	EndAt   time.Time `json:"end_at"`   // 搜索结束时间

	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}

type Request struct {
	EntID         int64  `json:"ent_id"`
	Start         string `json:"start"`
	End           string `json:"end"`
	ExportContent bool   `json:"export_content"`
	ShowClue      bool   `json:"show_clue"`
	MaxRow        int    `json:"max_row"`
	WorkerCount   int    `json:"worker_count"`
	DataSource    int    `json:"data_source"`
	StartEntId    int64  `json:"start_ent_id"`
	Asyn          bool   `json:"asyn"`
}

func MarshalReason(reason []int64) string {
	res, err := json.Marshal(reason)
	if err != nil || string(res) == "null" {
		return "[]"
	}
	return string(res)
}

// 二级词库表(分类表)
// nolint:maligned
type CategoryReq struct {
	Id              int64   `json:"id"`
	Ids             []int64 `json:"ids"`
	Name            string  `json:"name"`             // 词库名
	FirstCategory   int64   `json:"first_category"`   // 对应的一级词库名
	FirstCategories []int64 `json:"first_categories"` // 对应的一级词库名
	Product         string  `kok:"header.Captcha-Needed" json:"product"`
	Weight          int8    `json:"weight"` // 权重

	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}

type AuditWordReq struct {
	Ids             []int64  `json:"ids"`
	Product         string   `json:"product"`          // 产品线，来鼓还是美洽
	Names           []string `json:"names"`            // 敏感词
	Name            string   `json:"name"`             // 敏感词
	FirstCategory   int64    `json:"first_category"`   // 对应的一级词库名
	FirstCategories []int64  `json:"first_categories"` // 对应的一级词库名
	Category        int64    `json:"category"`         // 对应的二级词库名
	Status          string   `json:"status"`           // 当前状态，是否审批通过
	AddReason       string   `json:"add_reason"`       // 备注，添加信息

	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}

func UnmarshalReason(reason string) []int64 {
	list := make([]int64, 0)

	err := json.Unmarshal([]byte(reason), &list)
	if err != nil {
		return []int64{}
	}
	return list
}

func main() {
	c := CategoryReq{
		Id:              1,
		Ids:             []int64{1, 2, 3},
		Name:            "helloword",
		FirstCategory:   1,
		FirstCategories: []int64{1, 2, 3},
		Product:         "meiqia",
		Weight:          20,
	}

	w := AuditWordReq{
		Ids:             []int64{1, 2, 3},
		Product:         "meiqia",
		Names:           []string{"hellwword", "jhelloword"},
		Name:            "你好",
		FirstCategory:   1,
		FirstCategories: []int64{1, 2, 3},
		Category:        4,
		Status:          "pending",
		AddReason:       "就是想添加了",
	}

	cc, _ := json.Marshal(c)
	ww, _ := json.Marshal(w)
	fmt.Println(string(cc))
	fmt.Println(string(ww))

}
