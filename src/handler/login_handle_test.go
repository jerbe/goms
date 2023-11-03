package handler

import (
	"github.com/jerbe/goms/utils"
	"log"
	"testing"
	"time"

	"github.com/jerbe/goms/client"
	_ "github.com/jerbe/goms/config"
	"github.com/jerbe/goms/data/packet"
	_ "github.com/jerbe/goms/database"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/11/1 10:31
  @describe :
*/

func TestHandler_handleCharListRequest(t *testing.T) {
	utils.InitFuncQueue.Run()
	time.Sleep(time.Second)

	cli := client.Client{}
	cli.SetAccountID(1)
	cli.SetAccountName("犀利哥")
	cli.GetCharacterSlots()
	cli.SetWorld(0)
	cli.SetChannel(1)
	charList := client.LoadCharacterList(cli.GetWorld(), 1)

	p := packet.CharacterList(charList, false, cli.GetCharacterSlots())

	bytes := p.Bytes()
	log.Println(bytes)
}
