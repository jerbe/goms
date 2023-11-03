package handler

import (
	"github.com/jerbe/goms/constants"
	"github.com/jerbe/goms/data/packet"
	"github.com/jerbe/goms/server"
	"io"
	"log"

	"github.com/jerbe/goms/client"
	"github.com/jerbe/goms/data"
	"github.com/jerbe/goms/data/packet/code"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/26 22:52
  @describe :
*/

var (
	ChannelServerPool map[int]*server.ChannelServer
)

// -------------------------------------------------------------------------------------------------

func NewHandler(server server.IServer) *Handler {
	return &Handler{server: server}
}

// Handler 处理器
type Handler struct {
	// 基础服务
	server server.IServer
}

// OnClientConnect 当客户端连接时
func (h *Handler) OnClientConnect(cli *client.Client) error {
	return cli.SendMessageWithoutEncode(packet.SayHello(int16(constants.MapleVersion), cli.GetEncodeCypher().OriginalIv(), cli.GetDecodeCypher().OriginalIv()))
}

// OnClientReceiveMessage 当客户端接收到信息
func (h *Handler) OnClientReceiveMessage(cli *client.Client, in io.Reader) {
	dataReader, err := data.NewLittleEndianReaderFromReader(in)
	if err != nil {
		log.Printf("Handler OnClientReceiveMessage error. reason:%v ", err)
		return
	}

	opcode, err := dataReader.ReadShort()
	if err != nil {
		return
	}
	h.packetHandle(cli, code.NewReceiveOpcode(opcode), dataReader)
}

func (h *Handler) handle(client *client.Client, opcode code.ReceiveOpcode, reader *data.LittleEndianReader) {
	log.Printf("未实现[%s]的操作方法", opcode)
}

// packetHandle 数据包分发处理
func (h *Handler) packetHandle(client *client.Client, opcode code.ReceiveOpcode, reader *data.LittleEndianReader) {
	log.Printf("Receive Code:%v", opcode)
	switch opcode {
	case code.Receive.PONG:
		h.handle(client, opcode, reader)
	//case code.Receive.LOGIN:
	//	h.handle(client,opcode,reader)
	case code.Receive.LOGIN_PASSWORD:
		h.handleLoginPassword(client, reader)

	case code.Receive.HELLO_LOGIN:
		h.handleHelloLogin(client, reader)

	case code.Receive.HELLO_CHANNEL:
		h.handle(client, opcode, reader)

	case code.Receive.LICENSE_REQUEST:
		h.handle(client, opcode, reader)

	case code.Receive.SERVERLIST_REQUEST:
		h.handle(client, opcode, reader)

	case code.Receive.CHARLIST_REQUEST:
		h.handleCharListRequest(client, reader)

	case code.Receive.SERVERSTATUS_REQUEST:
		h.handleServerStatusRequest(client, reader)

	case code.Receive.CHECK_CHAR_NAME:
		h.handle(client, opcode, reader)

	case code.Receive.CREATE_CHAR:
		h.handle(client, opcode, reader)

	case code.Receive.DELETE_CHAR:
		h.handle(client, opcode, reader)
	case code.Receive.STRANGE_DATA:
		h.handle(client, opcode, reader)
	case code.Receive.CHAR_SELECT:
		h.handleSelectCharacter(client, reader)
	case code.Receive.AUTH_SECOND_PASSWORD:
		h.handle(client, opcode, reader)
	case code.Receive.SET_GENDER:
		h.handle(client, opcode, reader)
	case code.Receive.RSA_KEY:
		h.handle(client, opcode, reader)
	//case code.Receive.CHANNEL:
	//	h.handle(client,opcode,reader)
	case code.Receive.PLAYER_LOGGEDIN:
		h.handle(client, opcode, reader)
	case code.Receive.CHANGE_MAP:
		h.handle(client, opcode, reader)
	case code.Receive.CHANGE_CHANNEL:
		h.handle(client, opcode, reader)
	case code.Receive.ENTER_CASH_SHOP:
		h.handle(client, opcode, reader)
	case code.Receive.MOVE_PLAYER:
		h.handle(client, opcode, reader)
	case code.Receive.CANCEL_CHAIR:
		h.handle(client, opcode, reader)
	case code.Receive.USE_CHAIR:
		h.handle(client, opcode, reader)
	case code.Receive.CLOSE_RANGE_ATTACK:
		h.handle(client, opcode, reader)
	case code.Receive.RANGED_ATTACK:
		h.handle(client, opcode, reader)
	case code.Receive.MAGIC_ATTACK:
		h.handle(client, opcode, reader)
	case code.Receive.PASSIVE_ENERGY:
		h.handle(client, opcode, reader)
	case code.Receive.TAKE_DAMAGE:
		h.handle(client, opcode, reader)
	case code.Receive.GENERAL_CHAT:
		h.handle(client, opcode, reader)
	case code.Receive.CLOSE_CHALKBOARD:
		h.handle(client, opcode, reader)
	case code.Receive.FACE_EXPRESSION:
		h.handle(client, opcode, reader)
	case code.Receive.USE_ITEMEFFECT:
		h.handle(client, opcode, reader)
	case code.Receive.WHEEL_OF_FORTUNE:
		h.handle(client, opcode, reader)
	case code.Receive.MONSTER_BOOK_COVER:
		h.handle(client, opcode, reader)
	case code.Receive.NPC_TALK:
		h.handle(client, opcode, reader)
	case code.Receive.NPC_TALK_MORE:
		h.handle(client, opcode, reader)
	case code.Receive.NPC_SHOP:
		h.handle(client, opcode, reader)
	case code.Receive.STORAGE:
		h.handle(client, opcode, reader)
	case code.Receive.USE_HIRED_MERCHANT:
		h.handle(client, opcode, reader)
	case code.Receive.MERCH_ITEM_STORE:
		h.handle(client, opcode, reader)
	case code.Receive.DUEY_ACTION:
		h.handle(client, opcode, reader)
	case code.Receive.ITEM_SORT:
		h.handle(client, opcode, reader)
	case code.Receive.ITEM_GATHER:
		h.handle(client, opcode, reader)
	case code.Receive.ITEM_MOVE:
		h.handle(client, opcode, reader)
	case code.Receive.USE_ITEM:
		h.handle(client, opcode, reader)
	case code.Receive.CANCEL_ITEM_EFFECT:
		h.handle(client, opcode, reader)
	case code.Receive.USE_SUMMON_BAG:
		h.handle(client, opcode, reader)
	case code.Receive.PET_EXCEPTIONLIST:
		h.handle(client, opcode, reader)
	case code.Receive.PET_FOOD:
		h.handle(client, opcode, reader)
	case code.Receive.USE_MOUNT_FOOD:
		h.handle(client, opcode, reader)
	case code.Receive.USE_SCRIPTED_NPC_ITEM:
		h.handle(client, opcode, reader)
	case code.Receive.USE_CASH_ITEM:
		h.handle(client, opcode, reader)
	case code.Receive.USE_CATCH_ITEM:
		h.handle(client, opcode, reader)
	case code.Receive.USE_SKILL_BOOK:
		h.handle(client, opcode, reader)
	case code.Receive.USE_RETURN_SCROLL:
		h.handle(client, opcode, reader)
	case code.Receive.USE_UPGRADE_SCROLL:
		h.handle(client, opcode, reader)
	case code.Receive.DISTRIBUTE_AP:
		h.handle(client, opcode, reader)
	case code.Receive.AUTO_ASSIGN_AP:
		h.handle(client, opcode, reader)
	case code.Receive.HEAL_OVER_TIME:
		h.handle(client, opcode, reader)
	case code.Receive.DISTRIBUTE_SP:
		h.handle(client, opcode, reader)
	case code.Receive.SPECIAL_MOVE:
		h.handle(client, opcode, reader)
	case code.Receive.CANCEL_BUFF:
		h.handle(client, opcode, reader)
	case code.Receive.SKILL_EFFECT:
		h.handle(client, opcode, reader)
	case code.Receive.MESO_DROP:
		h.handle(client, opcode, reader)
	case code.Receive.GIVE_FAME:
		h.handle(client, opcode, reader)
	case code.Receive.CHAR_INFO_REQUEST:
		h.handle(client, opcode, reader)
	case code.Receive.SPAWN_PET:
		h.handle(client, opcode, reader)
	case code.Receive.CANCEL_DEBUFF:
		h.handle(client, opcode, reader)
	case code.Receive.CHANGE_MAP_SPECIAL:
		h.handle(client, opcode, reader)
	case code.Receive.USE_INNER_PORTAL:
		h.handle(client, opcode, reader)
	case code.Receive.TROCK_ADD_MAP:
		h.handle(client, opcode, reader)
	case code.Receive.QUEST_ACTION:
		h.handle(client, opcode, reader)
	case code.Receive.EFFECT_ON_OFF:
		h.handle(client, opcode, reader)
	case code.Receive.SKILL_MACRO:
		h.handle(client, opcode, reader)
	case code.Receive.ITEM_BAOWU:
		h.handle(client, opcode, reader)
	case code.Receive.ITEM_SUNZI:
		h.handle(client, opcode, reader)
	case code.Receive.ITEM_MAKER:
		h.handle(client, opcode, reader)
	case code.Receive.USE_TREASUER_CHEST:
		h.handle(client, opcode, reader)
	case code.Receive.PARTYCHAT:
		h.handle(client, opcode, reader)
	case code.Receive.PARTY_SS:
		h.handle(client, opcode, reader)
	case code.Receive.WHISPER:
		h.handle(client, opcode, reader)
	case code.Receive.MESSENGER:
		h.handle(client, opcode, reader)
	case code.Receive.PLAYER_INTERACTION:
		h.handle(client, opcode, reader)
	case code.Receive.PARTY_OPERATION:
		h.handle(client, opcode, reader)
	case code.Receive.DENY_PARTY_REQUEST:
		h.handle(client, opcode, reader)
	case code.Receive.GUILD_OPERATION:
		h.handle(client, opcode, reader)
	case code.Receive.DENY_GUILD_REQUEST:
		h.handle(client, opcode, reader)
	case code.Receive.BUDDYLIST_MODIFY:
		h.handle(client, opcode, reader)
	case code.Receive.NOTE_ACTION:
		h.handle(client, opcode, reader)
	case code.Receive.USE_DOOR:
		h.handle(client, opcode, reader)
	case code.Receive.CHANGE_KEYMAP:
		h.handle(client, opcode, reader)
	case code.Receive.UPDATE_CHAR_INFO:
		h.handle(client, opcode, reader)
	case code.Receive.ENTER_MTS:
		h.handle(client, opcode, reader)
	case code.Receive.ALLIANCE_OPERATION:
		h.handle(client, opcode, reader)
	case code.Receive.DENY_ALLIANCE_REQUEST:
		h.handle(client, opcode, reader)
	case code.Receive.REQUEST_FAMILY:
		h.handle(client, opcode, reader)
	case code.Receive.OPEN_FAMILY:
		h.handle(client, opcode, reader)
	case code.Receive.FAMILY_OPERATION:
		h.handle(client, opcode, reader)
	case code.Receive.DELETE_JUNIOR:
		h.handle(client, opcode, reader)
	case code.Receive.DELETE_SENIOR:
		h.handle(client, opcode, reader)
	case code.Receive.ACCEPT_FAMILY:
		h.handle(client, opcode, reader)
	case code.Receive.USE_FAMILY:
		h.handle(client, opcode, reader)
	case code.Receive.FAMILY_PRECEPT:
		h.handle(client, opcode, reader)
	case code.Receive.FAMILY_SUMMON:
		h.handle(client, opcode, reader)
	case code.Receive.CYGNUS_SUMMON:
		h.handle(client, opcode, reader)
	case code.Receive.ARAN_COMBO:
		h.handle(client, opcode, reader)
	case code.Receive.BBS_OPERATION:
		h.handle(client, opcode, reader)
	case code.Receive.TRANSFORM_PLAYER:
		h.handle(client, opcode, reader)
	case code.Receive.MOVE_PET:
		h.handle(client, opcode, reader)
	case code.Receive.PET_CHAT:
		h.handle(client, opcode, reader)
	case code.Receive.PET_COMMAND:
		h.handle(client, opcode, reader)
	case code.Receive.PET_LOOT:
		h.handle(client, opcode, reader)
	case code.Receive.PET_AUTO_POT:
		h.handle(client, opcode, reader)
	case code.Receive.MOVE_SUMMON:
		h.handle(client, opcode, reader)
	case code.Receive.SUMMON_ATTACK:
		h.handle(client, opcode, reader)
	case code.Receive.DAMAGE_SUMMON:
		h.handle(client, opcode, reader)
	case code.Receive.MOVE_LIFE:
		h.handle(client, opcode, reader)
	case code.Receive.AUTO_AGGRO:
		h.handle(client, opcode, reader)
	case code.Receive.FRIENDLY_DAMAGE:
		h.handle(client, opcode, reader)
	case code.Receive.MONSTER_BOMB:
		h.handle(client, opcode, reader)
	case code.Receive.HYPNOTIZE_DMG:
		h.handle(client, opcode, reader)
	case code.Receive.NPC_ACTION:
		h.handle(client, opcode, reader)
	case code.Receive.ITEM_PICKUP:
		h.handle(client, opcode, reader)
	case code.Receive.DAMAGE_REACTOR:
		h.handle(client, opcode, reader)
	case code.Receive.SNOWBALL:
		h.handle(client, opcode, reader)
	case code.Receive.LEFT_KNOCK_BACK:
		h.handle(client, opcode, reader)
	case code.Receive.COCONUT:
		h.handle(client, opcode, reader)
	case code.Receive.MONSTER_CARNIVAL:
		h.handle(client, opcode, reader)
	case code.Receive.SHIP_OBJECT:
		h.handle(client, opcode, reader)
	case code.Receive.CS_UPDATE:
		h.handle(client, opcode, reader)
	case code.Receive.BUY_CS_ITEM:
		h.handle(client, opcode, reader)
	case code.Receive.TOUCHING_CS:
		h.handle(client, opcode, reader)
	case code.Receive.COUPON_CODE:
		h.handle(client, opcode, reader)
	case code.Receive.MAPLETV:
		h.handle(client, opcode, reader)
	case code.Receive.MOVE_DRAGON:
		h.handle(client, opcode, reader)
	case code.Receive.REPAIR:
		h.handle(client, opcode, reader)
	case code.Receive.REPAIR_ALL:
		h.handle(client, opcode, reader)
	case code.Receive.TOUCHING_MTS:
		h.handle(client, opcode, reader)
	case code.Receive.USE_MAGNIFY_GLASS:
		h.handle(client, opcode, reader)
	case code.Receive.USE_POTENTIAL_SCROLL:
		h.handle(client, opcode, reader)
	case code.Receive.USE_EQUIP_SCROLL:
		h.handle(client, opcode, reader)
	case code.Receive.GAME_POLL:
		h.handle(client, opcode, reader)
	case code.Receive.OWL:
		h.handle(client, opcode, reader)
	case code.Receive.OWL_WARP:
		h.handle(client, opcode, reader)
	//case code.Receive.XMAS_SURPRISE//header->uniqueid(long)isentirestructure:
	//h.handle(client,opcode,reader)
	case code.Receive.USE_OWL_MINERVA:
		h.handle(client, opcode, reader)
	case code.Receive.RPS_GAME:
		h.handle(client, opcode, reader)
	case code.Receive.UPDATE_QUEST:
		h.handle(client, opcode, reader)
	case code.Receive.PLAYER_UPDATE:
		h.handle(client, opcode, reader)
	//case code.Receive.QUEST_ITEM//header->questid(int)->1/0(byteopenorclose):
	//	h.handle(client,opcode,reader)
	case code.Receive.USE_ITEM_QUEST:
		h.handle(client, opcode, reader)
	case code.Receive.FOLLOW_REQUEST:
		h.handle(client, opcode, reader)
	case code.Receive.FOLLOW_REPLY:
		h.handle(client, opcode, reader)
	case code.Receive.MOB_NODE:
		h.handle(client, opcode, reader)
	case code.Receive.DISPLAY_NODE:
		h.handle(client, opcode, reader)
	case code.Receive.TOUCH_REACTOR:
		h.handle(client, opcode, reader)
	case code.Receive.RING_ACTION:
		h.handle(client, opcode, reader)
	case code.Receive.MTS_TAB:
		h.handle(client, opcode, reader)
	case code.Receive.ChatRoom_SYSTEM:
		h.handle(client, opcode, reader)
	case code.Receive.Quest_KJ:
		h.handle(client, opcode, reader)
	case code.Receive.NEW_SX:
		h.handle(client, opcode, reader)
	case code.Receive.BOATS:
		h.handle(client, opcode, reader)
	case code.Receive.BEANS_GAME1:
		h.handle(client, opcode, reader)
	case code.Receive.BEANS_GAME2:
		h.handle(client, opcode, reader)
	case code.Receive.MOONRABBIT_HP:
		h.handle(client, opcode, reader)
	case code.Receive.MARRAGE_RECV:
		h.handle(client, opcode, reader)

	}
}
