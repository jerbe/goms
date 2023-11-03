package code

import (
	"embed"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/jerbe/goms/utils"
	"reflect"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/26 10:27
  @describe :
*/

type OpcodeType uint8

const (
	OpcodeTypeUnknown OpcodeType = iota
	OpcodeTypeReceive
	OpcodeTypeSend
)

type Opcoder interface {
	Type() OpcodeType
	Int16() int16
	String() string
}

// -------- Opcode

var opcodeStringMap = map[Opcode]string{
	Opcode_HELLO: "HELLO",
}

type Opcode int16

func (Opcode) Type() OpcodeType {
	return OpcodeTypeUnknown
}

func (c Opcode) Int16() int16 {
	return int16(c)
}

func (c Opcode) String() string {
	name, ok := opcodeStringMap[c]
	if !ok {
		name = "unknown"
	}
	return fmt.Sprintf("Opcode:%d => %s", c, name)
}

// Opcode_HELLO 连接上服务端就需要给客户端发送hello
const Opcode_HELLO = Opcode(13)

// ----------- ReceiveOpcode

func NewReceiveOpcode(v int16) ReceiveOpcode {
	return ReceiveOpcode(v)
}

var Receive receive

type receive struct {
	PONG ReceiveOpcode `opcode:"PONG"`
	//LOGIN ReceiveOpcode `opcode:"LOGIN"`
	LOGIN_PASSWORD       ReceiveOpcode `opcode:"LOGIN_PASSWORD"`
	HELLO_LOGIN          ReceiveOpcode `opcode:"HELLO_LOGIN"`
	HELLO_CHANNEL        ReceiveOpcode `opcode:"HELLO_CHANNEL"`
	LICENSE_REQUEST      ReceiveOpcode `opcode:"LICENSE_REQUEST"`
	SERVERLIST_REQUEST   ReceiveOpcode `opcode:"SERVERLIST_REQUEST"`
	CHARLIST_REQUEST     ReceiveOpcode `opcode:"CHARLIST_REQUEST"`
	SERVERSTATUS_REQUEST ReceiveOpcode `opcode:"SERVERSTATUS_REQUEST"`
	CHECK_CHAR_NAME      ReceiveOpcode `opcode:"CHECK_CHAR_NAME"`
	CREATE_CHAR          ReceiveOpcode `opcode:"CREATE_CHAR"`
	DELETE_CHAR          ReceiveOpcode `opcode:"DELETE_CHAR"`
	STRANGE_DATA         ReceiveOpcode `opcode:"STRANGE_DATA"`
	CHAR_SELECT          ReceiveOpcode `opcode:"CHAR_SELECT"`
	AUTH_SECOND_PASSWORD ReceiveOpcode `opcode:"AUTH_SECOND_PASSWORD"`
	SET_GENDER           ReceiveOpcode `opcode:"SET_GENDER"`
	RSA_KEY              ReceiveOpcode `opcode:"RSA_KEY"`
	//CHANNEL ReceiveOpcode `opcode:"CHANNEL"`
	PLAYER_LOGGEDIN       ReceiveOpcode `opcode:"PLAYER_LOGGEDIN"`
	CHANGE_MAP            ReceiveOpcode `opcode:"CHANGE_MAP"`
	CHANGE_CHANNEL        ReceiveOpcode `opcode:"CHANGE_CHANNEL"`
	ENTER_CASH_SHOP       ReceiveOpcode `opcode:"ENTER_CASH_SHOP"`
	MOVE_PLAYER           ReceiveOpcode `opcode:"MOVE_PLAYER"`
	CANCEL_CHAIR          ReceiveOpcode `opcode:"CANCEL_CHAIR"`
	USE_CHAIR             ReceiveOpcode `opcode:"USE_CHAIR"`
	CLOSE_RANGE_ATTACK    ReceiveOpcode `opcode:"CLOSE_RANGE_ATTACK"`
	RANGED_ATTACK         ReceiveOpcode `opcode:"RANGED_ATTACK"`
	MAGIC_ATTACK          ReceiveOpcode `opcode:"MAGIC_ATTACK"`
	PASSIVE_ENERGY        ReceiveOpcode `opcode:"PASSIVE_ENERGY"`
	TAKE_DAMAGE           ReceiveOpcode `opcode:"TAKE_DAMAGE"`
	GENERAL_CHAT          ReceiveOpcode `opcode:"GENERAL_CHAT"`
	CLOSE_CHALKBOARD      ReceiveOpcode `opcode:"CLOSE_CHALKBOARD"`
	FACE_EXPRESSION       ReceiveOpcode `opcode:"FACE_EXPRESSION"`
	USE_ITEMEFFECT        ReceiveOpcode `opcode:"USE_ITEMEFFECT"`
	WHEEL_OF_FORTUNE      ReceiveOpcode `opcode:"WHEEL_OF_FORTUNE"`
	MONSTER_BOOK_COVER    ReceiveOpcode `opcode:"MONSTER_BOOK_COVER"`
	NPC_TALK              ReceiveOpcode `opcode:"NPC_TALK"`
	NPC_TALK_MORE         ReceiveOpcode `opcode:"NPC_TALK_MORE"`
	NPC_SHOP              ReceiveOpcode `opcode:"NPC_SHOP"`
	STORAGE               ReceiveOpcode `opcode:"STORAGE"`
	USE_HIRED_MERCHANT    ReceiveOpcode `opcode:"USE_HIRED_MERCHANT"`
	MERCH_ITEM_STORE      ReceiveOpcode `opcode:"MERCH_ITEM_STORE"`
	DUEY_ACTION           ReceiveOpcode `opcode:"DUEY_ACTION"`
	ITEM_SORT             ReceiveOpcode `opcode:"ITEM_SORT"`
	ITEM_GATHER           ReceiveOpcode `opcode:"ITEM_GATHER"`
	ITEM_MOVE             ReceiveOpcode `opcode:"ITEM_MOVE"`
	USE_ITEM              ReceiveOpcode `opcode:"USE_ITEM"`
	CANCEL_ITEM_EFFECT    ReceiveOpcode `opcode:"CANCEL_ITEM_EFFECT"`
	USE_SUMMON_BAG        ReceiveOpcode `opcode:"USE_SUMMON_BAG"`
	PET_EXCEPTIONLIST     ReceiveOpcode `opcode:"PET_EXCEPTIONLIST"`
	PET_FOOD              ReceiveOpcode `opcode:"PET_FOOD"`
	USE_MOUNT_FOOD        ReceiveOpcode `opcode:"USE_MOUNT_FOOD"`
	USE_SCRIPTED_NPC_ITEM ReceiveOpcode `opcode:"USE_SCRIPTED_NPC_ITEM"`
	USE_CASH_ITEM         ReceiveOpcode `opcode:"USE_CASH_ITEM"`
	USE_CATCH_ITEM        ReceiveOpcode `opcode:"USE_CATCH_ITEM"`
	USE_SKILL_BOOK        ReceiveOpcode `opcode:"USE_SKILL_BOOK"`
	USE_RETURN_SCROLL     ReceiveOpcode `opcode:"USE_RETURN_SCROLL"`
	USE_UPGRADE_SCROLL    ReceiveOpcode `opcode:"USE_UPGRADE_SCROLL"`
	DISTRIBUTE_AP         ReceiveOpcode `opcode:"DISTRIBUTE_AP"`
	AUTO_ASSIGN_AP        ReceiveOpcode `opcode:"AUTO_ASSIGN_AP"`
	HEAL_OVER_TIME        ReceiveOpcode `opcode:"HEAL_OVER_TIME"`
	DISTRIBUTE_SP         ReceiveOpcode `opcode:"DISTRIBUTE_SP"`
	SPECIAL_MOVE          ReceiveOpcode `opcode:"SPECIAL_MOVE"`
	CANCEL_BUFF           ReceiveOpcode `opcode:"CANCEL_BUFF"`
	SKILL_EFFECT          ReceiveOpcode `opcode:"SKILL_EFFECT"`
	MESO_DROP             ReceiveOpcode `opcode:"MESO_DROP"`
	GIVE_FAME             ReceiveOpcode `opcode:"GIVE_FAME"`
	CHAR_INFO_REQUEST     ReceiveOpcode `opcode:"CHAR_INFO_REQUEST"`
	SPAWN_PET             ReceiveOpcode `opcode:"SPAWN_PET"`
	CANCEL_DEBUFF         ReceiveOpcode `opcode:"CANCEL_DEBUFF"`
	CHANGE_MAP_SPECIAL    ReceiveOpcode `opcode:"CHANGE_MAP_SPECIAL"`
	USE_INNER_PORTAL      ReceiveOpcode `opcode:"USE_INNER_PORTAL"`
	TROCK_ADD_MAP         ReceiveOpcode `opcode:"TROCK_ADD_MAP"`
	QUEST_ACTION          ReceiveOpcode `opcode:"QUEST_ACTION"`
	EFFECT_ON_OFF         ReceiveOpcode `opcode:"EFFECT_ON_OFF"`
	SKILL_MACRO           ReceiveOpcode `opcode:"SKILL_MACRO"`
	ITEM_BAOWU            ReceiveOpcode `opcode:"ITEM_BAOWU"`
	ITEM_SUNZI            ReceiveOpcode `opcode:"ITEM_SUNZI"`
	ITEM_MAKER            ReceiveOpcode `opcode:"ITEM_MAKER"`
	USE_TREASUER_CHEST    ReceiveOpcode `opcode:"USE_TREASUER_CHEST"`
	PARTYCHAT             ReceiveOpcode `opcode:"PARTYCHAT"`
	PARTY_SS              ReceiveOpcode `opcode:"PARTY_SS"`
	WHISPER               ReceiveOpcode `opcode:"WHISPER"`
	MESSENGER             ReceiveOpcode `opcode:"MESSENGER"`
	PLAYER_INTERACTION    ReceiveOpcode `opcode:"PLAYER_INTERACTION"`
	PARTY_OPERATION       ReceiveOpcode `opcode:"PARTY_OPERATION"`
	DENY_PARTY_REQUEST    ReceiveOpcode `opcode:"DENY_PARTY_REQUEST"`
	GUILD_OPERATION       ReceiveOpcode `opcode:"GUILD_OPERATION"`
	DENY_GUILD_REQUEST    ReceiveOpcode `opcode:"DENY_GUILD_REQUEST"`
	BUDDYLIST_MODIFY      ReceiveOpcode `opcode:"BUDDYLIST_MODIFY"`
	NOTE_ACTION           ReceiveOpcode `opcode:"NOTE_ACTION"`
	USE_DOOR              ReceiveOpcode `opcode:"USE_DOOR"`
	CHANGE_KEYMAP         ReceiveOpcode `opcode:"CHANGE_KEYMAP"`
	UPDATE_CHAR_INFO      ReceiveOpcode `opcode:"UPDATE_CHAR_INFO"`
	ENTER_MTS             ReceiveOpcode `opcode:"ENTER_MTS"`
	ALLIANCE_OPERATION    ReceiveOpcode `opcode:"ALLIANCE_OPERATION"`
	DENY_ALLIANCE_REQUEST ReceiveOpcode `opcode:"DENY_ALLIANCE_REQUEST"`
	REQUEST_FAMILY        ReceiveOpcode `opcode:"REQUEST_FAMILY"`
	OPEN_FAMILY           ReceiveOpcode `opcode:"OPEN_FAMILY"`
	FAMILY_OPERATION      ReceiveOpcode `opcode:"FAMILY_OPERATION"`
	DELETE_JUNIOR         ReceiveOpcode `opcode:"DELETE_JUNIOR"`
	DELETE_SENIOR         ReceiveOpcode `opcode:"DELETE_SENIOR"`
	ACCEPT_FAMILY         ReceiveOpcode `opcode:"ACCEPT_FAMILY"`
	USE_FAMILY            ReceiveOpcode `opcode:"USE_FAMILY"`
	FAMILY_PRECEPT        ReceiveOpcode `opcode:"FAMILY_PRECEPT"`
	FAMILY_SUMMON         ReceiveOpcode `opcode:"FAMILY_SUMMON"`
	CYGNUS_SUMMON         ReceiveOpcode `opcode:"CYGNUS_SUMMON"`
	ARAN_COMBO            ReceiveOpcode `opcode:"ARAN_COMBO"`
	BBS_OPERATION         ReceiveOpcode `opcode:"BBS_OPERATION"`
	TRANSFORM_PLAYER      ReceiveOpcode `opcode:"TRANSFORM_PLAYER"`
	MOVE_PET              ReceiveOpcode `opcode:"MOVE_PET"`
	PET_CHAT              ReceiveOpcode `opcode:"PET_CHAT"`
	PET_COMMAND           ReceiveOpcode `opcode:"PET_COMMAND"`
	PET_LOOT              ReceiveOpcode `opcode:"PET_LOOT"`
	PET_AUTO_POT          ReceiveOpcode `opcode:"PET_AUTO_POT"`
	MOVE_SUMMON           ReceiveOpcode `opcode:"MOVE_SUMMON"`
	SUMMON_ATTACK         ReceiveOpcode `opcode:"SUMMON_ATTACK"`
	DAMAGE_SUMMON         ReceiveOpcode `opcode:"DAMAGE_SUMMON"`
	MOVE_LIFE             ReceiveOpcode `opcode:"MOVE_LIFE"`
	AUTO_AGGRO            ReceiveOpcode `opcode:"AUTO_AGGRO"`
	FRIENDLY_DAMAGE       ReceiveOpcode `opcode:"FRIENDLY_DAMAGE"`
	MONSTER_BOMB          ReceiveOpcode `opcode:"MONSTER_BOMB"`
	HYPNOTIZE_DMG         ReceiveOpcode `opcode:"HYPNOTIZE_DMG"`
	NPC_ACTION            ReceiveOpcode `opcode:"NPC_ACTION"`
	ITEM_PICKUP           ReceiveOpcode `opcode:"ITEM_PICKUP"`
	DAMAGE_REACTOR        ReceiveOpcode `opcode:"DAMAGE_REACTOR"`
	SNOWBALL              ReceiveOpcode `opcode:"SNOWBALL"`
	LEFT_KNOCK_BACK       ReceiveOpcode `opcode:"LEFT_KNOCK_BACK"`
	COCONUT               ReceiveOpcode `opcode:"COCONUT"`
	MONSTER_CARNIVAL      ReceiveOpcode `opcode:"MONSTER_CARNIVAL"`
	SHIP_OBJECT           ReceiveOpcode `opcode:"SHIP_OBJECT"`
	CS_UPDATE             ReceiveOpcode `opcode:"CS_UPDATE"`
	BUY_CS_ITEM           ReceiveOpcode `opcode:"BUY_CS_ITEM"`
	TOUCHING_CS           ReceiveOpcode `opcode:"TOUCHING_CS"`
	COUPON_CODE           ReceiveOpcode `opcode:"COUPON_CODE"`
	MAPLETV               ReceiveOpcode `opcode:"MAPLETV"`
	MOVE_DRAGON           ReceiveOpcode `opcode:"MOVE_DRAGON"`
	REPAIR                ReceiveOpcode `opcode:"REPAIR"`
	REPAIR_ALL            ReceiveOpcode `opcode:"REPAIR_ALL"`
	TOUCHING_MTS          ReceiveOpcode `opcode:"TOUCHING_MTS"`
	USE_MAGNIFY_GLASS     ReceiveOpcode `opcode:"USE_MAGNIFY_GLASS"`
	USE_POTENTIAL_SCROLL  ReceiveOpcode `opcode:"USE_POTENTIAL_SCROLL"`
	USE_EQUIP_SCROLL      ReceiveOpcode `opcode:"USE_EQUIP_SCROLL"`
	GAME_POLL             ReceiveOpcode `opcode:"GAME_POLL"`
	OWL                   ReceiveOpcode `opcode:"OWL"`
	OWL_WARP              ReceiveOpcode `opcode:"OWL_WARP"`
	//XMAS_SURPRISE//header->uniqueid(long)isentirestructure
	USE_OWL_MINERVA ReceiveOpcode `opcode:"USE_OWL_MINERVA"`
	RPS_GAME        ReceiveOpcode `opcode:"RPS_GAME"`
	UPDATE_QUEST    ReceiveOpcode `opcode:"UPDATE_QUEST"`
	PLAYER_UPDATE   ReceiveOpcode `opcode:"PLAYER_UPDATE"`
	//QUEST_ITEM//header->questid(int)->1/0(byteopenorclose)
	USE_ITEM_QUEST  ReceiveOpcode `opcode:"USE_ITEM_QUEST"`
	FOLLOW_REQUEST  ReceiveOpcode `opcode:"FOLLOW_REQUEST"`
	FOLLOW_REPLY    ReceiveOpcode `opcode:"FOLLOW_REPLY"`
	MOB_NODE        ReceiveOpcode `opcode:"MOB_NODE"`
	DISPLAY_NODE    ReceiveOpcode `opcode:"DISPLAY_NODE"`
	TOUCH_REACTOR   ReceiveOpcode `opcode:"TOUCH_REACTOR"`
	RING_ACTION     ReceiveOpcode `opcode:"RING_ACTION"`
	MTS_TAB         ReceiveOpcode `opcode:"MTS_TAB"`
	ChatRoom_SYSTEM ReceiveOpcode `opcode:"ChatRoom_SYSTEM"`
	Quest_KJ        ReceiveOpcode `opcode:"quest_KJ"`
	NEW_SX          ReceiveOpcode `opcode:"NEW_SX"`
	BOATS           ReceiveOpcode `opcode:"BOATS"`
	BEANS_GAME1     ReceiveOpcode `opcode:"BEANS_GAME1"`
	BEANS_GAME2     ReceiveOpcode `opcode:"BEANS_GAME2"`
	MOONRABBIT_HP   ReceiveOpcode `opcode:"MOONRABBIT_HP"`
	MARRAGE_RECV    ReceiveOpcode `opcode:"MARRAGE_RECV"`
}

func initReceive() {
	Receive = receive{
		PONG: 0,
		//LOGIN
		LOGIN_PASSWORD:       1,
		HELLO_LOGIN:          2,
		HELLO_CHANNEL:        3,
		LICENSE_REQUEST:      4,
		SERVERLIST_REQUEST:   5,
		CHARLIST_REQUEST:     6,
		SERVERSTATUS_REQUEST: 7,
		CHECK_CHAR_NAME:      8,
		CREATE_CHAR:          9,
		DELETE_CHAR:          10,
		STRANGE_DATA:         11,
		CHAR_SELECT:          12,
		AUTH_SECOND_PASSWORD: 13,
		SET_GENDER:           14,
		RSA_KEY:              15,
		//CHANNEL
		PLAYER_LOGGEDIN:       16,
		CHANGE_MAP:            17,
		CHANGE_CHANNEL:        18,
		ENTER_CASH_SHOP:       19,
		MOVE_PLAYER:           20,
		CANCEL_CHAIR:          21,
		USE_CHAIR:             22,
		CLOSE_RANGE_ATTACK:    23,
		RANGED_ATTACK:         24,
		MAGIC_ATTACK:          25,
		PASSIVE_ENERGY:        26,
		TAKE_DAMAGE:           27,
		GENERAL_CHAT:          28,
		CLOSE_CHALKBOARD:      29,
		FACE_EXPRESSION:       30,
		USE_ITEMEFFECT:        31,
		WHEEL_OF_FORTUNE:      32,
		MONSTER_BOOK_COVER:    33,
		NPC_TALK:              34,
		NPC_TALK_MORE:         35,
		NPC_SHOP:              36,
		STORAGE:               37,
		USE_HIRED_MERCHANT:    38,
		MERCH_ITEM_STORE:      39,
		DUEY_ACTION:           40,
		ITEM_SORT:             41,
		ITEM_GATHER:           42,
		ITEM_MOVE:             43,
		USE_ITEM:              44,
		CANCEL_ITEM_EFFECT:    45,
		USE_SUMMON_BAG:        46,
		PET_EXCEPTIONLIST:     47,
		PET_FOOD:              48,
		USE_MOUNT_FOOD:        49,
		USE_SCRIPTED_NPC_ITEM: 50,
		USE_CASH_ITEM:         51,
		USE_CATCH_ITEM:        52,
		USE_SKILL_BOOK:        53,
		USE_RETURN_SCROLL:     54,
		USE_UPGRADE_SCROLL:    55,
		DISTRIBUTE_AP:         56,
		AUTO_ASSIGN_AP:        57,
		HEAL_OVER_TIME:        58,
		DISTRIBUTE_SP:         59,
		SPECIAL_MOVE:          60,
		CANCEL_BUFF:           61,
		SKILL_EFFECT:          62,
		MESO_DROP:             63,
		GIVE_FAME:             64,
		CHAR_INFO_REQUEST:     65,
		SPAWN_PET:             66,
		CANCEL_DEBUFF:         67,
		CHANGE_MAP_SPECIAL:    68,
		USE_INNER_PORTAL:      69,
		TROCK_ADD_MAP:         70,
		QUEST_ACTION:          71,
		EFFECT_ON_OFF:         72,
		SKILL_MACRO:           73,
		ITEM_BAOWU:            74,
		ITEM_SUNZI:            75,
		ITEM_MAKER:            76,
		USE_TREASUER_CHEST:    77,
		PARTYCHAT:             78,
		PARTY_SS:              79,
		WHISPER:               80,
		MESSENGER:             81,
		PLAYER_INTERACTION:    82,
		PARTY_OPERATION:       83,
		DENY_PARTY_REQUEST:    84,
		GUILD_OPERATION:       85,
		DENY_GUILD_REQUEST:    86,
		BUDDYLIST_MODIFY:      87,
		NOTE_ACTION:           88,
		USE_DOOR:              89,
		CHANGE_KEYMAP:         90,
		UPDATE_CHAR_INFO:      91,
		ENTER_MTS:             92,
		ALLIANCE_OPERATION:    93,
		DENY_ALLIANCE_REQUEST: 94,
		REQUEST_FAMILY:        95,
		OPEN_FAMILY:           96,
		FAMILY_OPERATION:      97,
		DELETE_JUNIOR:         98,
		DELETE_SENIOR:         99,
		ACCEPT_FAMILY:         100,
		USE_FAMILY:            101,
		FAMILY_PRECEPT:        102,
		FAMILY_SUMMON:         103,
		CYGNUS_SUMMON:         104,
		ARAN_COMBO:            105,
		BBS_OPERATION:         106,
		TRANSFORM_PLAYER:      107,
		MOVE_PET:              108,
		PET_CHAT:              109,
		PET_COMMAND:           110,
		PET_LOOT:              111,
		PET_AUTO_POT:          112,
		MOVE_SUMMON:           113,
		SUMMON_ATTACK:         114,
		DAMAGE_SUMMON:         115,
		MOVE_LIFE:             116,
		AUTO_AGGRO:            117,
		FRIENDLY_DAMAGE:       118,
		MONSTER_BOMB:          119,
		HYPNOTIZE_DMG:         120,
		NPC_ACTION:            121,
		ITEM_PICKUP:           122,
		DAMAGE_REACTOR:        123,
		SNOWBALL:              124,
		LEFT_KNOCK_BACK:       125,
		COCONUT:               126,
		MONSTER_CARNIVAL:      127,
		SHIP_OBJECT:           128,
		CS_UPDATE:             129,
		BUY_CS_ITEM:           130,
		TOUCHING_CS:           131,
		COUPON_CODE:           132,
		MAPLETV:               133,
		MOVE_DRAGON:           134,
		REPAIR:                135,
		REPAIR_ALL:            136,
		TOUCHING_MTS:          137,
		USE_MAGNIFY_GLASS:     138,
		USE_POTENTIAL_SCROLL:  139,
		USE_EQUIP_SCROLL:      140,
		GAME_POLL:             141,
		OWL:                   142,
		OWL_WARP:              143,
		//XMAS_SURPRISE//header->uniqueid(long)isentirestructure
		USE_OWL_MINERVA: 144,
		RPS_GAME:        145,
		UPDATE_QUEST:    146,
		PLAYER_UPDATE:   147,
		//QUEST_ITEM//header->questid(int)->1,/0,(byteopenorclose)
		USE_ITEM_QUEST:  148,
		FOLLOW_REQUEST:  149,
		FOLLOW_REPLY:    150,
		MOB_NODE:        151,
		DISPLAY_NODE:    152,
		TOUCH_REACTOR:   153,
		RING_ACTION:     154,
		MTS_TAB:         155,
		ChatRoom_SYSTEM: 156,
		Quest_KJ:        157,
		NEW_SX:          158,
		BOATS:           159,
		BEANS_GAME1:     160,
		BEANS_GAME2:     161,
		MOONRABBIT_HP:   162,
		MARRAGE_RECV:    163,
	}

	receiveTypeReflect := reflect.TypeOf(Receive)
	receiveValueReflect := reflect.ValueOf(Receive)
	mapValueReflect := reflect.ValueOf(receiveOpcodeStringMap)
	fieldNum := receiveTypeReflect.NumField()
	for i := 0; i < fieldNum; i++ {
		field := receiveTypeReflect.Field(i)
		value := receiveValueReflect.FieldByName(field.Name)
		tagName := field.Tag.Get("opcode")
		mapValueReflect.SetMapIndex(value, reflect.ValueOf(tagName))
	}
}

var receiveOpcodeStringMap = map[ReceiveOpcode]string{}

// ReceiveOpcode 接收到的数据包操作符
type ReceiveOpcode Opcode

func (code ReceiveOpcode) Type() OpcodeType {
	return OpcodeTypeReceive
}
func (code ReceiveOpcode) Int16() int16 {
	return int16(code)
}
func (code ReceiveOpcode) String() string {
	str := receiveOpcodeStringMap[code]
	return fmt.Sprintf("ReceiveOpcode:%d => %s", code, str)
}

// ----------- SendOpcode
var Send send

type send struct {
	PING SendOpcode `opcode:"PING"`
	//LOGIN
	LOGIN_STATUS         SendOpcode `opcode:"LOGIN_STATUS"`
	PIN_OPERATION        SendOpcode `opcode:"PIN_OPERATION"`
	SECONDPW_ERROR       SendOpcode `opcode:"SECONDPW_ERROR"`
	SERVERLIST           SendOpcode `opcode:"SERVERLIST"`
	SERVERSTATUS         SendOpcode `opcode:"SERVERSTATUS"`
	SERVER_IP            SendOpcode `opcode:"SERVER_IP"`
	CHARLIST             SendOpcode `opcode:"CHARLIST"`
	CHAR_NAME_RESPONSE   SendOpcode `opcode:"CHAR_NAME_RESPONSE"`
	RELOG_RESPONSE       SendOpcode `opcode:"RELOG_RESPONSE"`
	ADD_NEW_CHAR_ENTRY   SendOpcode `opcode:"ADD_NEW_CHAR_ENTRY"`
	DELETE_CHAR_RESPONSE SendOpcode `opcode:"DELETE_CHAR_RESPONSE"`
	CHANNEL_SELECTED     SendOpcode `opcode:"CHANNEL_SELECTED"`
	ALL_CHARLIST         SendOpcode `opcode:"ALL_CHARLIST"`
	CHOOSE_GENDER        SendOpcode `opcode:"CHOOSE_GENDER"`
	GENDER_SET           SendOpcode `opcode:"GENDER_SET"`
	CHAR_CASH            SendOpcode `opcode:"CHAR_CASH"`
	//CHANNEL
	CHANGE_CHANNEL               SendOpcode `opcode:"CHANGE_CHANNEL"`
	UPDATE_STATS                 SendOpcode `opcode:"UPDATE_STATS"`
	FAME_RESPONSE                SendOpcode `opcode:"FAME_RESPONSE"`
	UPDATE_SKILLS                SendOpcode `opcode:"UPDATE_SKILLS"`
	WARP_TO_MAP                  SendOpcode `opcode:"WARP_TO_MAP"`
	SERVERMESSAGE                SendOpcode `opcode:"SERVERMESSAGE"`
	AVATAR_MEGA                  SendOpcode `opcode:"AVATAR_MEGA"`
	SPAWN_NPC                    SendOpcode `opcode:"SPAWN_NPC"`
	REMOVE_NPC                   SendOpcode `opcode:"REMOVE_NPC"`
	SPAWN_NPC_REQUEST_CONTROLLER SendOpcode `opcode:"SPAWN_NPC_REQUEST_CONTROLLER"`
	SPAWN_MONSTER                SendOpcode `opcode:"SPAWN_MONSTER"`
	SPAWN_MONSTER_CONTROL        SendOpcode `opcode:"SPAWN_MONSTER_CONTROL"`
	MOVE_MONSTER_RESPONSE        SendOpcode `opcode:"MOVE_MONSTER_RESPONSE"`
	CHATTEXT                     SendOpcode `opcode:"CHATTEXT"`
	SHOW_STATUS_INFO             SendOpcode `opcode:"SHOW_STATUS_INFO"`
	SHOW_MESO_GAIN               SendOpcode `opcode:"SHOW_MESO_GAIN"`
	SHOW_QUEST_COMPLETION        SendOpcode `opcode:"SHOW_QUEST_COMPLETION"`
	WHISPER                      SendOpcode `opcode:"WHISPER"`
	SPAWN_PLAYER                 SendOpcode `opcode:"SPAWN_PLAYER"`
	ANNOUNCE_PLAYER_SHOP         SendOpcode `opcode:"ANNOUNCE_PLAYER_SHOP"`
	SHOW_SCROLL_EFFECT           SendOpcode `opcode:"SHOW_SCROLL_EFFECT"`
	SHOW_ITEM_GAIN_INCHAT        SendOpcode `opcode:"SHOW_ITEM_GAIN_INCHAT"`
	DOJO_WARP_UP                 SendOpcode `opcode:"DOJO_WARP_UP"`
	CURRENT_MAP_WARP             SendOpcode `opcode:"CURRENT_MAP_WARP"`
	KILL_MONSTER                 SendOpcode `opcode:"KILL_MONSTER"`
	DROP_ITEM_FROM_MAPOBJECT     SendOpcode `opcode:"DROP_ITEM_FROM_MAPOBJECT"`
	FACIAL_EXPRESSION            SendOpcode `opcode:"FACIAL_EXPRESSION"`
	MOVE_PLAYER                  SendOpcode `opcode:"MOVE_PLAYER"`
	MOVE_MONSTER                 SendOpcode `opcode:"MOVE_MONSTER"`
	CLOSE_RANGE_ATTACK           SendOpcode `opcode:"CLOSE_RANGE_ATTACK"`
	RANGED_ATTACK                SendOpcode `opcode:"RANGED_ATTACK"`
	MAGIC_ATTACK                 SendOpcode `opcode:"MAGIC_ATTACK"`
	ENERGY_ATTACK                SendOpcode `opcode:"ENERGY_ATTACK"`
	OPEN_NPC_SHOP                SendOpcode `opcode:"OPEN_NPC_SHOP"`
	CONFIRM_SHOP_TRANSACTION     SendOpcode `opcode:"CONFIRM_SHOP_TRANSACTION"`
	OPEN_STORAGE                 SendOpcode `opcode:"OPEN_STORAGE"`
	MODIFY_INVENTORY_ITEM        SendOpcode `opcode:"MODIFY_INVENTORY_ITEM"`
	REMOVE_PLAYER_FROM_MAP       SendOpcode `opcode:"REMOVE_PLAYER_FROM_MAP"`
	REMOVE_ITEM_FROM_MAP         SendOpcode `opcode:"REMOVE_ITEM_FROM_MAP"`
	UPDATE_CHAR_LOOK             SendOpcode `opcode:"UPDATE_CHAR_LOOK"`
	SHOW_FOREIGN_EFFECT          SendOpcode `opcode:"SHOW_FOREIGN_EFFECT"`
	GIVE_FOREIGN_BUFF            SendOpcode `opcode:"GIVE_FOREIGN_BUFF"`
	CANCEL_FOREIGN_BUFF          SendOpcode `opcode:"CANCEL_FOREIGN_BUFF"`
	DAMAGE_PLAYER                SendOpcode `opcode:"DAMAGE_PLAYER"`
	CHAR_INFO                    SendOpcode `opcode:"CHAR_INFO"`
	UPDATE_QUEST_INFO            SendOpcode `opcode:"UPDATE_QUEST_INFO"`
	GIVE_BUFF                    SendOpcode `opcode:"GIVE_BUFF"`
	CANCEL_BUFF                  SendOpcode `opcode:"CANCEL_BUFF"`
	PLAYER_INTERACTION           SendOpcode `opcode:"PLAYER_INTERACTION"`
	UPDATE_CHAR_BOX              SendOpcode `opcode:"UPDATE_CHAR_BOX"`
	NPC_TALK                     SendOpcode `opcode:"NPC_TALK"`
	KEYMAP                       SendOpcode `opcode:"KEYMAP"`
	SHOW_MONSTER_HP              SendOpcode `opcode:"SHOW_MONSTER_HP"`
	PARTY_OPERATION              SendOpcode `opcode:"PARTY_OPERATION"`
	UPDATE_PARTYMEMBER_HP        SendOpcode `opcode:"UPDATE_PARTYMEMBER_HP"`
	MULTICHAT                    SendOpcode `opcode:"MULTICHAT"`
	APPLY_MONSTER_STATUS         SendOpcode `opcode:"APPLY_MONSTER_STATUS"`
	CANCEL_MONSTER_STATUS        SendOpcode `opcode:"CANCEL_MONSTER_STATUS"`
	CLOCK                        SendOpcode `opcode:"CLOCK"`
	SPAWN_PORTAL                 SendOpcode `opcode:"SPAWN_PORTAL"`
	SPAWN_DOOR                   SendOpcode `opcode:"SPAWN_DOOR"`
	REMOVE_DOOR                  SendOpcode `opcode:"REMOVE_DOOR"`
	SPAWN_SUMMON                 SendOpcode `opcode:"SPAWN_SUMMON"`
	REMOVE_SUMMON                SendOpcode `opcode:"REMOVE_SUMMON"`
	SUMMON_ATTACK                SendOpcode `opcode:"SUMMON_ATTACK"`
	MOVE_SUMMON                  SendOpcode `opcode:"MOVE_SUMMON"`
	SPAWN_MIST                   SendOpcode `opcode:"SPAWN_MIST"`
	REMOVE_MIST                  SendOpcode `opcode:"REMOVE_MIST"`
	DAMAGE_SUMMON                SendOpcode `opcode:"DAMAGE_SUMMON"`
	DAMAGE_MONSTER               SendOpcode `opcode:"DAMAGE_MONSTER"`
	BUDDYLIST                    SendOpcode `opcode:"BUDDYLIST"`
	SHOW_ITEM_EFFECT             SendOpcode `opcode:"SHOW_ITEM_EFFECT"`
	SHOW_CHAIR                   SendOpcode `opcode:"SHOW_CHAIR"`
	CANCEL_CHAIR                 SendOpcode `opcode:"CANCEL_CHAIR"`
	SKILL_EFFECT                 SendOpcode `opcode:"SKILL_EFFECT"`
	CANCEL_SKILL_EFFECT          SendOpcode `opcode:"CANCEL_SKILL_EFFECT"`
	BOSS_ENV                     SendOpcode `opcode:"BOSS_ENV"`
	REACTOR_SPAWN                SendOpcode `opcode:"REACTOR_SPAWN"`
	REACTOR_HIT                  SendOpcode `opcode:"REACTOR_HIT"`
	REACTOR_DESTROY              SendOpcode `opcode:"REACTOR_DESTROY"`
	MAP_EFFECT                   SendOpcode `opcode:"MAP_EFFECT"`
	GUILD_OPERATION              SendOpcode `opcode:"GUILD_OPERATION"`
	ALLIANCE_OPERATION           SendOpcode `opcode:"ALLIANCE_OPERATION"`
	BBS_OPERATION                SendOpcode `opcode:"BBS_OPERATION"`
	FAMILY                       SendOpcode `opcode:"FAMILY"`
	EARN_TITLE_MSG               SendOpcode `opcode:"EARN_TITLE_MSG"`
	SHOW_MAGNET                  SendOpcode `opcode:"SHOW_MAGNET"`
	MERCH_ITEM_MSG               SendOpcode `opcode:"MERCH_ITEM_MSG"`
	MERCH_ITEM_STORE             SendOpcode `opcode:"MERCH_ITEM_STORE"`
	MESSENGER                    SendOpcode `opcode:"MESSENGER"`
	NPC_ACTION                   SendOpcode `opcode:"NPC_ACTION"`
	SPAWN_PET                    SendOpcode `opcode:"SPAWN_PET"`
	MOVE_PET                     SendOpcode `opcode:"MOVE_PET"`
	PET_CHAT                     SendOpcode `opcode:"PET_CHAT"`
	PET_COMMAND                  SendOpcode `opcode:"PET_COMMAND"`
	PET_NAMECHANGE               SendOpcode `opcode:"PET_NAMECHANGE"`
	PET_FLAG_CHANGE              SendOpcode `opcode:"PET_FLAG_CHANGE"`
	COOLDOWN                     SendOpcode `opcode:"COOLDOWN"`
	PLAYER_HINT                  SendOpcode `opcode:"PLAYER_HINT"`
	SUMMON_HINT                  SendOpcode `opcode:"SUMMON_HINT"`
	SUMMON_HINT_MSG              SendOpcode `opcode:"SUMMON_HINT_MSG"`
	CYGNUS_INTRO_DISABLE_UI      SendOpcode `opcode:"CYGNUS_INTRO_DISABLE_UI"`
	CYGNUS_INTRO_LOCK            SendOpcode `opcode:"CYGNUS_INTRO_LOCK"`
	USE_SKILL_BOOK               SendOpcode `opcode:"USE_SKILL_BOOK"`
	SHOW_EQUIP_EFFECT            SendOpcode `opcode:"SHOW_EQUIP_EFFECT"`
	SKILL_MACRO                  SendOpcode `opcode:"SKILL_MACRO"`
	CS_OPEN                      SendOpcode `opcode:"CS_OPEN"`
	CS_UPDATE                    SendOpcode `opcode:"CS_UPDATE"`
	CS_OPERATION                 SendOpcode `opcode:"CS_OPERATION"`
	MTS_OPEN                     SendOpcode `opcode:"MTS_OPEN"`
	PLAYER_NPC                   SendOpcode `opcode:"PLAYER_NPC"`
	SHOW_NOTES                   SendOpcode `opcode:"SHOW_NOTES"`
	SUMMON_SKILL                 SendOpcode `opcode:"SUMMON_SKILL"`
	ARIANT_PQ_START              SendOpcode `opcode:"ARIANT_PQ_START"`
	CATCH_MONSTER                SendOpcode `opcode:"CATCH_MONSTER"`
	CATCH_ARIANT                 SendOpcode `opcode:"CATCH_ARIANT"`
	ARIANT_SCOREBOARD            SendOpcode `opcode:"ARIANT_SCOREBOARD"`
	ZAKUM_SHRINE                 SendOpcode `opcode:"ZAKUM_SHRINE"`
	BOAT_EFFECT                  SendOpcode `opcode:"BOAT_EFFECT"`
	CHALKBOARD                   SendOpcode `opcode:"CHALKBOARD"`
	DUEY                         SendOpcode `opcode:"DUEY"`
	TROCK_LOCATIONS              SendOpcode `opcode:"TROCK_LOCATIONS"`
	MONSTER_CARNIVAL_START       SendOpcode `opcode:"MONSTER_CARNIVAL_START"`
	MONSTER_CARNIVAL_OBTAINED_CP SendOpcode `opcode:"MONSTER_CARNIVAL_OBTAINED_CP"`
	MONSTER_CARNIVAL_PARTY_CP    SendOpcode `opcode:"MONSTER_CARNIVAL_PARTY_CP"`
	MONSTER_CARNIVAL_SUMMON      SendOpcode `opcode:"MONSTER_CARNIVAL_SUMMON"`
	MONSTER_CARNIVAL_SUMMON1     SendOpcode `opcode:"MONSTER_CARNIVAL_SUMMON1"`
	MONSTER_CARNIVAL_DIED        SendOpcode `opcode:"MONSTER_CARNIVAL_DIED"`
	SPAWN_HIRED_MERCHANT         SendOpcode `opcode:"SPAWN_HIRED_MERCHANT"`
	UPDATE_HIRED_MERCHANT        SendOpcode `opcode:"UPDATE_HIRED_MERCHANT"`
	SEND_TITLE_BOX               SendOpcode `opcode:"SEND_TITLE_BOX"`
	DESTROY_HIRED_MERCHANT       SendOpcode `opcode:"DESTROY_HIRED_MERCHANT"`
	UPDATE_MOUNT                 SendOpcode `opcode:"UPDATE_MOUNT"`
	MONSTERBOOK_ADD              SendOpcode `opcode:"MONSTERBOOK_ADD"`
	MONSTERBOOK_CHANGE_COVER     SendOpcode `opcode:"MONSTERBOOK_CHANGE_COVER"`
	FAIRY_PEND_MSG               SendOpcode `opcode:"FAIRY_PEND_MSG"`
	VICIOUS_HAMMER               SendOpcode `opcode:"VICIOUS_HAMMER"`
	FISHING_BOARD_UPDATE         SendOpcode `opcode:"FISHING_BOARD_UPDATE"`
	FISHING_CAUGHT               SendOpcode `opcode:"FISHING_CAUGHT"`
	OX_QUIZ                      SendOpcode `opcode:"OX_QUIZ"`
	ROLL_SNOWBALL                SendOpcode `opcode:"ROLL_SNOWBALL"`
	HIT_SNOWBALL                 SendOpcode `opcode:"HIT_SNOWBALL"`
	SNOWBALL_MESSAGE             SendOpcode `opcode:"SNOWBALL_MESSAGE"`
	LEFT_KNOCK_BACK              SendOpcode `opcode:"LEFT_KNOCK_BACK"`
	FINISH_SORT                  SendOpcode `opcode:"FINISH_SORT"`
	FINISH_GATHER                SendOpcode `opcode:"FINISH_GATHER"`
	SEND_PEDIGREE                SendOpcode `opcode:"SEND_PEDIGREE"`
	OPEN_FAMILY                  SendOpcode `opcode:"OPEN_FAMILY"`
	FAMILY_MESSAGE               SendOpcode `opcode:"FAMILY_MESSAGE"`
	FAMILY_INVITE                SendOpcode `opcode:"FAMILY_INVITE"`
	FAMILY_JUNIOR                SendOpcode `opcode:"FAMILY_JUNIOR"`
	SENIOR_MESSAGE               SendOpcode `opcode:"SENIOR_MESSAGE"`
	REP_INCREASE                 SendOpcode `opcode:"REP_INCREASE"`
	FAMILY_LOGGEDIN              SendOpcode `opcode:"FAMILY_LOGGEDIN"`
	FAMILY_BUFF                  SendOpcode `opcode:"FAMILY_BUFF"`
	FAMILY_USE_REQUEST           SendOpcode `opcode:"FAMILY_USE_REQUEST"`
	YELLOW_CHAT                  SendOpcode `opcode:"YELLOW_CHAT"`
	PIGMI_REWARD                 SendOpcode `opcode:"PIGMI_REWARD"`
	GM_EFFECT                    SendOpcode `opcode:"GM_EFFECT"`
	HIT_COCONUT                  SendOpcode `opcode:"HIT_COCONUT"`
	COCONUT_SCORE                SendOpcode `opcode:"COCONUT_SCORE"`
	LEVEL_UPDATE                 SendOpcode `opcode:"LEVEL_UPDATE"`
	MARRIAGE_UPDATE              SendOpcode `opcode:"MARRIAGE_UPDATE"`
	JOB_UPDATE                   SendOpcode `opcode:"JOB_UPDATE"`
	HORNTAIL_SHRINE              SendOpcode `opcode:"HORNTAIL_SHRINE"`
	STOP_CLOCK                   SendOpcode `opcode:"STOP_CLOCK"`
	MESOBAG_SUCCESS              SendOpcode `opcode:"MESOBAG_SUCCESS"`
	MESOBAG_FAILURE              SendOpcode `opcode:"MESOBAG_FAILURE"`
	SERVER_BLOCKED               SendOpcode `opcode:"SERVER_BLOCKED"`
	DRAGON_MOVE                  SendOpcode `opcode:"DRAGON_MOVE"`
	DRAGON_REMOVE                SendOpcode `opcode:"DRAGON_REMOVE"`
	DRAGON_SPAWN                 SendOpcode `opcode:"DRAGON_SPAWN"`
	ARAN_COMBO                   SendOpcode `opcode:"ARAN_COMBO"`
	TOP_MSG                      SendOpcode `opcode:"TOP_MSG"`
	TEMP_STATS                   SendOpcode `opcode:"TEMP_STATS"`
	TEMP_STATS_RESET             SendOpcode `opcode:"TEMP_STATS_RESET"`
	TUTORIAL_SUMMON              SendOpcode `opcode:"TUTORIAL_SUMMON"`
	REPAIR_WINDOW                SendOpcode `opcode:"REPAIR_WINDOW"`
	PYRAMID_UPDATE               SendOpcode `opcode:"PYRAMID_UPDATE"`
	PYRAMID_RESULT               SendOpcode `opcode:"PYRAMID_RESULT"`
	ENERGY                       SendOpcode `opcode:"ENERGY"`
	GET_MTS_TOKENS               SendOpcode `opcode:"GET_MTS_TOKENS"`
	MTS_OPERATION                SendOpcode `opcode:"MTS_OPERATION"`
	SHOW_POTENTIAL_EFFECT        SendOpcode `opcode:"SHOW_POTENTIAL_EFFECT"`
	SHOW_POTENTIAL_RESET         SendOpcode `opcode:"SHOW_POTENTIAL_RESET"`
	CHAOS_ZAKUM_SHRINE           SendOpcode `opcode:"CHAOS_ZAKUM_SHRINE"`
	CHAOS_HORNTAIL_SHRINE        SendOpcode `opcode:"CHAOS_HORNTAIL_SHRINE"`
	GAME_POLL_QUESTION           SendOpcode `opcode:"GAME_POLL_QUESTION"`
	GAME_POLL_REPLY              SendOpcode `opcode:"GAME_POLL_REPLY"`
	GMEVENT_INSTRUCTIONS         SendOpcode `opcode:"GMEVENT_INSTRUCTIONS"`
	BOAT_EFF                     SendOpcode `opcode:"BOAT_EFF"`
	OWL_OF_MINERVA               SendOpcode `opcode:"OWL_OF_MINERVA"`
	XMAS_SURPRISE                SendOpcode `opcode:"XMAS_SURPRISE"`
	CASH_SONG                    SendOpcode `opcode:"CASH_SONG"`
	UPDATE_INVENTORY_SLOT        SendOpcode `opcode:"UPDATE_INVENTORY_SLOT"`
	FOLLOW_REQUEST               SendOpcode `opcode:"FOLLOW_REQUEST"`
	FOLLOW_EFFECT                SendOpcode `opcode:"FOLLOW_EFFECT"`
	FOLLOW_MOVE                  SendOpcode `opcode:"FOLLOW_MOVE"`
	FOLLOW_MSG                   SendOpcode `opcode:"FOLLOW_MSG"`
	FOLLOW_MESSAGE               SendOpcode `opcode:"FOLLOW_MESSAGE"`
	TALK_MONSTER                 SendOpcode `opcode:"TALK_MONSTER"`
	REMOVE_TALK_MONSTER          SendOpcode `opcode:"REMOVE_TALK_MONSTER"`
	MONSTER_PROPERTIES           SendOpcode `opcode:"MONSTER_PROPERTIES"`
	MOVE_PLATFORM                SendOpcode `opcode:"MOVE_PLATFORM"`
	MOVE_ENV                     SendOpcode `opcode:"MOVE_ENV"`
	UPDATE_ENV                   SendOpcode `opcode:"UPDATE_ENV"`
	ENGAGE_REQUEST               SendOpcode `opcode:"ENGAGE_REQUEST"`
	GHOST_POINT                  SendOpcode `opcode:"GHOST_POINT"`
	GHOST_STATUS                 SendOpcode `opcode:"GHOST_STATUS"`
	ENGAGE_RESULT                SendOpcode `opcode:"ENGAGE_RESULT"`
	ENGLISH_QUIZ                 SendOpcode `opcode:"ENGLISH_QUIZ"`
	ARIANT_SCORE_UPDATE          SendOpcode `opcode:"ARIANT_SCORE_UPDATE"`
	RPS_GAME                     SendOpcode `opcode:"RPS_GAME"`
	UPDATE_BEANS                 SendOpcode `opcode:"UPDATE_BEANS"`
	BLOCK_MSG                    SendOpcode `opcode:"BLOCK_MSG"`
	AUTO_HP_POT                  SendOpcode `opcode:"AUTO_HP_POT"`
	AUTO_MP_POT                  SendOpcode `opcode:"AUTO_MP_POT"`
	LICENSE_RESULT               SendOpcode `opcode:"LICENSE_RESULT"`
	SPAWN_LOVE                   SendOpcode `opcode:"SPAWN_LOVE"`
	REMOVE_LOVE                  SendOpcode `opcode:"REMOVE_LOVE"`
	FORCED_MAP_EQUIP             SendOpcode `opcode:"FORCED_MAP_EQUIP"`
	SHOW_PREDICT_CARD            SendOpcode `opcode:"SHOW_PREDICT_CARD"`
	BEANS_TIPS                   SendOpcode `opcode:"BEANS_TIPS"`
	BEANS_GAME1                  SendOpcode `opcode:"BEANS_GAME1"`
	BEANS_GAME2                  SendOpcode `opcode:"BEANS_GAME2"`
}

func initSend() {
	Send = send{
		PING: 0,
		//LOGIN
		LOGIN_STATUS:         1,
		PIN_OPERATION:        2,
		SECONDPW_ERROR:       3,
		SERVERLIST:           4,
		SERVERSTATUS:         5,
		SERVER_IP:            6,
		CHARLIST:             7,
		CHAR_NAME_RESPONSE:   8,
		RELOG_RESPONSE:       9,
		ADD_NEW_CHAR_ENTRY:   10,
		DELETE_CHAR_RESPONSE: 11,
		CHANNEL_SELECTED:     12,
		ALL_CHARLIST:         13,
		CHOOSE_GENDER:        14,
		GENDER_SET:           15,
		CHAR_CASH:            16,
		//CHANNEL
		CHANGE_CHANNEL:               17,
		UPDATE_STATS:                 18,
		FAME_RESPONSE:                19,
		UPDATE_SKILLS:                20,
		WARP_TO_MAP:                  21,
		SERVERMESSAGE:                22,
		AVATAR_MEGA:                  23,
		SPAWN_NPC:                    24,
		REMOVE_NPC:                   25,
		SPAWN_NPC_REQUEST_CONTROLLER: 26,
		SPAWN_MONSTER:                27,
		SPAWN_MONSTER_CONTROL:        28,
		MOVE_MONSTER_RESPONSE:        29,
		CHATTEXT:                     3,
		SHOW_STATUS_INFO:             31,
		SHOW_MESO_GAIN:               32,
		SHOW_QUEST_COMPLETION:        33,
		WHISPER:                      34,
		SPAWN_PLAYER:                 35,
		ANNOUNCE_PLAYER_SHOP:         36,
		SHOW_SCROLL_EFFECT:           37,
		SHOW_ITEM_GAIN_INCHAT:        38,
		DOJO_WARP_UP:                 39,
		CURRENT_MAP_WARP:             40,
		KILL_MONSTER:                 41,
		DROP_ITEM_FROM_MAPOBJECT:     42,
		FACIAL_EXPRESSION:            43,
		MOVE_PLAYER:                  44,
		MOVE_MONSTER:                 45,
		CLOSE_RANGE_ATTACK:           46,
		RANGED_ATTACK:                47,
		MAGIC_ATTACK:                 48,
		ENERGY_ATTACK:                49,
		OPEN_NPC_SHOP:                50,
		CONFIRM_SHOP_TRANSACTION:     51,
		OPEN_STORAGE:                 52,
		MODIFY_INVENTORY_ITEM:        53,
		REMOVE_PLAYER_FROM_MAP:       54,
		REMOVE_ITEM_FROM_MAP:         55,
		UPDATE_CHAR_LOOK:             56,
		SHOW_FOREIGN_EFFECT:          57,
		GIVE_FOREIGN_BUFF:            58,
		CANCEL_FOREIGN_BUFF:          59,
		DAMAGE_PLAYER:                60,
		CHAR_INFO:                    61,
		UPDATE_QUEST_INFO:            62,
		GIVE_BUFF:                    63,
		CANCEL_BUFF:                  64,
		PLAYER_INTERACTION:           65,
		UPDATE_CHAR_BOX:              66,
		NPC_TALK:                     67,
		KEYMAP:                       68,
		SHOW_MONSTER_HP:              69,
		PARTY_OPERATION:              70,
		UPDATE_PARTYMEMBER_HP:        71,
		MULTICHAT:                    72,
		APPLY_MONSTER_STATUS:         73,
		CANCEL_MONSTER_STATUS:        74,
		CLOCK:                        75,
		SPAWN_PORTAL:                 76,
		SPAWN_DOOR:                   77,
		REMOVE_DOOR:                  78,
		SPAWN_SUMMON:                 79,
		REMOVE_SUMMON:                80,
		SUMMON_ATTACK:                81,
		MOVE_SUMMON:                  82,
		SPAWN_MIST:                   83,
		REMOVE_MIST:                  84,
		DAMAGE_SUMMON:                85,
		DAMAGE_MONSTER:               86,
		BUDDYLIST:                    87,
		SHOW_ITEM_EFFECT:             88,
		SHOW_CHAIR:                   89,
		CANCEL_CHAIR:                 90,
		SKILL_EFFECT:                 91,
		CANCEL_SKILL_EFFECT:          92,
		BOSS_ENV:                     93,
		REACTOR_SPAWN:                94,
		REACTOR_HIT:                  95,
		REACTOR_DESTROY:              96,
		MAP_EFFECT:                   97,
		GUILD_OPERATION:              98,
		ALLIANCE_OPERATION:           99,
		BBS_OPERATION:                100,
		FAMILY:                       101,
		EARN_TITLE_MSG:               102,
		SHOW_MAGNET:                  103,
		MERCH_ITEM_MSG:               104,
		MERCH_ITEM_STORE:             105,
		MESSENGER:                    106,
		NPC_ACTION:                   107,
		SPAWN_PET:                    108,
		MOVE_PET:                     109,
		PET_CHAT:                     110,
		PET_COMMAND:                  111,
		PET_NAMECHANGE:               112,
		PET_FLAG_CHANGE:              113,
		COOLDOWN:                     114,
		PLAYER_HINT:                  115,
		SUMMON_HINT:                  116,
		SUMMON_HINT_MSG:              117,
		CYGNUS_INTRO_DISABLE_UI:      118,
		CYGNUS_INTRO_LOCK:            119,
		USE_SKILL_BOOK:               120,
		SHOW_EQUIP_EFFECT:            121,
		SKILL_MACRO:                  122,
		CS_OPEN:                      123,
		CS_UPDATE:                    124,
		CS_OPERATION:                 125,
		MTS_OPEN:                     126,
		PLAYER_NPC:                   127,
		SHOW_NOTES:                   128,
		SUMMON_SKILL:                 129,
		ARIANT_PQ_START:              130,
		CATCH_MONSTER:                131,
		CATCH_ARIANT:                 132,
		ARIANT_SCOREBOARD:            133,
		ZAKUM_SHRINE:                 134,
		BOAT_EFFECT:                  135,
		CHALKBOARD:                   136,
		DUEY:                         137,
		TROCK_LOCATIONS:              138,
		MONSTER_CARNIVAL_START:       139,
		MONSTER_CARNIVAL_OBTAINED_CP: 140,
		MONSTER_CARNIVAL_PARTY_CP:    141,
		MONSTER_CARNIVAL_SUMMON:      142,
		MONSTER_CARNIVAL_SUMMON1:     143,
		MONSTER_CARNIVAL_DIED:        144,
		SPAWN_HIRED_MERCHANT:         145,
		UPDATE_HIRED_MERCHANT:        146,
		SEND_TITLE_BOX:               147,
		DESTROY_HIRED_MERCHANT:       148,
		UPDATE_MOUNT:                 149,
		MONSTERBOOK_ADD:              150,
		MONSTERBOOK_CHANGE_COVER:     151,
		FAIRY_PEND_MSG:               152,
		VICIOUS_HAMMER:               153,
		FISHING_BOARD_UPDATE:         154,
		FISHING_CAUGHT:               155,
		OX_QUIZ:                      156,
		ROLL_SNOWBALL:                157,
		HIT_SNOWBALL:                 158,
		SNOWBALL_MESSAGE:             159,
		LEFT_KNOCK_BACK:              160,
		FINISH_SORT:                  161,
		FINISH_GATHER:                162,
		SEND_PEDIGREE:                163,
		OPEN_FAMILY:                  164,
		FAMILY_MESSAGE:               165,
		FAMILY_INVITE:                166,
		FAMILY_JUNIOR:                167,
		SENIOR_MESSAGE:               168,
		REP_INCREASE:                 169,
		FAMILY_LOGGEDIN:              170,
		FAMILY_BUFF:                  171,
		FAMILY_USE_REQUEST:           172,
		YELLOW_CHAT:                  173,
		PIGMI_REWARD:                 174,
		GM_EFFECT:                    175,
		HIT_COCONUT:                  176,
		COCONUT_SCORE:                177,
		LEVEL_UPDATE:                 178,
		MARRIAGE_UPDATE:              179,
		JOB_UPDATE:                   180,
		HORNTAIL_SHRINE:              181,
		STOP_CLOCK:                   182,
		MESOBAG_SUCCESS:              183,
		MESOBAG_FAILURE:              184,
		SERVER_BLOCKED:               185,
		DRAGON_MOVE:                  186,
		DRAGON_REMOVE:                187,
		DRAGON_SPAWN:                 188,
		ARAN_COMBO:                   189,
		TOP_MSG:                      190,
		TEMP_STATS:                   191,
		TEMP_STATS_RESET:             192,
		TUTORIAL_SUMMON:              193,
		REPAIR_WINDOW:                194,
		PYRAMID_UPDATE:               195,
		PYRAMID_RESULT:               196,
		ENERGY:                       197,
		GET_MTS_TOKENS:               198,
		MTS_OPERATION:                199,
		SHOW_POTENTIAL_EFFECT:        200,
		SHOW_POTENTIAL_RESET:         201,
		CHAOS_ZAKUM_SHRINE:           202,
		CHAOS_HORNTAIL_SHRINE:        203,
		GAME_POLL_QUESTION:           204,
		GAME_POLL_REPLY:              205,
		GMEVENT_INSTRUCTIONS:         206,
		BOAT_EFF:                     207,
		OWL_OF_MINERVA:               208,
		XMAS_SURPRISE:                209,
		CASH_SONG:                    210,
		UPDATE_INVENTORY_SLOT:        211,
		FOLLOW_REQUEST:               212,
		FOLLOW_EFFECT:                213,
		FOLLOW_MOVE:                  214,
		FOLLOW_MSG:                   215,
		FOLLOW_MESSAGE:               216,
		TALK_MONSTER:                 217,
		REMOVE_TALK_MONSTER:          218,
		MONSTER_PROPERTIES:           219,
		MOVE_PLATFORM:                220,
		MOVE_ENV:                     221,
		UPDATE_ENV:                   222,
		ENGAGE_REQUEST:               223,
		GHOST_POINT:                  224,
		GHOST_STATUS:                 225,
		ENGAGE_RESULT:                226,
		ENGLISH_QUIZ:                 227,
		ARIANT_SCORE_UPDATE:          228,
		RPS_GAME:                     229,
		UPDATE_BEANS:                 230,
		BLOCK_MSG:                    231,
		AUTO_HP_POT:                  232,
		AUTO_MP_POT:                  233,
		LICENSE_RESULT:               234,
		SPAWN_LOVE:                   235,
		REMOVE_LOVE:                  236,
		FORCED_MAP_EQUIP:             237,
		SHOW_PREDICT_CARD:            238,
		BEANS_TIPS:                   239,
		BEANS_GAME1:                  240,
		BEANS_GAME2:                  241,
	}

	sendTypeReflect := reflect.TypeOf(Send)
	sendValueReflect := reflect.ValueOf(Send)
	mapValueReflect := reflect.ValueOf(sendOpcodeStringMap)
	fieldNum := sendTypeReflect.NumField()
	for i := 0; i < fieldNum; i++ {
		field := sendTypeReflect.Field(i)
		value := sendValueReflect.FieldByName(field.Name)
		tagName := field.Tag.Get("opcode")
		mapValueReflect.SetMapIndex(value, reflect.ValueOf(tagName))
	}
}

// SendOpcode 接收到的数据包操作符
type SendOpcode Opcode

func (code SendOpcode) Type() OpcodeType {
	return OpcodeTypeReceive
}
func (code SendOpcode) Int16() int16 {
	return int16(code)
}
func (code SendOpcode) String() string {
	//str := sendOpcodeStringMap[code]
	//return fmt.Sprintf("SendOpcode:%d => %s", code, str)
	return ""
}

var sendOpcodeStringMap = map[SendOpcode]string{}

//go:embed recvops.properties
//go:embed sendops.properties
var embedData embed.FS

func init() {
	initSend()
	initReceive()

	readData, err := embedData.ReadFile("sendops.properties")
	utils.PanicError(err)

	file, err := ini.Load(readData)
	utils.PanicError(err)
	section := file.Section("")
	sendReflect := reflect.ValueOf(&Send).Elem()
	for _, key := range section.Keys() {
		if value := sendReflect.FieldByName(key.Name()); value.IsValid() && !value.IsZero() && value.CanSet() {
			code := SendOpcode(key.MustInt())
			value.Set(reflect.ValueOf(code))
			sendOpcodeStringMap[code] = key.Name()

		}
	}

	readData, err = embedData.ReadFile("recvops.properties")
	utils.PanicError(err)
	file, err = ini.LooseLoad(readData)
	utils.PanicError(err)
	section = file.Section("")
	//receiveOpcodeStringMap = make(map[int16]string)
	receiveReflect := reflect.ValueOf(&Receive).Elem()
	for _, key := range section.Keys() {
		if value := receiveReflect.FieldByName(key.Name()); value.IsValid() && !value.IsZero() && value.CanSet() {
			code := ReceiveOpcode(key.MustInt())
			value.Set(reflect.ValueOf(code))
			receiveOpcodeStringMap[code] = key.Name()
		}
	}
}
