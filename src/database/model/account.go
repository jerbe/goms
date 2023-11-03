package model

import (
	"time"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/22 23:42
  @describe :
*/

// ------------------------------------------------
const (

	// TableAccounts 表格:账户基础信息
	TableAccounts = "`accounts`"
	// FieldsAccounts 所有字段:账户基础信息所有字段
	FieldsAccounts = "`id`,`name`,`password`,`salt`,`2ndpassword`,`salt2`,`loggedin`,`lastlogin`,`createdat`,`birthday`,`banned`,`banreason`,`gm`,`email`,`macs`,`tempban`,`greason`,`ACash`,`mPoints`,`gender`,`SessionIP`,`points`,`vpoints`,`lastlogon`,`facebook_id`,`access_token`,`password_otp`,`expiration`,`VIP`,`money`,`moneyb`,`lastGainHM`,`paypalNX`"

	// TableAccountsInfo 表格:账户附属信息
	TableAccountsInfo = "`accounts_info`"
	// FieldsAccountsInfo 所有字段:账户附属信息所有字段
	FieldsAccountsInfo = "`id`,`accId`,`worldId`,`cardSlots`,`gamePoints`,`updateTime`,`gamePointspd`,`gamePointsps`,`sjrw`,`sgrw`,`fbrw`,`sbossrw`,`sgrwa`,`fbrwa`,`sbossrwa`,`lb`"
)

//------------------------------------------------

// Account 账户信息
type Account struct {
	// ID 账户ID,自增
	ID uint64 `json:"id" db:"id"`

	// Name 账户名
	Name string `json:"name" db:"name"`

	// Password 密码
	Password string `json:"password" db:"password"`

	// Salt 可能是密码加密盐
	Salt string `json:"salt" db:"salt"`

	// SecondPassword 二级密码
	SecondPassword string `json:"second_password" db:"2ndpassword"`

	// Salt2 二级密码加密盐
	Salt2 string `json:"salt_2" db:"salt_2"`

	// LoggedIn 登录状态
	LoggedIn uint8 `json:"logged_in" db:"loggedin"`

	// LastLogin 上次登录时间? 怎么有两个?
	LastLogin *time.Time `json:"last_login" db:"lastlogin"`

	// LastLogon 上次登录时间? 怎么有两个?
	LastLogon *time.Time `json:"last_logon" db:"lastlogon"`

	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" db:"createdat"`

	// Birthday 生日
	Birthday time.Time `json:"birthday" db:"birthday"`

	// Banned 禁用状态
	Banned uint `json:"banned" db:"banned"`

	// BanReason 禁用原因
	BanReason string `json:"ban_reason" db:"banreason"`

	// GM 是否是游戏管理员
	GM uint `json:"gm" db:"gm"`

	// Email 邮箱地址
	Email string `json:"email" db:"email"`

	// Macs  登录时的硬件地址
	Macs string `json:"macs" db:"macs"`

	// TempBanedAt 临时禁用时间
	TempBanedAt time.Time `json:"temp_baned_at" db:"tempban"`

	// GReason  暂时不知道是什么
	GReason *uint `json:"g_reason" db:"greason"`

	// ACash 暂时不知道是什么,点券相关?
	ACash int `json:"a_cash" db:"acash"`

	// MPoints 暂时不知道是什么,积分相关?
	MPoints int `json:"m_points" db:"mpoints"`

	// Gender 性别
	Gender uint8 `json:"gender" db:"gender"`

	// SessionIP 会话IP
	SessionIP string `json:"session_ip" db:"SessionIP"`

	// Points 积分相关?
	Points int `json:"points" db:"points"`

	// VPoints 会员积分?
	VPoints int `json:"v_points" db:"vpoints"`

	// FacebookID Facebook的账户ID
	FacebookID string `json:"facebook_id" db:"facebook_id"`

	// FbAccessToken 访问代币?哪个平台的?facebook?
	FbAccessToken string `json:"access_token" db:"access_token"`

	// FbPasswordOtp 一次性密码?哪个平台的?Facebook
	FbPasswordOtp string `json:"password_otp" db:"password_otp"`

	// Expiration 到期时间 Facebook 的token到期时间?
	Expiration *time.Time `json:"expiration" db:"expiration"`

	// Vip 会员的标记,有等级之分?
	Vip int `json:"vip" db:"vip"`

	// Money 钱?
	Money int `json:"money" db:"money"`

	// MoneyB 钱B?
	MoneyB int `json:"money_b" db:"moneyb"`

	// LastGainHM 最后雇用商人进行开店的时间戳
	LastGainHM int `json:"last_gain_hm" db:"last_gain_hm"`

	// PaypalNX paypal相关的,应该是现金支付相关?
	PaypalNX int `json:"paypal_nx" db:"paypalNX"`
}

// AccountInfo 账户扩展信息
type AccountInfo struct {
}
