package database

import (
	"fmt"
	"github.com/jerbe/goms/database/model"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/31 10:41
  @describe :
*/

// -----------------------------------------------------------------------------------------
// ------------------------------------ Characters -----------------------------------------
// -----------------------------------------------------------------------------------------

// Character 获取角色信息
func Character(charId int64) (*model.Character, error) {
	return Default.Character(charId)
}

// Character 获取角色信息
func (db *database) Character(charId int64) (*model.Character, error) {
	sqlQuery := fmt.Sprintf("SELECT %s FROM %s WHERE `id` = ? ", model.FieldsCharacters, model.TableCharacters)

	result := new(model.Character)
	err := db.sqlDB.Get(&result, sqlQuery, charId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CharactersList 获取账户的所有角色信息
func CharactersList(accID int64, worldID int) ([]*model.Character, error) {
	return Default.CharactersList(accID, worldID)
}

// CharactersList 获取账户的所有角色信息
func (db *database) CharactersList(accID int64, worldID int) ([]*model.Character, error) {
	sqlQuery := fmt.Sprintf("SELECT %s FROM %s WHERE `accountid` = ? AND `world` = ? ", model.FieldsCharacters, model.TableCharacters)

	result := make([]*model.Character, 0)
	err := db.sqlDB.Select(&result, sqlQuery, accID, worldID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// -----------------------------------------------------------------------------------------
// --------------------------------- CharacterSlots ----------------------------------------
// -----------------------------------------------------------------------------------------

// CharactersSlots 获取账户的指定世界可用角色卡槽数量
func CharactersSlots(accountID int64, worldID int) (int, error) {
	return Default.CharactersSlots(accountID, worldID)
}

// CharactersSlots 获取账户的指定世界可用角色卡槽数量
func (db *database) CharactersSlots(accountID int64, worldID int) (int, error) {
	sqlQuery := fmt.Sprintf("SELECT %s FROM %s WHERE accid = ? AND worldid = ?", "`charslots`", model.TableCharacterSlots)

	result := 0
	err := db.sqlDB.Get(&result, sqlQuery, accountID, worldID)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// AddCharactersSlots 新增账户的该世界可用卡槽数量
func AddCharactersSlots(accountID int64, worldID int, slots int) (int64, error) {
	return Default.AddCharactersSlots(accountID, worldID, slots)
}

// AddCharactersSlots 新增账户的该世界可用卡槽数量
func (db *database) AddCharactersSlots(accountID int64, worldID int, slots int) (int64, error) {
	sqlQuery := fmt.Sprintf("INSERT INTO %s (`accid`, `worldid`, `charslots`) VALUES (?, ?, ?)", model.TableCharacterSlots)

	result, err := db.sqlDB.Exec(sqlQuery, accountID, worldID, slots)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// -----------------------------------------------------------------------------------------
// -------------------------------------- Buddies ------------------------------------------
// -----------------------------------------------------------------------------------------

// BuddiesWithCharacterInfo 含有角色信息的好友资料
func BuddiesWithCharacterInfo(characterID int64) ([]*model.BuddyWithCharacterInfo, error) {
	return Default.BuddiesWithCharacterInfo(characterID)
}

// BuddiesWithCharacterInfo 含有角色信息的好友资料
func (db *database) BuddiesWithCharacterInfo(characterID int64) ([]*model.BuddyWithCharacterInfo, error) {
	sqlQuery := fmt.Sprintf("SELECT b.buddyid, b.pending, c.name as buddyname, c.job as buddyjob, c.level as buddylevel, b.groupname FROM %s as b, %s as c WHERE c.id = b.buddyid AND b.characterid = ?", model.TableBuddies, model.TableCharacters)

	result := make([]*model.BuddyWithCharacterInfo, 0)
	err := db.sqlDB.Select(&result, sqlQuery, characterID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
