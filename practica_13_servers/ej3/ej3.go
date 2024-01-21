package main

import (
	"fmt"
)

func Ej3() {
	// Desarrollar una API en Go que maneje saludos personalizados y mantenga un registro de los nombres ya saludados. La API deberá ofrecer respuestas específicas basadas en si el nombre ha sido previamente saludado o no.

	// ### Endpoints de la API

	// 1. **POST /hello**:
	// 		- Funcionalidad: Saluda al nombre proporcionado en el cuerpo de la solicitud.
	// 		- Comportamiento: Si el nombre ya ha sido saludado antes, la API debe ofrecer una respuesta de "hola de nuevo" en lugar de un simple saludo.
	// 2. **GET /hello**:
	// 		- Funcionalidad: Devuelve una lista de todos los nombres que han sido saludados.

	// ### Requisitos Específicos

	// 1. **Manejo de Estado**:
	// 		- La API debe ser capaz de recordar los nombres que ya ha saludado. Considera el uso de una estructura de datos adecuada para almacenar estos nombres entre solicitudes.
	// 2. **Validaciones**:
	// 		- Asegúrate de que la API maneje adecuadamente las entradas inválidas y proporcione mensajes de error claros.
	// 3. **Paquetes**
	// 		- Se pide intencionalmente que NO se escriban las rutinas de código dentro del proyecto del servidor, sino que se importen usando los mecanismo de dependencias que tiene el lenguaje.
	// 💡 Ten especial cuidado con las condiciones de carrera al acceder o modificar el registro de nombres. Asegúrate de que las operaciones sobre la estructura de datos sean seguras.
	fmt.Print("EJERCICIO 3\n\n")
}
