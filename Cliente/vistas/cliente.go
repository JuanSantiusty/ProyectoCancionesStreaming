package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	pb "servidor/grpc-servidor/serviciosCancion" // ruta generada por protoc

	"google.golang.org/grpc"
)

func main() {
	// Conectar al servidor gRPC (localhost:50053)
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("No se pudo conectar %v", err)
	}
	defer conn.Close()

	// Crear cliente
	c := pb.NewServiciosCancionesClient(conn)
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
		fmt.Printf("\nEliga un genero:")
		//Se captura el genero
		fmt.Printf("\nIngrese el titulo de la canción a buscar:")
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
			infoCancion(tituloLeido, c, ctx)
		case 2:
			fmt.Printf("\nHasta luego")
			break menuprincipal
		}
	}

}

// Funcion que se encargara de pedir y mostrar la informacion de una cancion
func infoCancion(tituloLeido string, c pb.ServiciosCancionesClient, ctx context.Context) {
	//Se crea un objeto de tipo DTO que contiene el título de la canción a buscar
	objPeticion := &pb.PeticionDTO{Titulo: tituloLeido}

	// Llamada al procedimiento remoto buscarCancion
	res, err := c.BuscarCancion(ctx, objPeticion)
	if err != nil {
		fmt.Printf("Error al pedir canción %v", err)
	}

	// Impresión de la respuesta

	if res.Codigo == 200 {
		fmt.Print(res.Mensaje)
		objcancion := res.ObjCancion
		fmt.Printf("\nTitulo de la cancion: %s", objcancion.Titulo)
		fmt.Printf("\nArtista/Banda: %s", objcancion.Artista_Banda)
		fmt.Printf("\nAño lanzamiento: %d", objcancion.Lanzamiento)
		fmt.Printf("\nDuracion: %s", objcancion.Duracion)
		fmt.Printf("\nGenero musical: %s", objcancion.ObjGenero.Nombre)
		fmt.Printf("\n")
	} else {
		fmt.Printf("%s", res.Mensaje)
	}
}
