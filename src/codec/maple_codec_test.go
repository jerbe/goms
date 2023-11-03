package codec

import (
	"bytes"
	"fmt"
	"github.com/jerbe/goms/constants"
	"io"
	"log"
	"testing"
	"time"

	"github.com/jerbe/goms/crypt"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/25 12:33
  @describe :
*/

type M struct {
	P int
}

func (m *M) Clone() *M {
	x := *m
	return &x
}

func TestMClone(t *testing.T) {
	m := &M{P: 1}
	fmt.Printf("%p", m)

	n := m.Clone()
	fmt.Printf("%p", n)

}

func TestNewMapleCodec(t *testing.T) {
	encodeCypher, err := crypt.NewMapleCypher(constants.OFBKey, []byte{1, 0x5F, 4, 0x3F}, 0xFFFF-constants.MapleVersion)
	if err != nil {
		panic(err)
	}

	decodeCypher, err := crypt.NewMapleCypher(constants.OFBKey, []byte{9, 0, 0x5, 0x5F}, constants.MapleVersion)
	if err != nil {
		panic(err)
	}

	codec := NewMapleCodec(encodeCypher, decodeCypher)

	//data := []byte("我他喵的是你大爷😈")
	data := "hello"
	//data := 9999999999999999999999999199999999999999999999999991.000000000000001
	//data := "\n"
	//encodeInput := &bytes.Buffer{}
	encodeOutput := &bytes.Buffer{}
	decodeOutput := &bytes.Buffer{}
	go func() {
		for {
			for encodeOutput.Len() == 0 {

			}
			codec.Decode(encodeOutput, decodeOutput)
			{
				res, err := io.ReadAll(decodeOutput)
				log.Println(string(res), err)
			}
		}

	}()
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		//encodeInput.Write(data)
		if err := codec.Encode(data, encodeOutput); err != nil {
			log.Panicln("encode:", err)
		}
	}
}

func TestMapleEncoder_Encode(t *testing.T) {
	encodeCypher, err := crypt.NewMapleCypher(constants.OFBKey, []byte{1, 0x5F, 4, 0x3F}, 0xFFFF-constants.MapleVersion)
	if err != nil {
		panic(err)
	}

	data := []byte{9, 0, 0, 8, 0, 196, 227, 186, 195, 202, 192, 189, 231, 3, 0, 0, 100, 0, 100, 0, 3, 244, 1, 0, 0, 10, 0, 196, 227, 186, 195, 202, 192, 189, 231, 45, 49, 110, 0, 0, 0, 0, 0, 0, 10, 0, 196, 227, 186, 195, 202, 192, 189, 231, 45, 50, 120, 0, 0, 0, 0, 1, 0, 10, 0, 196, 227, 186, 195, 202, 192, 189, 231, 45, 51, 130, 0, 0, 0, 0, 2, 0, 0, 0}

	buffer := &bytes.Buffer{}
	encoder := NewMapleEncoder(encodeCypher)
	encoder.Encode(data, buffer)
	result, err := io.ReadAll(buffer)
	if err != nil {
		panic(err)
	}
	log.Println(result)
}
