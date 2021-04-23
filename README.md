# ticket-crud-golang
Esta es una prueba tecnica para la compañia Double V Partners.

## Planteamiento de la prueba:
Queremos un API que nos permita crear, eliminar, editar y recuperar tickets. Que se pueda recuperar todos o filtrar por uno específico.
Los ticket tienen un id, un usuario, una fecha de creación, una fecha de actualización y un estatus (abierto/cerrado). Debes usar C# (Netocore) o Golang.

## Tecnologias:
Se usó el lenguaje Golang para la solución del problema y para la base de datos, se usó PostgreSQL.
Para evitar problemas de dependencias se optó por usar Docker para crear un contenedor de la aplicación.

## Aclaraciones:
Se dejó el archivo `.env` en el repositorio por cuestiones de la prueba.

## Iniciar proyecto:
Como es un proyecto dockerizado, se requiere instanciar los contenedores con el siguiente comando: `docker-compose -up -d --build`.
