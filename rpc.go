package neo

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func InitJsonRpc(rcvr interface{}, address string) error {

	rpc.Register(rcvr)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("start connection " + address)

	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("get a new client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}

}

func JsonRequest(address string, method string, args interface{}, reply interface{}) error {
	conn, err := jsonrpc.Dial("tcp", address)
	if err != nil {
		return err
	}
	return conn.Call(method, args, reply)
}
