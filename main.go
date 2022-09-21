package main

import (
	"time"

	"github.com/gogogoghost/ws2812/ws2812d"
	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

func main() {
	_, err := host.Init()
	if err != nil {
		panic(err)
	}
	_, err = driverreg.Init()
	if err != nil {
		panic(err)
	}
	s, err := spireg.Open("/dev/spidev1.0")
	if err != nil {
		panic(err)
	}
	defer s.Close()

	conn, err := s.Connect(physic.KiloHertz*2500, spi.Mode3, 8)
	if err != nil {
		panic(err)
	}

	startColor := 0
	for {
		buf := ws2812d.ResetBytes()
		color := startColor
		for i := 0; i < 20; i++ {
			switch color {
			case 0:
				buf = append(buf, ws2812d.EncodeRGB(0xff, 0x00, 0x00, 0x44)...)
			case 1:
				buf = append(buf, ws2812d.EncodeRGB(0x00, 0xff, 0x00, 0x44)...)
			case 2:
				buf = append(buf, ws2812d.EncodeRGB(0x00, 0x00, 0xff, 0x44)...)
			}
			color++
			if color == 3 {
				color = 0
			}
		}

		err = conn.Tx(buf, nil)
		if err != nil {
			panic(err)
		}
		startColor++
		if startColor == 3 {
			startColor = 0
		}
		time.Sleep(time.Duration(1) * time.Second)
	}

}
