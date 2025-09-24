package capaFachadaServices

import (
	"fmt"
	"reflect"
	rp "servidor/grpc-servidor/CapaAccesoDatos"
	mo "servidor/grpc-servidor/CapaAccesoDatos/modelos"
)

func CargarCanciones(r *rp.CancionesRepo) {
	r.AgregarCancion(mo.NewCancion(1, "lamentoboliviano", "Enanitos Verdes", 2004, "3:10", mo.NewGenero(1, "rock")))
	fmt.Print("Canciones cargadas")
}

func BuscarCancion(titulo string, r *rp.CancionesRepo) mo.RespuestaDTO[mo.Cancion] {
	objvacio := mo.Cancion{}
	objCancion := r.BuscarPorTitulo(titulo)
	if !reflect.DeepEqual(objvacio, objCancion) {
		fmt.Print("Cancione encontrada")
		return mo.RespuestaDTO[mo.Cancion]{
			Data:    objCancion,
			Codigo:  200,
			Mensaje: "Canción encontrada",
		}
	} else {
		fmt.Print("Cancione no encontrada")
		return mo.RespuestaDTO[mo.Cancion]{
			Data:    objCancion,
			Codigo:  400,
			Mensaje: "Cancion no encontrada",
		}
	}
}

// Implementacion para buscar canciones por genero, por implementar
func BuscarPorGenero(titulo string, r *rp.CancionesRepo) {
	// implementacion
}
