package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B float64
}

type Quotient struct {
	Quo, Rem float64
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer conn.Close()

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	args := Args{A: 15, B: 3}
	var reply float64
	var quotient Quotient

	// Qo'shish
	err = client.Call("Arith.Add", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Add: %f + %f = %f\n", args.A, args.B, reply)

	// Ayirish
	err = client.Call("Arith.Subtract", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Subtract: %f - %f = %f\n", args.A, args.B, reply)

	// Ko'paytirish
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Multiply: %f * %f = %f\n", args.A, args.B, reply)

	// Bo'lish
	err = client.Call("Arith.Divide", args, &quotient)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Divide: %f / %f = %f remainder %f\n", args.A, args.B, quotient.Quo, quotient.Rem)
}
