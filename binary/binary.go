package binary

import (
	"errors"
	"io"
	"math"
)

const (
	ByteSize  = 1
	ShortSize = 2
	TriadSize = 3
	IntSize   = 4
	LongSize  = 8

	FloatSize  = 4
	DoubleSize = 8
)

type Triad uint32

/*

//MaxTriadSize max size of the Triad. range: 0 - 16777215 (2^24 = 16777216)
const MaxTriadSize int = 0xffffff

*/

//NotEnoughErr .
var NotEnoughErr = errors.New("Not enough")

//TODO: adds comments

//Binary

func ReadByte(v []byte) byte {
	return v[0]
}

func WriteByte(v byte) []byte {
	return []byte{
		v,
	}
}

func ReadSByte(v []byte) int8 {
	return int8(v[0])
}

func WriteSByte(v int8) []byte {
	return []byte{
		byte(v),
	}
}

func ReadShort(v []byte) int16 {
	return int16(v[0])<<8 | int16(v[1])
}

func WriteShort(v int16) []byte {
	return []byte{
		byte(v >> 8),
		byte(v),
	}
}

func ReadLShort(v []byte) int16 {
	return int16(v[0]) | int16(v[1])<<8
}

func WriteLShort(v int16) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
	}
}

func ReadUShort(v []byte) uint16 {
	return uint16(v[0])<<8 | uint16(v[1])
}

func WriteUShort(v uint16) []byte {
	return []byte{
		byte(v >> 8),
		byte(v),
	}
}

func ReadLUShort(v []byte) uint16 {
	return uint16(v[0]) | uint16(v[1])<<8
}

func WriteLUShort(v uint16) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
	}
}

func ReadTriad(v []byte) Triad {
	return Triad(v[0])<<16 | Triad(v[1])<<8 | Triad(v[2])
}

func WriteTriad(v Triad) []byte {
	return []byte{
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}

func ReadLTriad(v []byte) Triad {
	return Triad(v[0]) | Triad(v[1])<<8 | Triad(v[2])<<16
}

func WriteLTriad(v Triad) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
	}
}

func ReadInt(v []byte) int32 {
	return int32(v[0])<<24 | int32(v[1])<<16 | int32(v[2])<<8 | int32(v[3])
}

func WriteInt(v int32) []byte {
	return []byte{
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}

func ReadUInt(v []byte) uint32 {
	return uint32(v[0])<<24 | uint32(v[1])<<16 | uint32(v[2])<<8 | uint32(v[3])
}

func WriteUInt(v uint32) []byte {
	return []byte{
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}

func ReadLInt(v []byte) int32 {
	return int32(v[0]) | int32(v[1])<<8 | int32(v[2])<<16 | int32(v[3])<<24
}

func WriteLInt(v int32) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
	}
}

func ReadLUInt(v []byte) uint32 {
	return uint32(v[0]) | uint32(v[1])<<8 | uint32(v[2])<<16 | uint32(v[3])<<24
}

func WriteLUInt(v uint32) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
	}
}

func ReadLong(v []byte) int64 {
	return int64(v[0])<<56 | int64(v[1])<<48 | int64(v[2])<<40 | int64(v[3])<<32 |
		int64(v[4])<<24 | int64(v[5])<<16 | int64(v[6])<<8 | int64(v[7])
}

func WriteLong(v int64) []byte {
	return []byte{
		byte(v >> 56),
		byte(v >> 48),
		byte(v >> 40),
		byte(v >> 32),
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}

func ReadULong(v []byte) uint64 {
	return uint64(v[0])<<56 | uint64(v[1])<<48 | uint64(v[2])<<40 | uint64(v[3])<<32 |
		uint64(v[4])<<24 | uint64(v[5])<<16 | uint64(v[6])<<8 | uint64(v[7])
}

func WriteULong(v uint64) []byte {
	return []byte{
		byte(v >> 56),
		byte(v >> 48),
		byte(v >> 40),
		byte(v >> 32),
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}

func ReadLLong(v []byte) int64 {
	return int64(v[0]) | int64(v[1])<<8 | int64(v[2])<<16 | int64(v[3])<<24 |
		int64(v[4])<<32 | int64(v[5])<<40 | int64(v[6])<<48 | int64(v[7])<<56
}

func WriteLLong(v int64) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
		byte(v >> 32),
		byte(v >> 40),
		byte(v >> 48),
		byte(v >> 56),
	}
}

func ReadLULong(v []byte) uint64 {
	return uint64(v[0]) | uint64(v[1])<<8 | uint64(v[2])<<16 | uint64(v[3])<<24 |
		uint64(v[4])<<32 | uint64(v[5])<<40 | uint64(v[6])<<48 | uint64(v[7])<<56
}

func WriteLULong(v uint64) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
		byte(v >> 32),
		byte(v >> 40),
		byte(v >> 48),
		byte(v >> 56),
	}
}

func ReadFloat(v []byte) float32 { //TODO: umm... right method ?
	return math.Float32frombits(ReadUInt(v))
}

func WriteFloat(v float32) []byte {
	return WriteUInt(math.Float32bits(v))
}

func ReadLFloat(v []byte) float32 {
	return math.Float32frombits(ReadLUInt(v))
}

func WriteLFloat(v float32) []byte {
	return WriteLUInt(math.Float32bits(v))
}

func ReadDouble(v []byte) float64 {
	return math.Float64frombits(ReadULong(v))
}

func WriteDouble(v float64) []byte {
	return WriteULong(math.Float64bits(v))
}

func ReadLDouble(v []byte) float64 {
	return math.Float64frombits(ReadLULong(v))
}

func WriteLDouble(v float64) []byte {
	return WriteULong(math.Float64bits(v))
}

//buf

//Read reads data into b by order
func Read(reader io.Reader, order Order, data interface{}) error {
	size := dataSize(data)

	bytes := make([]byte, size)

	n, err := reader.Read(bytes)
	if err != nil {
		return err
	}

	if n < size {
		return NotEnoughErr
	}

	switch value := data.(type) {
	case *int8:
		*value = order.SByte(bytes)
	case *uint8:
		*value = order.Byte(bytes)
	case *int16:
		*value = order.Short(bytes)
	case *uint16:
		*value = order.UShort(bytes)
	case *Triad:
		*value = order.Triad(bytes)
	case *int32:
		*value = order.Int(bytes)
	case *uint32:
		*value = order.UInt(bytes)
	case *int64:
		*value = order.Long(bytes)
	case *uint64:
		*value = order.ULong(bytes)
	case *float32:
		*value = order.Float(bytes)
	case *float64:
		*value = order.Double(bytes)
	}

	return nil
}

//Write writes the contents of data into buffer by order
func Write(writer io.Writer, order Order, data interface{}) error {
	var value []byte
	switch v := data.(type) {
	case int8:
		value = order.PutSByte(v)
	case *int8:
		value = order.PutSByte(*v)
	case uint8:
		value = order.PutByte(v)
	case *uint8:
		value = order.PutByte(*v)
	case int16:
		value = order.PutShort(v)
	case *int16:
		value = order.PutShort(*v)
	case uint16:
		value = order.PutUShort(v)
	case *uint16:
		value = order.PutUShort(*v)
	case Triad:
		value = order.PutTriad(v)
	case *Triad:
		value = order.PutTriad(*v)
	case int32:
		value = order.PutInt(v)
	case *int32:
		value = order.PutInt(*v)
	case uint32:
		value = order.PutUInt(v)
	case *uint32:
		value = order.PutUInt(*v)
	case int64:
		value = order.PutLong(v)
	case *int64:
		value = order.PutLong(*v)
	case uint64:
		value = order.PutULong(v)
	case *uint64:
		value = order.PutULong(*v)
	}

	_, err := writer.Write(value)

	return err
}

//dataSize returns byte size of type
func dataSize(data interface{}) int {
	size := 0
	switch data.(type) {
	case int8, *int8, uint8, *uint8:
		size = ByteSize
	case int16, *int16, uint16, *uint16:
		size = ShortSize
	case Triad, *Triad:
		size = TriadSize
	case int32, *int32, uint32, *uint32:
		size = IntSize
	case int64, *int64, uint64, *uint64:
		size = LongSize
	case float32, *float32:
		size = FloatSize
	case float64, *float64:
		size = DoubleSize
	}
	return size
}

//Order for Raknet Protocol
type Order interface {
	Byte(v []byte) byte
	SByte(v []byte) int8
	Short(v []byte) int16
	UShort(v []byte) uint16
	Triad(v []byte) Triad
	Int(v []byte) int32
	UInt(v []byte) uint32
	Long(v []byte) int64
	ULong(v []byte) uint64
	Float(v []byte) float32
	Double(v []byte) float64
	PutByte(v byte) []byte
	PutSByte(v int8) []byte
	PutShort(v int16) []byte
	PutUShort(v uint16) []byte
	PutTriad(v Triad) []byte
	PutInt(v int32) []byte
	PutUInt(v uint32) []byte
	PutLong(v int64) []byte
	PutULong(v uint64) []byte
	PutFloat(v float32) []byte
	PutDouble(v float64) []byte
}

//BigEndian .
var BigEndian bigEndian

//LittleEndian .
var LittleEndian littleEndian

type bigEndian struct {
}

func (bigEndian) Byte(v []byte) byte {
	return ReadByte(v)
}

func (bigEndian) PutByte(v byte) []byte {
	return WriteByte(v)
}

func (bigEndian) SByte(v []byte) int8 {
	return ReadSByte(v)
}

func (bigEndian) PutSByte(v int8) []byte {
	return WriteSByte(v)
}

func (bigEndian) Short(v []byte) int16 {
	return ReadShort(v)
}

func (bigEndian) PutShort(v int16) []byte {
	return WriteShort(v)
}

func (bigEndian) UShort(v []byte) uint16 {
	return ReadUShort(v)
}

func (bigEndian) PutUShort(v uint16) []byte {
	return WriteUShort(v)
}

func (bigEndian) Triad(v []byte) Triad {
	return ReadTriad(v)
}

func (bigEndian) PutTriad(v Triad) []byte {
	return WriteTriad(v)
}

func (bigEndian) Int(v []byte) int32 {
	return ReadInt(v)
}

func (bigEndian) PutInt(v int32) []byte {
	return WriteInt(v)
}

func (bigEndian) UInt(v []byte) uint32 {
	return ReadUInt(v)
}

func (bigEndian) PutUInt(v uint32) []byte {
	return WriteUInt(v)
}

func (bigEndian) Long(v []byte) int64 {
	return ReadLong(v)
}

func (bigEndian) PutLong(v int64) []byte {
	return WriteLong(v)
}

func (bigEndian) ULong(v []byte) uint64 {
	return ReadULong(v)
}

func (bigEndian) PutULong(v uint64) []byte {
	return WriteULong(v)
}

func (bigEndian) Float(v []byte) float32 {
	return ReadFloat(v)
}

func (bigEndian) PutFloat(v float32) []byte {
	return WriteFloat(v)
}

func (bigEndian) Double(v []byte) float64 {
	return ReadDouble(v)
}

func (bigEndian) PutDouble(v float64) []byte {
	return WriteDouble(v)
}

type littleEndian struct {
}

func (littleEndian) Byte(v []byte) byte {
	return ReadByte(v)
}

func (littleEndian) PutByte(v byte) []byte {
	return WriteByte(v)
}

func (littleEndian) SByte(v []byte) int8 {
	return ReadSByte(v)
}

func (littleEndian) PutSByte(v int8) []byte {
	return WriteSByte(v)
}

func (littleEndian) Short(v []byte) int16 {
	return ReadLShort(v)
}

func (littleEndian) PutShort(v int16) []byte {
	return WriteLShort(v)
}

func (littleEndian) UShort(v []byte) uint16 {
	return ReadLUShort(v)
}

func (littleEndian) PutUShort(v uint16) []byte {
	return WriteLUShort(v)
}

func (littleEndian) Triad(v []byte) Triad {
	return ReadLTriad(v)
}

func (littleEndian) PutTriad(v Triad) []byte {
	return WriteLTriad(v)
}

func (littleEndian) Int(v []byte) int32 {
	return ReadLInt(v)
}

func (littleEndian) PutInt(v int32) []byte {
	return WriteLInt(v)
}

func (littleEndian) UInt(v []byte) uint32 {
	return ReadLUInt(v)
}

func (littleEndian) PutUInt(v uint32) []byte {
	return WriteLUInt(v)
}

func (littleEndian) Long(v []byte) int64 {
	return ReadLLong(v)
}

func (littleEndian) PutLong(v int64) []byte {
	return WriteLLong(v)
}

func (littleEndian) ULong(v []byte) uint64 {
	return ReadLULong(v)
}

func (littleEndian) PutULong(v uint64) []byte {
	return WriteLULong(v)
}

func (littleEndian) Float(v []byte) float32 {
	return ReadLFloat(v)
}

func (littleEndian) PutFloat(v float32) []byte {
	return WriteLFloat(v)
}

func (littleEndian) Double(v []byte) float64 {
	return ReadLDouble(v)
}

func (littleEndian) PutDouble(v float64) []byte {
	return WriteLDouble(v)
}
