package main

import (
    pb "github.com/adanrs/tarea2/booksapp"
    "google.golang.org/grpc"
    "log"
    "net"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    lis, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterBookInfoServer(s, &server{})

    log.Printf("GRPC ON  " + port)

    if err := s.Serve(lis); err != nil {
        log.Fatalf("POS NO SIRVE: %v", err)
    }
}
