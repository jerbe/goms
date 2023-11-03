package packet

import (
	"github.com/jerbe/goms/constants"
	"sort"
	"time"

	"github.com/jerbe/goms/client"
	"github.com/jerbe/goms/client/inventory"
	"github.com/jerbe/goms/data"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/11/1 23:16
  @describe :
*/

// appendCharacterEntryList 在writer中追加角色入口列表信息
func appendCharacterEntryList(writer *data.LittleEndianWriter, charList []*client.Character, showAll bool) *data.LittleEndianWriter {
	for i := 0; i < len(charList); i++ {
		char := charList[i]
		appendCharacterEntry(writer, char, char.GetGm() != 1 && char.GetLevel() >= 10, showAll)
	}
	return writer
}

// appendCharacterEntry 在writer中追加角色入口
func appendCharacterEntry(writer *data.LittleEndianWriter, char *client.Character, ranking, showAll bool) *data.LittleEndianWriter {
	appendCharacterBaseInfo(writer, char)
	appendCharacterLook(writer, char, true, showAll).
		WriteByte(0)

	// 这里判断职业? 900是GM? 143版本是判断Job是否是GM
	if char.GetJob() == 900 {
		writer.WriteByte(2)
	}

	return writer
}

// appendCharacterBaseInfo 在writer中追加角色基础信息
func appendCharacterBaseInfo(writer *data.LittleEndianWriter, char *client.Character) *data.LittleEndianWriter {
	writer.WriteInt(int(char.GetID())). //character id
						WriteAsciiString(char.GetName(), 13). // character name
						WriteByte(byte(char.GetGender())).    // gender (0 = male, 1 = female)
						WriteByte(byte(char.GetSkinColor())). // skin color
						WriteInt(char.GetFace()).             // face
						WriteInt(char.GetHair()).             //hair
						WriteZeroBytes(24).
						WriteByte(byte(char.GetLevel())).              //level
						WriteShort(int16(char.GetJob())).              //job 职业
						WriteShort(int16(char.GetStr())).              // str 力量
						WriteShort(int16(char.GetDex())).              // dex 敏捷
						WriteShort(int16(char.GetInt())).              // int(intelligence) 智力
						WriteShort(int16(char.GetLuk())).              // luk 幸运
						WriteShort(int16(char.GetHp())).               // hp 健康值
						WriteShort(int16(char.GetMaxHp())).            // maxhp 最大健康值
						WriteShort(int16(char.GetMp())).               // mp 魔法值
						WriteShort(int16(char.GetMaxMp())).            // maxmp 魔法值
						WriteShort(int16(char.GetAp())).               // Ap
						WriteShort(int16(char.SkillPointRemaining())). // Sp
						WriteInt(char.GetExp()).                       // Exp 经验
						WriteShort(int16(char.GetFame())).             //fame 声望
						WriteInt(0).                                   // Gachapon exp 扭蛋经验值?
						WriteLong(time.Now().UnixNano()).
						WriteInt(char.GetMapID()). // current map id
						WriteByte(byte(char.GetSpawnPoint()))
	return writer
}

// appendCharacterLook 在writer中追加角色外观信息
// mega 百宝袋?(Megaphone?)
// showAll 显示所有信息包括宠物
func appendCharacterLook(writer *data.LittleEndianWriter, char *client.Character, mega, showAll bool) *data.LittleEndianWriter {

	megaByte := byte(0)
	if mega {
		megaByte = 1
	}
	writer.WriteByte(byte(char.GetGender())).
		WriteByte(byte(char.GetSkinColor())).
		WriteInt(char.GetFace()).
		WriteByte(megaByte).
		WriteInt(char.GetHair())

	// 开始填充装备
	// 获取已装备在身的库存物品
	equipped := char.GetInventory(inventory.TypeEquipped)

	myEquip := make(map[int8]int)
	maskedEquip := make(map[int8]int)

	// weapon 武器
	var weapon inventory.IItem

	/*
		01[1,101]. 头饰/帽子
		02[2,102]. 脸饰
		03[3,103]. 眼饰
		04[4,104. 耳饰
		05[5,105]. 上衣(套服?)
		06[6,106]. 裤子/裙子
		07[7,107]. 鞋子
		08[8,108]. 手套
		09[9,109]. 披风
		10[10,110]. 未知
		11[11,111]. 点装武器(可以装备到主武器上给武器加成?)
		12[12,112]. 戒指
		13[13,113]. 戒指
		14[14,114]. (未知)
		15[15,115]. 戒指
		17[16,116]. 坠子
		18[17,117]. 骑兽
		19[18,118]. 鞍子

		26. 勋章
		27. 戒指
		28. 戒指
		29. 腰带
	*/

	/*
		for _, item := range equipped.Items() {
			if item.Position() < 0 {
				continue
			}
			pos := byte(item.Position())
			if _, ok := myEquip[pos]; !ok && pos < 100 {
				myEquip[pos] = item.ItemId()
			}

			if pos > 100 && pos != 111 {
				pos -= 100
				if itemId, ok := myEquip[pos]; ok {
					maskedEquip[pos] = itemId
				}
				myEquip[pos] = item.ItemId()
			}

			if pos == 111 {
				weapon = item
			}
		}
	*/
	for _, item := range equipped.Items() {
		if item.Position() < -128 {
			continue
		}

		pos := int8(item.Position() * -1)
		if _, ok := myEquip[pos]; !ok && pos < 100 {
			myEquip[pos] = item.ItemId()
		} else if (pos > 100 || pos == -128) && pos != 111 {
			if pos == -128 {
				pos = 28
			} else {
				pos -= 100
			}

			if i, ok := myEquip[pos]; ok {
				maskedEquip[pos] = i
			}
			myEquip[pos] = item.ItemId()
		} else if _, ok := myEquip[pos]; ok {
			maskedEquip[pos] = item.ItemId()
		}
	}

	// 可见项展示
	for b, i := range myEquip {
		writer.WriteByte(uint8(b)).
			WriteInt(i)
	}
	writer.WriteByte(0xFF)
	// 可见项展示完毕

	// 遮罩(面具?)类展示
	for b, i := range maskedEquip {
		writer.WriteByte(byte(b)).WriteInt(i)
	}
	writer.WriteByte(0xFF)
	// 结束(面具?)遮罩类

	// 武器展示
	if weapon != nil {
		writer.WriteInt(weapon.ItemId())
	} else {
		writer.WriteInt(0)
	}
	// 结束武器展示

	// 宠物展示
	for i := 0; i < 3; i++ {
		if showAll {
			itemId := 0
			if pet := char.GetSlotPet(i); pet != nil {
				itemId = pet.ItemId()
			}
			writer.WriteInt(itemId)
		} else {
			writer.WriteInt(0)
		}
	}

	return writer
}

// appendCharacter 追加角色信息
func appendCharacterInfo(writer *data.LittleEndianWriter, char *client.Character) {
	// (Basic)基础信息, 好友数量
	writer.WriteLong(-1).Write(0)
	appendCharacterBaseInfo(writer, char).Write(byte(char.BuddyCapacity)).
		// (Bless)保佑?
		Write(1)
	// (Buddy)好友信息
	appendInventoryInfo(writer, char)

	// (Inventory)库存信息
	// (Skill)技能信息
	// (CoolDown)冷却信息
	// (Quest)任务信息
	// (Ring)戒指信息
	// (Rocks)岩石信息?
	// (MonsterBook)怪物树?
}

func appendInventoryInfo(writer *data.LittleEndianWriter, char *client.Character) {
	writer.WriteMapleAsciiString(char.Name).
		WriteInt(char.Meso).
		WriteInt(int(char.ID)).
		WriteInt(char.Beans).
		WriteInt(0).
		Write(byte(char.GetInventory(inventory.TypeEquip).GetSlotLimit())).
		Write(byte(char.GetInventory(inventory.TypeUse).GetSlotLimit())).
		Write(byte(char.GetInventory(inventory.TypeSetup).GetSlotLimit())).
		Write(byte(char.GetInventory(inventory.TypeEtc).GetSlotLimit())).
		Write(byte(char.GetInventory(inventory.TypeCash).GetSlotLimit())).
		WriteLong(time.Now().UnixNano())

	equippedItems := char.GetInventory(inventory.TypeEquipped).Items()

	sort.Sort(inventory.ItemList(equippedItems))

	for i := 0; i < len(equippedItems); i++ {
		item := equippedItems[i]
		if -100 < item.Position() && item.Position() < 0 {

		}
	}
}

func appendItemInfo(writer *data.LittleEndianWriter, item inventory.IItem, zeroPosition bool, leaveOut bool, cs bool) {
	if item.UniqueId() > 0 { // 唯一的特殊物品
		if constants.IsPet(item.ItemId()) { //判断是否是宠物

		} else if constants.IsRing(item.ItemId()) { // 判断是否是戒指

		} else { // 否则是点装

		}
	} else { //平常的物品
		appendNormalItemInfo(writer, item, zeroPosition, leaveOut, cs)
	}
}

func appendNormalItemInfo(writer *data.LittleEndianWriter, item inventory.IItem, zeroPosition bool, leaveOut bool, cs bool) {

}
