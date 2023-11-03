package database

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jerbe/goms/database/model"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/31 12:05
  @describe :
*/

const (
	itemTypeInventory = iota
	itemTypeStorage
	itemTypeCashShopExplorer
	itemTypeCashShopCygnus
	itemTypeCashShopAran
	itemTypeHiredMerchant
	itemTypeDuey
	itemTypeCashShopEvan
	itemTypeMTS
	itemTypeMTSTransfer
	itemTypeCashShopDB
	itemTypeCashShopResist
)

type fieldKV struct {
	Key   string
	Value any
}

// WithEquipmentItemListToMap 将包含已装备库存列表转换成映射表
func WithEquipmentItemListToMap(list []*model.ItemWithEquipment) map[int8][]*model.ItemWithEquipment {
	result := make(map[int8][]*model.ItemWithEquipment)
	for _, item := range list {
		t := int8(item.InventoryType)
		if _, ok := result[t]; !ok {
			result[t] = make([]*model.ItemWithEquipment, 0)
		}
		subList := result[t]
		subList = append(subList, item)
		result[t] = subList
	}
	return result
}

// WithEquipmentItems 包含已装备的物品列表
func WithEquipmentItems(table, equipTable string, typ int, equippedOnly bool, filterFields []fieldKV) ([]*model.ItemWithEquipment, error) {
	return Default.WithEquipmentItems(table, equipTable, typ, equippedOnly, filterFields)
}

// WithEquipmentItems 包含已装备的物品列表
func (db *database) WithEquipmentItems(table, equipTable string, typ int, equippedOnly bool, filterFields []fieldKV) ([]*model.ItemWithEquipment, error) {

	if len(filterFields) == 0 {
		return nil, errors.New("database.WithEquipmentItems filterFields not set")
	}

	whereSqls := make([]string, 0, len(filterFields))
	whereArgs := make([]any, 0, len(filterFields)+1)
	whereArgs = append(whereArgs, typ)
	for i := 0; i < len(filterFields); i++ {
		field := filterFields[i]
		whereSqls = append(whereSqls, fmt.Sprintf(" `%s` = ? ", field.Key))
		whereArgs = append(whereArgs, field.Value)
	}

	if equippedOnly {
		whereSqls = append(whereSqls, " `inventorytype` = ? ")
		whereArgs = append(whereArgs, -1) // -1 => client.inventory.TypeEquipped
	}

	sqlQuery := fmt.Sprintf("SELECT %s FROM %s LEFT JOIN %s USING (`inventoryitemid`) WHERE `type` = ? AND %s", model.FieldsItems, table, equipTable, strings.Join(whereSqls, " AND "))

	result := make([]*model.ItemWithEquipment, 0)
	err := db.sqlDB.Select(&result, sqlQuery, whereArgs...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ------
// -----------------------------------  CharacterInventoryItemsWithEquipment -----------------------------------
// ------

// CharacterInventoryWithEquipmentItems 获取角色的所有库存
// charID 角色ID
func CharacterInventoryWithEquipmentItems(equippedOnly bool, charID int64) ([]*model.ItemWithEquipment, error) {
	return Default.CharacterInventoryWithEquipmentItems(equippedOnly, charID)
}

// CharacterInventoryWithEquipmentItems 获取角色的所有库存
// charID 角色ID
func (db *database) CharacterInventoryWithEquipmentItems(equippedOnly bool, charID int64) ([]*model.ItemWithEquipment, error) {
	return Default.WithEquipmentItems(model.TableInventoryItems,
		model.TableInventoryEquipment,
		itemTypeInventory,
		equippedOnly,
		[]fieldKV{
			{
				Key:   "characterid",
				Value: charID,
			},
		})
}

// ------
// -----------------------------------  AccountInventoryItemsWithEquipment -----------------------------------
// ------

// AccountInventoryWithEquipmentItems 获取账户已战备的的所有库存包含
// accountID 账户ID
func AccountInventoryWithEquipmentItems(equippedOnly bool, accountID int64) ([]*model.ItemWithEquipment, error) {
	return Default.AccountInventoryWithEquipmentItems(equippedOnly, accountID)
}

// AccountInventoryWithEquipmentItems 获取账户已战备的的所有库存包含
// accountID 账户ID
func (db *database) AccountInventoryWithEquipmentItems(equippedOnly bool, accountID int64) ([]*model.ItemWithEquipment, error) {
	return Default.WithEquipmentItems(model.TableInventoryItems,
		model.TableInventoryEquipment,
		itemTypeInventory,
		equippedOnly,
		[]fieldKV{
			{
				Key:   "accountid",
				Value: accountID,
			},
		})
}
