package main

// inputName preserves the original Pascal control flow.
func (m *gameState) inputName() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (778) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu(opSub, m.sp, 778, 16)) & 65535
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((2) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((8) & 255)
	m.clock(16)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((255) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (960) & 65535
	m.push(m.ax)
	m.ax = (m.ax & 65280) | ((7) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1550)
	m.clock(16)
	m.di = ((m.bp + -272) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.di = (4190) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.runtime(0x6ea, 0x612)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (5) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.di = ((m.bp + -272) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.di = (4190) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.clock(16)
	m.runtime(0x6ea, 0x612)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (4) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (14) & 65535
	m.push(m.ax)
	m.ax = (14) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.clock(16)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.w8(m.ss, (m.bp+-264)&65535, 0)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-2)&65535, m.ax)
L1131:
	m.clock(16)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
	m.w8(m.ds, (5570)&65535, 7)
L113b:
	m.clock(16)
	m.w16(m.ss, (m.bp+-4)&65535, m.step(m.r16(m.ss, (m.bp+-4)&65535), 1, 16))
	m.alu(opSub, m.r16(m.ss, (m.bp+-4)&65535), 10, 16)
	m.clock(16)
	if m.s == m.o {
		goto L114b
	}
	m.w16(m.ss, (m.bp+-6)&65535, 12)
	goto L1150
L114b:
	m.clock(16)
	m.w16(m.ss, (m.bp+-6)&65535, 7)
L1150:
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.bp+-4)&65535), 20, 16)
	if m.z || m.s != m.o {
		goto L115b
	}
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
L115b:
	m.clock(16)
	m.di = ((m.bp + -268) & 65535) & 65535
	m.clock(16)
	m.push(m.ss)
	m.push(m.di)
	m.di = (4196) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.runtime(0x6ea, 0x612)
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 11, 16)) & 65535
	m.push(m.ax)
	m.ax = (7) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.push(m.r16(m.ss, (m.bp+-6)&65535))
	m.clock(16)
	m.push(m.r16(m.ss, (m.bp+-6)&65535))
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.ax = (20) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.w16(m.ds, (5536)&65535, 3)
	m.ax = (m.ax & 65280) | ((51) & 255)
	m.push(m.ax)
	m.di = (5536) & 65535
	m.clock(16)
	m.push(m.ds)
	m.push(m.di)
	m.runtime(0x681, 0xb)
	m.runtime(0x688, 0x2fa)
	m.ax = (m.ax & 65280) | ((m.alu(opOr, (m.ax & 255), (m.ax & 255), 8)) & 255)
	if !m.z {
		goto L11bf
	}
	m.alu(opSub, m.r16(m.ds, (5538)&65535), 0, 16)
	if !m.c && !m.z {
		goto L11bf
	}
	goto L113b
L11bf:
	m.clock(16)
	m.w16(m.ss, (m.bp+-8)&65535, 1)
	goto L11c9
L11c6:
	m.clock(16)
	m.w16(m.ss, (m.bp+-8)&65535, m.step(m.r16(m.ss, (m.bp+-8)&65535), 1, 16))
L11c9:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-8)&65535)) & 65535
	m.product(opImul, m.r16(m.ss, (m.bp+-8)&65535))
	m.cx = (2) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 20, 16)) & 65535
	m.product(opImul, m.r16(m.ds, (5518)&65535))
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
	m.alu(opSub, m.r16(m.ss, (m.bp+-8)&65535), 20, 16)
	if !m.z {
		goto L11c6
	}
	m.runtime(0x688, 0x2f3)
	m.di = ((m.bp + -268) & 65535) & 65535
	m.push(m.ss)
	m.clock(16)
	m.push(m.di)
	m.di = (4196) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.runtime(0x6ea, 0x612)
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 11, 16)) & 65535
	m.push(m.ax)
	m.ax = (7) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (7) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (7) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.runtime(0x688, 0x2fa)
	m.ax = (m.ax & 65280) | ((m.alu(opOr, (m.ax & 255), (m.ax & 255), 8)) & 255)
	if m.z {
		goto L123a
	}
	m.runtime(0x688, 0x30c)
	m.push(m.ax)
	m.runtime(0x6ea, 0x1564)
	m.clock(16)
	m.w8(m.ds, (5570)&65535, (m.ax & 255))
L123a:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (5570)&65535)) & 255)
	m.push(m.ax)
	m.di = (4198) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.runtime(0x6ea, 0x8c7)
	if !m.z {
		goto L124d
	}
	goto L12f2
L124d:
	m.clock(16)
	m.di = ((m.bp + -778) & 65535) & 65535
	m.push(m.ss)
	m.clock(16)
	m.push(m.di)
	m.di = ((m.bp + -264) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.runtime(0x6ea, 0x612)
	m.di = ((m.bp + -522) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (5570)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x742)
	m.runtime(0x6ea, 0x69f)
	m.di = ((m.bp + -264) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.clock(16)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.di = ((m.bp + -522) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (5570)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x742)
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 11, 16)) & 65535
	m.push(m.ax)
	m.ax = (5) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.di = ((m.bp + -522) & 65535) & 65535
	m.push(m.ss)
	m.clock(16)
	m.push(m.di)
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (5570)&65535)) & 255)
	m.push(m.ax)
	m.runtime(0x6ea, 0x742)
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 11, 16)) & 65535
	m.push(m.ax)
	m.ax = (4) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (12) & 65535
	m.push(m.ax)
	m.ax = (12) & 65535
	m.clock(16)
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.w16(m.ss, (m.bp+-2)&65535, m.ax)
L12f2:
	m.clock(16)
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 8, 8)
	if m.z {
		goto L12fc
	}
	m.clock(16)
	goto L13a3
L12fc:
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.bp+-2)&65535), 0, 16)
	if !m.z && m.s == m.o {
		goto L1305
	}
	goto L13a3
L1305:
	m.clock(16)
	m.di = ((m.bp + -522) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.di = ((m.bp + -264) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.ax = (1) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x65e)
	m.di = ((m.bp + -264) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.w16(m.ss, (m.bp+-2)&65535, m.ax)
	m.di = ((m.bp + -268) & 65535) & 65535
	m.clock(16)
	m.push(m.ss)
	m.push(m.di)
	m.di = (4230) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.runtime(0x6ea, 0x612)
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 11, 16)) & 65535
	m.push(m.ax)
	m.ax = (5) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (7) & 65535
	m.clock(16)
	m.push(m.ax)
	m.ax = (7) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
	m.di = ((m.bp + -268) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.di = (4230) & 65535
	m.push(m.cs)
	m.clock(16)
	m.push(m.di)
	m.runtime(0x6ea, 0x612)
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.ax = (m.alu(opAdd, m.ax, 11, 16)) & 65535
	m.push(m.ax)
	m.ax = (4) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (7) & 65535
	m.push(m.ax)
	m.ax = (7) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.clock(16)
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawTextCells()
L13a3:
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.bp+-2)&65535), 9, 16)
	if !m.z && m.s == m.o {
		goto L13c7
	}
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 13, 8)
	if m.z {
		goto L13b7
	}
	m.alu(opSub, m.r16(m.ds, (5538)&65535), 0, 16)
	if m.c || m.z {
		goto L13bd
	}
L13b7:
	m.clock(16)
	m.alu(opSub, m.r16(m.ss, (m.bp+-2)&65535), 2, 16)
	if !m.z && m.s == m.o {
		goto L13c7
	}
L13bd:
	m.clock(16)
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 27, 8)
	if m.z {
		goto L13c7
	}
	goto L1131
L13c7:
	m.clock(16)
	m.alu(opSub, m.r8(m.ds, (5570)&65535), 27, 8)
	if !m.z {
		goto L13d2
	}
	m.push(m.cs)
	m.push(0)
	m.exitGame()
L13d2:
	m.clock(16)
	m.di = ((m.bp + -264) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.di = (4482) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.ax = (m.alu(opXor, m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.w16(m.ss, (m.bp+-268)&65535, m.ax)
	m.ax = (1) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-268)&65535), 16)
	if !m.z && m.s == m.o {
		goto L142e
	}
	m.clock(16)
	m.w16(m.ss, (m.bp+-2)&65535, m.ax)
	goto L1403
L1400:
	m.clock(16)
	m.w16(m.ss, (m.bp+-2)&65535, m.step(m.r16(m.ss, (m.bp+-2)&65535), 1, 16))
L1403:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.dx = (17) & 65535
	m.product(opMul, m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu(opAdd, m.di, 106, 16)) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.di = (4482) & 65535
	m.clock(16)
	m.push(m.ds)
	m.push(m.di)
	m.runtime(0x6ea, 0x717)
	if !m.z {
		goto L1425
	}
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
L1425:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-2)&65535)) & 65535
	m.alu(opSub, m.ax, m.r16(m.ss, (m.bp+-268)&65535), 16)
	if !m.z {
		goto L1400
	}
L142e:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-4)&65535)) & 65535
	m.w16(m.ds, (4470)&65535, m.ax)
	m.clock(16)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 4) & 65535
	return
}
