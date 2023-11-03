package model

import "time"

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/31 23:13
  @describe :
*/

// -----------------------------------------------------------------------------------------
// ------------------------------------ Characters -----------------------------------------
// -----------------------------------------------------------------------------------------

const (
	// TableCharacters 表格:角色信息
	TableCharacters = "`characters`"
	// FieldsCharacters 所有字段:账户的所有角色
	FieldsCharacters = "`id`,`accountid`,`world`,`name`,`level`,`exp`,`str`,`dex`,`luk`,`int`,`hp`,`mp`,`maxhp`,`maxmp`,`meso`,`hpApUsed`,`job`,`skincolor`,`gender`,`fame`,`hair`,`face`,`ap`,`map`,`spawnpoint`,`gm`,`party`,`buddyCapacity`,`createdate`,`guildid`,`guildrank`,`allianceRank`,`monsterbookcover`,`dojo_pts`,`dojoRecord`,`pets`,`sp`,`subcategory`,`Jaguar`,`rank`,`rankMove`,`jobRank`,`jobRankMove`,`marriageId`,`familyid`,`seniorid`,`junior1`,`junior2`,`currentrep`,`totalrep`,`charmessage`,`expression`,`constellation`,`blood`,`month`,`day`,`beans`,`prefix`,`skillzq`,`grname`,`jzname`,`mrfbrw`,`mrsjrw`,`mrsgrw`,`mrsbossrw`,`hythd`,`mrsgrwa`,`mrfbrwa`,`mrsbossrwa`,`mrsgrws`,`mrsbossrws`,`mrfbrws`,`mrsgrwas`,`mrsbossrwas`,`mrfbrwas`,`ddj`,`vip`,`bosslog`,`djjl`,`qiandao`,`mountid`,`sg`"
)

// ICharacter 角色接口
type ICharacter interface {
	GetID() int64
	GetAccountID() int
	GetWorld() int8
	GetName() string
	GetLevel() int16
	GetExp() int
	GetStr() int
	GetDex() int
	GetLuk() int
	GetInt() int
	GetHp() int
	GetMp() int
	GetMaxHp() int
	GetMaxMp() int
	GetMeso() int
	GetHpApUsed() int
	GetJob() int
	GetSkinColor() int8
	GetGender() int8
	GetFame() int
	GetHair() int
	GetFace() int
	GetAp() int
	GetMapID() int
	GetSpawnPoint() int16
	GetGm() int8
	GetParty() int
	GetBuddyCapacity() int16
	GetCreateDate() time.Time
	GetGuildID() int
	GetGuildRank() int8
	GetAllianceRank() int8
	GetMonsterBookCover() int
	GetDojoPts() int
	GetDojoRecord() int8
	GetPets() string
	GetSp() string
	GetSubcategory() int
	GetJaguar() int
	GetRank() int
	GetRankMove() int
	GetJobRank() int
	GetJobRankMove() int
	GetMarriageID() int
	GetFamilyID() int
	GetSeniorID() int
	GetJunior1() int
	GetJunior2() int
	GetCurrentRep() int
	GetTotalRep() int
	GetCharMessage() string
	GetExpression() int
	GetConstellation() int
	GetBlood() int
	GetMonth() int
	GetDay() int
	GetBeans() int
	GetPrefix() int
	GetSkillZQ() int
	GetGrname() int
	GetJzname() int
	GetMrfbrw() int
	GetMrsjrw() int
	GetMrsgrw() int
	GetMrsbossrw() int
	GetHythd() int
	GetMrsgrwa() int
	GetMrfbrwa() int
	GetMrsbossrwa() int
	GetMrsgrws() int
	GetMrsbossrws() int
	GetMrfbrws() int
	GetMrsgrwas() int
	GetMrsbossrwas() int
	GetMrfbrwas() int
	GetDdj() int
	GetVip() int
	GetBossLog() int
	GetDjjl() int
	GetQiandao() int
	GetMountID() int
	GetSg() int
}

var _ ICharacter = new(Character)

// Character 账户角色信息
type Character struct {
	// ID 自增ID
	ID int64 `json:"id" db:"id"`

	// AccountID 账户ID
	AccountID int `json:"accountid" db:"accountid"`

	// World 世界？大区？
	World int8 `json:"world" db:"world"`

	// Name 角色名
	Name string `json:"name" db:"name"`

	// Level 等级
	Level int16 `json:"level" db:"level"`

	// Exp 经验值
	Exp int `json:"exp" db:"exp"`

	// Str 力量
	Str int `json:"str" db:"str"`

	// Dex 敏捷
	Dex int `json:"dex" db:"dex"`

	// Luk 运气
	Luk int `json:"luk" db:"luk"`

	// Int (intelligence) 智力
	Int int `json:"int" db:"int"`

	// Hp 生命值
	Hp int `json:"hp" db:"hp"`

	// Mp 魔法值
	Mp int `json:"mp" db:"mp"`

	// MaxHp 最大生命值
	MaxHp int `json:"maxhp" db:"maxhp"`

	// MaxMp 最大魔法值
	MaxMp    int `json:"maxmp" db:"maxmp"`
	Meso     int `json:"meso" db:"meso"`
	HpapUsed int `json:"hpApUsed" db:"hpApUsed"`
	Job      int `json:"job" db:"job"`

	// SkinColor 皮肤颜色
	SkinColor int8 `json:"skincolor" db:"skincolor"`

	// Gender 性别
	Gender int8 `json:"gender" db:"gender"`

	// Fame 人气
	Fame int `json:"fame" db:"fame"`

	// Hair 发型
	Hair int `json:"hair" db:"hair"`

	// Face 脸型
	Face int `json:"face" db:"face"`

	// Ap (Ability Points) 能力点. 用于提升玩家角色的主要属性，如力量（Strength）、敏捷（Dexterity）、智力（Intelligence）和幸运（Luck）
	Ap    int `json:"ap" db:"ap"`
	MapID int `json:"map" db:"map"`

	// SpawnPoint !!! 解释待定:角色、怪物或物品等在地图上出现的特定位置或点。SpawnPoint 用于确定角色或物体在游戏世界中的生成位置。
	SpawnPoint int16 `json:"spawnpoint" db:"spawnpoint"`

	// Gm 是否GM
	Gm               int8      `json:"gm" db:"gm"`
	Party            int       `json:"party" db:"party"`
	BuddyCapacity    int16     `json:"buddyCapacity" db:"buddyCapacity"`
	CreateDate       time.Time `json:"createdate" db:"createdate"`
	GuildID          int       `json:"guildid" db:"guildid"`
	GuildRank        int8      `json:"guildrank" db:"guildrank"`
	AllianceRank     int8      `json:"allianceRank" db:"allianceRank"`
	MonsterBookCover int       `json:"monsterbookcover" db:"monsterbookcover"`
	DojoPts          int       `json:"dojo_pts" db:"dojo_pts"`
	DojoRecord       int8      `json:"dojoRecord" db:"dojoRecord"`
	Pets             string    `json:"pets" db:"pets"`
	// Sp (Skill Point) 技能点. 用于表示玩家在游戏中获得的技能点数，这些技能点可以用于提升和学习各种技能，包括职业技能、特殊技能等。
	Sp string `json:"sp" db:"sp"`

	Subcategory int `json:"subcategory" db:"subcategory"`
	Jaguar      int `json:"Jaguar" db:"Jaguar"`
	Rank        int `json:"rank" db:"rank"`
	RankMove    int `json:"rankMove" db:"rankMove"`
	JobRank     int `json:"jobRank" db:"jobRank"`
	JobRankMove int `json:"jobRankMove" db:"jobRankMove"`

	// MarriageID 伴侣角色ID
	MarriageID int `json:"marriageId" db:"marriageId"`
	FamilyID   int `json:"familyid" db:"familyid"`
	SeniorID   int `json:"seniorid" db:"seniorid"`
	Junior1    int `json:"junior1" db:"junior1"`
	Junior2    int `json:"junior2" db:"junior2"`
	CurrentRep int `json:"currentrep" db:"currentrep"`
	TotalRep   int `json:"totalrep" db:"totalrep"`

	// CharMessage 角色消息?
	CharMessage   string `json:"charmessage" db:"charmessage"`
	Expression    int    `json:"expression" db:"expression"`
	Constellation int    `json:"constellation" db:"constellation"`
	Blood         int    `json:"blood" db:"blood"`
	Month         int    `json:"month" db:"month"`
	Day           int    `json:"day" db:"day"`
	Beans         int    `json:"beans" db:"beans"`
	Prefix        int    `json:"prefix" db:"prefix"`
	SkillZQ       int    `json:"skillzq" db:"skillzq"`
	Grname        int    `json:"grname" db:"grname"`
	Jzname        int    `json:"jzname" db:"jzname"`
	Mrfbrw        int    `json:"mrfbrw" db:"mrfbrw"`
	Mrsjrw        int    `json:"mrsjrw" db:"mrsjrw"`
	Mrsgrw        int    `json:"mrsgrw" db:"mrsgrw"`
	Mrsbossrw     int    `json:"mrsbossrw" db:"mrsbossrw"`
	Hythd         int    `json:"hythd" db:"hythd"`
	Mrsgrwa       int    `json:"mrsgrwa" db:"mrsgrwa"`
	Mrfbrwa       int    `json:"mrfbrwa" db:"mrfbrwa"`
	Mrsbossrwa    int    `json:"mrsbossrwa" db:"mrsbossrwa"`
	Mrsgrws       int    `json:"mrsgrws" db:"mrsgrws"`
	Mrsbossrws    int    `json:"mrsbossrws" db:"mrsbossrws"`
	Mrfbrws       int    `json:"mrfbrws" db:"mrfbrws"`
	Mrsgrwas      int    `json:"mrsgrwas" db:"mrsgrwas"`
	Mrsbossrwas   int    `json:"mrsbossrwas" db:"mrsbossrwas"`
	Mrfbrwas      int    `json:"mrfbrwas" db:"mrfbrwas"`
	Ddj           int    `json:"ddj" db:"ddj"`
	Vip           int    `json:"vip" db:"vip"`
	BossLog       int    `json:"bosslog" db:"bosslog"`
	Djjl          int    `json:"djjl" db:"djjl"`
	Qiandao       int    `json:"qiandao" db:"qiandao"`
	Mountid       int    `json:"mountid" db:"mountid"`
	Sg            int    `json:"sg" db:"sg"`
}

func (c *Character) GetID() int64 {
	return c.ID
}

func (c *Character) GetAccountID() int {
	return c.AccountID
}

func (c *Character) GetWorld() int8 {
	return c.World
}

func (c *Character) GetName() string {
	return c.Name
}

func (c *Character) GetLevel() int16 {
	return c.Level
}

func (c *Character) GetExp() int {
	return c.Exp
}

func (c *Character) GetStr() int {
	return c.Str
}

func (c *Character) GetDex() int {
	return c.Dex
}

func (c *Character) GetLuk() int {
	return c.Luk
}

func (c *Character) GetInt() int {
	return c.Int
}

func (c *Character) GetHp() int {
	return c.Hp
}

func (c *Character) GetMp() int {
	return c.Mp
}

func (c *Character) GetMaxHp() int {
	return c.MaxHp
}

func (c *Character) GetMaxMp() int {
	return c.MaxMp
}

func (c *Character) GetMeso() int {
	return c.Meso
}

func (c *Character) GetHpApUsed() int {
	return c.HpapUsed
}

func (c *Character) GetJob() int {
	return c.Job
}

func (c *Character) GetSkinColor() int8 {
	return c.SkinColor
}

func (c *Character) GetGender() int8 {
	return c.Gender
}

func (c *Character) GetFame() int {
	return c.Fame
}

func (c *Character) GetHair() int {
	return c.Hair
}

func (c *Character) GetFace() int {
	return c.Face
}

func (c *Character) GetAp() int {
	return c.Ap
}

func (c *Character) GetMapID() int {
	return c.MapID
}

func (c *Character) GetSpawnPoint() int16 {
	return c.SpawnPoint
}

func (c *Character) GetGm() int8 {
	return c.Gm
}

func (c *Character) GetParty() int {
	return c.Party
}

func (c *Character) GetBuddyCapacity() int16 {
	return c.BuddyCapacity
}

func (c *Character) GetCreateDate() time.Time {
	return c.CreateDate
}

func (c *Character) GetGuildID() int {
	return c.GuildID
}

func (c *Character) GetGuildRank() int8 {
	return c.GuildRank
}

func (c *Character) GetAllianceRank() int8 {
	return c.AllianceRank
}

func (c *Character) GetMonsterBookCover() int {
	return c.MonsterBookCover
}

func (c *Character) GetDojoPts() int {
	return c.DojoPts
}

func (c *Character) GetDojoRecord() int8 {
	return c.DojoRecord
}

func (c *Character) GetPets() string {
	return c.Pets
}

func (c *Character) GetSp() string {
	return c.Sp
}

func (c *Character) GetSubcategory() int {
	return c.Subcategory
}

func (c *Character) GetJaguar() int {
	return c.Jaguar
}

func (c *Character) GetRank() int {
	return c.Rank
}

func (c *Character) GetRankMove() int {
	return c.RankMove
}

func (c *Character) GetJobRank() int {
	return c.JobRank
}

func (c *Character) GetJobRankMove() int {
	return c.JobRankMove
}

func (c *Character) GetMarriageID() int {
	return c.MarriageID
}

func (c *Character) GetFamilyID() int {
	return c.FamilyID
}

func (c *Character) GetSeniorID() int {
	return c.SeniorID
}

func (c *Character) GetJunior1() int {
	return c.Junior1
}

func (c *Character) GetJunior2() int {
	return c.Junior2
}

func (c *Character) GetCurrentRep() int {
	return c.CurrentRep
}

func (c *Character) GetTotalRep() int {
	return c.TotalRep
}

func (c *Character) GetCharMessage() string {
	return c.CharMessage
}

func (c *Character) GetExpression() int {
	return c.Expression
}

func (c *Character) GetConstellation() int {
	return c.Constellation
}

func (c *Character) GetBlood() int {
	return c.Blood
}

func (c *Character) GetMonth() int {
	return c.Month
}

func (c *Character) GetDay() int {
	return c.Day
}

func (c *Character) GetBeans() int {
	return c.Beans
}

func (c *Character) GetPrefix() int {
	return c.Prefix
}

func (c *Character) GetSkillZQ() int {
	return c.SkillZQ
}

func (c *Character) GetGrname() int {
	return c.Grname
}

func (c *Character) GetJzname() int {
	return c.Jzname
}

func (c *Character) GetMrfbrw() int {
	return c.Mrfbrw
}

func (c *Character) GetMrsjrw() int {
	return c.Mrsjrw
}

func (c *Character) GetMrsgrw() int {
	return c.Mrsgrw
}

func (c *Character) GetMrsbossrw() int {
	return c.Mrsbossrw
}

func (c *Character) GetHythd() int {
	return c.Hythd
}

func (c *Character) GetMrsgrwa() int {
	return c.Mrsgrwa
}

func (c *Character) GetMrfbrwa() int {
	return c.Mrfbrwa
}

func (c *Character) GetMrsbossrwa() int {
	return c.Mrsbossrwa
}

func (c *Character) GetMrsgrws() int {
	return c.Mrsgrws
}

func (c *Character) GetMrsbossrws() int {
	return c.Mrsbossrws
}

func (c *Character) GetMrfbrws() int {
	return c.Mrfbrws
}

func (c *Character) GetMrsgrwas() int {
	return c.Mrsgrwas
}

func (c *Character) GetMrsbossrwas() int {
	return c.Mrsbossrwas
}

func (c *Character) GetMrfbrwas() int {
	return c.Mrfbrwas
}

func (c *Character) GetDdj() int {
	return c.Ddj
}

func (c *Character) GetVip() int {
	return c.Vip
}

func (c *Character) GetBossLog() int {
	return c.BossLog
}

func (c *Character) GetDjjl() int {
	return c.Djjl
}

func (c *Character) GetQiandao() int {
	return c.Qiandao
}

func (c *Character) GetMountID() int {
	return c.Mountid
}

func (c *Character) GetSg() int {
	return c.Sg
}

// -----------------------------------------------------------------------------------------
// --------------------------------- CharacterSlots ----------------------------------------
// -----------------------------------------------------------------------------------------

const (
	// TableCharacterSlots 表格:角色卡槽数据表
	TableCharacterSlots = "`character_slots`"

	// FieldsCharacterSlots 所有字段:账户的角色卡槽数量信息表
	FieldsCharacterSlots = "`id`,`accid`,`worldid`,`charslots`"
)

// CharacterSlots 账户的角色卡槽数据
type CharacterSlots struct {
	ID int64 `json:"id" db:"id"`
	// AccountID 账户ID
	AccountID int64 `json:"acc_id" db:"accid"`

	// WorldID 世界ID
	WorldID int `json:"world_id" db:"worldid"`

	// CharacterSlots 角色卡槽数量
	CharacterSlots int `json:"char_slots" db:"charslots"`
}

// -----------------------------------------------------------------------------------------
// -------------------------------------- Buddies ------------------------------------------
// -----------------------------------------------------------------------------------------

const (
	// TableBuddies 好友列表表格
	TableBuddies = "`buddies`"

	// FieldsBuddies 好友列表所有字段
	FieldsBuddies = "`id`,`characterid`,`buddyid`,`pending`,`groupname`"
)

// Buddy 好友信息
type Buddy struct {
	Id int64 `json:"id" db:"id"`

	CharacterId int64 `json:"characterid" db:"characterid"`

	BuddyId int64 `json:"buddyid" db:"buddyid"`

	Pending uint16 `json:"pending" db:"pending"`

	GroupName string `json:"groupname" db:"groupname"`
}

// BuddyWithCharacterInfo 包含名称等信息的好友信息
type BuddyWithCharacterInfo struct {
	Buddy

	BuddyName string `json:"buddyname" db:"buddyname"`

	BuddyJob int `json:"buddyjob" db:"buddyjob"`

	BuddyLevel int `json:"buddylevel" db:"buddylevel"`
}
