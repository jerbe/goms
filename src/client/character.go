package client

import (
	"errors"
	"github.com/jerbe/goms/client/inventory"
	"github.com/jerbe/goms/constants"
	"github.com/jerbe/goms/database"
	"github.com/jerbe/goms/database/model"
	"strconv"
	"strings"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/29 17:42
  @describe :
*/

type Character struct {
	model.Character

	// loadFromDB 是否已经从数据库中读取数据了
	loadFromDB bool

	// skillPointList 技能点列表
	skillPointList []int

	// 宠物列表
	petList []*inventory.Pet

	// inventoryMap 库存映射表
	inventoryMap map[inventory.Type]*inventory.Inventory
}

// NewCharacter 根据id返回一个空的角色
func NewCharacter(id int64) *Character {
	char := &Character{}
	char.ID = id
	return char
}

// NewCharacterFromModel 从model中返回一个角色信息
func NewCharacterFromModel(m *model.Character) *Character {
	char := &Character{
		Character: *m,
	}
	return char
}

// LoadFromDB 从数据库中加载数据
// loggedin 登入游戏场景界面
func (char *Character) LoadFromDB(loggedin bool) error {
	if char.loadFromDB && loggedin {
		return errors.New("already load from database")
	}

	// @TODO 从数据库中读出来所有数据
	if loggedin {
		// 加载地图 map

		// 生成点SpawnPoint并给角色指定位置

		// 加载派对

		// 加载宠物信息

		// 加载成就

		// 加载任务信息

		// 加载库存物品
		// 设置各个库存信息

		// 加载技能数据

		// 处理仙女的祝福

		// 设置技能宏

		// 设置键盘映射

		// 设置保存位置

		// 设置荣誉

		// 设置愿望清单

		// 设置干燥地点?

		// 设置重置地点?

		// 设置装载数据?
	} else {
		err := char.loadInventory(true)
		if err != nil {
			return err
		}
	}
	char.loadFromDB = true

	return nil
}

// resetDefault 重置到默认,就是清空掉所有数据
func (char *Character) resetDefault() {

}

// loadInventory 加载库存物品
func (char *Character) loadInventory(equippedOnly bool) error {
	char.inventoryMap = make(map[inventory.Type]*inventory.Inventory)
	for i := 0; i < len(inventory.TypeList); i++ {
		t := inventory.TypeList[i]
		char.inventoryMap[t] = inventory.NewInventory(t, 100)
	}

	// 加载库存物品
	itemsList, err := database.CharacterInventoryWithEquipmentItems(equippedOnly, char.ID)
	if err != nil {
		return err
	}
	itemMap := database.WithEquipmentItemListToMap(itemsList)

	for key, items := range itemMap {
		inventoryType := inventory.GetType(key)
		myInventory := char.GetInventory(inventoryType)

		for i := 0; i < len(items); i++ {
			var item inventory.IItem
			if inventoryType == inventory.TypeEquip || inventoryType == inventory.TypeEquipped {
				item = inventory.NewEquipFromModel(items[i])
			} else {
				item = inventory.NewItem(items[i].Base())
				if constants.IsPet(item.ItemId()) {
					// 包含宠物
					if item.UniqueId() > -1 {
						char.petList = append(char.petList, inventory.NewPetFromModel(items[i].Base()))
					} else {

					}
				}
			}
			myInventory.AddItem(item)
		}
	}

	return nil
}

// SkillPointList 获取已设置的技能点列表
func (char *Character) SkillPointList() []int {
	if char.skillPointList == nil && char.Sp != "" {
		sps := strings.Split(char.Sp, ",")
		char.skillPointList = make([]int, 10)
		for i := 0; i < len(sps); i++ {
			v, _ := strconv.ParseInt(sps[i], 10, 64)
			char.skillPointList[i] = int(v)
		}
	}
	return char.skillPointList
}

// SkillPointRemaining 剩余的技能点数
func (char *Character) SkillPointRemaining() int {
	return char.SpRemainingBySkillBook(constants.GetSkillBook(char.Job))
}

// SpRemainingBySkillBook 根据技能书剩余的
func (char *Character) SpRemainingBySkillBook(bookID int) int {
	return char.SkillPointList()[bookID]
}
