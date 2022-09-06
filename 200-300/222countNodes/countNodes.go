package main

import (
	"fmt"

	. "qianliout/leetcode/common/treenode"
)

func main() {
	// reverse(8, 5)
	// set1(8, 4)
	set0(8, 3)

}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func dfs(root *TreeNode) int {

}

// 反一位
func reverse(pre uint64, flag int64) uint64 {
	fmt.Println("1<<flag", 1<<flag)
	after := pre ^ uint64(1<<flag)
	fmt.Println("after", after)
	return after
}

func set0(pre uint64, flag int64) uint64 {
	/*
		binar |=  (1<<n)把n位设置为1
		binar &= ^(1 << n) // 把n位设置为0
	*/
	pre &= ^(1 << flag)
	fmt.Println("after ", pre)
	return pre
}

type ImageListResponse struct {
	ID                int64             `json:"id"`
	Digest            string            `json:"digest"`
	NodeIP            string            `json:"nodeIP"`
	ScanStatus        string            `json:"scanStatus"`
	Online            bool              `json:"online"`        // 在线 "true",离线："false"
	FromType          string            `json:"fromType"`      // 节点镜像："node" 仓库镜像:"registry"
	SecurityIssue     []SecurityIssue   `json:"securityIssue"` // 安全问题
	ImageAttr         ImageAttrResponse `json:"imageAttr"`     // 镜像属性
	UUID              uint32            `json:"uuid"`          // 镜像uuid
	CompleteTime      int64             `json:"completeTime"`  // 扫描完成时间戳
	FullRepoName      string            `json:"fullRepoName"`
	Tag               string            `json:"tag"`
	RiskScore         float64           `json:"riskScore"` // 镜像评分
	RegistryID        int64             `json:"registryId"`
	RegistryName      string            `json:"registryName"`
	RegistryUrl       string            `json:"registryUrl"`
	RegistryDeletedAt int64             `json:"registryDeletedAt"`
	Os                string            `json:"os"`
	NodeHostname      string            `json:"nodeHostname"`
	Flag              uint64            `json:"flag"`
	Project           string            `json:"project"`
	LastSyncAt        int64             `json:"lastSyncAt"`
}

type SecurityIssue struct {
	Value int64  `json:"value"` // 安全问题的ID
	Label string `json:"label"` // 安全问题
	Info  string `json:"info"`  // 详细信息
}

// 镜像属性
type ImageAttrResponse struct {
	Trusted      bool   `json:"trusted"`      // 是否是可信镜像：是：true，否：false
	ImageType    string `json:"imageType"`    // 镜像类型,基础镜像："base",应用镜像："app"
	HasFixedVuln bool   `json:"hasFixedVuln"` // 是否包含可修复漏洞: 是:true,否：false
	Reinforced   bool   `json:"reinforced"`   // 镜像是否已固:是：true,否：false
}
