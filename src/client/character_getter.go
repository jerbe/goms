package client

import (
	"github.com/jerbe/goms/client/inventory"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/31 09:21
  @describe :
*/

// GetInventory 根据库存类型获取指定的库存数据
func (char *Character) GetInventory(inventoryType inventory.Type) *inventory.Inventory {
	return char.inventoryMap[inventoryType]
}

// GetPetList 获取所有宠物列表
func (char *Character) GetPetList() []*inventory.Pet {
	return char.petList
}

// GetSlotPet 获取指定宠物栏内的宠物
func (char *Character) GetSlotPet(slot int) *inventory.Pet {
	if slot < 0 || slot >= len(char.petList) {
		return nil
	}
	return char.petList[slot]
}
