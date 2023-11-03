package database

import (
	"errors"
	"fmt"
	"github.com/jerbe/goms/database/model"
)

/*
*

	@author : Jerbe - The porter from Earth
	@time : 2023/10/27 15:46
	@describe :
*/

// FindAccount 根据用户ID获取一个账户信息
func (db *database) FindAccount(id int64) (*model.Account, error) {
	sqlQuery := fmt.Sprintf("SELECT %s FROM `%s` WHERE `id` = ? ", model.FieldsAccounts, model.TableAccounts)
	result := new(model.Account)

	err := db.sqlDB.Get(result, sqlQuery, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindAccountByUsername 根据用户名获取一个账户信息
func (db *database) FindAccountByUsername(username string) (*model.Account, error) {
	sqlQuery := fmt.Sprintf("SELECT %s FROM `%s` WHERE `name` = ? ", model.FieldsAccounts, model.TableAccounts)
	result := new(model.Account)

	err := db.sqlDB.Get(result, sqlQuery, username)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateAccountFilter 需要更新账户的过滤器
type UpdateAccountFilter struct {
}

// UpdateAccountData 需要更新账户信息的数据
type UpdateAccountData struct {
}

// UpdateAccount 更新账户信息
func UpdateAccount(filter *UpdateAccountData, data *UpdateAccountData) (int64, error) {
	return Default.UpdateAccount(filter, *data)
}

// UpdateAccount 更新账户信息
func (db *database) UpdateAccount(filter *UpdateAccountData, data UpdateAccountData) (int64, error) {
	return 0, nil
}

var resetAllAccountStatus = false
var resetAllAccountLoginStatus = false
var resetLastGainHiredMerchant = false

// ResetAllAccount 服务启动时重置所有人的账户状态
func ResetAllAccount() (int64, error) {
	return Default.ResetAllAccount()
}

// ResetAllAccount 服务启动时重置所有人的账户状态
func (db *database) ResetAllAccount() (int64, error) {
	if resetAllAccountStatus {
		return 0, errors.New("already reset")
	}

	result, err := db.sqlDB.Exec("UPDATE `accounts` SET `loggedin` = 0, `lastGainHM` = 0")
	if err != nil {
		return 0, err
	}
	resetAllAccountStatus = true
	return result.RowsAffected()
}

// ResetAllLoginStatus 服务启动时重置所有人登录状态
func ResetAllLoginStatus() (int64, error) {
	return Default.ResetAllLoginStatus()
}

// ResetAllLoginStatus 服务启动时重置所有人登录状态
func (db *database) ResetAllLoginStatus() (int64, error) {
	if resetAllAccountLoginStatus {
		return 0, errors.New("already reset")
	}
	result, err := db.sqlDB.Exec(fmt.Sprintf("UPDATE %s SET `loggedin` = 0", model.TableAccounts))
	if err != nil {
		return 0, err
	}
	resetAllAccountLoginStatus = true
	return result.RowsAffected()
}

// ResetAllLastGainHiredMerchant 服务启动时重置最后获得雇佣商人的时间
func ResetAllLastGainHiredMerchant() (int64, error) {
	return Default.ResetAllLastGainHiredMerchant()
}

// ResetAllLastGainHiredMerchant 服务启动时重置最后获得雇佣商人的时间
func (db *database) ResetAllLastGainHiredMerchant() (int64, error) {
	if resetLastGainHiredMerchant {
		return 0, errors.New("already reset")
	}
	result, err := db.sqlDB.Exec(fmt.Sprintf("UPDATE %s SET `lastGainHM` = 0", model.TableAccounts))
	if err != nil {
		return 0, err
	}
	resetLastGainHiredMerchant = true
	return result.RowsAffected()
}

// CheckAccountExist 检测账户是否存在
func CheckAccountExist(username string) (bool, error) {
	return Default.CheckAccountExist(username)
}

// CheckAccountExist 检测账户是否存在
func (db *database) CheckAccountExist(username string) (bool, error) {
	count := int64(0)
	err := db.sqlDB.Get(&count, fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE `name` = ?", model.TableAccounts), username)
	if err != nil {
		return false, err
	}
	return count > 1, nil
}
