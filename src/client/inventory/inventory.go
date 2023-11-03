package inventory

import (
	"sort"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/30 09:19
  @describe : 库存库
*/

// Type 库存类型
type Type int8

// BitfieldEncoding 位域编码
func (t Type) BitfieldEncoding() int16 {
	return 2 << t
}

// TypeList 所有已经定义的库存类型
// 注意,不要修改
var TypeList = []Type{TypeUndefined, TypeEquip, TypeUse, TypeSetup, TypeEtc, TypeCash, TypeEquipped}

// TypeEquipped 表示物品已经装备在角色身上。这是一个状态而不是库存类型，用于指示某个物品当前是否被角色装备。
const TypeEquipped Type = -1 //
const (
	// TypeUndefined 表示未定义或未分类的库存类型。这通常用作默认值或错误情况下的类型。
	TypeUndefined Type = iota

	// TypeEquip 表示装备库存类型。这包括玩家可以穿戴的物品，如武器、防具等。
	TypeEquip

	// TypeUse 表示消耗品库存类型。这包括玩家可以使用一次或多次的物品，如药水、卷轴等。
	TypeUse

	// TypeSetup 表示设置库存类型。这包括物品，通常用于角色的设置或定制，如技能书或角色特定的物品。
	TypeSetup

	// TypeEtc 表示装饰品库存类型。这包括一些不影响角色属性但可以穿戴的物品，如戒指、耳环等。
	TypeEtc

	// TypeCash 表示现金库存类型。这包括通过游戏商城或购买获得的物品，通常需要虚拟货币或真实货币购买。
	TypeCash
)

// GetType 根据数字获取库存类型
func GetType(kind int8) Type {
	for _, b := range TypeList {
		if int8(b) == kind {
			return b
		}
	}
	return TypeUndefined
}

// GetTypeByWZName 根据wz名称获取类型
func GetTypeByWZName(name string) Type {
	switch name {
	case "Install":
		return TypeSetup
	case "Consume":
		return TypeUse
	case "Etc":
		return TypeEtc
	case "Cash":
		return TypeCash
	case "Pet":
		return TypeCash
	}
	return TypeUndefined
}

// ------ Inventory 库存

const (
	// 最大的库存槽位数量限制
	maxInventorySlotLimit = 1024
)

// NewInventory 返回一个库存(或则背包)
func NewInventory(typ Type, slotsLimit int) *Inventory {
	return &Inventory{
		typ:        typ,
		slotsLimit: uint(slotsLimit),
		slotsMap:   make(map[int16]IItem),
	}
}

// Inventory 库存信息
type Inventory struct {
	// typ 库存类型
	typ Type

	// slotsLimit 槽位限制
	slotsLimit uint

	// 槽位映射表
	slotsMap map[int16]IItem
}

// Type 获取背包的类型
func (ivt *Inventory) Type() Type {
	return ivt.typ
}

// AddItem 增加物品
func (ivt *Inventory) AddItem(items ...IItem) {
	for i := 0; i < len(items); i++ {
		item := items[i]
		ivt.slotsMap[item.Position()] = item
	}
}

// Items 返回所有数据
func (ivt *Inventory) Items() []IItem {
	items := make([]IItem, 0, len(ivt.slotsMap))
	var i = 0
	for _, item := range ivt.slotsMap {
		items = append(items, item)
		i++
	}
	return items
}

// AddSlot 增加槽位
func (ivt *Inventory) AddSlot(count uint) {
	ivt.slotsLimit += count
	if ivt.slotsLimit > maxInventorySlotLimit {
		ivt.slotsLimit = maxInventorySlotLimit
	}
}

// SetSlotLimit 设置槽位限制数
func (ivt *Inventory) SetSlotLimit(limit uint) {
	ivt.slotsLimit = limit
	if ivt.slotsLimit > maxInventorySlotLimit {
		ivt.slotsLimit = maxInventorySlotLimit
	}
}

// GetSlotLimit 获取槽位限制数
func (ivt *Inventory) GetSlotLimit() uint {
	return ivt.slotsLimit
}

// GetItemById 根据物品ID获取指定物品.如果库存中存在该物品，则返回该物品，否则返回 null
func (ivt *Inventory) GetItemById(itemID int) (IItem, bool) {
	for _, item := range ivt.slotsMap {
		if item.ItemId() == itemID {
			return item, true
		}
	}
	return nil, false
}

// GetItemQuantitiesById 根据物品ID获取指定物品的持有总数量
func (ivt *Inventory) GetItemQuantitiesById(itemID int) int {
	total := 0
	for _, item := range ivt.slotsMap {
		if item.ItemId() == itemID {
			total += int(item.Quantity())
		}
	}
	return total
}

// GetItemListByItemId 根据物品ID获取指定物品列表,因为每个物品在每个卡槽都都有存放数量显示
func (ivt *Inventory) GetItemListByItemId(itemID int) []IItem {
	itemList := make(ItemList, 0)
	for _, item := range ivt.slotsMap {
		if item.ItemId() == itemID {
			itemList = append(itemList, item)
		}
	}

	// 需要排序
	sort.Sort(itemList)

	return itemList
}

// GetItemByUniqueId  根据物品唯一ID获取指定物品.如果库存中存在该物品，则返回该物品，否则返回 null
func (ivt *Inventory) GetItemByUniqueId(uniqueId int) (IItem, bool) {
	for _, item := range ivt.slotsMap {
		if item.UniqueId() == uniqueId {
			return item, true
		}
	}
	return nil, false
}

// IsFull 是否已经满了
func (ivt *Inventory) IsFull() bool {
	return len(ivt.slotsMap) >= int(ivt.slotsLimit)
}

// CheckFull 提前检测背包是否会满
func (ivt *Inventory) CheckFull(margin int) bool {
	return len(ivt.slotsMap) >= margin+int(ivt.slotsLimit)
}

// GetNextFreeSlot 获取下一个有空位的卡槽位置
func (ivt *Inventory) GetNextFreeSlot() int {
	if ivt.IsFull() {
		return -1
	}
	for j := uint(0); j < ivt.slotsLimit; j++ {
		if _, ok := ivt.slotsMap[int16(j)]; !ok {
			return int(j)
		}
	}
	return -1
}

// GetFreeSlotNum 获取空闲槽位的数量
func (ivt *Inventory) GetFreeSlotNum() int {
	if ivt.IsFull() {
		return 0
	}
	freeSlot := 0
	for j := uint(0); j < ivt.slotsLimit; j++ {
		if _, ok := ivt.slotsMap[int16(j)]; !ok {
			freeSlot++
		}
	}
	return freeSlot
}
