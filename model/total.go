type Total struct {
    CardNumber          string  `gorm:"type:varchar(200);sensitive" json:"card_number"` // 假设卡号需要特殊处理
	FBAccount           string  `gorm:"type:varchar(100)" json:"fb_account"`
    Date                string  `gorm:"type:varchar(20)" json:"date"`
    Balance             float64 `gorm:"type:decimal(10,2)" json:"balance"`
    Deplete             float64 `gorm:"type:decimal(10,2)" json:"deplete"`
}

// 增

// 删

// 改

// 查