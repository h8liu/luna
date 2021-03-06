package eg

import (
	"bytes"
	"encoding/binary"
	"fmt"

	. "github.com/h8liu/luna/finger"
)

func Img() (code, data []byte) {
	cbuf := new(bytes.Buffer)
	b4 := make([]byte, 4)

	for _, b := range []uint32{
		// start:
		Movi(R5, 9),      // r5 = 9
		Sllv(R5, R5, 12), // r5 = r5 << 12 // data page offset
		Ld(R0, R5, 0),    // r0 = [r5]

		// Ld(R0, PC, 68),   // r0 = [pc + 68]
		Movi(R1, 1),      // r1 = 1
		Sllv(R1, R1, 18), // r1 = r1 << 18
		St(R1, R0, 4),    // [r0 + 4] = r1
		Movi(R1, 1),      // r1 = 1
		Sllv(R1, R1, 16), // r1 = r1 << 16
		Movi(R4, 0x3f),   // r4 = 0x3f

		// loop:
		St(R1, R0, 40),   // [r0 + 40] = r1
		Sllv(R2, R4, 16), // r2 = r4 << 16

		// wait1:
		Subi(R2, R2, 1), // r2 = r2 - 1
		Cmpi(R2, 0),     // if r2 != r3:
		Bne(-4),         // goto wait1

		St(R1, R0, 28),   // [r0 + 28] = r1
		Sllv(R2, R4, 16), // r2 = r4 << 16

		// wait2:
		Subi(R2, R2, 1), // r2 = r2 - 1
		Cmpi(R2, 0),     // if r2 != r3:
		Bne(-4),         // goto wait2

		J(-12), // goto loop

		0x20200000, // add the data
	} {
		binary.LittleEndian.PutUint32(b4, b)
		cbuf.Write(b4)
		fmt.Printf("d(0x%08x)\n", b)
	}

	dbuf := new(bytes.Buffer)
	binary.LittleEndian.PutUint32(b4, 0x20200000)
	dbuf.Write(b4)

	return cbuf.Bytes(), dbuf.Bytes()
}
