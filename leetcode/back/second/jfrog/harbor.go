package main

const (
	RejectPolicyIgnore = "ignore" // 忽略
	RejectPolicyAlarm  = "alarm"  // 报警
	RejectPolicyReject = "reject" // 阻断

	RejectPolicyBaseModel = "base" // 基本模式
	RejectPolicySafeModel = "safe" // 安全模式

	// 漏洞级别
	NegligibleVuln = "Negligible" // 可忽略
	UnknownVuln    = "Unknown"    // 未知
	LowVuln        = "Low"        // 低
	MediumVuln     = "Medium"     // 中
	HighVuln       = "High"       // 危
	CriticalVuln   = "Critical"   // 高危

)
const (
	// TOP5 统计类别
	RejectReasonVuluScore        = 1 // 漏洞评分低于设置值
	RejectReasonHasSensitiveFile = 2 // 存在敏感文件
	RejectReasonHasMalicious     = 3 // 存在恶意文件

	RejectReasonHasCustomizeVulu = 4 // 存在自定义漏洞

	RejectReasonHasNegligible  = 5  // 存在可忽略漏洞
	RejectReasonHasUnknown     = 6  // 存在末知漏洞
	RejectReasonHasLow         = 7  // 存在低危漏洞
	RejectReasonHasMedium      = 8  // 存在中危漏洞
	RejectReasonHasHigh        = 9  // 存在危险漏洞
	RejectReasonHasCritical    = 10 // 存在高危漏洞
	RejectNoLibrary            = 11 // 来源镜像不在本地仓库（安全模式）
	RejectScanFailure          = 12 // 镜像扫描失败
	RejectScanNotScanned       = 13 // 镜像未扫描
	RejectReasonUntrustedImage = 14 // 不信任的镜像

	RejectReasonUntrustedBaseImage = 15 // 基础镜像不可信
	RejectReasonWebshellScore      = 16 // webshell 评分低于设置值
)

const (
	LangEn = "en"
	LangZh = "ch"
)

const (
	UsePatternForCICD   = "for_cicd"
	UsePatternForK8s    = "for_k8s"
	UsePatternForOnline = "for_online"
)
const RegistryUseTypeBuff = 2 // 表示CICD的中转仓库

const (
	ImageFromTypeCICD   = 2
	ImageFromTypeNormal = 1
)

var reasonZHMap = map[int64]string{
	RejectReasonVuluScore:          "漏洞评分低于设置值",
	RejectReasonHasSensitiveFile:   "存在敏感文件",
	RejectReasonHasMalicious:       "存在恶意文件",
	RejectReasonHasCustomizeVulu:   "存在自定义漏洞",
	RejectReasonHasNegligible:      "存在可忽略漏洞",
	RejectReasonHasUnknown:         "存在末知漏洞",
	RejectReasonHasLow:             "存在低危漏洞",
	RejectReasonHasMedium:          "存在中危漏洞",
	RejectReasonHasHigh:            "存在危险漏洞",
	RejectReasonHasCritical:        "存在高危漏洞",
	RejectNoLibrary:                "来源镜像不在本地仓库",
	RejectScanFailure:              "镜像扫描失败",
	RejectScanNotScanned:           "镜像未扫描",
	RejectReasonUntrustedImage:     "非可信镜像",
	RejectReasonUntrustedBaseImage: "基础镜像不可信",
	RejectReasonWebshellScore:      "webshell评分高于设置值",
}
var reasonENMap = map[int64]string{
	RejectReasonVuluScore:          "Vulnerability score lower than set value",
	RejectReasonHasSensitiveFile:   "Exist sensitive file",
	RejectReasonHasMalicious:       "Exist malicious file",
	RejectReasonHasCustomizeVulu:   "Exist custom vulnerability file",
	RejectReasonHasNegligible:      "Exist Negligible vulnerability file",
	RejectReasonHasUnknown:         "Exist Unknown vulnerability file",
	RejectReasonHasLow:             "Exist Low vulnerability file",
	RejectReasonHasMedium:          "Exist Medium vulnerability file",
	RejectReasonHasHigh:            "Exist High vulnerability file",
	RejectReasonHasCritical:        "Exist Critical vulnerability file",
	RejectNoLibrary:                "image not in config registry",
	RejectScanFailure:              "image scan failure",
	RejectScanNotScanned:           "image not scanned",
	RejectReasonUntrustedImage:     "untrusted image",
	RejectReasonUntrustedBaseImage: "untrusted base image",
	RejectReasonWebshellScore:      "webshell score more than set value",
}
var reasonChMap = map[string]string{
	NegligibleVuln: GetRejectReason(LangZh)[RejectReasonHasNegligible],
	UnknownVuln:    GetRejectReason(LangZh)[RejectReasonHasUnknown],
	LowVuln:        GetRejectReason(LangZh)[RejectReasonHasLow],
	MediumVuln:     GetRejectReason(LangZh)[RejectReasonHasMedium],
	HighVuln:       GetRejectReason(LangZh)[RejectReasonHasHigh],
	CriticalVuln:   GetRejectReason(LangZh)[RejectReasonHasCritical],
}
var reasonEnMap = map[string]string{
	NegligibleVuln: GetRejectReason(LangEn)[RejectReasonHasNegligible],
	UnknownVuln:    GetRejectReason(LangEn)[RejectReasonHasUnknown],
	LowVuln:        GetRejectReason(LangEn)[RejectReasonHasLow],
	MediumVuln:     GetRejectReason(LangEn)[RejectReasonHasMedium],
	HighVuln:       GetRejectReason(LangEn)[RejectReasonHasHigh],
	CriticalVuln:   GetRejectReason(LangEn)[RejectReasonHasCritical],
}

func GetRejectReason(lag string) map[int64]string {
	if lag == LangZh {
		return reasonZHMap
	} else if lag == LangEn {
		return reasonENMap
	}
	return make(map[int64]string)
}

func GetSeverityRejectReason(severity string) int64 {
	subScore := map[string]int64{
		CriticalVuln:   RejectReasonHasCritical,
		HighVuln:       RejectReasonHasHigh,
		MediumVuln:     RejectReasonHasMedium,
		LowVuln:        RejectReasonHasLow,
		NegligibleVuln: RejectReasonHasNegligible,
		UnknownVuln:    RejectReasonHasUnknown,
	}
	return subScore[severity]
}

func GetVuluRuleKey(vuleLeve string, lag string) string {
	switch lag {
	case LangEn:
		return reasonEnMap[vuleLeve]
	case LangZh:
		return reasonChMap[vuleLeve]
	}
	return ""
}
