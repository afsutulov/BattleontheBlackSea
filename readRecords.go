package main

// readRecords preserves the original Pascal control flow.
func (m *gameState) readRecords() {
	m.push(m.bp)
	m.bp = (m.sp) & 65535
	m.ax = (388) & 65535
	m.runtime(0x6ea, 0x244)
	m.sp = (m.alu("sub", m.sp, 388, 16)) & 65535
	m.di = (3302) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.di = ((m.bp + -256) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.ax = (255) & 65535
	m.clock(16)
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.di = ((m.bp + -384) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.di = ((m.bp + -256) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.runtime(0x6ea, 0x12f0)
	m.di = ((m.bp + -384) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.ax = (17) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x131e)
	m.clock(16)
	m.runtime(0x6ea, 0x207)
	m.ax = (m.alu("or", m.ax, m.ax, 16)) & 65535
	if !m.z {
		goto Ld71
	}
	goto Lfa9
Ld71:
	m.clock(16)
	m.di = (7068) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.di = (3310) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x12b2)
	m.clock(16)
	m.di = ((m.bp + -256) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x12b2)
	m.di = (3316) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x12b2)
	m.runtime(0x6ea, 0x1235)
	m.runtime(0x6ea, 0x20e)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.clock(16)
	m.w16(m.ds, (4468)&65535, m.ax)
	goto Ldb3
Ldaf:
	m.clock(16)
	m.w16(m.ds, (4468)&65535, m.step(m.r16(m.ds, (4468)&65535), 1, 16))
Ldb3:
	m.clock(16)
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, 106, 16)) & 65535
	m.w16(m.ss, (m.bp+-388)&65535, m.di)
	m.w16(m.ss, (m.bp+-386)&65535, m.ds)
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.clock(16)
	m.alu("sub", m.ax, 0, 16)
	if !m.z {
		goto Le0e
	}
	m.di = (3328) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w16(m.es, (m.di+11)&65535, 3)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w16(m.es, (m.di+13)&65535, 10)
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+15)&65535, 27)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+16)&65535, 0)
	goto Lf96
Le0e:
	m.clock(16)
	m.alu("sub", m.ax, 1, 16)
	if !m.z {
		goto Le50
	}
	m.di = (3333) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.clock(16)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w16(m.es, (m.di+11)&65535, 10)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w16(m.es, (m.di+13)&65535, 6)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+15)&65535, 27)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+16)&65535, 0)
	goto Lf96
Le50:
	m.clock(16)
	m.alu("sub", m.ax, 2, 16)
	m.clock(16)
	if !m.z {
		goto Le92
	}
	m.di = (3339) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w16(m.es, (m.di+11)&65535, 10)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w16(m.es, (m.di+13)&65535, 5)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.clock(16)
	m.w8(m.es, (m.di+15)&65535, 32)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+16)&65535, 0)
	goto Lf96
Le92:
	m.clock(16)
	m.alu("sub", m.ax, 3, 16)
	if !m.z {
		goto Led4
	}
	m.di = (3328) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (10) & 65535
	m.clock(16)
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w16(m.es, (m.di+11)&65535, 8)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w16(m.es, (m.di+13)&65535, 4)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+15)&65535, 35)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+16)&65535, 0)
	goto Lf96
Led4:
	m.clock(16)
	m.alu("sub", m.ax, 4, 16)
	if !m.z {
		goto Lf16
	}
	m.clock(16)
	m.di = (3344) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w16(m.es, (m.di+11)&65535, 6)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w16(m.es, (m.di+13)&65535, 3)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+15)&65535, 38)
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+16)&65535, 0)
	goto Lf96
Lf16:
	m.clock(16)
	m.alu("sub", m.ax, 5, 16)
	if !m.z {
		goto Lf57
	}
	m.di = (3351) & 65535
	m.push(m.cs)
	m.push(m.di)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.clock(16)
	m.runtime(0x6ea, 0x62c)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w16(m.es, (m.di+11)&65535, 4)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w16(m.es, (m.di+13)&65535, 2)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+15)&65535, 40)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+16)&65535, 0)
	goto Lf96
Lf57:
	m.clock(16)
	m.alu("sub", m.ax, 6, 16)
	if !m.z {
		goto Lf96
	}
	m.di = (3357) & 65535
	m.clock(16)
	m.push(m.cs)
	m.push(m.di)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.push(m.es)
	m.push(m.di)
	m.ax = (10) & 65535
	m.push(m.ax)
	m.runtime(0x6ea, 0x62c)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w16(m.es, (m.di+11)&65535, 2)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w16(m.es, (m.di+13)&65535, 1)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+15)&65535, 48)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.clock(16)
	m.w8(m.es, (m.di+16)&65535, 0)
Lf96:
	m.clock(16)
	m.alu("sub", m.r16(m.ds, (4468)&65535), 6, 16)
	if m.z {
		goto Lfa0
	}
	goto Ldaf
Lfa0:
	m.clock(16)
	m.w16(m.ds, (4468)&65535, 6)
	goto L105a
Lfa9:
	m.clock(16)
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.ds, (4468)&65535, m.ax)
Lfae:
	m.clock(16)
	m.alu("sub", m.r16(m.ds, (4468)&65535), 255, 16)
	if m.s != m.o {
		goto Lfb9
	}
	goto L1046
Lfb9:
	m.clock(16)
	m.di = ((m.bp + -384) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.runtime(0x6ea, 0x14d3)
	m.runtime(0x6ea, 0x20e)
	m.ax = (m.ax & 65280) | ((m.alu("or", (m.ax & 255), (m.ax & 255), 8)) & 255)
	if !m.z {
		goto L1046
	}
	m.di = ((m.bp + -384) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.clock(16)
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, 106, 16)) & 65535
	m.push(m.ds)
	m.push(m.di)
	m.runtime(0x6ea, 0x13d3)
	m.sp = (m.alu("add", m.sp, 4, 16)) & 65535
	m.runtime(0x6ea, 0x20e)
	m.push(m.r16(m.ds, (4468)&65535))
	m.push(m.cs)
	m.push(0)
	m.validRecord()
	m.ax = (m.ax & 65280) | ((m.alu("or", (m.ax & 255), (m.ax & 255), 8)) & 255)
	if !m.z {
		goto L103f
	}
	m.clock(16)
	m.alu("sub", m.r16(m.ds, (4468)&65535), 0, 16)
	if m.z || m.s != m.o {
		goto L103f
	}
	m.ax = (m.r16(m.ds, (4468)&65535)) & 65535
	m.dx = (17) & 65535
	m.product("mul", m.dx)
	m.di = (m.ax) & 65535
	m.di = (m.alu("add", m.di, 106, 16)) & 65535
	m.w16(m.ss, (m.bp+-388)&65535, m.di)
	m.w16(m.ss, (m.bp+-386)&65535, m.ds)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.es, (m.di+11)&65535, m.ax)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.ax = (m.alu("xor", m.ax, m.ax, 16)) & 65535
	m.w16(m.es, (m.di+13)&65535, m.ax)
	m.clock(16)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+15)&65535, 99)
	m.temp = m.r16(m.ss, (m.bp+-388)&65535)
	m.es = m.r16(m.ss, ((m.bp+-388)&65535)+2)
	m.di = (m.temp) & 65535
	m.w8(m.es, (m.di+16)&65535, 0)
L103f:
	m.clock(16)
	m.w16(m.ds, (4468)&65535, m.step(m.r16(m.ds, (4468)&65535), 1, 16))
	goto Lfae
L1046:
	m.clock(16)
	m.w16(m.ds, (4468)&65535, m.step(m.r16(m.ds, (4468)&65535), -1, 16))
	m.di = ((m.bp + -384) & 65535) & 65535
	m.push(m.ss)
	m.push(m.di)
	m.runtime(0x6ea, 0x139f)
	m.clock(16)
	m.runtime(0x6ea, 0x20e)
L105a:
	m.clock(16)
	m.sp = (m.bp) & 65535
	m.bp = (m.pop()) & 65535
	m.recordsReady = true
	m.sp = (m.sp + 4) & 65535
	return
}
