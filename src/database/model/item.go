package model

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/31 12:05
  @describe :
*/

// ------------------------------------------------
const (

	// --------- TABLE ------------------------------------------------

	// TableInventoryItems 表格:库存物品表
	TableInventoryItems = "`inventoryitems`"

	// TableInventoryEquipment 表格:装备栏表
	TableInventoryEquipment = "`inventoryequipment`"

	// TableCSItems (cash shop items) 表格:现金商店物品表
	TableCSItems = "`csitems`"

	// TableCSEquipment  表格:现金商店已装备物品表
	TableCSEquipment = "`csequipment`"

	// TableHiredMerchItems 表格:雇佣商人物品表
	TableHiredMerchItems = "`hiredmerchitems`"

	// TableHiredMerchEquipment 表格:雇佣商人已装备物品表
	TableHiredMerchEquipment = "`hiredmerchequipment`"

	// TableDueyItems 表格:Duey物品表
	TableDueyItems = "`dueyitems`"

	// TableDueyEquipment 表格:Duey已装备物品表
	TableDueyEquipment = "`dueyequipment`"

	// TableMTSItems 表格:MTS物品表
	TableMTSItems = "`mtsitems`"

	// TableMTSEquipment 表格:MTS已装备物品表
	TableMTSEquipment = "`mtsequipment`"

	// TableMTSTransferItems 表格:MTS Transfer物品表
	TableMTSTransferItems = "`mtstransfer`"

	// TableMTSTransferEquipment 表格:MTS Transfer已装备物品表
	TableMTSTransferEquipment = "`mtstransferequipment`"

	// --------- FIELDS ------------------------------------------------

	// FieldsItems 所有字段:库存/背包/xxx等等有物品的都公用同样的字段
	FieldsItems = "`inventoryitemid`,`characterid`,`accountid`,`packageid`,`itemid`,`inventorytype`,`position`,`quantity`,`owner`,`GM_Log`,`uniqueid`,`flag`,`expiredate`,`type`,`sender`"
)

// IItem 个人持有的物品
type IItem interface {
	// isItem 是否是个人持有的物品
	isItem() bool

	// Base 基础库存物品
	Base() *Item
}

// Item 持有的物品
type Item struct {
	// ID 自增ID,索引用吧,没其他效果
	ID int64 `json:"inventoryitemid" db:"inventoryitemid"`

	// CharacterID 角色ID
	CharacterID *int64 `json:"characterid" db:"characterid"`

	// AccountID 账户ID
	AccountID *int `json:"accountid" db:"accountid"`

	// PackageID ??
	PackageID *int `json:"packageid" db:"packageid"`

	// ItemID 物品ID
	ItemID int `json:"itemid" db:"itemid"`

	// InventoryType 库存类型
	InventoryType int `json:"inventorytype" db:"inventorytype"`

	// Position 物品摆放位置
	Position int `json:"position" db:"position"`

	// Quantity 数量
	Quantity int `json:"quantity" db:"quantity"`

	// Owner 拥有者
	Owner *string `json:"owner" db:"owner"`

	// GMLog 管理员日志
	GMLog *string `json:"GM_Log" db:"GM_Log"`

	// UniqueID 唯一标识
	UniqueID int `json:"uniqueid" db:"uniqueid"`

	// Flag 标记
	Flag int `json:"flag" db:"flag"`

	// ExpireDate 到期日期
	ExpireDate int64 `json:"expiredate" db:"expiredate"`

	// Type 类型
	Type int `json:"type" db:"type"`

	// Sender 发送人/来演
	Sender string `json:"sender" db:"sender"`
}

func (*Item) isItem() bool {
	return true
}

// Base 基础物品
func (i *Item) Base() *Item {
	return i
}

// ItemWithEquipment 已配备/已使用的物品
type ItemWithEquipment struct {
	Item

	InventoryEquipmentID int   `json:"inventoryequipmentid" db:"inventoryequipmentid"`
	InventoryItemID      int   `json:"inventoryitemid" db:"inventoryitemid"`
	UpgradeSlots         uint8 `json:"upgradeslots" db:"upgradeslots"`
	Level                uint8 `json:"level" db:"level"`
	Str                  int16 `json:"str" db:"str"`
	Dex                  int16 `json:"dex" db:"dex"`
	Intelligence         int16 `json:"int" db:"int"`
	Luk                  int16 `json:"luk" db:"luk"`
	HP                   int16 `json:"hp" db:"hp"`
	MP                   int16 `json:"mp" db:"mp"`
	Watk                 int16 `json:"watk" db:"watk"`
	Matk                 int16 `json:"matk" db:"matk"`
	Wdef                 int16 `json:"wdef" db:"wdef"`
	Mdef                 int16 `json:"mdef" db:"mdef"`
	Acc                  int16 `json:"acc" db:"acc"`
	Avoid                int16 `json:"avoid" db:"avoid"`
	Hands                int16 `json:"hands" db:"hands"`
	Speed                int16 `json:"speed" db:"speed"`
	Jump                 int16 `json:"jump" db:"jump"`
	ViciousHammer        int8  `json:"ViciousHammer" db:"ViciousHammer"`
	ItemEXP              int   `json:"itemEXP" db:"itemEXP"`
	Durability           int   `json:"durability" db:"durability"`
	Enhance              uint8 `json:"enhance" db:"enhance"`
	Potential1           int16 `json:"potential1" db:"potential1"`
	Potential2           int16 `json:"potential2" db:"potential2"`
	Potential3           int16 `json:"potential3" db:"potential3"`
	HpR                  int16 `json:"hpR" db:"hpR"`
	MpR                  int16 `json:"mpR" db:"mpR"`
	ItemLevel            int16 `json:"itemlevel" db:"itemlevel"`
}

// Base 基础物品
func (i *ItemWithEquipment) Base() *Item {
	return &i.Item
}

// ---------------------------------------------------------------------------------------
// ------------------------------------ Pet ----------------------------------------------
// ---------------------------------------------------------------------------------------

// Pet 宠物信息
type Pet struct {
	// PetID 宠物ID
	PetID uint `json:"petid" db:"petid"`

	// Name 宠物名称
	Name string `json:"name" db:"name"`

	// Level 宠物等级
	Level uint `json:"level" db:"level"`

	// Closeness 亲密度
	Closeness uint `json:"closeness" db:"closeness"`

	// Fullness 饥饿度
	Fullness uint `json:"fullness" db:"fullness"`

	// Seconds 秒
	Seconds int `json:"seconds" db:"seconds"`

	// Flags 标志
	Flags int16 `json:"flags" db:"flags"`
}
