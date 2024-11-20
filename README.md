# git
echo "# new_wb" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/ShaoYanHang/new_wb.git
git push -u origin main

# mysql
mysql -uroot -p0695f508cccae5fb

# sql
create database db2 default character set utf8mb4 collate utf8mb4_unicode_ci;

# 服务器信息 
http://fozadmin.com/

# 表 FB
Date	
Transaction ID	
Transaction Description
Payment Method	
Amount	
Currency
Account

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


# 表 CARD
交易编号
交易时间
卡号
昵称
账单名称
交易类型
订单金额
订单币种
交易金额
交易费
交易币种
交易状态
授权码
结果码
结果描述
清算状态

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

# 表 total
type Total struct {
    CardNumber          string  `gorm:"type:varchar(200);sensitive" json:"card_number"` // 假设卡号需要特殊处理
	FBAccount           string  `gorm:"type:varchar(100)" json:"fb_account"`
    Date                string  `gorm:"type:varchar(20)" json:"date"`
    Balance             float64 `gorm:"type:decimal(10,2)" json:"balance"`
    Deplete             float64 `gorm:"type:decimal(10,2)" json:"deplete"`
}


# 表 USER

type User struct {
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
    Email     string `gorm:"type:varchar(200)" json:"email"`
    Iphone    string `gorm:"type:varchar(20)" json:"iphone"`
}

