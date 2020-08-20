package main

import (
    "context"
    pb "github.com/adanrs/tarea2/booksapp"
    "google.golang.org/grpc"
    "log"
    "os"
    "time"
)

func main() {
    address := os.Getenv("ADDRESS")
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewBookInfoClient(conn)

    // Agregar libro
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.AddBook(ctx, &pb.Book{
        Id:        "1",
        Title:     "Operating System Concepts",
        Edition:   "9th",
        Copyright: "2012",
        Language:  "ENGLISH",
        Pages:     "976",
        Author:    "Abraham Silberschatz",
        Publisher: "John Wiley & Sons"})
    if err != nil {
        log.Fatalf("\n\n\nImposible de agregar libro: %v", err)
    }

    log.Printf("\n\n\nLibro creado con el ID: %s", r.Value)
    book, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
    if err != nil {
        log.Fatalf("\n\n\nLibro inexiste: %v", err)
    }
    log.Printf("\n\n\nConsulta Realizada: ", book.String())

    bookDel, err := c.DeleteBook(ctx, &pb.BookID{Value: r.Value})
    if err != nil {
        log.Fatalf("\n\n\nNo se pudo eliminar el libro: %v", err)
    }
    log.Printf("\n\n\nÑLibro Eliminado: ", bookDel.String())

    
    bookGet, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
    if err != nil {
        //log.Fatalf("\n\nEl libro consultado no existe: %v", err)
        log.Printf("\n\n\nEl libro consultado no existe: %v", err)
    } else {
	    log.Printf("\n\n\nLibro consultado: ", bookGet.String())
    }

    upd, err := c.UpdateBook(ctx, &pb.Book{
        Id:        "1",
        Title:     "RENOVADO",
        Edition:   "19th",
        Copyright: "2020",
        Language:  "ESPAÑOL",
        Pages:     "976",
        Author:    "Abraham Silberschatz",
        Publisher: "John Wiley & Sons"})
    if err != nil {
        log.Fatalf("\n\n\nImposible de actualizar el libro: %v", err)
    } else {
	    bookGetUpdated, err := c.GetBook(ctx, &pb.BookID{Value: upd.Value})
	    if err != nil {
		    log.Fatalf("\n\n\nEl libro consultado no existe: %v", err)
	    }
	    log.Printf("\n\n\nSe Actualizo Correctamente: ", bookGetUpdated.String())
    }
    c.ReadCSV(ctx,&pb.File{Value: "books.csv"})
    log.Printf("\n\n\nVERIFICANDO VERSION CSV")

}
