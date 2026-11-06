package main

// placePlayerFleet preserves the original Pascal control flow.
func (m *gameState) placePlayerFleet() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (30) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu("sub", m.sp, 30, 16)) & 65535
	m.w16(m.ds, (4472)&65535, 20)
	m.push(m.bp)
	m.push(m.cs)
	m.push(0)
	m.drawPlacementCursor()
	m.push(m.cs)
	m.push(0)
	m.screenTransition()
	m.w16(m.ss, (m.bp+-16)&65535, 1)
	m.clock(16)
	goto L3348
L3345:
	m.clock(16)
	m.w16(m.ss, (m.bp+-16)&65535, m.step(m.r16(m.ss, (m.bp+-16)&65535), 1, 16))
L3348:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+5351)&65535)) & 255)
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.w8(m.ds, (m.di+5353)&65535, (m.ax & 255))
	m.w16(m.ss, (m.bp+-10)&65535, 1)
	m.clock(16)
	m.w16(m.ss, (m.bp+-12)&65535, 1)
	m.w16(m.ss, (m.bp+-14)&65535, 1)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-2)&65535, m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (5530)&65535, m.ax)
L3377:
	m.clock(16)
	m.w16(m.ds, (5530)&65535, m.step(m.r16(m.ds, (5530)&65535), 1, 16))
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((1) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.clock(16)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (8000) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x152d)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.clock(16)
	m.runtime(0x6ea, 0xcc7)
	m.ax = (m.alu("add", m.ax, 53, 16)) & 65535
	m.di = (m.ax) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.temp = m.r16(m.ds, (m.di+5572)&65535)
	m.es = m.r16(m.ds, ((m.di+5572)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (20) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
	m.w16(m.ds, (5536)&65535, 3)
	m.ax = (m.ax & 65280) | ((51) & 255)
	m.push(m.ax)
	m.di = (5536) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.runtime(0x681, 0xb)
	m.alu("sub", m.r16(m.ss, (m.bp+-14)&65535), 0, 16)
	if !m.z {
		goto L3420
	}
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.cx = (14) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 48, 16)) & 65535
	m.w16(m.ss, (m.bp+-6)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.cx = (14) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 5, 16)) & 65535
	m.w16(m.ss, (m.bp+-8)&65535, m.ax)
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+5351)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.clock(16)
	m.ax = (m.alu("add", m.ax, 4, 16)) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
	goto L344a
L3420:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.cx = (14) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 49, 16)) & 65535
	m.w16(m.ss, (m.bp+-6)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.cx = (14) & 65535
	m.product("imul", m.cx)
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.w16(m.ss, (m.bp+-8)&65535, m.ax)
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+5351)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
L344a:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5530)&65535)) & 65535
	m.dx = (m.alu("xor", m.dx, m.dx, 16)) & 65535
	m.cx = (5) & 65535
	m.product("div", m.cx)
	m.temp = m.dx
	m.dx = (m.ax) & 65535
	m.ax = (m.temp) & 65535
	m.alu("sub", m.ax, 4, 16)
	if m.c {
		goto L345d
	}
	m.clock(16)
	goto L3545
L345d:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+5351)&65535)) & 255)
	m.alu("sub", (m.ax & 255), 1, 8)
	if !m.z {
		goto L3481
	}
	m.di = (13014) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.clock(16)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	goto L34ca
L3481:
	m.clock(16)
	m.alu("sub", (m.ax & 255), 2, 8)
	if !m.z {
		goto L349a
	}
	m.di = (13027) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.clock(16)
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	goto L34ca
L349a:
	m.clock(16)
	m.alu("sub", (m.ax & 255), 3, 8)
	if !m.z {
		goto L34b3
	}
	m.di = (13037) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.clock(16)
	m.runtime(0x6ea, 0x62c)
	goto L34ca
L34b3:
	m.clock(16)
	m.alu("sub", (m.ax & 255), 4, 8)
	if !m.z {
		goto L34ca
	}
	m.di = (13045) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (255) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
L34ca:
	m.clock(16)
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (221) & 65535
	m.push(m.ax)
	m.ax = (3) & 65535
	m.push(m.ax)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (8) & 65535
	m.push(m.ax)
	m.ax = (8) & 65535
	m.clock(16)
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.di = (6056) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.ax = (220) & 65535
	m.push(m.ax)
	m.ax = (2) & 65535
	m.push(m.ax)
	m.clock(16)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.ax = (15) & 65535
	m.push(m.ax)
	m.ax = (15) & 65535
	m.push(m.ax)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.cs)
	m.push(0)
	m.drawText()
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-4)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.temp = m.r16(m.ds, (m.di+5572)&65535)
	m.es = m.r16(m.ds, ((m.di+5572)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.r16(m.ss, (m.bp+-6)&65535))
	m.push(m.r16(m.ss, (m.bp+-8)&65535))
	m.ax = (1) & 65535
	m.push(m.ax)
	m.clock(16)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
L3545:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((1) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (24000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.clock(16)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (8000) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x152d)
	m.runtime(0x688, 0x2fa)
	m.ax = (m.ax & 65280) | ((m.alu("or", (m.ax & 255), (m.ax & 255), 8)) & 255)
	if !m.z {
		goto L357a
	}
	goto L3932
L357a:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, 1)
	m.clock(16)
	goto L3586
L3582:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L3586:
	m.clock(16)
	m.ax = (m.r16(m.ds, (5558)&65535)) & 65535
	m.product("imul", m.r16(m.ds, (5558)&65535))
	m.cx = (2) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.alu("add", m.ax, 20, 16)) & 65535
	m.product("imul", m.r16(m.ds, (5518)&65535))
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
	m.alu("sub", m.r16(m.ds, (5558)&65535), 20, 16)
	m.clock(16)
	if !m.z {
		goto L3582
	}
	m.runtime(0x688, 0x2f3)
	m.runtime(0x688, 0x30c)
	m.w8(m.ds, (5570)&65535, (m.ax & 255))
	m.alu("sub", m.r8(m.ds, (5570)&65535), 0, 8)
	if m.z {
		goto L35bd
	}
	goto L3643
L35bd:
	m.clock(16)
	m.runtime(0x688, 0x30c)
	m.w8(m.ds, (5570)&65535, (m.ax & 255))
	m.alu("sub", m.r8(m.ds, (5570)&65535), 72, 8)
	if !m.z {
		goto L35cf
	}
	m.w16(m.ss, (m.bp+-12)&65535, m.step(m.r16(m.ss, (m.bp+-12)&65535), -1, 16))
L35cf:
	m.clock(16)
	m.alu("sub", m.r8(m.ds, (5570)&65535), 80, 8)
	if !m.z {
		goto L35d9
	}
	m.w16(m.ss, (m.bp+-12)&65535, m.step(m.r16(m.ss, (m.bp+-12)&65535), 1, 16))
L35d9:
	m.clock(16)
	m.alu("sub", m.r8(m.ds, (5570)&65535), 75, 8)
	if !m.z {
		goto L35e3
	}
	m.w16(m.ss, (m.bp+-10)&65535, m.step(m.r16(m.ss, (m.bp+-10)&65535), -1, 16))
L35e3:
	m.clock(16)
	m.alu("sub", m.r8(m.ds, (5570)&65535), 77, 8)
	if !m.z {
		goto L35ed
	}
	m.w16(m.ss, (m.bp+-10)&65535, m.step(m.r16(m.ss, (m.bp+-10)&65535), 1, 16))
L35ed:
	m.clock(16)
	m.alu("sub", m.r8(m.ds, (5570)&65535), 59, 8)
	if !m.z {
		goto L35fe
	}
	m.ax = (1) & 65535
	m.ax = (m.alu("sub", m.ax, m.r16(m.ds, (5518)&65535), 16)) & 65535
	m.w16(m.ds, (5518)&65535, m.ax)
L35fe:
	m.clock(16)
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (5570)&65535)) & 255)
	m.push(m.ax)
	m.di = (13057) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.runtime(0x6ea, 0x8c7)
	if m.z {
		goto L3617
	}
	m.ax = (1) & 65535
	m.ax = (m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-14)&65535), 16)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-14)&65535, m.ax)
L3617:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-10)&65535), 1, 16)
	if m.s == m.o {
		goto L3622
	}
	m.w16(m.ss, (m.bp+-10)&65535, 10)
L3622:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-10)&65535), 10, 16)
	if m.z || m.s != m.o {
		goto L362d
	}
	m.w16(m.ss, (m.bp+-10)&65535, 1)
L362d:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-12)&65535), 1, 16)
	if m.s == m.o {
		goto L3638
	}
	m.clock(16)
	m.w16(m.ss, (m.bp+-12)&65535, 10)
L3638:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-12)&65535), 10, 16)
	if m.z || m.s != m.o {
		goto L3643
	}
	m.w16(m.ss, (m.bp+-12)&65535, 1)
L3643:
	m.clock(16)
	m.alu("sub", m.r8(m.ds, (5570)&65535), 27, 8)
	if !m.z {
		goto L364e
	}
	m.push(m.cs)
	m.push(0)
	m.exitGame()
L364e:
	m.clock(16)
	m.alu("sub", m.r8(m.ds, (5570)&65535), 13, 8)
	m.clock(16)
	if m.z {
		goto L3658
	}
	goto L392d
L3658:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+5351)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.w16(m.ss, (m.bp+-24)&65535, m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-26)&65535, m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-28)&65535, m.ax)
	m.alu("sub", m.r16(m.ss, (m.bp+-14)&65535), 0, 16)
	m.clock(16)
	if m.z || m.s != m.o {
		goto L367f
	}
	m.w16(m.ss, (m.bp+-28)&65535, 1)
	goto L3684
L367f:
	m.clock(16)
	m.w16(m.ss, (m.bp+-26)&65535, 1)
L3684:
	m.clock(16)
	m.w16(m.ss, (m.bp+-2)&65535, 1)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.w16(m.ss, (m.bp+-20)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-30)&65535, m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if !m.z && m.s == m.o {
		goto L36f2
	}
	m.w16(m.ss, (m.bp+-18)&65535, m.ax)
	goto L36ab
L36a8:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, m.step(m.r16(m.ss, (m.bp+-18)&65535), 1, 16))
L36ab:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-20)&65535), 11, 16)
	if m.s == m.o {
		goto L36d3
	}
	m.alu("sub", m.r16(m.ss, (m.bp+-22)&65535), 11, 16)
	if m.s == m.o {
		goto L36d3
	}
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-20)&65535)) & 65535
	m.dx = (24) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, m.cx, 16)) & 65535
	m.alu("sub", m.r16(m.ds, (m.di+4494)&65535), 0, 16)
	if !m.z {
		goto L36d3
	}
	goto L36d8
L36d3:
	m.clock(16)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-2)&65535, m.ax)
L36d8:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-20)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-26)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-20)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-28)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-18)&65535)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if !m.z {
		goto L36a8
	}
L36f2:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-2)&65535), 0, 16)
	if !m.z && m.s == m.o {
		goto L36fb
	}
	goto L38f1
L36fb:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-14)&65535), 0, 16)
	if !m.z && m.s == m.o {
		goto L3704
	}
	goto L37ab
L3704:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.w16(m.ss, (m.bp+-30)&65535, m.ax)
	m.ax = (65535) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if m.z || m.s != m.o {
		goto L3715
	}
	goto L37a8
L3715:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, m.ax)
	m.clock(16)
	goto L371d
L371a:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, m.step(m.r16(m.ss, (m.bp+-18)&65535), 1, 16))
L371d:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-18)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.alu("sub", m.ax, 11, 16)
	if m.s == m.o {
		goto L374f
	}
	m.alu("sub", m.r16(m.ss, (m.bp+-22)&65535), 11, 16)
	if m.s == m.o {
		goto L374f
	}
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.dx = (24) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4494)&65535, 65535)
L374f:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-10)&65535), 11, 16)
	if m.s == m.o {
		goto L3774
	}
	m.alu("sub", m.r16(m.ss, (m.bp+-22)&65535), 11, 16)
	m.clock(16)
	if m.s == m.o {
		goto L3774
	}
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.dx = (24) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4494)&65535, 65535)
L3774:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.alu("sub", m.ax, 11, 16)
	m.clock(16)
	if m.s == m.o {
		goto L379d
	}
	m.alu("sub", m.r16(m.ss, (m.bp+-22)&65535), 11, 16)
	if m.s == m.o {
		goto L379d
	}
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.dx = (24) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4494)&65535, 65535)
L379d:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-18)&65535)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if m.z {
		goto L37a8
	}
	goto L371a
L37a8:
	m.clock(16)
	goto L384f
L37ab:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.w16(m.ss, (m.bp+-30)&65535, m.ax)
	m.ax = (65535) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if m.z || m.s != m.o {
		goto L37bc
	}
	goto L384f
L37bc:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, m.ax)
	goto L37c4
L37c1:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, m.step(m.r16(m.ss, (m.bp+-18)&65535), 1, 16))
L37c4:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-18)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-20)&65535, m.ax)
	m.alu("sub", m.r16(m.ss, (m.bp+-20)&65535), 11, 16)
	if m.s == m.o {
		goto L37f6
	}
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.alu("sub", m.ax, 11, 16)
	m.clock(16)
	if m.s == m.o {
		goto L37f6
	}
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-20)&65535)) & 65535
	m.dx = (24) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4494)&65535, 65535)
L37f6:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-20)&65535), 11, 16)
	if m.s == m.o {
		goto L381b
	}
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-12)&65535), 11, 16)
	if m.s == m.o {
		goto L381b
	}
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-20)&65535)) & 65535
	m.dx = (24) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4494)&65535, 65535)
L381b:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-20)&65535), 11, 16)
	if m.s == m.o {
		goto L3844
	}
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.alu("sub", m.ax, 11, 16)
	if m.s == m.o {
		goto L3844
	}
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-20)&65535)) & 65535
	m.dx = (24) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4494)&65535, 65535)
L3844:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-18)&65535)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if m.z {
		goto L384f
	}
	goto L37c1
L384f:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.w16(m.ss, (m.bp+-20)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-24)&65535)) & 65535
	m.ax = (m.step(m.ax, -1, 16)) & 65535
	m.w16(m.ss, (m.bp+-30)&65535, m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	m.clock(16)
	if !m.z && m.s == m.o {
		goto L38a5
	}
	m.w16(m.ss, (m.bp+-18)&65535, m.ax)
	goto L3871
L386e:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, m.step(m.r16(m.ss, (m.bp+-18)&65535), 1, 16))
L3871:
	m.clock(16)
	m.bx = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.shift(m.ax, 1, 16, true)) & 65535
	m.cx = (m.ax) & 65535
	m.ax = (m.r16(m.ss, (m.bp+-20)&65535)) & 65535
	m.dx = (24) & 65535
	m.product("mul", m.dx)
	m.clock(16)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, m.cx, 16)) & 65535
	m.w16(m.ds, (m.di+4494)&65535, m.bx)
	m.ax = (m.r16(m.ss, (m.bp+-20)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-26)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-20)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-22)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-28)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-22)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-18)&65535)) & 65535
	m.alu("sub", m.ax, m.r16(m.ss, (m.bp+-30)&65535), 16)
	if !m.z {
		goto L386e
	}
L38a5:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, 1)
	m.clock(16)
	goto L38af
L38ac:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, m.step(m.r16(m.ss, (m.bp+-18)&65535), 1, 16))
L38af:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-18)&65535)) & 65535
	m.cx = (10) & 65535
	m.product("imul", m.cx)
	m.dx = (m.ax) & 65535
	m.ax = (340) & 65535
	m.ax = (m.alu("sub", m.ax, m.dx, 16)) & 65535
	m.product("imul", m.r16(m.ds, (5518)&65535))
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-18)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (6) & 65535
	m.product("idiv", m.cx)
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.runtime(0x688, 0x2f3)
	m.ax = (m.r16(m.ss, (m.bp+-18)&65535)) & 65535
	m.ax = (m.alu("add", m.ax, 2, 16)) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.alu("sub", m.r16(m.ss, (m.bp+-18)&65535), 20, 16)
	if !m.z {
		goto L38ac
	}
	goto L392d
L38f1:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, 1)
	goto L38fb
L38f8:
	m.clock(16)
	m.w16(m.ss, (m.bp+-18)&65535, m.step(m.r16(m.ss, (m.bp+-18)&65535), 1, 16))
L38fb:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-18)&65535)) & 65535
	m.cx = (40) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 340, 16)) & 65535
	m.product("imul", m.r16(m.ds, (5518)&65535))
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
	m.clock(16)
	m.ax = (4) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.runtime(0x688, 0x2f3)
	m.ax = (4) & 65535
	m.push(m.ax)
	m.runtime(0x688, 0x29e)
	m.alu("sub", m.r16(m.ss, (m.bp+-18)&65535), 20, 16)
	if !m.z {
		goto L38f8
	}
L392d:
	m.clock(16)
	m.runtime(0x688, 0x2f3)
L3932:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-2)&65535), 1, 16)
	m.clock(16)
	if m.z {
		goto L393b
	}
	goto L3377
L393b:
	m.clock(16)
	m.push(m.bp)
	m.push(m.cs)
	m.push(0)
	m.drawPlacementCursor()
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+-14)&65535)) & 255)
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.w8(m.ds, (m.di+5352)&65535, (m.ax & 255))
	m.ax = (m.r16(m.ss, (m.bp+-10)&65535)) & 65535
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.clock(16)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.w16(m.ds, (m.di+5354)&65535, m.ax)
	m.ax = (m.r16(m.ss, (m.bp+-12)&65535)) & 65535
	m.di = (m.r16(m.ss, (m.bp+-16)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.w16(m.ds, (m.di+5356)&65535, m.ax)
	m.w16(m.ds, (5558)&65535, 1)
	goto L3976
L3972:
	m.clock(16)
	m.w16(m.ds, (5558)&65535, m.step(m.r16(m.ds, (5558)&65535), 1, 16))
L3976:
	m.clock(16)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.clock(16)
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.alu("sub", m.r8(m.ds, (m.di+5353)&65535), 0, 8)
	if !m.c && !m.z {
		goto L3988
	}
	goto L3a3d
L3988:
	m.clock(16)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.alu("sub", m.r8(m.ds, (m.di+5352)&65535), 0, 8)
	if !m.z {
		goto L39db
	}
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.clock(16)
	m.ax = (m.r16(m.ds, (m.di+5354)&65535)) & 65535
	m.cx = (14) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 48, 16)) & 65535
	m.w16(m.ss, (m.bp+-6)&65535, m.ax)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.r16(m.ds, (m.di+5356)&65535)) & 65535
	m.cx = (14) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 5, 16)) & 65535
	m.w16(m.ss, (m.bp+-8)&65535, m.ax)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.clock(16)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+5351)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.ax = (m.alu("add", m.ax, 4, 16)) & 65535
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
	goto L3a18
L39db:
	m.clock(16)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.r16(m.ds, (m.di+5354)&65535)) & 65535
	m.cx = (14) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, 49, 16)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-6)&65535, m.ax)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.r16(m.ds, (m.di+5356)&65535)) & 65535
	m.cx = (14) & 65535
	m.product("imul", m.cx)
	m.ax = (m.step(m.ax, 1, 16)) & 65535
	m.w16(m.ss, (m.bp+-8)&65535, m.ax)
	m.di = (m.r16(m.ds, (5558)&65535)) & 65535
	m.cx = (m.cx & 65280) | ((3) & 255)
	m.di = (m.shift(m.di, (m.cx & 255), 16, true)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ds, (m.di+5351)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.w16(m.ss, (m.bp+-4)&65535, m.ax)
L3a18:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-4)&65535)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.di = (m.shift(m.di, 1, 16, true)) & 65535
	m.temp = m.r16(m.ds, (m.di+5572)&65535)
	m.es = m.r16(m.ds, ((m.di+5572)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (48000) & 65535
	m.push(m.es)
	m.push(m.di)
	m.push(m.r16(m.ss, (m.bp+-6)&65535))
	m.push(m.r16(m.ss, (m.bp+-8)&65535))
	m.ax = (1) & 65535
	m.clock(16)
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
L3a3d:
	m.clock(16)
	m.alu("sub", m.r16(m.ds, (5558)&65535), 10, 16)
	if m.z {
		goto L3a47
	}
	goto L3972
L3a47:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-16)&65535), 10, 16)
	if m.z {
		goto L3a50
	}
	goto L3345
L3a50:
	m.clock(16)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 4) & 65535
	return
}
