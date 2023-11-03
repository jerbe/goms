package inventory

import "github.com/jerbe/goms/database/model"

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/31 09:36
  @describe :
*/

// NewPetFromModel 从一个model生成一个宠物
func NewPetFromModel(base *model.Item) *Pet {
	item := NewItem(base.Base())
	pet := &Pet{Item: *item}

	return pet
}

// Pet 宠物
type Pet struct {
	Item

	// summoned 是否召唤
	summoned bool
}

// SetSummoned 设置是否召唤
func (p *Pet) SetSummoned(val bool) {
	p.summoned = val
}

// GetSummoned 获取是否被召唤
func (p *Pet) GetSummoned() bool {
	return p.summoned
}
