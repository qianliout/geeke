package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	image := Image{
		ID:   12,
		Name: "(test)https://harbor-prod.tensorsecurity.com",
		FixedVuln: VulnSeverityCount{
			Critical: 2,
			High:     3,
			Medium:   4,
			Low:      5,
			Unknown:  6,
		},
		UnFixedVuln: VulnSeverityCount{
			Critical: 7,
			High:     9,
			Medium:   12,
			Low:      2,
			Unknown:  3,
		},
		Malicious: 12,
		RiskScore: 23,
		Flag:      0,
	}
	res := ImageResponse{
		Images:  []Image{image},
		End:     false,
		StartID: 12,
	}
	marshal, _ := json.Marshal(res)
	fmt.Println(string(marshal))

	rov := RiskOverView{
		ImageCount:             1,
		OnlineImageCount:       3,
		TrustedImageCount:      4,
		HasFixedVulnImageCount: 5,
		ReinforcedImageCount:   3,
		VulnSum:                3,
		VulnSeverity: VulnSeverityCount{
			Critical: 3,
			High:     4,
			Medium:   3,
			Low:      4,
			Unknown:  3,
		},
		ImageSecurity: ImageSecurity{
			HasVuln:          2,
			HasMalicious:     3,
			HasSensitive:     3,
			HasWebshell:      2,
			HasSoftware:      4,
			HasExceptEnv:     3,
			PrivilegedBoot:   3,
			HasExceptLicense: 2,
		},
	}

	by, _ := json.Marshal(rov)

	fmt.Println(string(by))
	vuln := VulnWithImage{
		Images: []string{},
		VulnDetail: VulnDetail{
			Name:          "cev-2020",
			Severity:      "high",
			SeverityInt:   2,
			FixedBy:       "v2",
			Description:   "hello/word",
			Cvssv3score:   "9.8",
			Cvssv3vector:  "cddfd/dfdfe/d",
			CnvdTitle:     "本地攻击",
			FixSuggestion: "fdfld",
			PkgName:       "centos",
			PkgVersion:    "v3",
			UniqueVuln:    1232332,
		},
	}

	vulns := VulnWithImageResponse{
		Vulns:   []VulnWithImage{vuln},
		End:     false,
		StartID: 12,
	}
	bys, _ := json.Marshal(vulns)
	fmt.Println(string(bys))

	vi := VirusInfo{
		FileName:  "hello",
		FilePath:  "/var/lib/hello",
		VirusName: "word",
	}
	by2, _ := json.Marshal(vi)
	fmt.Println(string(by2))

	imr := ImageRiskOverView{
		ImageID:       12,
		ImageName:     "(test)https://harbor-prod.tensorsecurity.com/hell/word:v1",
		FixSuggestion: []string{"请升级"},
		Vulns: VulnDetailSeverityGroup{
			Critical: []VulnWithImage{vuln},
			High:     []VulnWithImage{vuln},
			Medium:   []VulnWithImage{vuln},
			Low:      []VulnWithImage{vuln},
			Unknown:  []VulnWithImage{vuln},
		},
		VulnSeverityCount: VulnSeverityCount{
			Critical: 2,
			High:     3,
			Medium:   4,
			Low:      5,
			Unknown:  7,
		},
	}
	by4, _ := json.Marshal(imr)
	fmt.Println(string(by4))

}

type VulnWithImageResponse struct {
	Vulns   []VulnWithImage `json:"vulns"`
	End     bool            `json:"end"`
	StartID int64           `json:"startID"`
}

// 漏洞及影响镜像
type VulnWithImage struct {
	Images     []string   `json:"images"`
	VulnDetail VulnDetail `json:"vulnDetail"`
}

type VirusInfo struct {
	FileName  string `json:"filename"`
	FilePath  string `json:"filepath"`
	VirusName string `json:"virusname"`
}

// 单个镜像的扫描报告
type ImageRiskOverView struct {
	ImageID           int64                   `json:"imageID"`
	ImageName         string                  `json:"imageName"`
	FixSuggestion     []string                `json:"fixSuggestion"`
	Vulns             VulnDetailSeverityGroup `json:"vulns"`             // 漏洞层级列表
	VulnSeverityCount VulnSeverityCount       `json:"vulnSeverityCount"` // 漏洞层级分布
}

// 漏洞详情按层级统计
type VulnDetailSeverityGroup struct {
	Critical []VulnWithImage `json:"critical"`
	High     []VulnWithImage `json:"high"`
	Medium   []VulnWithImage `json:"medium"`
	Low      []VulnWithImage `json:"low"`
	Unknown  []VulnWithImage `json:"unknown"`
}

// 漏洞详情
type VulnDetail struct {
	Name          string `json:"name"`          // 形如CVE-2021-28831
	Severity      string `json:"severity"`      // 威胁等级
	SeverityInt   int    `json:"-"`             // 威胁等级
	FixedBy       string `json:"fixedby"`       // 修复版本
	Description   string `json:"description"`   // 漏洞介绍
	Cvssv3score   string `json:"cvssv3Score"`   // 漏洞评分
	Cvssv3vector  string `json:"cvssv3Vector"`  // 攻击维度
	CnvdTitle     string `json:"cnvdTitle"`     // 漏洞类型 // 取值方式:vuln.Metadata.CNVDs[0].Title
	FixSuggestion string `json:"fixSuggestion"` // 修复建议
	PkgName       string `json:"pkgName"`       // 软件包来源
	PkgVersion    string `json:"pkgVersion"`    // 软件包版本
	UniqueVuln    uint64 `json:"uniqueVuln;string"`
}

type ImageResponse struct {
	Images  []Image `json:"images"`
	End     bool    `json:"end"`
	StartID int64   `json:"startID"`
}

// 镜像列表
type Image struct {
	ID          int64             `json:"id"`
	Name        string            `json:"name"`
	FixedVuln   VulnSeverityCount `json:"fixedVuln"`
	UnFixedVuln VulnSeverityCount `json:"unFixedVuln"`
	Malicious   int64             `json:"malicious"` // 病毒
	RiskScore   float64           `json:"riskScore"`
	Flag        uint64            `json:"flag"`
}

// 漏洞个数按层级统计
type VulnSeverityCount struct {
	Critical int64 `json:"critical"`
	High     int64 `json:"high"`
	Medium   int64 `json:"medium"`
	Low      int64 `json:"low"`
	Unknown  int64 `json:"unknown"`
}

// 风险信息总揽
type RiskOverView struct {
	ImageCount             int64             `json:"imageCount"`
	OnlineImageCount       int64             `json:"onlineImageCount"`
	TrustedImageCount      int64             `json:"trustedImageCount"`
	HasFixedVulnImageCount int64             `json:"hasFixedVulnImageCount"`
	ReinforcedImageCount   int64             `json:"reinforcedImageCount"`
	VulnSum                int64             `json:"vulnSum"` // 漏洞总数
	VulnSeverity           VulnSeverityCount `json:"vulnSeverity"`
	ImageSecurity          ImageSecurity     `json:"imageSecurity"`
}

// 镜像安全问题统计
type ImageSecurity struct {
	HasVuln          int64 `json:"hasVuln"`
	HasMalicious     int64 `json:"hasMalicious"`
	HasSensitive     int64 `json:"hasSensitive"`
	HasWebshell      int64 `json:"hasWebshell"`
	HasSoftware      int64 `json:"hasSoftware"`
	HasExceptEnv     int64 `json:"hasExceptEnv"`
	PrivilegedBoot   int64 `json:"privilegedBoot"`
	HasExceptLicense int64 `json:"hasExceptLicense"`
}
