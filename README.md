# ProyectoCancionesStreaming
Proyecto realizado en Go y utilizando grpc de gogle para la reproducción por streaming de canciones

Servidor de canciones y streaming :
Intalar dependencias con los comandos:
go get google.golang.org/grpc
go get google.golang.org/protobuf

Cliente:
go mod tidy

Comando para los Stub de cliente y servidor:
protoc --go_out=. --go-grpc_out=. servicios.proto

Puertos:
El Servidor de canciones utiliza el puerto 50053
El servidor de streaming utiliza el puerto 50051