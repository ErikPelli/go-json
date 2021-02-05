package json

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"unsafe"
)

type structFieldSet struct {
	dec         decoder
	offset      uintptr
	isTaggedKey bool
}

type structDecoder struct {
	fieldMap        map[string]*structFieldSet
	stringDecoder   *stringDecoder
	structName      string
	fieldName       string
	isTriedOptimize bool
	keyBitmapInt8   [][256]int8
	keyBitmapInt16  [][256]int16
	sortedFieldSets []*structFieldSet
	keyDecoder      func(*structDecoder, *sliceHeader, int64) (int64, *structFieldSet, error)
}

var (
	bitHashTable      [64]int
	largeToSmallTable [256]byte
)

func init() {
	hash := uint64(0x03F566ED27179461)
	for i := 0; i < 64; i++ {
		bitHashTable[hash>>58] = i
		hash <<= 1
	}
	for i := 0; i < 256; i++ {
		c := i
		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}
		largeToSmallTable[i] = byte(c)
	}
}

func newStructDecoder(structName, fieldName string, fieldMap map[string]*structFieldSet) *structDecoder {
	return &structDecoder{
		fieldMap:      fieldMap,
		stringDecoder: newStringDecoder(structName, fieldName),
		structName:    structName,
		fieldName:     fieldName,
		keyDecoder:    decodeKey,
	}
}

const (
	allowOptimizeMaxKeyLen   = 64
	allowOptimizeMaxFieldLen = 16
)

func (d *structDecoder) tryOptimize() {
	if d.isTriedOptimize {
		return
	}
	fieldMap := map[string]*structFieldSet{}
	for k, v := range d.fieldMap {
		k := strings.ToLower(k)
		fieldMap[k] = v
	}

	if len(fieldMap) > allowOptimizeMaxFieldLen {
		d.isTriedOptimize = true
		return
	}

	var maxKeyLen int
	sortedKeys := []string{}
	for key := range fieldMap {
		keyLen := len(key)
		if keyLen > allowOptimizeMaxKeyLen {
			d.isTriedOptimize = true
			return
		}
		if maxKeyLen < keyLen {
			maxKeyLen = keyLen
		}
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)
	if len(sortedKeys) <= 8 {
		keyBitmap := make([][256]int8, maxKeyLen)
		for i, key := range sortedKeys {
			for j := 0; j < len(key); j++ {
				c := key[j]
				keyBitmap[j][c] |= (1 << uint(i))
			}
			d.sortedFieldSets = append(d.sortedFieldSets, fieldMap[key])
		}
		d.keyBitmapInt8 = keyBitmap
		d.keyDecoder = decodeKeyByBitmapInt8
	} else {
		keyBitmap := make([][256]int16, maxKeyLen)
		for i, key := range sortedKeys {
			for j := 0; j < len(key); j++ {
				c := key[j]
				keyBitmap[j][c] |= (1 << uint(i))
			}
			d.sortedFieldSets = append(d.sortedFieldSets, fieldMap[key])
		}
		d.keyBitmapInt16 = keyBitmap
		d.keyDecoder = decodeKeyByBitmapInt16
	}
}

func decodeKeyByBitmapInt8(d *structDecoder, buf *sliceHeader, cursor int64) (int64, *structFieldSet, error) {
	var (
		field  *structFieldSet
		curBit int8 = math.MaxInt8
	)
	for {
		switch char(buf.data, cursor) {
		case ' ', '\n', '\t', '\r':
			cursor++
		case '"':
			cursor++
			c := char(buf.data, cursor)
			switch c {
			case '"':
				cursor++
				return cursor, field, nil
			case nul:
				return 0, nil, errUnexpectedEndOfJSON("string", cursor)
			}
			keyIdx := 0
			bitmap := d.keyBitmapInt8
			keyBitmapLen := len(bitmap)
			for {
				c := char(buf.data, cursor)
				switch c {
				case '"':
					x := uint64(curBit & -curBit)
					fieldSetIndex := bitHashTable[(x*0x03F566ED27179461)>>58]
					field = d.sortedFieldSets[fieldSetIndex]
					cursor++
					return cursor, field, nil
				case nul:
					return 0, nil, errUnexpectedEndOfJSON("string", cursor)
				default:
					if keyIdx >= keyBitmapLen {
						for {
							cursor++
							switch char(buf.data, cursor) {
							case '"':
								cursor++
								return cursor, field, nil
							case '\\':
								cursor++
								if char(buf.data, cursor) == nul {
									return 0, nil, errUnexpectedEndOfJSON("string", cursor)
								}
							case nul:
								return 0, nil, errUnexpectedEndOfJSON("string", cursor)
							}
						}
					}
					curBit &= bitmap[keyIdx][largeToSmallTable[c]]
					if curBit == 0 {
						for {
							cursor++
							switch char(buf.data, cursor) {
							case '"':
								cursor++
								return cursor, field, nil
							case '\\':
								cursor++
								if char(buf.data, cursor) == nul {
									return 0, nil, errUnexpectedEndOfJSON("string", cursor)
								}
							case nul:
								return 0, nil, errUnexpectedEndOfJSON("string", cursor)
							}
						}
					}
					keyIdx++
				}
				cursor++
			}
		default:
			return cursor, nil, errNotAtBeginningOfValue(cursor)
		}
	}
}

func char(ptr unsafe.Pointer, offset int64) byte {
	return *(*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(offset)))
}

func decodeKeyByBitmapInt16(d *structDecoder, buf *sliceHeader, cursor int64) (int64, *structFieldSet, error) {
	var (
		field  *structFieldSet
		curBit int16 = math.MaxInt16
	)
	for {
		switch char(buf.data, cursor) {
		case ' ', '\n', '\t', '\r':
			cursor++
		case '"':
			cursor++
			c := char(buf.data, cursor)
			switch c {
			case '"':
				cursor++
				return cursor, field, nil
			case nul:
				return 0, nil, errUnexpectedEndOfJSON("string", cursor)
			}
			keyIdx := 0
			bitmap := d.keyBitmapInt16
			keyBitmapLen := len(bitmap)
			for {
				c := char(buf.data, cursor)
				switch c {
				case '"':
					x := uint64(curBit & -curBit)
					fieldSetIndex := bitHashTable[(x*0x03F566ED27179461)>>58]
					field = d.sortedFieldSets[fieldSetIndex]
					cursor++
					return cursor, field, nil
				case nul:
					return 0, nil, errUnexpectedEndOfJSON("string", cursor)
				default:
					if keyIdx >= keyBitmapLen {
						for {
							cursor++
							switch char(buf.data, cursor) {
							case '"':
								cursor++
								return cursor, field, nil
							case '\\':
								cursor++
								if char(buf.data, cursor) == nul {
									return 0, nil, errUnexpectedEndOfJSON("string", cursor)
								}
							case nul:
								return 0, nil, errUnexpectedEndOfJSON("string", cursor)
							}
						}
					}
					curBit &= bitmap[keyIdx][largeToSmallTable[c]]
					if curBit == 0 {
						for {
							cursor++
							switch char(buf.data, cursor) {
							case '"':
								cursor++
								return cursor, field, nil
							case '\\':
								cursor++
								if char(buf.data, cursor) == nul {
									return 0, nil, errUnexpectedEndOfJSON("string", cursor)
								}
							case nul:
								return 0, nil, errUnexpectedEndOfJSON("string", cursor)
							}
						}
					}
					keyIdx++
				}
				cursor++
			}
		default:
			return cursor, nil, errNotAtBeginningOfValue(cursor)
		}
	}
}

func decodeKey(d *structDecoder, buf *sliceHeader, cursor int64) (int64, *structFieldSet, error) {
	key, c, err := d.stringDecoder.decodeByte(buf, cursor)
	if err != nil {
		return 0, nil, err
	}
	cursor = c
	k := *(*string)(unsafe.Pointer(&key))
	field, exists := d.fieldMap[k]
	if !exists {
		return cursor, nil, nil
	}
	return cursor, field, nil
}

func (d *structDecoder) decodeStream(s *stream, p unsafe.Pointer) error {
	s.skipWhiteSpace()
	switch s.char() {
	case 'n':
		if err := nullBytes(s); err != nil {
			return err
		}
		return nil
	case nul:
		s.read()
	default:
		if s.char() != '{' {
			return errNotAtBeginningOfValue(s.totalOffset())
		}
	}
	s.cursor++
	if s.char() == '}' {
		s.cursor++
		return nil
	}
	for {
		s.reset()
		key, err := d.stringDecoder.decodeStreamByte(s)
		if err != nil {
			return err
		}
		s.skipWhiteSpace()
		if s.char() == nul {
			s.read()
		}
		if s.char() != ':' {
			return errExpected("colon after object key", s.totalOffset())
		}
		s.cursor++
		if s.char() == nul {
			if !s.read() {
				return errExpected("object value after colon", s.totalOffset())
			}
		}
		k := *(*string)(unsafe.Pointer(&key))
		field, exists := d.fieldMap[k]
		if exists {
			if err := field.dec.decodeStream(s, unsafe.Pointer(uintptr(p)+field.offset)); err != nil {
				return err
			}
		} else if s.disallowUnknownFields {
			return fmt.Errorf("json: unknown field %q", k)
		} else {
			if err := s.skipValue(); err != nil {
				return err
			}
		}
		s.skipWhiteSpace()
		if s.char() == nul {
			s.read()
		}
		c := s.char()
		if c == '}' {
			s.cursor++
			return nil
		}
		if c != ',' {
			return errExpected("comma after object element", s.totalOffset())
		}
		s.cursor++
	}
}

func (d *structDecoder) decode(buf *sliceHeader, cursor int64, p unsafe.Pointer) (int64, error) {
	buflen := int64(buf.len)
	cursor = skipWhiteSpace(buf, cursor)
	switch char(buf.data, cursor) {
	case 'n':
		if cursor+3 >= buflen {
			return 0, errUnexpectedEndOfJSON("null", cursor)
		}
		if char(buf.data, cursor+1) != 'u' {
			return 0, errInvalidCharacter(char(buf.data, cursor+1), "null", cursor)
		}
		if char(buf.data, cursor+2) != 'l' {
			return 0, errInvalidCharacter(char(buf.data, cursor+2), "null", cursor)
		}
		if char(buf.data, cursor+3) != 'l' {
			return 0, errInvalidCharacter(char(buf.data, cursor+3), "null", cursor)
		}
		cursor += 4
		return cursor, nil
	case '{':
	default:
		return 0, errNotAtBeginningOfValue(cursor)
	}
	if buflen < 2 {
		return 0, errUnexpectedEndOfJSON("object", cursor)
	}
	cursor++
	for ; cursor < buflen; cursor++ {
		c, field, err := d.keyDecoder(d, buf, cursor)
		if err != nil {
			return 0, err
		}
		cursor = c
		cursor = skipWhiteSpace(buf, cursor)
		if char(buf.data, cursor) != ':' {
			return 0, errExpected("colon after object key", cursor)
		}
		cursor++
		if cursor >= buflen {
			return 0, errExpected("object value after colon", cursor)
		}
		if field != nil {
			c, err := field.dec.decode(buf, cursor, unsafe.Pointer(uintptr(p)+field.offset))
			if err != nil {
				return 0, err
			}
			cursor = c
		} else {
			c, err := skipValue(buf, cursor)
			if err != nil {
				return 0, err
			}
			cursor = c
		}
		cursor = skipWhiteSpace(buf, cursor)
		if char(buf.data, cursor) == '}' {
			cursor++
			return cursor, nil
		}
		if char(buf.data, cursor) != ',' {
			return 0, errExpected("comma after object element", cursor)
		}
	}
	return cursor, nil
}
