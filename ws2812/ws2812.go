package ws2812

const (
	//0b0100
	ZERO = 0x04
	//0b0110
	ONE = 0x06
)

func EncodeByte(b byte) []byte {
	//一个int 4字节32位 用低24位做缓冲区
	tmp := 0
	for i := 0; i < 8; i++ {
		if (b>>i)&0x01 == 1 {
			//缓冲区写入1
			tmp |= ONE << (i * 3)
		} else {
			//缓冲区写入0
			tmp |= ZERO << (i * 3)
		}
	}
	return []byte{
		byte((tmp >> 16) & 0xff),
		byte((tmp >> 8) & 0xff),
		byte(tmp & 0xff),
	}
}

func EncodeRGB(r, g, b, a byte) []byte {
	var tmp []byte
	tmp = append(tmp, EncodeByte(byte(int(g)*int(a)/0xff))...)
	tmp = append(tmp, EncodeByte(byte(int(r)*int(a)/0xff))...)
	tmp = append(tmp, EncodeByte(byte(int(b)*int(a)/0xff))...)
	return tmp
}

func ResetBytes() []byte {
	return []byte{0x00}
}
