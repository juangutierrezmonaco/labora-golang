package practica

import (
	"fmt"
	"sort"
	"strings"

	"github.com/labora/labora-golang/util"
)

// Person struct & methods
type Person struct {
	Name   string
	Age    int
	Height float64
	Weight float64
}

func (p *Person) enterData() error {
	var err error

	fmt.Printf("Ingrese el nombre: ")
	p.Name = util.ScanSentence()

	fmt.Printf("Ingrese la edad: ")
	ageInterface, err := util.ScanValue(new(int))
	if err != nil {
		return fmt.Errorf("Debe ingresar un número entero para la edad.")
	}
	p.Age = ageInterface.(int)
	if p.Age <= 0 {
		return fmt.Errorf("Debe ingresar un número entero, mayor a 0 para la edad.")
	}

	fmt.Printf("Ingrese la altura (en cm): ")
	heightInterface, err := util.ScanValue(new(float64))
	if err != nil {
		return fmt.Errorf("Debe ingresar un número para la altura.")
	}
	p.Height = heightInterface.(float64)
	if p.Height <= 0 {
		return fmt.Errorf("Debe ingresar un número mayor a 0 para la altura.")
	}

	fmt.Printf("Ingrese el peso (en kg): ")
	weightInterface, err := util.ScanValue(new(float64))
	if err != nil {
		return fmt.Errorf("Error al ingresar el peso:")
	}
	p.Weight = weightInterface.(float64)
	if p.Weight <= 0 {
		return fmt.Errorf("Debe ingresar un número mayor a 0 para el peso.")
	}

	return nil
}

func (p *Person) print() {
	fmt.Printf("Nombre: %s\n", p.Name)
	fmt.Printf("Edad: %d años\n", p.Age)
	fmt.Printf("Altura: %.2f cm\n", p.Height)
	fmt.Printf("Peso: %.2f kg\n", p.Weight)
	fmt.Printf("IMC: %.2f kg/m2 (%s)\n", p.calcBmi(), getBmiMessage(p))
}

func (p *Person) calcBmi() float64 {
	heightInMeters := p.Height / 100
	return p.Weight / (heightInMeters * heightInMeters)
}

func getBmiMessage(p *Person) string {
	bmi := p.calcBmi()
	switch true {
	case bmi <= 18.5:
		return "Bajo peso"
	case bmi > 18.5 && bmi <= 24.9:
		return "Peso normal"
	case bmi > 24.9 && bmi <= 29.9:
		return "Sobrepeso"
	case bmi > 29.9:
		return "Obesidad"
	}
	return ""
}

// Person array struct & methods
type Persons []Person

func (persons *Persons) enter(quantity int) {
	var err error
	for i := 0; i < quantity; i++ {
		fmt.Printf("Ingreso de datos de %d personas.\n", quantity)
		fmt.Printf("Complete los datos de la persona %d.\n\n", i+1)
		var p Person
		err = p.enterData()
		if err != nil {
			fmt.Println()
			fmt.Println(err.Error())
			util.WaitForEnter()			

			util.ClearConsole()
			i--
			continue
		}
		*persons = append(*persons, p)
		util.ClearConsole()
	}
	fmt.Printf("Se registraron correctamente %d personas.\n\n", quantity)
}

func (persons *Persons) find(findBy int, value string) Persons {
	var personsFound Persons = nil

	// Map with functions to sort
	searchByFunc := map[int]func(p Person) bool{
		1: func(p Person) bool {

			return strings.Contains(strings.ToUpper(p.Name), strings.ToUpper(value))
		},
		2: func(p Person) bool {
			return util.Itoa(p.Age) == value
		},
		3: func(p Person) bool {
			return util.Itoa(int(p.Height)) == value
		},
		4: func(p Person) bool {
			return util.Itoa(int(p.Age)) == value
		},
	}

	// Search
	searchFunc, funcExists := searchByFunc[findBy]
	if funcExists {
		for _, p := range *persons {
			if searchFunc(p) {
				personsFound = append(personsFound, p)
			}
		}
	}

	return personsFound
}

func (persons *Persons) print() {
	if len(*persons) == 0 {
		fmt.Println("No hay personas para mostrar.")
	} else {
		for _, p := range *persons {
			p.print()
			fmt.Println()
		}
	}
}

func (persons *Persons) sort(sortBy int) Persons {
	// Make a copy
	personsCopy := make(Persons, len(*persons))
	copy(personsCopy, *persons)

	// Map with functions to sort
	sortByFunc := map[int]func(p1, p2 Person) bool{
		1: func(p1, p2 Person) bool {
			return p1.Name < p2.Name
		},
		2: func(p1, p2 Person) bool {
			return p1.Age < p2.Age
		},
		3: func(p1, p2 Person) bool {
			return p1.Height < p2.Height
		},
		4: func(p1, p2 Person) bool {
			return p1.Weight < p2.Weight
		},
	}

	// Sort
	sortFunc, funcExists := sortByFunc[sortBy]
	if funcExists {
		sort.Slice(personsCopy, func(i, j int) bool {
			return sortFunc(personsCopy[i], personsCopy[j])
		})
	}

	return personsCopy
}

// Menu funcs
func sortPersons(persons *Persons) Persons {
	// Get the sort criteria
	options := []string{
		"Nombre.",
		"Edad.",
		"Altura.",
		"Peso.",
	}

	menu := util.NewMenu(nil, options, "Ordenar por: ")
	op := menu.Run()

	// Sort the array
	return persons.sort(op)
}

func findPersons(persons *Persons) Persons {
	// Get the search criteria
	options := []string{
		"Nombre.",
		"Edad.",
		"Altura.",
		"Peso.",
	}

	menu := util.NewMenu(nil, options, "Buscar por: ")
	op := menu.Run()

	// Get the value to find
	fmt.Print("Ingrese el valor a buscar: ")
	value := util.ScanSentence()

	// Search
	return persons.find(op, value)
}

func EjPairProgramming() {
	// Descripción del problema
	// Escribir un programa que declare una estructura llamada Persona que tenga los siguientes campos: nombre (cadena), edad (entero), altura (float) y peso (float).

	// A continuación, el programa deberá permitir el ingreso de los datos de 5 personas por teclado, en el orden que se desee.
	// Crear una función llamada ordenarPersonas que ordene las personas de acuerdo a un criterio de ordenamiento elegido por el usuario.
	// El criterio de ordenamiento puede ser por nombre, edad, altura o peso, y el usuario debe seleccionarlo ingresando un número del 1 al 4.
	// La función deberá devolver el slice de personas ordenado según el criterio elegido.

	// Luego, crear otra función llamada buscarPersona que reciba el slice de personas y un valor a buscar, y devuelva la persona encontrada o nil si no se encuentra.

	// Deberás mostrar en pantalla: las personas creadas, sin ordenar, las personas ordenadas según el criterio elegido por el usuario, y el resultado de una búsqueda.

	// Además, se deberán aplicar las siguientes consideraciones:
	// El programa debe verificar que los datos ingresados sean válidos y no permitir la creación de personas con valores incorrectos. Por ejemplo, no se
	// permitirá ingresar una edad negativa o una altura igual a cero.
	// La función de búsqueda debe permitir buscar por cualquier campo de la estructura Persona, no solo por nombre.
	// Se deberá implementar una función que calcule el índice de masa corporal (IMC) de una persona, que se define como el peso de la persona dividido por el
	// cuadrado de su altura (expresado en kg/m2).

	// Al mostrar las personas, se deberá incluir su índice de masa corporal (IMC) y clasificarlo en las siguientes categorías:
	// Bajo peso: IMC menor a 18.5
	// Peso normal: IMC entre 18.5 y 24.9
	// Sobrepeso: IMC entre 25 y 29.9
	// Obesidad: IMC mayor a 30

	// El programa deberá permitir al usuario realizar la opción de ordenar o buscar varias veces, sin tener que volver a ingresar los datos de las personas.
	// Ejemplo de menú del programa:

	// ======= MENÚ =======
	// 1. Registrar personas
	// 2. Buscar persona
	// 3. Listar personas
	// 4. Ordenar lista de personas
	// 5. Salir
	// Insira la opción:
	fmt.Print("EJERCICIO 3\n\n")

	// Declaring necessary variables
	var persons Persons
	enterPersonsFunc := func() {
		persons.enter(1)
	}
	findPersonsFunc := func() {
		res := findPersons(&persons)
		fmt.Println(res)
		fmt.Print("Los resultados de la búsqueda son: \n\n")
		res.print()
	}
	printPersonsFunc := func() {
		persons.print()
	}
	sortPersonsFunc := func() {
		persons = sortPersons(&persons)
		fmt.Println("Personas ordenadas correctamente.")
	}

	// Running the menu
	functions := []func(){
		enterPersonsFunc,
		findPersonsFunc,
		printPersonsFunc,
		sortPersonsFunc,
	}

	options := []string{
		"Registrar personas.",
		"Buscar persona.",
		"Listar personas.",
		"Ordenar lista de personas.",
	}

	menu := util.NewMenu(functions, options, "MENÚ")
	menu.Run()
}
