package items

type NameCode struct {
	ID         int64
	Name       string
	Code       string `gorm:"uniqueIndex:uidex_code,priority:1"`
	Profile    int64
	CashFlow   int64
	Balance    int64
	StockPrice float64 // 当前股价
	CrawlDate  int64
	SHSZ       string `gorm:"column:shsz"`
}

type Industry struct {
	Name string `gorm:"uniqueIndex:uidex_indeustry_code,priority:2"`
	Code string `gorm:"uniqueIndex:uidex_indeustry_code,priority:1"`
}

// 利润表
type Profile struct {
	ID               int64
	Name             string
	Code             string `gorm:"uniqueIndex:uidex_profile,priority:1"`
	ReportingPeriod  string `gorm:"uniqueIndex:uidex_profile,priority:2"`
	OperateAllIncome int64
	OperateIncome    int64
	OperateAllCost   int64
	OperateCost      int64
	Tax              int64
	SalesExpense     int64
	ManageExpense    int64
	RDExpense        int64 `json:"column:rd_expense" json:"rd_expenses"`
	FinancialExpense int64
	DilutedEarn      float64 // 每股收益
}

// 上海股票的代码爬取
type NubSh struct {
	Result []DataSH `json:"result"`
}
type DataSH struct {
	Code string `json:"SECURITY_CODE_A"`
	Name string `json:"COMPANY_ABBR"`
}

type CashFlow struct {
	ID              int64
	Code            string `gorm:"uniqueIndex:uidex_cash,priority:1"`
	Name            string
	ReportingPeriod string `gorm:"uniqueIndex:uidex_cash,priority:2"`
	SalesCash       float64
	SumInFow        float64
	BuyCash         float64
	SumOutFow       float64
	NetCashFlow     float64
}

type Balance struct {
	ID              int64
	Name            string
	Code            string `gorm:"uniqueIndex:uidex_balance,priority:1"`
	ReportingPeriod string `gorm:"uniqueIndex:uidex_balance,priority:2"`
	MoneyFunds      int64
	TransFinance    int64
	Stock           int64
	ShortLoan       int64
	LongLoan        int64
	Capital         int64
}
