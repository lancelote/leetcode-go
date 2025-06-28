package p443

import (
	"strconv"
)

type Compressor struct {
	w    int
	Data []byte
}

func (c *Compressor) writeRune(r rune) {
	c.Data[c.w] = byte(r)
	c.w++
}

func (c *Compressor) writeByte(b byte) {
	c.Data[c.w] = b
	c.w++
}

func (c *Compressor) writeNum(n int) {
	if n == 1 {
		return
	}

	for _, r := range strconv.Itoa(n) {
		c.writeRune(r)
	}
}

func (c *Compressor) write(b byte, n int) {
	c.writeByte(b)
	c.writeNum(n)
}

func compress(chars []byte) int {
	c := Compressor{Data: chars}

	lastByte := chars[0]
	count := 1

	for i := 1; i < len(chars); i++ {
		currByte := chars[i]
		if currByte == lastByte {
			count++
		} else {
			c.write(lastByte, count)
			lastByte = currByte
			count = 1
		}
	}

	c.write(lastByte, count)
	return c.w
}
