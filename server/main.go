package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"github.com/MuhammadSarimin/belajar-grpc.git/student"
	"google.golang.org/grpc"
)

type dataStudent struct {
	student.UnimplementedDataStudentServer
	mu       sync.Mutex
	students []*student.Student
}

func (d *dataStudent) FindStudentByEmail(c context.Context, s *student.Student) (*student.Student, error) {
	return nil, nil
}

func (d *dataStudent) load() {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatalln("failed load data", err.Error())
	}

	if err := json.Unmarshal(data, &d.students); err != nil {
		log.Fatalln("failed unmarshall", err.Error())
	}
}

func newServer() *dataStudent {
	s := dataStudent{}
	s.load()
	return &s
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln("failed listen to server", err.Error())
	}

	grpcServer := grpc.NewServer()
	student.RegisterDataStudentServer(grpcServer, newServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("error server grpc", err.Error())
	}
}
