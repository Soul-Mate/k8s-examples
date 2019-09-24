package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Arith struct {
}

func (a *Arith) Add(args *Args, reply *ArithResult) error {
	reply.Value = args.X + args.Y
	return nil
}

func (a *Arith) Sub(args *Args, reply *ArithResult) error {
	reply.Value = args.X - args.Y
	return nil
}

func (a *Arith) Mul(args *Args, reply *ArithResult) error {
	reply.Value = args.X * args.Y
	return nil
}

func (a *Arith) Div(args *Args, reply *ArithResult) error {
	if args.Y == 0 {
		return errors.New("divide by zero")
	}
	reply.Quotient = float64(args.X / args.Y)
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":30001")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	log.Fatal("rpc server start error:", http.Serve(l, nil))
}
