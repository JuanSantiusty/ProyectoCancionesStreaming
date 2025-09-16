package CapaAccesoDatos

import (
	mo "servidor/grpc-servidor/modelos"
)

type CancionesRepo struct {
	canciones []mo.Cancion
}

func newCancionesRepo() *CancionesRepo {
	return &CancionesRepo{
		canciones: []mo.Cancion{},
	}
}

func (r *CancionesRepo) agregarCancion(cancion mo.Cancion) {
	r.canciones = append(r.canciones, cancion)
}
