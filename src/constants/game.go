package constants

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/29 11:37
  @describe :
*/

// GetSkillBook 获取技能书
// jobId 职业ID
func GetSkillBook(jobId int) int {
	if jobId >= 2210 && jobId <= 2218 {
		return jobId - 2209
	}
	switch jobId {
	case 3210, 3310, 3510:
		return 1
	case 3211, 3311, 3511:
		return 2
	case 3212, 3312, 3512:
		return 3
	}
	return 0
}

// IsPet 判断是否是宠物
// itemId 物品ID
func IsPet(itemId int) bool {
	return itemId/10000 == 500
}

// IsRing 判断是否是戒指
// itemId 物品ID
func IsRing(itemId int) bool {
	return itemId/10000 == 500
}

// IsWeapon 是否是武器
func IsWeapon(itemId int) bool {
	return itemId >= 1300000 && itemId < 1500000
}

// IsTimelessItem 是否永久的物品
func IsTimelessItem(itemId int) bool {
	switch itemId {
	case 1032031, //shield earring, but technically
		1102172,
		1002776,
		1002777,
		1002778,
		1002779,
		1002780,
		1082234,
		1082235,
		1082236,
		1082237,
		1082238,
		1052155,
		1052156,
		1052157,
		1052158,
		1052159,
		1072355,
		1072356,
		1072357,
		1072358,
		1072359,
		1092057,
		1092058,
		1092059,
		1122011,
		1122012,
		1302081,
		1312037,
		1322060,
		1332073,
		1332074,
		1372044,
		1382057,
		1402046,
		1412033,
		1422037,
		1432047,
		1442063,
		1452057,
		1462050,
		1472068,
		1482023,
		1492023,
		1342011:
		return true
	default:
		return false
	}
}

// IsReverseItem 是个反向的物品
func IsReverseItem(itemId int) bool {
	switch itemId {
	case 1002790,
		1002791,
		1002792,
		1002793,
		1002794,
		1082239,
		1082240,
		1082241,
		1082242,
		1082243,
		1052160,
		1052161,
		1052162,
		1052163,
		1052164,
		1072361,
		1072362,
		1072363,
		1072364,
		1072365,
		1302086,
		1312038,
		1322061,
		1332075,
		1332076,
		1372045,
		1382059,
		1402047,
		1412034,
		1422038,
		1432049,
		1442067,
		1452059,
		1462051,
		1472071,
		1482024,
		1492025,
		1342012:
		return true
	default:
		return false
	}
}

// ItemMaxLevel 物品的最大等级
func ItemMaxLevel(itemId int) int {
	if IsTimelessItem(itemId) {
		return 5
	} else if IsReverseItem(itemId) {
		return 3
	} else {
		switch itemId {
		case 1302109,
			1312041,
			1322067,
			1332083,
			1372048,
			1382064,
			1402055,
			1412037,
			1422041,
			1432052,
			1442073,
			1452064,
			1462058,
			1472079,
			1482035,
			1302108,
			1312040,
			1322066,
			1332082,
			1372047,
			1382063,
			1402054,
			1412036,
			1422040,
			1432051,
			1442072,
			1452063,
			1462057,
			1472078,
			1482036:
			return 1
		case 1072376:
			return 2
		}
	}
	return 0
}

//
//func StatFromWeapon(itemId int) (status.MonsterStatus, bool) {
//	switch itemId {
//	case 1302109,
//		1312041,
//		1322067,
//		1332083,
//		1372048,
//		1382064,
//		1402055,
//		1412037,
//		1422041,
//		1432052,
//		1442073,
//		1452064,
//		1462058,
//		1472079,
//		1482035:
//		return MonsterStatus.命中
//	case 1302108:
//	case 1312040:
//	case 1322066:
//	case 1332082:
//	case 1372047:
//	case 1382063:
//	case 1402054:
//	case 1412036:
//	case 1422040:
//	case 1432051:
//	case 1442072:
//	case 1452063:
//	case 1462057:
//	case 1472078:
//	case 1482036:
//		return MonsterStatus.速度
//	}
//	return null
//}
