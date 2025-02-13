package deadline

import (
	"net"
	"time"

	"github.com/sagernet/sing/common/buf"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
)

type PacketConn struct {
	N.NetPacketConn
	reader *PacketReader
}

func NewPacketConn(conn N.NetPacketConn) *PacketConn {
	return &PacketConn{NetPacketConn: conn, reader: NewPacketReader(conn)}
}

func (c *PacketConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	return c.reader.ReadFrom(p)
}

func (c *PacketConn) ReadPacket(buffer *buf.Buffer) (destination M.Socksaddr, err error) {
	return c.reader.ReadPacket(buffer)
}

func (c *PacketConn) SetReadDeadline(t time.Time) error {
	return c.NetPacketConn.SetReadDeadline(t)
}

func (c *PacketConn) ReaderReplaceable() bool {
	return c.reader.ReaderReplaceable()
}

func (c *PacketConn) WriterReplaceable() bool {
	return true
}

func (c *PacketConn) Upstream() any {
	return c.NetPacketConn
}
