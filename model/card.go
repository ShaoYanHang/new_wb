type Transaction struct {
	TransactionID       string  `gorm:"type:varchar(50);primaryKey" json:"transaction_id"`
	Date                   string  `gorm:"type:varchar(20)" json:"date"`
	CardNumber          string  `gorm:"type:varchar(200);sensitive" json:"card_number"` // 假设卡号需要特殊处理
	FBAccount            string  `gorm:"type:varchar(100)" json:"fb_account"`
	BillName            string  `gorm:"type:varchar(255)" json:"bill_name"`
	TransactionType     string  `gorm:"type:varchar(100)" json:"transaction_type"`
	OrderAmount         float64 `gorm:"type:decimal(10,2)" json:"order_amount"` // 假设金额为小数
	OrderCurrency       string  `gorm:"type:varchar(10)" json:"order_currency"`
	TransactionAmount   float64 `gorm:"type:decimal(10,2)" json:"transaction_amount"`
	TransactionFee      float64 `gorm:"type:decimal(10,2)" json:"transaction_fee"` // 假设费用也为小数
	TransactionCurrency string  `gorm:"type:varchar(10)" json:"transaction_currency"`
	TransactionStatus   string  `gorm:"type:varchar(100)" json:"transaction_status"`
	AuthorizationCode   string  `gorm:"type:varchar(100)" json:"authorization_code,omitempty"` // 如果可能为空，使用omitempty
	ResultCode          string  `gorm:"type:varchar(100)" json:"result_code"`
	ResultDescription   string  `gorm:"type:varchar(255)" json:"result_description"`
	SettlementStatus    string  `gorm:"type:varchar(100)" json:"settlement_status"`
	IsJudge             bool    `gorm:"type:boolean" json:"is_judge"`
}

// 增

// 删

// 改

// 查