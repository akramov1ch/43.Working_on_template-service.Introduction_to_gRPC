package main

import (
	"errors"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Arith struct{}

type Args struct {
	A, B float64
}

type Quotient struct {
	Quo, Rem float64
}

func (t *Arith) Add(args *Args, reply *float64) error {
	*reply = args.A + args.B
	return nil
}

func (t *Arith) Subtract(args *Args, reply *float64) error {
	*reply = args.A - args.B
	return nil
}

func (t *Arith) Multiply(args *Args, reply *float64) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = float64(int(args.A) % int(args.B))
	return nil
}

func main() {
	arith := new(Arith)
	server := rpc.NewServer()
	server.Register(arith)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
