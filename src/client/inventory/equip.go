package inventory

import (
	"github.com/jerbe/goms/constants"
	"github.com/jerbe/goms/database/model"
	"log"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/31 15:52
  @describe :
*/

const (
	ArmorExpRatio  = 350000
	WeaponExpRatio = 700000
)

// NewEquipFromModel 从一个model生成一个装备
func NewEquipFromModel(base *model.ItemWithEquipment) *Equip {
	item := NewItem(base.Base())
	equip := &Equip{Item: *item}
	equip.equipLevel = byte(base.ItemLevel)
	return equip
}

// Equip 装备
type Equip struct {
	Item

	upgradeSlots  int8
	level         int8
	vicioushammer int8
	enhance       int8
	itemLevel     int8

	str int8
	dex int8

	//inte (intelligence) 智力
	inte       int8
	luk        int8
	hp         int8
	mp         int8
	watk       int8
	matk       int8
	wdef       int8
	mdef       int8
	acc        int8
	avoid      int8
	hands      int8
	speed      int8
	jump       int8
	potential1 int8
	potential2 int8
	potential3 int8
	hpR        int8
	mpR        int8
	charmExp   int8
	pvpDamage  int8

	itemEXP    int
	durability int
}

func (e *Equip) UpgradeSlots() int8 {
	return e.upgradeSlots
}

func (e *Equip) SetUpgradeSlots(upgradeSlots int8) {
	e.upgradeSlots = upgradeSlots
}

func (e *Equip) Level() int8 {
	return e.level
}

func (e *Equip) SetLevel(level int8) {
	e.level = level
}

func (e *Equip) Vicioushammer() int8 {
	return e.vicioushammer
}

func (e *Equip) SetVicioushammer(vicioushammer int8) {
	e.vicioushammer = vicioushammer
}

func (e *Equip) Enhance() int8 {
	return e.enhance
}

func (e *Equip) SetEnhance(enhance int8) {
	e.enhance = enhance
}

func (e *Equip) ItemLevel() int8 {
	return e.itemLevel
}

func (e *Equip) SetItemLevel(itemLevel int8) {
	e.itemLevel = itemLevel
}

func (e *Equip) Str() int8 {
	return e.str
}

func (e *Equip) SetStr(str int8) {
	if str < 0 {
		str = 0
	}
	e.str = str
}

func (e *Equip) Dex() int8 {

	return e.dex
}

func (e *Equip) SetDex(dex int8) {
	if dex < 0 {
		dex = 0
	}
	e.dex = dex
}

func (e *Equip) Inte() int8 {
	return e.inte
}

func (e *Equip) SetInte(inte int8) {
	if inte < 0 {
		inte = 0
	}
	e.inte = inte
}

func (e *Equip) Luk() int8 {
	return e.luk
}

func (e *Equip) SetLuk(luk int8) {
	if luk < 0 {
		luk = 0
	}
	e.luk = luk
}

func (e *Equip) Hp() int8 {
	return e.hp
}

func (e *Equip) SetHp(hp int8) {
	if hp < 0 {
		hp = 0
	}
	e.hp = hp
}

func (e *Equip) Mp() int8 {
	return e.mp
}

func (e *Equip) SetMp(mp int8) {
	if mp < 0 {
		mp = 0
	}
	e.mp = mp
}

func (e *Equip) Watk() int8 {
	return e.watk
}

func (e *Equip) SetWatk(watk int8) {
	if watk < 0 {
		watk = 0
	}
	e.watk = watk
}

func (e *Equip) Matk() int8 {
	return e.matk
}

func (e *Equip) SetMatk(matk int8) {
	if matk < 0 {
		matk = 0
	}
	e.matk = matk
}

func (e *Equip) Wdef() int8 {
	return e.wdef
}

func (e *Equip) SetWdef(wdef int8) {
	if wdef < 0 {
		wdef = 0
	}
	e.wdef = wdef
}

func (e *Equip) Mdef() int8 {
	return e.mdef
}

func (e *Equip) SetMdef(mdef int8) {
	if mdef < 0 {
		mdef = 0
	}
	e.mdef = mdef
}

func (e *Equip) Acc() int8 {
	return e.acc
}

func (e *Equip) SetAcc(acc int8) {
	if acc < 0 {
		acc = 0
	}
	e.acc = acc
}

func (e *Equip) Avoid() int8 {
	return e.avoid
}

func (e *Equip) SetAvoid(avoid int8) {
	if avoid < 0 {
		avoid = 0
	}
	e.avoid = avoid
}

func (e *Equip) Hands() int8 {

	return e.hands
}

func (e *Equip) SetHands(hands int8) {
	if hands < 0 {
		hands = 0
	}
	e.hands = hands
}

func (e *Equip) Speed() int8 {
	return e.speed
}

func (e *Equip) SetSpeed(speed int8) {
	if speed < 0 {
		speed = 0
	}
	e.speed = speed
}

func (e *Equip) Jump() int8 {
	return e.jump
}

func (e *Equip) SetJump(jump int8) {
	if jump < 0 {
		jump = 0
	}
	e.jump = jump
}

func (e *Equip) Potential1() int8 {
	return e.potential1
}

func (e *Equip) SetPotential1(potential1 int8) {
	e.potential1 = potential1
}

func (e *Equip) Potential2() int8 {
	return e.potential2
}

func (e *Equip) SetPotential2(potential2 int8) {
	e.potential2 = potential2
}

func (e *Equip) Potential3() int8 {
	return e.potential3
}

func (e *Equip) SetPotential3(potential3 int8) {
	e.potential3 = potential3
}

func (e *Equip) HpR() int8 {
	return e.hpR
}

func (e *Equip) SetHpR(hpR int8) {
	e.hpR = hpR
}

func (e *Equip) MpR() int8 {
	return e.mpR
}

func (e *Equip) SetMpR(mpR int8) {
	e.mpR = mpR
}

func (e *Equip) CharmExp() int8 {
	return e.charmExp
}

func (e *Equip) SetCharmExp(charmExp int8) {
	e.charmExp = charmExp
}

func (e *Equip) PvpDamage() int8 {
	return e.pvpDamage
}

func (e *Equip) SetPvpDamage(pvpDamage int8) {
	e.pvpDamage = pvpDamage
}

func (e *Equip) ItemEXP() int {
	return e.itemEXP
}

func (e *Equip) SetItemEXP(itemEXP int) {
	if itemEXP < 0 {
		itemEXP = 0
	}
	e.itemEXP = itemEXP
}

func (e *Equip) EquipExp() int {
	if e.itemEXP <= 0 {
		return 0
	}
	//aproximate value
	if constants.IsWeapon(e.itemId) {
		return e.itemEXP / WeaponExpRatio
	} else {
		return e.itemEXP / ArmorExpRatio
	}
}

func (e *Equip) EquipExpForLevel() int {
	//if e.EquipExp() <= 0 {
	//	return 0
	//}
	//expz := e.EquipExp()
	//for (i := e.BaseLevel(); i <= GameConstants.getMaxLevel(getItemId()); i++) {
	//	if (expz >= GameConstants.getExpForLevel(i, getItemId())) {
	//		expz -= GameConstants.getExpForLevel(i, getItemId());
	//	} else { //for 0, dont continue;
	//		break;
	//	}
	//}
	//return expz;
	return 0
}

func BaseLevel() int {
	//return (GameConstants.getStatFromWeapon(getItemId()) == null ? 1 : 0);
	return 0
}

func (e *Equip) Durability() int {
	return e.durability
}

func (e *Equip) SetDurability(durability int) {
	e.durability = durability
}

// getType 返回一个类型,1=装备
func (e *Equip) getType() byte {
	return 1
}

// SetQuantity 设置数量
func (s *Equip) SetQuantity(quantity int16) {
	if quantity < 0 || quantity > 1 {
		log.Println("notable quantity")
	}
}
