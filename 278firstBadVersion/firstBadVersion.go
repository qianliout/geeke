package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	reg := Registry{
		ID:           0,
		Name:         "",
		RegType:      "",
		Url:          "",
		Username:     "",
		Password:     "",
		Description:  "",
		SyncInterval: 0,
		AccessKey:    "",
		AccessSecret: "",
		InstanceID:   "",
		RegionID:     "",
	}

	bys, _ := json.Marshal(reg)
	fmt.Println(string(bys))

}

func firstBadVersion(n int) int {
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(version int) bool { return true }

type Registry struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`     // 仓库名字,仓库名是仓库的唯一标识,一个仓库名称  对应一个用户
	RegType      string `json:"regType"`  // 仓库类型
	Url          string `json:"url"`      // 如:docker.io/v2, quay.io/v2
	Username     string `json:"username"` // user for login registry
	Password     string `json:"password"` // DES加密
	Description  string `json:"description"`
	SyncInterval int64  `json:"syncInterval"` // 单位：分钟
	AccessKey    string `json:"accessKey"`    // 阿里云仓库的AccessKey
	AccessSecret string `json:"accessSecret"` // 阿里云仓库的AccessSecret
	InstanceID   string `json:"instanceId"`   // 阿里云仓库企业版实例ID
	RegionID     string `json:"regionId"`     // 阿里云仓库企业版地域ID
}
