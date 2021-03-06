package main

import (
    "context"
    pb "github.com/adanrs/tarea2/booksapp"
    "github.com/gofrs/uuid"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "fmt"
    "encoding/csv"
    //"encoding/json"
    //"io"
    "os"
    "log"
)

type server struct {
    bookMap map[string]*pb.Book
}

func (s *server) AddBook(ctx context.Context, in *pb.Book) (*pb.BookID, error) {
    //out, err := uuid.NewV4()
    _, err := uuid.NewV4()
    if err != nil {
        return nil, status.Errorf(codes.Internal,
            "Error while generating Book ID", err)
    }
    //in.Id = out.String()
    if s.bookMap == nil {
        s.bookMap = make(map[string]*pb.Book)
    }
    s.bookMap[in.Id] = in
    return &pb.BookID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) GetBook(ctx context.Context, in *pb.BookID) (*pb.Book, error) {
    value, exists := s.bookMap[in.Value]
    if exists {
        return value, status.New(codes.OK, "").Err()
    }
    return nil, status.Errorf(codes.NotFound, "Book does not exist.", in.Value)
}

func (s *server) DeleteBook(ctx context.Context, in *pb.BookID) (*pb.BookID, error) {
    _, exists := s.bookMap[in.Value]
    if exists {
	delete(s.bookMap, in.Value)
	return &pb.BookID{Value: in.Value}, status.New(codes.OK, "").Err()
    }
    return nil, status.Errorf(codes.NotFound, "Book does not exist.", in.Value)
}

func (s *server) UpdateBook(ctx context.Context, in *pb.Book) (*pb.BookID, error) {
 
    if s.bookMap == nil {
        s.bookMap = make(map[string]*pb.Book)
    }
    s.bookMap[in.Id] = in
    return &pb.BookID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) ReadCSV(ctx context.Context, in *pb.File) (*pb.BookID, error)  {
    f, _ := os.Open(in.Value)
   bookline, err := csv.NewReader(f).ReadAll()
   if err != nil {
      log.Fatal(err)
   }
   if s.bookMap == nil {
        s.bookMap = make(map[string]*pb.Book)
    }


   // Ciclo a traves de todos los registros, crear libro por cada uno
   for _, book := range bookline {
      fmt.Printf("\n\nLeyendo y creando libro con ID: ", book[0])
      newBook, err := s.AddBook(ctx, &pb.Book{
        Id:        book[0],
        Title:     book[1],
        Edition:   book[2],
        Copyright: book[3],
        Language:  book[4],
        Pages:     book[5],
        Author:    book[6],
        Publisher: book[7]})

	if err != nil {
		log.Fatalf("\n\n\nNo se pudo agregar el libro: %v", err)
	}
	fmt.Printf("\n\n\nLibro agregado: ", newBook.String())

	bookGet, err := s.GetBook(ctx, &pb.BookID{Value: book[0]})
	if err != nil {
		log.Printf("\n\n\nEl libro consultado no existe: %v", err)
	} else {
		log.Printf("\n\n\nVerificando insert:\n\n", bookGet.String())
	}
   }

    return &pb.BookID{Value: "1"}, status.New(codes.OK, "").Err()
}
