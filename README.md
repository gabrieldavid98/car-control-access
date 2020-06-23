# car-control-access
Control de acceso para vehículos de un parqueadero

Para correr el proyecto se requiere:

- Mysql 5.7 
- Go 1.14.4

1. Crear una base de datos llamada `access_control` 
2. Agregar el string de conexión al arcivo .env
3. En la carpeta raiz del proyecto ejecutar el comando `go run main.go`
4. Listo :D

## TODO
- Agregar pruebas unitarias
- Agregar api endpoints
- Implemnetar logica para la creacion del respote de pagos