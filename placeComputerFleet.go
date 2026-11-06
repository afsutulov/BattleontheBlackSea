package main

// placeComputerFleet preserves the original Pascal control flow.
func (m *gameState) placeComputerFleet() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (30) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu(opSub, m.sp, 30, 16)) & 65535
	m.w16(m.ds, (4474)&65535, 20)
	m.w16(m.ss, (m.bp+-16)&65535, 11)
	goto L3a72
L3a6f:
	m.clock(16)
	m.w16(m.ss, (m.bp+-16)&65535, m.step(m.r16(m.ss, (m.bp+-16)&65535), 1, 16))
L3a72:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+5351)&65535)) & 255)
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.w8(m.ds, (m.di+5353)&65535, (m.ax & 255))
L3a88:
	m.clock(16)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.w16(m.ss, (m.bp+-10)&65535, m.ax)
	m.clock(16)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.w16(m.ss, (m.bp+-12)&65535, m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.w16(m.ss, (m.bp+-14)&65535, m.ax)
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+5351)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu(opXor, (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.w16(m.ss, (m.bp+-24)&65535, m.ax)
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-26)&65535, m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-28)&65535, m.ax)
	m.alu(opSub, m.r16(m.ss, (m.bp+-14)&65535), 0, 16)
	if m.z || m.s != m.o {
		goto L3ad5
	}
	m.w16(m.ss, (m.bp+-28)&65535, 1)
	goto L3ada
L3ad5:
	m.clock(16)
	m.w16(m.ss, (m.bp+-26)&65535, 1)
L3ada:
	m.clock(16)
	m.w16(m.ss, (m.bp+-2)&65535, 1)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-20)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.w16(m.ss, (m.bp+-30)&65535, m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if !m.z && m.s == m.o {
		goto L3b48
	}
	m.w16(m.ss, (m.bp+-18)&65535, m.ax)
	goto L3b01
L3afe:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, m.step(m.r16(m.ss, (m.bp+-18)&65535), 1, 16))
L3b01:
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.bp+-20)&65535), 11, 16)
	if m.s == m.o {
		goto L3b29
	}
	m.alu(opSub, m.r16(m.ss, (m.bp+-22)&65535), 11, 16)
	if m.s == m.o {
		goto L3b29
	}
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-20)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.alu(opSub, m.r16(m.ds, (m.di+4782)&65535), 0, 16)
	if !m.z {
		goto L3b29
	}
	goto L3b2e
L3b29:
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-2)&65535, m.ax)
L3b2e:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-20)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.r16(m.ss, (m.bp+-26)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-20)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.r16(m.ss, (m.bp+-28)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-18)&65535)) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if !m.z {
		goto L3afe
	}
L3b48:
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.bp+-2)&65535), 1, 16)
	if m.z {
		goto L3b51
	}
	goto L3a88
L3b51:
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.bp+-14)&65535), 0, 16)
	if !m.z && m.s == m.o {
		goto L3b5a
	}
	goto L3c01
L3b5a:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.w16(m.ss, (m.bp+-30)&65535, m.ax)
	m.ax = (65535) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if m.z || m.s != m.o {
		goto L3b6b
	}
	m.clock(16)
	goto L3bfe
L3b6b:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, m.ax)
	goto L3b73
L3b70:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, m.step(m.r16(m.ss, (m.bp+-18)&65535), 1, 16))
L3b73:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.r16(m.ss, (m.bp+-18)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.clock(16)
	m.alu(opSub, m.ax, 11, 16)
	if m.s == m.o {
		goto L3ba5
	}
	m.alu(opSub, m.r16(m.ss, (m.bp+-22)&65535), 11, 16)
	if m.s == m.o {
		goto L3ba5
	}
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4782)&65535, 65535)
L3ba5:
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.bp+-10)&65535), 11, 16)
	if m.s == m.o {
		goto L3bca
	}
	m.alu(opSub, m.r16(m.ss, (m.bp+-22)&65535), 11, 16)
	if m.s == m.o {
		goto L3bca
	}
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4782)&65535, 65535)
L3bca:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.alu(opSub, m.ax, 11, 16)
	if m.s == m.o {
		goto L3bf3
	}
	m.alu(opSub, m.r16(m.ss, (m.bp+-22)&65535), 11, 16)
	if m.s == m.o {
		goto L3bf3
	}
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4782)&65535, 65535)
L3bf3:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-18)&65535)) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if m.z {
		goto L3bfe
	}
	goto L3b70
L3bfe:
	m.clock(16)
	goto L3ca5
L3c01:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.w16(m.ss, (m.bp+-30)&65535, m.ax)
	m.ax = (65535) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if m.z || m.s != m.o {
		goto L3c12
	}
	m.clock(16)
	goto L3ca5
L3c12:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, m.ax)
	goto L3c1a
L3c17:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, m.step(m.r16(m.ss, (m.bp+-18)&65535), 1, 16))
L3c1a:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.r16(m.ss, (m.bp+-18)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-20)&65535, m.ax)
	m.alu(opSub, m.r16(m.ss, (m.bp+-20)&65535), 11, 16)
	if m.s == m.o {
		goto L3c4c
	}
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.alu(opSub, m.ax, 11, 16)
	if m.s == m.o {
		goto L3c4c
	}
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-20)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4782)&65535, 65535)
L3c4c:
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.bp+-20)&65535), 11, 16)
	if m.s == m.o {
		goto L3c71
	}
	m.alu(opSub, m.r16(m.ss, (m.bp+-12)&65535), 11, 16)
	if m.s == m.o {
		goto L3c71
	}
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-20)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4782)&65535, 65535)
L3c71:
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.bp+-20)&65535), 11, 16)
	if m.s == m.o {
		goto L3c9a
	}
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.alu(opSub, m.ax, 11, 16)
	if m.s == m.o {
		goto L3c9a
	}
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-20)&65535)) & 65535
	m.dx = (24) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4782)&65535, 65535)
L3c9a:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-18)&65535)) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if m.z {
		goto L3ca5
	}
	goto L3c17
L3ca5:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.w16(m.ss, (m.bp+-20)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.w16(m.ss, (m.bp+-30)&65535, m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.clock(16)
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if !m.z && m.s == m.o {
		goto L3cfb
	}
	m.w16(m.ss, (m.bp+-18)&65535, m.ax)
	goto L3cc7
L3cc4:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, m.step(m.r16(m.ss, (m.bp+-18)&65535), 1, 16))
L3cc7:
	m.clock(16)
	m.bx = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-20)&65535)) & 65535
	m.dx = (24) & 65535
	m.clock(16)
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4782)&65535, m.bx)
	m.ax = (m.r16(m.ss, (m.bp+-20)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.r16(m.ss, (m.bp+-26)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-20)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.alu(opAdd, m.ax, m.r16(m.ss, (m.bp+-28)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-18)&65535)) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if !m.z {
		goto L3cc4
	}
L3cfb:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-14)&65535)) & 255)
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.w8(m.ds, (m.di+5352)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.w16(m.ds, (m.di+5354)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.w16(m.ds, (m.di+5356)&65535, m.ax)
	m.alu(opSub, m.r16(m.ss, (m.bp+-16)&65535), 20, 16)
	m.clock(16)
	if m.z {
		goto L3d2e
	}
	goto L3a6f
L3d2e:
	m.clock(16)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 4) & 65535
	return
}
