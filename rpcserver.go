package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"time"
)

const (
	RPC_OK = iota
	RPC_ERR
	RPC_STATUS_OK
	RPC_STATUS_FAIL
)

func NewRPCServer(addr string) (*RaftRpcService, error) {
	rpc := &RaftRpcService{}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer()
	RegisterRaftRpcServer(grpcServer, rpc)
	return rpc, grpcServer.Serve(listener)
}

type RaftRpcService struct{}

//Handler rpc request
func (rrs *RaftRpcService) OpRPC(ctx context.Context, r *OpRequest) (*OpReply, error) {
	var err error
	var value []byte
	switch r.Op {
	case CmdGet:
		value, err = localServer.FSM.Get([]byte(r.Bucket), []byte(r.Key))
	//CmdSet
	default:
		raftState := localServer.Raft
		reqCmd, _ := json.Marshal(r)
		err = raftState.Apply(reqCmd, time.Second).Error()
	}

	if err != nil {
		return &OpReply{
			Status:  RPC_STATUS_FAIL,
			ErrCode: RPC_ERR,
			Msg:     fmt.Sprintf("%v", err),
			Value:   nil,
		}, nil
	}

	return &OpReply{
		Status:  RPC_STATUS_OK,
		ErrCode: RPC_OK,
		Msg:     "",
		Value:   value,
	}, nil
}
