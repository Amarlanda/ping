// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"log"
	"net"

	"github.com/kelseyhightower/ping"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)

const (
	version = "v2.0.0"
)

var (
	listenAddr string
)

type server struct{}

func main() {
	flag.StringVar(&listenAddr, "listen-addr", ":50052", "GRPC listen address")
	flag.Parse()

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	ping.RegisterPingServer(s, &server{})
	reflection.Register(s)

	log.Fatal(s.Serve(ln))
}

func (s *server) Ping(ctx context.Context, in *ping.Request) (*ping.Response, error) {
	p, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("Ping request from %s", p.Addr)
	}
	return &ping.Response{Message: "Pong", Version: version}, nil
}
