package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	s := map[string]interface{}{

		"product":       "meiqia",
		"ent_id":        97567,
		"word_id":       0,
		"msg_id":        345,
		"track_id":      "1lp2cMA805FVWcek7mlpI6zeX8B",
		"visit_id":      "1lp2eSIXvNAsQ6mr7mDXrvfqKIx",
		"agent_id":      123,
		"created_at":    "2020-12-18T08:24:16Z",
		"created_at_ts": 1608279856,
		"msg":           "xxp",
		"word":          "",
	}
	bytes, _ := json.Marshal(s)
	w := new(WordTriggerCH)

	json.Unmarshal(bytes, w)
	fmt.Println(w.CreatedAt.IsZero())
	fmt.Println(validate(w))

}

func validate(data *WordTriggerCH) bool {
	if data.Product == "" || data.CreatedAtTs <= 0 || data.MsgId <= 0 || data.EntId <= 0 ||
		data.TrackId == "" || data.VisitId == "" || data.MSG == "" || data.AgentId <= 0 || data.CreatedAt.IsZero() {
		return false
	}
	return true
}

// clickhouse里面所存的触发信息
type WordTriggerCH struct {
	Product     string    `db:"product" json:"product"`
	EntId       int64     `db:"ent_id" json:"ent_id"`     // 公司ID
	WordId      int64     `db:"word_id" json:"word_id"`   // 词语ID
	MsgId       int64     `db:"msg_id" json:"msg_id"`     // 消息ID
	TrackId     string    `db:"track_id" json:"track_id"` // Track ID
	VisitId     string    `db:"visit_id" json:"visit_id"` // 访客ID
	AgentId     int64     `db:"agent_id" json:"agent_id"` // 客服ID
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	CreatedAtTs int64     `db:"-" json:"created_at_ts"` // ribbon发送的时间，以时间戳传递，保险

	MSG  string `db:"-" json:"msg"`
	Word string `db:"word" json:"word"`
}
