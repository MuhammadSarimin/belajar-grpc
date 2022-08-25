package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/MuhammadSarimin/belajar-grpc.git/student"
	"google.golang.org/grpc"
)

func get(c student.DataStudentClient, e string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := student.Student{Email: e}

	st, err := c.FindStudentByEmail(ctx, &s)
	if err != nil {
		log.Fatalln("error get student", err.Error())
		return
	}

	js, _ := json.Marshal(st)
	log.Println(string(js))
}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	con, err := grpc.Dial("9000", opts...)

	if err != nil {
		log.Fatalln("failed dial", err.Error())
	}

	defer con.Close()

	client := student.NewDataStudentClient(con)

	get(client, "sarimin@google.com")
	get(client, "msarimin@google.com")
	get(client, "msarimin@google")
}
