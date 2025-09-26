# ProyectoCancionesStreaming
Proyecto realizado en Go y utilizando grpc de gogle para la reproducción por streaming de canciones

Comando para instalar compilador de Protocol Buffers:
apt install -y protobuf-compiler

Comandos para instalar los plugins de gRPC para Go:
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

Comandos para agregar $GOPATH/bin al PATH con los comandos;
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc

Para poder utilizar gRPC ejecutamos los siguientes comandos que instalar las dependencias necesarias para utilizar gRPC (Ejecutamos en servidor canciones y streaming):
go get google.golang.org/grpc
go get google.golang.org/protobuf

Para el cliente ejecutamos comandos para forzar la instalacion de algunas dependencia y importar las librerias de faiface/beep para trabajar con audio:
go mod tidy
go get github.com/faiface/beep
go get github.com/faiface/beep/mp3
go get github.com/faiface/beep/speaker

Comando para generar los Stub de cliente y servidor:
protoc --go_out=. --go-grpc_out=. servicios.proto

Puertos:
El Servidor de canciones utiliza el puerto 50053
El servidor de streaming utiliza el puerto 50051