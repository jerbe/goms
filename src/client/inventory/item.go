package inventory

import (
	"github.com/jerbe/goms/database/model"
	"github.com/jerbe/goms/utils"
	"math"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/29 21:18
  @describe :
*/

// ItemList 物品列表类,实现 sort.Interface 接口
type ItemList []IItem

func (list ItemList) Len() int {
	return len(list)
}

func (list ItemList) Less(i, j int) bool {
	a := list[i]
	b := list[j]
	if math.Abs(float64(a.Position())) < math.Abs(float64(b.Position())) {
		return true
	}
	return false
}

func (list ItemList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// IItem 物品基础接口
type IItem interface {
	utils.Comparable

	Type() byte
	ItemId() int

	Position() int16
	SetPosition(position int16)

	SetFlag(flag byte)
	Flag() byte

	SetLocked(flag byte)
	IsLocked() bool

	SetQuantity(quantity int16)
	Quantity() int16

	SetOwner(owner string)
	Owner() string

	SetGMLog(GameMasterLog string)
	GMLog() string

	SetUniqueId(id int)
	UniqueId() int

	SetExpiration(expire int64)
	Expiration() int64

	SetGiftFrom(gf string)
	GiftFrom() string

	SetEquipLevel(j byte)
	EquipLevel() byte

	Copy() IItem
}

// ----------- Item

func NewItem(data *model.Item) *Item {
	item := &Item{
		position:   int16(data.Position),
		flag:       byte(data.Flag),
		quantity:   int16(data.Quantity),
		itemId:     data.ItemID,
		uniqueId:   data.UniqueID,
		expiration: data.ExpireDate,
		giftFrom:   data.Sender,
	}
	if data.Owner != nil {
		item.owner = *data.Owner
	}
	if data.GMLog != nil {
		item.gmLog = *data.GMLog
	}
	return item
}

var _ IItem = new(Item)

// Item 物品
type Item struct {
	// position 位置
	position int16

	// flag 标记
	flag byte

	// isLocked 是否已经锁定
	isLocked bool

	// quantity 数量
	quantity int16

	// owner 拥有人
	owner string

	// gmLog 用于gm显示的日志文案
	gmLog string

	// itemId 物品ID
	itemId int

	// uniqueId 唯一标识
	uniqueId int

	// expiration 有效时长还是到期日期?
	expiration int64

	// giftFrom(sender) 物品来源
	giftFrom string

	// equipLevel 装备等级
	equipLevel byte
}

// Equals 进行相等比较
func (i *Item) Equals(obj any) bool {
	item, ok := obj.(IItem)
	if !ok {
		return false
	}

	return i.uniqueId == item.UniqueId() && i.itemId == item.ItemId() && i.quantity == item.Quantity() && i.position == item.Position()
}

// Type 获取类型 2 = 物品
func (i *Item) Type() byte {
	return 2
}

// SetPosition 设置物品位置(背包的?)
func (i *Item) SetPosition(position int16) {
	i.position = position
}

// Position 获取位置(背包的?)
func (i *Item) Position() int16 {
	return i.position
}

// SetFlag 设置标记
func (i *Item) SetFlag(flag byte) {
	i.flag = flag
}

// Flag 获取标记
func (i *Item) Flag() byte {
	return i.flag
}

// SetLocked 锁定
func (i *Item) SetLocked(flag byte) {
	i.flag = flag
}

// IsLocked 判断是否已经锁定
func (i *Item) IsLocked() bool {
	return i.isLocked
}

// SetQuantity 设置数量
func (i *Item) SetQuantity(quantity int16) {
	i.quantity = quantity
}

// Quantity 获取数量
func (i *Item) Quantity() int16 {
	return i.quantity
}

// SetOwner 设置拥有人
func (i *Item) SetOwner(owner string) {
	i.owner = owner
}

// Owner  拥有人
func (i *Item) Owner() string {
	return i.owner
}

// SetGMLog 设置用于gm显示的日志文案
func (i *Item) SetGMLog(logStr string) {
	i.gmLog = logStr
}

// GMLog 用于gm显示的日志文案
func (i *Item) GMLog() string {
	return i.gmLog
}

// ItemId 物品ID
func (i *Item) ItemId() int {
	return i.itemId
}

func (i *Item) SetUniqueId(id int) {
	i.uniqueId = id
}

// UniqueId 唯一标识
func (i *Item) UniqueId() int {
	return i.uniqueId
}

// SetExpiration 设置有效时长还是到期日期
func (i *Item) SetExpiration(expire int64) {
	i.expiration = expire
}

// Expiration 有效时长还是到期日期
func (i *Item) Expiration() int64 {
	return i.expiration
}

// SetGiftFrom 设置物品来源
func (i *Item) SetGiftFrom(sender string) {
	i.giftFrom = sender
}

// GiftFrom 获取物品来源
func (i *Item) GiftFrom() string {
	return i.giftFrom
}

// SetEquipLevel 设置装备等级
func (i *Item) SetEquipLevel(level byte) {
	i.equipLevel = level
}

// EquipLevel 获取装备等级
func (i *Item) EquipLevel() byte {
	return i.equipLevel
}

func (i *Item) Copy() IItem {
	panic("未实现")
}

// IsPet 判断一个物品是否是宠物
func IsPet(item IItem) (*Pet, bool) {
	pet, ok := item.(*Pet)
	return pet, ok
}

// IsRing 判断一个物品是否是戒指
func IsRing(item IItem) (*Ring, bool) {
	ring, ok := item.(*Ring)
	return ring, ok
}
