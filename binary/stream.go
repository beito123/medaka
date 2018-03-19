package binary

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

import (
	"bytes"
	"encoding/hex"
	"strconv"
	"strings"
)

//BinaryStream is basic binary stream.
type BinaryStream struct {
	Buffer *bytes.Buffer
}

//Reset resets Buffer
func (bs *BinaryStream) Reset() error {
	bs.Buffer.Reset()
	return nil
}

//Get returns n bytes from Buffer with []byte
func (bs *BinaryStream) Get(n int) []byte {
	return bs.Buffer.Next(n)
}

//Put puts value to buffer
func (bs *BinaryStream) Put(value []byte) error {
	return Write(bs.Buffer, BigEndian, value)
}

//Bytes returns the bytes left from Buffer.
func (bs *BinaryStream) Bytes() []byte {
	return bs.Buffer.Bytes()
}

//Len returns len the bytes left
func (bs *BinaryStream) Len() int {
	return bs.Buffer.Len()
}

//Skip skips n bytes on buffer
func (bs *BinaryStream) Skip(n int) {
	_ = bs.Buffer.Next(n)
}

/*
 * Data types
 * | name  | size | encode |                   range                   |
 *  Byte    1byte   Big                        0 - 255
 *  SByte   1byte   Big                     -128 - 127
 *  Short   2bytes  Big                        0 - 65535
 *  SShort  2bytes  Big                   -32768 - 32767
 *  LShort  2bytes  Little                     0 - 65535
 *  LSShort 2bytes  Little                -32768 - 32767
 *  Triad   3bytes  Little                     0 - 16777215
 *  Int     4bytes  Big              -2147483648 - 2147483647
 *  Long    8bytes  Big     -9223372036854775808 - 9223372036854775807
 *  String  ?bytes  Big                        ? - ?
 */

/*
 * Byte
 * SignedByte
 * Short
 * SignedShort
 * LShort
 * SignedLShort
 * Triad
 * LTriad
 * Int
 * Float
 * LFloat
 * Double
 * LDouble
 * Long
 * /////////////
 * Bool
 * String
 * HexString
 * Address
 * NBT
 * Item
 * UUID
 * Position
 * BlockPosition
 * EntityMetadata
 */

//Byte sets byte(unsign) got from buffer to value
func (bs *BinaryStream) Byte(value *byte) error {
	return Read(bs.Buffer, BigEndian, value)
}

//SByte sets byte(sign) got from buffer to value
func (bs *BinaryStream) SByte(value *int8) error {
	return Read(bs.Buffer, BigEndian, value)
}

//PutByte puts byte(unsign) from value to buffer
func (bs *BinaryStream) PutByte(value byte) error {
	return Write(bs.Buffer, BigEndian, value)
}

//PutSByte puts byte(sign) from value to buffer
func (bs *BinaryStream) PutSByte(value int8) error {
	return Write(bs.Buffer, BigEndian, value)
}

//Short sets short(unsign) got from buffer to value
func (bs *BinaryStream) Short(value *uint16) error {
	return Read(bs.Buffer, BigEndian, value)
}

//SShort sets short(sign) got from buffer to value
func (bs *BinaryStream) SShort(value *int16) error {
	return Read(bs.Buffer, BigEndian, value)
}

//LShort sets short(unsign) got from buffer as LittleEndian to value
func (bs *BinaryStream) LShort(value *uint16) error {
	return Read(bs.Buffer, LittleEndian, value)
}

//LSShort sets short(sign) got from buffer as LittleEndian to value
func (bs *BinaryStream) LSShort(value *int16) error {
	return Read(bs.Buffer, LittleEndian, value)
}

//PutShort puts short(unsign) from value to buffer
func (bs *BinaryStream) PutShort(value uint16) error {
	return Write(bs.Buffer, BigEndian, value)
}

//PutSShort puts short(sign) from value to buffer
func (bs *BinaryStream) PutSShort(value int16) error {
	return Write(bs.Buffer, BigEndian, value)
}

//PutLShort puts short(unsign) from value to buffer as LittleEndian
func (bs *BinaryStream) PutLShort(value uint16) error {
	return Write(bs.Buffer, LittleEndian, value)
}

//PutLSShort puts short(sign) from value to buffer as LittleEndian
func (bs *BinaryStream) PutLSShort(value int16) error {
	return Write(bs.Buffer, LittleEndian, value)
}

//Triad sets triad got from buffer to value
func (bs *BinaryStream) Triad(value *Triad) error {
	return Read(bs.Buffer, BigEndian, value)
}

//PutTriad puts triad from value to buffer
func (bs *BinaryStream) PutTriad(value Triad) error {
	return Write(bs.Buffer, BigEndian, value)
}

//LTriad sets triad got from buffer as LittleEndian to value
func (bs *BinaryStream) LTriad(value *Triad) error {
	return Read(bs.Buffer, LittleEndian, value)
}

//PutLTriad puts triad from value to buffer as LittleEndian
func (bs *BinaryStream) PutLTriad(value Triad) error {
	return Write(bs.Buffer, LittleEndian, value)
}

//Int sets int got from buffer to value
func (bs *BinaryStream) Int(value *int32) error {
	return Read(bs.Buffer, BigEndian, value)
}

//PutInt puts int from value to buffer
func (bs *BinaryStream) PutInt(value int32) error {
	return Write(bs.Buffer, BigEndian, value)
}

//Long sets long got from buffer to value
func (bs *BinaryStream) Long(value *int64) error {
	return Read(bs.Buffer, BigEndian, value)
}

//PutLong puts long from value to buffer
func (bs *BinaryStream) PutLong(value int64) error {
	return Write(bs.Buffer, BigEndian, value)
}

//String sets string(len short, str string) got from buffer to value
func (bs *BinaryStream) String(value *string) error {
	var n uint16
	err := bs.Short(&n)
	if err != nil {
		return err
	}

	*value = string(bs.Get(int(n)))
	return nil
}

//PutString puts string(len short, str string) to Buffer
func (bs *BinaryStream) PutString(value string) error {
	n := uint16(len(value))
	err := bs.PutShort(n)
	if err != nil {
		return err
	}
	return bs.Put([]byte(value))
}

//HexString gets hex string from Buffer (for Magic)
func (bs *BinaryStream) HexString(n int, value *string) {
	*value = hex.EncodeToString(bs.Buffer.Bytes())
}

//PutHexString puts hex string to Buffer
func (bs *BinaryStream) PutHexString(value string) error {
	bytes, err := hex.DecodeString(value)
	if err != nil {
		return err
	}
	return bs.Put(bytes)
}

//Address sets address got from Buffer to addr and port
//address(version byte, address byte x4, port ushort)
func (bs *BinaryStream) Address(addr *string, port *uint16) error {
	var version byte
	err := bs.Byte(&version)
	if err != nil {
		return err
	}

	var address string

	if version == 4 {
		var bytes byte
		for i := 0; i < 4; i++ {
			err = bs.Byte(&bytes)
			if err != nil {
				return err
			}

			address = address + strconv.Itoa(int(^bytes&0xff))
			if i < 3 {
				address = address + "."
			}
		}
		addr = &address

		err = bs.Short(port)
		if err != nil {
			return err
		}
	} else {
		//IPv6
	}

	return nil
}

//PutAddress puts address to Buffer
//address(version byte, address byte x4, port ushort)
func (bs *BinaryStream) PutAddress(addr string, port, version uint16) error {
	err := bs.PutByte(byte(version))
	if err != nil {
		return err
	}

	if version == 4 {
		for _, str := range strings.Split(addr, ".") {
			i, _ := strconv.Atoi(str)
			err = bs.PutByte(^byte(i) & 0xff)
			if err != nil {
				return err
			}
		}
		err = bs.PutShort(port)
		if err != nil {
			return err
		}
	} else {
		//ipv6
	}

	return nil
}
