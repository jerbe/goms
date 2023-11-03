package wz

import (
	"flag"
	"log"
	"os"

	"github.com/jerbe/goms/utils"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/23 10:30
  @describe :
*/

var (
	WzManage *wzManage
	wzPath   = flag.String(WZKey, DefaultWZPath, "wz root path")
)

func init() {
	log.Println("[wz] 开始初始")
	utils.InitFuncQueue.Push(initWzManage)
}

// initWzManage 初始化管理器
func initWzManage() {
	log.Println("[wz] 初始中...")
	var err error
	WzDir, err = os.Open(*wzPath)
	utils.PanicError(err)

	WzBaseDir, err = utils.ChildFile(WzDir, "Base.wz")
	WzCharacterDir, err = utils.ChildFile(WzDir, "Character.wz")
	WzEffectDir, err = utils.ChildFile(WzDir, "Effect.wz")
	WzEtcDir, err = utils.ChildFile(WzDir, "Etc.wz")
	WzItemDir, err = utils.ChildFile(WzDir, "Item.wz")
	WzMapDir, err = utils.ChildFile(WzDir, "Map.wz")
	WzMobDir, err = utils.ChildFile(WzDir, "Mob.wz")
	WzMorphDir, err = utils.ChildFile(WzDir, "Morph.wz")
	WzNpcDir, err = utils.ChildFile(WzDir, "Npc.wz")
	WzQuestDir, err = utils.ChildFile(WzDir, "Quest.wz")
	WzReactorDir, err = utils.ChildFile(WzDir, "Reactor.wz")
	WzSkillDir, err = utils.ChildFile(WzDir, "Skill.wz")
	WzSoundDir, err = utils.ChildFile(WzDir, "Sound.wz")
	WzStringDir, err = utils.ChildFile(WzDir, "String.wz")
	WzTamingMobDir, err = utils.ChildFile(WzDir, "TamingMob.wz")
	WzUiDir, err = utils.ChildFile(WzDir, "UI.wz")

}

const (
	// WZKey 通过环境变量获取到WZ根目录的键
	WZKey = "wz.path"

	// DefaultWZPath 默认WZ的目录
	DefaultWZPath = "wz"
)

var (
	// WzDir 根目录
	WzDir *os.File

	// WzBaseDir 基础文件夹
	WzBaseDir *os.File

	// WzCharacterDir 角色文件夹
	WzCharacterDir *os.File

	// WzEffectDir 特效文件夹
	WzEffectDir *os.File

	// WzEtcDir 装饰文件夹
	WzEtcDir *os.File

	// WzItemDir 物品文件夹
	WzItemDir *os.File

	// WzMapDir 地图文件夹
	WzMapDir *os.File

	// WzMobDir Mod文件夹(模型?)
	WzMobDir *os.File

	// WzMorphDir 变形文件夹(什么鬼?)
	WzMorphDir *os.File

	// WzNpcDir NPC文件夹
	WzNpcDir *os.File

	// WzQuestDir 任务文件夹?
	WzQuestDir *os.File

	// WzReactorDir 反应堆文件夹(什么东西?)
	WzReactorDir *os.File

	// WzSkillDir 技能文件夹
	WzSkillDir *os.File

	// WzSoundDir 声音文件夹
	WzSoundDir *os.File

	// WzStringDir 文本文件夹
	WzStringDir *os.File

	// WzTamingMobDir 驯服模型文件夹(宠物?)
	WzTamingMobDir *os.File

	// WzUiDir UI画面文件夹
	WzUiDir *os.File
)

type wzManage struct {
}
