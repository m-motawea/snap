package snap

import (
	"log"
	"reflect"
	"testing"
)

func TestSNAP(t *testing.T) {
	s := SNAP{
		OUI:        OUI(0), // Default EthernetTypes
		ProtocolID: ProtocolID(0x0800),
		Data: []byte{
			00, 00, 0x45, 00, 0x04, 0x1f, 0xa7, 0x29, 0x40, 0x00, 0x40, 0x11, 0x62, 0xf7, 0xc0, 0xa8, 0x01, 0x08,
			00, 0x10, 0x9c, 0xc6, 0xce, 0x36,
			00, 0x00, 0x98, 0xe1, 0x57, 0x65, 0x04, 0x0b, 0x30, 0xca,
		},
	}
	b, err := s.MarshalBinary()
	if err != nil {
		t.Errorf("marashaling header failed due to error %v", err)
	}
	log.Printf("Marshal result:\n%v", b)
	s2 := SNAP{}
	err = s2.UnmarshalBinary(b)
	if err != nil {
		t.Errorf("unmarashaling header failed due to error %v", err)
	}
	if !reflect.DeepEqual(s, s2) {
		t.Errorf("two SNAP are not equal. test1:\n%v\ntest2:\n%v\n", s, s2)
	}
}
