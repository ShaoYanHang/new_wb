type FB struct {
	FBAccount                string  `gorm:"type:varchar(20)" json:"fb_account"`
	Date                   string  `gorm:"type:varchar(20)" json:"date"`
	TransactionID          string  `gorm:"type:varchar(100);primaryKey" json:"transaction_id"` // 假设Transaction ID是唯一的，可以用作主键
	TransactionDescription string  `gorm:"type:varchar(100)" json:"transaction_description"`
	CardNumber          string  `gorm:"type:varchar(100)" json:"card_number"`
	Amount                 float64 `gorm:"type:decimal(10,2)" json:"amount"` // 假设金额是浮点数，可以调整精度
	Currency               string  `gorm:"type:varchar(10)" json:"currency"`
	IsTicked               bool    `gorm:"type:boolean" json:"is_ticked"`                // 假设这是一个布尔字段，表示交易是否授权
	IsTradingAuthorization bool    `gorm:"type:boolean" json:"is_trading_authorization"` // 假设这是一个布尔字段，表示交易是否授权
	Note                   string  `gorm:"type:varchar(500)" json:"note"`
}

// 增

// 删

// 改

// 查