package main

// screenTransition preserves the original Pascal control flow.
func (m *gameState) screenTransition() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (48) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu("sub", m.sp, 48, 16)) & 65535
	m.w8(m.ds, (5570)&65535, 7)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((1) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.clock(16)
	m.w16(m.ss, (m.bp+-46)&65535, 8000)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.ss, (m.bp+-42)&65535, m.ax)
	goto L1466
L1463:
	m.clock(16)
	m.w16(m.ss, (m.bp+-42)&65535, m.step(m.r16(m.ss, (m.bp+-42)&65535), 1, 16))
L1466:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-42)&65535)) & 65535
	m.w8(m.ss, (m.bp+m.di+-40)&65535, 0)
	m.alu("sub", m.r16(m.ss, (m.bp+-42)&65535), 39, 16)
	if !m.z {
		goto L1463
	}
L1473:
	m.clock(16)
	m.ax = (40) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0xcc7)
	m.w16(m.ss, (m.bp+-42)&65535, m.ax)
	m.runtime(0x688, 0x2fa)
	m.ax = (m.ax & 65280) | ((m.alu("or", (m.ax & 255), (m.ax & 255), 8)) & 255)
	if m.z {
		goto L14bb
	}
	m.runtime(0x688, 0x30c)
	m.w8(m.ds, (5570)&65535, (m.ax & 255))
	m.alu("sub", m.r8(m.ds, (5570)&65535), 0, 8)
	if !m.z {
		goto L14b0
	}
	m.runtime(0x688, 0x30c)
	m.w8(m.ds, (5570)&65535, (m.ax & 255))
	m.alu("sub", m.r8(m.ds, (5570)&65535), 59, 8)
	if !m.z {
		goto L14b0
	}
	m.ax = (1) & 65535
	m.clock(16)
	m.ax = (m.alu("sub", m.ax, m.r16(m.ds, (5518)&65535), 16)) & 65535
	m.w16(m.ds, (5518)&65535, m.ax)
L14b0:
	m.clock(16)
	m.alu("sub", m.r8(m.ds, (5570)&65535), 27, 8)
	if !m.z {
		goto L14bb
	}
	m.push(m.cs)
	m.push(0)
	m.exitGame()
L14bb:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-42)&65535)) & 65535
	m.alu("sub", m.r8(m.ss, (m.bp+m.di+-40)&65535), 200, 8)
	if !m.c {
		goto L1473
	}
	m.ax = (m.r16(m.ss, (m.bp+-46)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.clock(16)
	m.cx = (50) & 65535
	m.product("idiv", m.cx)
	m.temp = m.dx
	m.dx = (m.ax) & 65535
	m.ax = (m.temp) & 65535
	m.ax = (m.alu("or", m.ax, m.ax, 16)) & 65535
	if !m.z {
		goto L14ee
	}
	m.runtime(0x688, 0x2f3)
	m.ax = (m.r16(m.ss, (m.bp+-46)&65535)) & 65535
	m.dx = 0
	if m.ax&32768 != 0 {
		m.dx = 65535
	}
	m.cx = (200) & 65535
	m.product("idiv", m.cx)
	m.temp = m.dx
	m.dx = (m.ax) & 65535
	m.ax = (m.temp) & 65535
	m.ax = (m.alu("add", m.ax, 30, 16)) & 65535
	m.product("imul", m.r16(m.ds, (5518)&65535))
	m.push(m.ax)
	m.runtime(0x688, 0x2c6)
L14ee:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-42)&65535)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+m.di+-40)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.cx = (40) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-42)&65535), 16)) & 65535
	m.w16(m.ss, (m.bp+-44)&65535, m.ax)
	m.ax = (m.ax & 65280) | ((5) & 255)
	m.dx = (974) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.ax & 65280) | ((1) & 255)
	m.dx = (975) & 65535
	m.videoOut(m.dx, (m.ax & 255))
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.clock(16)
	m.w16(m.ss, (m.bp+-48)&65535, m.ax)
	goto L1517
L1514:
	m.clock(16)
	m.w16(m.ss, (m.bp+-48)&65535, m.step(m.r16(m.ss, (m.bp+-48)&65535), 1, 16))
L1517:
	m.clock(16)
	m.ax = (m.r16(m.ss, (m.bp+-48)&65535)) & 65535
	m.cx = (40) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-44)&65535), 16)) & 65535
	m.alu("sub", m.ax, 8000, 16)
	if m.s == m.o {
		goto L155c
	}
	m.ax = (m.r16(m.ss, (m.bp+-48)&65535)) & 65535
	m.cx = (40) & 65535
	m.clock(16)
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-44)&65535), 16)) & 65535
	m.di = (m.ax) & 65535
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (m.alu("add", m.di, 48000, 16)) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (m.r16(m.ss, (m.bp+-48)&65535)) & 65535
	m.cx = (40) & 65535
	m.product("imul", m.cx)
	m.ax = (m.alu("add", m.ax, m.r16(m.ss, (m.bp+-44)&65535), 16)) & 65535
	m.di = (m.ax) & 65535
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.clock(16)
	m.push(m.es)
	m.push(m.di)
	m.ax = (1) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x152d)
L155c:
	m.clock(16)
	m.alu("sub", m.r16(m.ss, (m.bp+-48)&65535), 11, 16)
	if !m.z {
		goto L1514
	}
	m.di = (m.r16(m.ss, (m.bp+-42)&65535)) & 65535
	m.alu("sub", m.r8(m.ss, (m.bp+m.di+-40)&65535), 189, 8)
	if !m.c {
		goto L1597
	}
	m.temp = m.r16(m.ds, (5636)&65535)
	m.es = m.r16(m.ds, ((5636)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.clock(16)
	m.si = (40960) & 65535
	m.es = (m.si) & 65535
	m.di = (0) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (m.r16(m.ss, (m.bp+-42)&65535)) & 65535
	m.cx = (3) & 65535
	m.ax = (m.shift(m.ax, (m.cx & 255), 16, true)) & 65535
	m.push(m.ax)
	m.di = (m.r16(m.ss, (m.bp+-42)&65535)) & 65535
	m.ax = (m.ax & 65280) | ((m.r8(m.ss, (m.bp+m.di+-40)&65535)) & 255)
	m.ax = (m.ax & 255) | (((m.alu("xor", (m.ax >> 8), (m.ax >> 8), 8)) & 255) << 8)
	m.ax = (m.alu("add", m.ax, 4, 16)) & 65535
	m.push(m.ax)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.clock(16)
	m.push(m.ax)
	m.push(m.cs)
	m.push(0)
	m.drawSprite()
L1597:
	m.clock(16)
	m.di = (m.r16(m.ss, (m.bp+-42)&65535)) & 65535
	m.w8(m.ss, (m.bp+m.di+-40)&65535, m.alu("add", m.r8(m.ss, (m.bp+m.di+-40)&65535), 4, 8))
	m.w16(m.ss, (m.bp+-46)&65535, m.alu("sub", m.r16(m.ss, (m.bp+-46)&65535), 4, 16))
	m.alu("sub", m.r16(m.ss, (m.bp+-46)&65535), 0, 16)
	if m.z || m.s != m.o {
		goto L15ab
	}
	goto L1473
L15ab:
	m.clock(16)
	m.runtime(0x688, 0x2f3)
	m.sp = (m.bp) & 65535
	m.clock(16)
	m.bp = (m.pop()) & 65535
	m.sp = (m.sp + 4) & 65535
	return
}
