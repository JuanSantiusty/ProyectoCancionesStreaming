package capaFachadaServices

import (
	"reflect"
	rp "servidor/grpc-servidor/CapaAccesoDatos"
	mo "servidor/grpc-servidor/CapaAccesoDatos/modelos"
)

func CargarCanciones(r *rp.CancionesRepo) {

}

func BuscarCancion(titulo string, r *rp.CancionesRepo) mo.RespuestaDTO[mo.Cancion] {
	objvacio := mo.Cancion{}
	objCancion := r.BuscarPorTitulo(titulo)
	if !reflect.DeepEqual(objvacio, objCancion) {
		return mo.RespuestaDTO[mo.Cancion]{
			Data:    objCancion,
			Codigo:  200,
			Mensaje: "Canción encontrada",
		}
	} else {
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
