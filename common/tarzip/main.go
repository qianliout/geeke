package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	s := WordTrigger{
		MsgInfo: &MsgInfo{
			ID:          1,
			ContentType: "text",
			Content:     "你发的哈",
			TriggerWord: "你发",
			CreatedOn:   time.Now(),
		},
		VisitInfo: &VisitInfo{
			FirstVisitPageSourceBySession: "www",
			FirstVisitPageUrlBySession:    "www",
			Url:                           "www",
			Title:                         "标题",
			IP:                            "183",
		},
		AgentInfo: &AgentInfo{
			FromType: "fdhj",
			AgentId:  12,
			Email:    "l@s",
			Realname: "就是好",
		},
		EntInfo: &Ent{
			EntId:         12,
			Email:         "e@edfd",
			FullName:      "颇感兴趣",
			RegPhone:      "183",
			ContactPhone:  "193",
			ValidatePhone: "3984",
			EntToken:      "dfkdsl",
		},
	}
	marshal, _ := json.Marshal(s)

	fmt.Println(string(marshal))
}

// 返回前端的数据
type WordTrigger struct {
	Id        int64  `json:"id"`
	EntId     int64  `json:"ent_id"` // 公司ID
	EntName   string `json:"ent_name"`
	EntStatus string `json:"ent_status"`

	WordId   int64  `json:"word_id"`   // 词语ID
	WordName string `json:"word_name"` // 敏感词

	MsgId   int64  `json:"msg_id"`  // 消息ID
	Content string `json:"content"` // 消息内容

	TrackId   string    `json:"track_id"` // 访客ID
	AgentId   int64     `json:"agent_id"` // 客服ID
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 详情页下的内容
	MsgInfo   *MsgInfo   `json:"msg_info"`
	VisitInfo *VisitInfo `json:"visit_info"`
	AgentInfo *AgentInfo `json:"agent_info"`
	EntInfo   *Ent       `json:"ent_info"`
}

type VisitInfo struct {
	FirstVisitPageSourceBySession string `json:"first_visit_page_source_by_session"` // 来源页面
	FirstVisitPageUrlBySession    string `json:"first_visit_page_url_by_session"`    // 着路页
	Url                           string `json:"url"`                                // 对话页
	Title                         string `json:"title"`
	IP                            string `json:"ip"`
}

type Extra struct {
}

type MsgInfo struct {
	ID          int64     `json:"id"`
	ContentType string    `json:"content_type"`
	Content     string    `json:"content"`
	TriggerWord string    `json:"trigger_word"`
	CreatedOn   time.Time `json:"created_on"`
}

type AgentInfo struct {
	FromType string `json:"from_type"`
	AgentId  int64  `json:"agent_id"`
	Ip       string `json:"ip"`
	Email    string `json:"email"`
	Realname string `json:"realname"`
}

type Ent struct {
	EntId         int64  `json:"ent_id"`
	Email         string `json:"email"` // 注册邮箱
	FullName      string `json:"full_name"`
	RegPhone      string `json:"reg_phone"`      // 注册手机
	ContactPhone  string `json:"contact_phone"`  // 联系人手机
	ValidatePhone string `json:"validate_phone"` // 验证手机
	EntToken      string `json:"ent_token"`
}
