package ga

import "crypto/rand"

type BitArray struct {
	bits []uint32
	size int32
}

func NewBitArray(size int) *BitArray {
	a := new(BitArray)
	a.size = int32(size)
	a.bits = make([]uint32, (size+31)/32)
	return a
}

func (a *BitArray) Size() int {
	return int(a.size)
}

func (a *BitArray) SizeInBytes() int32 {
	return (a.size + 7) >> 3
}

func (a *BitArray) EnsureCapacity(size int) {
	if size > len(a.bits)<<5 {
		oldBits := a.bits
		a.bits = make([]uint32, (size+31)/32)
		copy(a.bits, oldBits)
	}
}

func (a *BitArray) Get(i int) bool {
	return (a.bits[i>>5] & (1 << uint(i&0x1f))) != 0
}

func (a *BitArray) Set(i int) {
	a.bits[i>>5] |= (1 << uint(i&0x1f))
}

func (a *BitArray) Flip(i int) {
	a.bits[i>>5] ^= (1 << uint(i&0x1f))
}

// func (a *BitArray) NextSetFrom(from int32) {

// }

// func (a *BitArray) NextUnsetFrom(from int32) {

// }

func (a *BitArray) SetBulk(bulkIndex int, bulk int) {
	a.bits[bulkIndex] = uint32(bulk)
}

// func (a *BitArray) SetRange(start, end int32) { }

func (a *BitArray) Clear() {
	bulkNum := len(a.bits)
	for i := 0; i < bulkNum; i++ {
		a.bits[i] = 0
	}
}

// func (a *BitArray) IsRange() { }
// func (a *BitArray) AppendBit() { }
// func (a *BitArray) AppendBits() { }
// func (a *BitArray) AppendBitArray() { }
// func (a *BitArray) Xor() { }
// func (a *BitArray) And() { }
// func (a *BitArray) Or() { }

func (a *BitArray) Copy() *BitArray {
	dup := NewBitArray(int(a.size))
	copy(dup.bits, a.bits)
	return dup
}

func (a *BitArray) CrossAt(crossPos int, other *BitArray) (res *BitArray) {
	res = a.Copy()

	crossPosInt := crossPos >> 5

	for i := 0; i < crossPosInt; i++ {
		res.bits[i] = other.bits[i]
	}

	crossPosIntRemainBits := crossPos & 0x1f
	if crossPosIntRemainBits <= 0 {
		return res
	}

	x := res.bits[crossPosInt]
	y := other.bits[crossPosInt]

	mask := bitMask(crossPosIntRemainBits)

	res.bits[crossPosInt] = ((y & mask) | (x & ^mask))

	return res
}

// func (a *BitArray) ToBytes(){}
func (a *BitArray) GetInt(offset int, bitsNum int) int {
	if bitsNum < 0 {
		bitsNum = int(a.size) - offset
	}

	if offset == 0 && bitsNum == int(a.size) {
		return int(a.bits[0] & bitMask(bitsNum))
	}

	offsetInt := offset >> 5
	offsetBitPos := offset & 0x1f
	if offsetBitPos == 0 {
		return int(a.bits[offsetInt] & bitMask(bitsNum))
	}

	res := a.bits[offsetInt] >> uint(offsetBitPos)
	if bitsNum <= 32-offsetBitPos || offsetInt+1 >= len(a.bits) {
		return int(res & bitMask(bitsNum))
	}

	res |= (((a.bits[offsetInt+1]) & bitMask(min(bitsNum-32, 0)+offsetBitPos)) << uint(32-offsetBitPos))
	return int(res & bitMask(bitsNum))
}

// func (a *BitArray) GetBulks() { }
// func (a *BitArray) Reverse() { }
func (a *BitArray) String() string {
	res := make([]byte, (a.size/8+1)*9)

	j := 0
	for i := 0; i < int(a.size); i++ {
		if (i & 0x07) == 0 {
			res[j] = ' '
			j++
		}

		if a.Get(i) {
			res[j] = 'X'
		} else {
			res[j] = '.'
		}
		j++
	}

	return string(res[:j])
}

func (a *BitArray) FillRandBits() {
	bulkNum := len(a.bits)
	buf := make([]byte, bulkNum*4)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}

	for i := 0; i < bulkNum; i++ {
		a.bits[i] = (uint32(buf[i*4]) << 24) | (uint32(buf[i*4+1]) << 16) | (uint32(buf[i*4+2]) << 8) | (uint32(buf[i*4+3]))
	}
}
