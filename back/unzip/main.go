package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	param := ImageListParam{
		Online:        "true",
		Keyword:       "hello",
		FromType:      "registry",
		SecurityIssue: make([]SecurityIssue, 0),
		ImageAttr: ImageAttrParam{
			Trusted:        "true",
			ImageType:      "app",
			HasFixedVuln:   "true",
			Reinforced:     "true",
			PrivilegedBoot: "true",
		},
		ImageIds:        []int64{1, 2, 3},
		ScanStatus:      []int64{1, 2, 3},
		JustReturnImage: false,
		UUIDs:           []uint32{1, 2, 3},
		Intersection:    true,
		Projects:        []string{"1,hello"},
		NodeHostname:    "hello",
		UpdateTaskInfo: UpdateTaskInfo{
			Scope:        1,
			TriggerType:  2,
			StrategyID:   3,
			StrategyName: "test01",
			Operator:     "liuqianli@outback.com",
		},
	}
	param.SecurityIssue = append(param.SecurityIssue, SecurityIssue{
		ID:   1,
		Name: "vuln",
		Info: "has vuln",
	})

	byts, _ := json.Marshal(param)
	fmt.Println(string(byts))

}

func MkdirIfNotExist(path string) error {
	stat, err := os.Stat(path)
	if err == nil {
		if stat.IsDir() {
			return nil
		} else {
			// 先删除这个文件再创建

			if err := os.Remove(path); err != nil {
				return err
			}
			return os.Mkdir(path, os.ModePerm)
		}
	}
	if os.IsNotExist(err) {
		return os.Mkdir(path, os.ModePerm)
	}
	return err
}

// 镜像属性
type ImageAttrParam struct {
	Trusted        string `json:"trusted"`        // 是否是可信镜像：是："true"，否："false"
	ImageType      string `json:"imageType"`      // 镜像类型,基础镜像："base",应用镜像："app"
	HasFixedVuln   string `json:"hasFixedVuln"`   // 是否包含可修复漏洞: 是:"true",否："false"
	Reinforced     string `json:"reinforced"`     // 镜像是否已固:是："true",否："false"
	PrivilegedBoot string `json:"privilegedBoot"` // 是否特权启动：是："true",否："false"
}

type SecurityIssue struct {
	ID   int64  `json:"id"`   // 安全问题的ID
	Name string `json:"name"` // 安全问题Name
	Info string `json:"info"` // 详细信息
}

// 镜像属性
type ImageAttrResponse struct {
	Trusted      bool   `json:"trusted"`      // 是否是可信镜像：是：true，否：false
	ImageType    string `json:"imageType"`    // 镜像类型,基础镜像："base",应用镜像："app"
	HasFixedVuln bool   `json:"hasFixedVuln"` // 是否包含可修复漏洞: 是:true,否：false
	Reinforced   bool   `json:"reinforced"`   // 镜像是否已固:是：true,否：false
}

type ImageListParam struct {
	Online          string          `json:"online"`          // 在线 "true",离线："false"
	Keyword         string          `json:"keyword"`         // 关键字搜索
	FromType        string          `json:"fromType"`        // 节点镜像："node" 仓库镜像:"registry"
	SecurityIssue   []SecurityIssue `json:"securityIssue"`   // 安全问题
	ImageAttr       ImageAttrParam  `json:"imageAttr"`       // 镜像属性
	ImageIds        []int64         `json:"imageIds"`        // 镜像ID列表
	ScanStatus      []int64         `json:"scanStatus"`      // 扫描状态
	JustReturnImage bool            `json:"justReturnImage"` // 只返回镜像信息
	UUIDs           []uint32        `json:"uuids"`           // 镜像uuid
	Intersection    bool            `json:"intersection"`    // 交集还是并集
	Projects        []string        `json:"projects"`        // 仓库和repo的筛选
	NodeHostname    string          `json:"nodeHostname"`    // 节点名精确匹配

	//  以下是镜像扫描时的参数
	UpdateTaskInfo UpdateTaskInfo `json:"updateTaskInfo"`

	// 后端处理数据的中间结构
	RegistryIds []int64 `json:"-"`
	Repos       []Repo  `json:"-"`
}

type UpdateTaskInfo struct {
	Scope        int    `json:"scope"`
	TriggerType  int    `json:"triggerType"`  // 扫描类型
	StrategyID   int64  `json:"strategyId"`   // 扫描策略ID
	StrategyName string `json:"strategyName"` // 扫描策略名字 用于openapi
	Operator     string `json:"operator"`     // 操作人
}

type Repo struct {
	RegistryID int64  `json:"registryID"`
	RepoName   string `json:"repoName"`
}
