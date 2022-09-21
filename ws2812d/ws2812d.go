package ws2812d

import "github.com/gogogoghost/ws2812/ws2812"

func EncodeRGB(r, g, b, a byte) []byte {
	return ws2812.EncodeRGB(g, r, b, a)
}

func ResetBytes() []byte {
	return ws2812.ResetBytes()
}
