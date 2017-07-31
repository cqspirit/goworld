package dispatcher_client

import (
	"net"

	"time"

	"github.com/xiaonanln/goworld/gwlog"
	"github.com/xiaonanln/goworld/netutil"
	"github.com/xiaonanln/goworld/proto"
)

type DispatcherClient struct {
	*proto.GoWorldConnection
}

func newDispatcherClient(conn net.Conn) *DispatcherClient {
	gwc := proto.NewGoWorldConnection(netutil.NewBufferedReadConnection(netutil.NetConnection{conn}), false)

	dc := &DispatcherClient{
		GoWorldConnection: gwc,
	}
	go func() {
		defer gwlog.Debug("%s: auto flush routine quited", gwc)
		for !gwc.IsClosed() {
			time.Sleep(time.Millisecond * 10)
			dispatcherClientDelegate.HandleDispatcherClientBeforeFlush()

			err := gwc.Flush()
			if err != nil {
				break
			}
		}
	}()
	return dc
}

func (dc *DispatcherClient) Close() error {
	return dc.GoWorldConnection.Close()
}
