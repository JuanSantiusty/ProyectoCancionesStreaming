package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	inf "cliente/infraestructura"
	pb "servidor/grpc-servidor/serviciosCancion"

	"google.golang.org/grpc"
)

func main() {
	// Conectar al servidor gRPC (localhost:50053)
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("No se pudo conectar %v", err)
	}
	defer conn.Close()

	// Crear
	var servicio inf.ServicioCanciones = inf.NewClienteGRPC(conn)
	// Contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*25)
	defer cancel()
	// lectura
	reader := bufio.NewReader(os.Stdin)
menuprincipal:
	for 1 < 2 {
		fmt.Printf("\nGeneros musicales disponibles")
		fmt.Printf("\n1: Rock")
		fmt.Printf("\n2: Salir")
		fmt.Printf("\nEliga un genero musical:")
		//Se captura el genero
		numG, _ := reader.ReadString('\n')
		numG = strings.TrimSpace(numG)
		numero, err := strconv.Atoi(numG)
		if err != nil {
			fmt.Printf("\nSe debe ingresar un numero valido")
			continue
		}
		switch numero {
		case 1:
			tituloLeido := "lamentoboliviano"
			buscarCancion(tituloLeido, servicio, ctx)
		case 2:
			fmt.Printf("\nHasta luego")
			break menuprincipal
		}
	}

}

// Función que usa la interfaz, recibe titulo de la cancion variable que implemente la interfaz y el contexto
func buscarCancion(titulo string, servicio inf.ServicioCanciones, ctx context.Context) {
	peticion := &pb.PeticionDTO{Titulo: titulo}

	respuesta, err := servicio.BuscarCancion(ctx, peticion)
	if err != nil {
		fmt.Printf("❌ Error al buscar canción: %v\n", err)
		return
	}

	mostrarCancion(respuesta)
}

// Funcion para mostrar informacion de una cancion
func mostrarCancion(respuesta *pb.RespuestaCancionDTO) {
	if respuesta.Codigo == 200 {
		fmt.Printf("✅ %s\n", respuesta.Mensaje)
		cancion := respuesta.ObjCancion
		fmt.Printf("🎵 Título: %s\n", cancion.Titulo)
		fmt.Printf("🎤 Artista/Banda: %s\n", cancion.Artista_Banda)
		fmt.Printf("📅 Año lanzamiento: %d\n", cancion.Lanzamiento)
		fmt.Printf("⏱️ Duración: %s\n", cancion.Duracion)
		fmt.Printf("🎸 Género musical: %s\n", cancion.ObjGenero.Nombre)
	} else {
		fmt.Printf("❌ %s\n", respuesta.Mensaje)
	}
}
