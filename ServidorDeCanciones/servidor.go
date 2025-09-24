package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	rp "servidor/grpc-servidor/CapaAccesoDatos"
	se "servidor/grpc-servidor/capaFachadaServices"
	pb "servidor/grpc-servidor/serviciosCancion"
)

var repo = rp.NewCancionesRepo()

// Estructura que implementa el servicio
type servidorCanciones struct {
	pb.UnimplementedServiciosCancionesServer
}

// Implementación del método buscarCancion
func (s *servidorCanciones) BuscarCancion(ctx context.Context, req *pb.PeticionDTO) (*pb.RespuestaCancionDTO, error) {

	titulo := req.GetTitulo()
	fmt.Printf("%s", titulo)
	fmt.Print("Buscando Cancion")
	resp := se.BuscarCancion(titulo, repo)
	fmt.Print("Terminar de buscar cancion")

	var respuesta pb.RespuestaCancionDTO
	respuesta.Codigo = resp.Codigo
	respuesta.Mensaje = resp.Mensaje

	if resp.Codigo == 200 {
		respuesta.ObjCancion = new(pb.Cancion)
		respuesta.ObjCancion.Id = resp.Data.Id
		respuesta.ObjCancion.Titulo = resp.Data.Titulo
		respuesta.ObjCancion.Artista_Banda = resp.Data.Artista_Banda
		respuesta.ObjCancion.Lanzamiento = resp.Data.Lanzamiento
		respuesta.ObjCancion.Duracion = resp.Data.Duracion
		respuesta.ObjCancion.ObjGenero = new(pb.Genero)
		respuesta.ObjCancion.ObjGenero.Id = resp.Data.Genero.Id
		respuesta.ObjCancion.ObjGenero.Nombre = resp.Data.Genero.Nombre
	}

	return &respuesta, nil
}

func main() {

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Error al abrir el puerto: %v", err)
	}

	// Crear servidor gRPC
	grpcServer := grpc.NewServer()

	// Registrar el servicio
	pb.RegisterServiciosCancionesServer(grpcServer, &servidorCanciones{})

	// Cargar canciones
	se.CargarCanciones(repo)

	// Iniciar el servidor
	log.Println("Servidor gRPC escuchando en puerto 50053...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
