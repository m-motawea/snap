package snap

import (
	"encoding/binary"
	"errors"
)

type OUI uint32 // should be 24 bit only max 0xFFFFFF
type ProtocolID uint16

func (o OUI) IsValid() bool {
	if o > 0xFFFFFF {
		return false
	}
	return true
}

type SNAP struct {
	OUI OUI
	ProtocolID ProtocolID
	Data []byte
}

func (s *SNAP) MarshalBinary() ([]byte, error) {
	if len(s.Data) > 1492 {
		return []byte{}, errors.New("invalid payload length")
	}
	b := make([]byte, 5 + len(s.Data))
	if !s.OUI.IsValid() {
		return b, errors.New("invalid value for OUI")
	}
	temp := uint32(s.OUI) << 8
	binary.BigEndian.PutUint32(b[0:4], temp)
	binary.BigEndian.PutUint16(b[3:], uint16(s.ProtocolID))
	if len(s.Data) > 0 {
		copy(b[5:], s.Data)	
	}
	return b, nil
}

func (s *SNAP) UnmarshalBinary(b []byte) error {
	if len(b) < 5 || len(b) > 1497 {
		return errors.New("invalid header length")
	}
	o := OUI(binary.BigEndian.Uint32(b[0:4]) >> 8)
	s.OUI = o
	s.ProtocolID = ProtocolID(binary.BigEndian.Uint16(b[3:5]))
	if len(b) > 5 {
		s.Data = make([]byte, len(b) - 5)
		copy(s.Data, b[5:])
	}
	return nil
}