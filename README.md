# ProyectoCancionesStreaming
Proyecto realizado en Go y utilizando grpc de gogle para la reproducción por streaming de canciones

Servidor de canciones:
Intalar dependencias con los comandos:
go get google.golang.org/grpc
go get google.golang.org/protobuf

Comando para los Stub de cliente y servidor:
protoc --go_out=. --go-grpc_out=. servicios.proto
