package client

import "github.com/jerbe/goms/database"

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/11/2 19:54
  @describe :
*/

type Buddy struct {
	// id 角色ID
	id int

	// name 好友名称
	name string

	// group 组名
	group string

	// level 角色等级
	level int

	// job 角色职业
	job int

	// visible 是否可见. 隐身/离线为不可见
	visible bool

	// channel 所在的频道
	channel int
}

type buddyOperate uint8

const (
	BuddyOperateAdded buddyOperate = iota
	BuddyOperateDeleted
)

type BuddyAddResult uint8

const (
	BuddyAddNil BuddyAddResult = iota
	BuddyAddOK
	BuddyAddExists
	BuddyAddFull
)

func NewBuddyList(capacity uint8) *BuddyList {
	return &BuddyList{
		capacity:     capacity,
		list:         make([]*Buddy, 0, capacity),
		idBuddyMap:   make(map[int]*Buddy),
		nameBuddyMap: make(map[string]*Buddy),
	}
}

type BuddyList struct {
	capacity     uint8
	list         []*Buddy
	idBuddyMap   map[int]*Buddy
	nameBuddyMap map[string]*Buddy
}

// GetBuddyById 根据角色ID获取好友
func (bl *BuddyList) GetBuddyById(characterId int) (*Buddy, bool) {
	buddy, ok := bl.idBuddyMap[characterId]
	return buddy, ok
}

// GetBuddyByName 根据角色名字获取好友
func (bl *BuddyList) GetBuddyByName(characterName string) (*Buddy, bool) {
	buddy, ok := bl.nameBuddyMap[characterName]
	return buddy, ok
}

// AddBuddy 添加好友
func (bl *BuddyList) AddBuddy(buddy *Buddy) (int, BuddyAddResult) {
	return bl.addBuddy(buddy)
}

// addBuddy 添加好友
func (bl *BuddyList) addBuddy(buddy *Buddy) (int, BuddyAddResult) {
	if buddy == nil {
		return len(bl.list), BuddyAddNil
	}

	if _, ok := bl.idBuddyMap[buddy.id]; ok {
		return len(bl.list), BuddyAddExists
	}

	if uint8(len(bl.list)) == bl.capacity {
		return len(bl.list), BuddyAddFull
	}

	bl.list = append(bl.list, buddy)
	bl.idBuddyMap[buddy.id] = buddy
	bl.nameBuddyMap[buddy.name] = buddy
	return len(bl.list), BuddyAddOK
}

// RemoveBuddyById 根据角色ID将好友从队列中删除
func (bl *BuddyList) RemoveBuddyById(id int) int {
	if _, ok := bl.idBuddyMap[id]; !ok {
		return len(bl.list)
	}

	for i := 0; i < len(bl.list); i++ {
		if id == bl.list[i].id {
			buddy := bl.list[i]
			bl.list = append(bl.list[0:i+1], bl.list[i:]...)
			delete(bl.idBuddyMap, buddy.id)
			delete(bl.nameBuddyMap, buddy.name)
			return len(bl.list)
		}
	}
	return len(bl.list)
}

// SetCapacity 设置容量
func (bl *BuddyList) SetCapacity(capacity uint8) {
	bl.capacity = capacity
}

// GetCapacity 获取容量
func (bl *BuddyList) GetCapacity() uint8 {
	return bl.capacity
}

// BuddiesList 获取伙伴列表
func (bl *BuddyList) BuddiesList() []*Buddy {
	n := make([]*Buddy, len(bl.list))
	copy(n, bl.list)
	return n
}

// BuddiesIds 获取伙伴id列表
func (bl *BuddyList) BuddiesIds() []int {
	n := make([]int, 0, len(bl.list))
	for i := 0; i < len(bl.list); i++ {
		n = append(n, bl.list[i].id)
	}
	return n
}

// IsFull 是否已经满了
func (bl *BuddyList) IsFull() bool {
	return len(bl.list) >= int(bl.capacity)
}

// LoadFromDB 从数据库中加载
func (bl *BuddyList) LoadFromDB(characterID int64) {
	infos, err := database.BuddiesWithCharacterInfo(characterID)
	if err != nil {
		return
	}

	bl.list = make([]*Buddy, 0, bl.capacity)
	bl.idBuddyMap = make(map[int]*Buddy)
	bl.nameBuddyMap = make(map[string]*Buddy)

	for i := 0; i < len(infos); i++ {
		info := infos[i]
		buddy := &Buddy{
			id:      int(info.BuddyId),
			name:    info.BuddyName,
			group:   info.GroupName,
			level:   info.BuddyLevel,
			job:     info.BuddyJob,
			visible: false,
			channel: 0,
		}
		bl.addBuddy(buddy)
	}
}
