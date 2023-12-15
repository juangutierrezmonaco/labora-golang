package practica

import (
	"fmt"
	"strings"

	"github.com/labora/labora-golang/util"
)

// Grade
type Grades map[string]float64 // courseName -> note

func (gg *Grades) alreadyExists(courseName string) bool {
	_, exists := (*gg)[courseName]
	return exists
}

func (gg *Grades) avg() float64 {
	var acc float64
	for _, v := range *gg {
		acc += v
	}
	return acc / float64(len(*gg))
}

func (gg *Grades) toString() string {
	var str string
	if len(*gg) == 0 {
		str = "No hay materias cargadas."
		return str
	}

	underline := strings.Repeat("-", 30)
	str += fmt.Sprintf("%-30s\t\t%-30s\n", "MATERIA", "NOTA")
	str += fmt.Sprintf("%-30s\t\t%-30s\n", underline, underline)
	for courseName, note := range *gg {
		str += fmt.Sprintf("%-30s\t\t%-30.2f\n", util.CapitalizeWords(courseName), note)
	}
	str += fmt.Sprintf("\nPromedio: %.2f", gg.avg())
	return str
}

func (gg *Grades) scan() error {
	fmt.Print("Ingrese cuántas materias desea cargar: ")
	n, err := strToInt(util.ScanSentence())
	if err != nil {
		return fmt.Errorf("Ingresó una cantidad de materias no numérica.")
	}

	for i := 0; i < n; i++ {
		fmt.Print("Nombre de la materia: ")
		name := util.ScanSentence()
		nameUppercase := strings.ToUpper(name)

		fmt.Printf("Nota de la materia [%s]: ", nameUppercase)
		note, err := util.StrToFloat(util.ScanSentence())
		if err != nil {
			return fmt.Errorf("Ingresó una nota no numérica.")
		}

		err = gg.Add(name, note)
		if err != nil {
			return err
		}
	}

	return nil
}

func (gg *Grades) update() error {
	fmt.Println(gg.toString())
	fmt.Print("\nIngrese el nombre de la materia que desea modificar: ")
	name := util.ScanSentence()

	if !gg.alreadyExists(name) {
		return fmt.Errorf("No existe la materia que está tratando de modificar.")
	}

	var err error
	functions := []func(){
		func() {
			fmt.Printf("Ingrese la nota para la materia [%s]: ", strings.ToUpper(name))
			note, err := util.StrToFloat(util.ScanSentence())
			if err != nil {
				err = fmt.Errorf("Ingresó una nota no numérica.")
				return
			}
			err = gg.Update(name, note)
		},
		func() {
			err = gg.Delete(name)
		},
	}

	options := []string{
		"Modificar",
		"Eliminar",
	}

	title := fmt.Sprintf("¿Desea modificar o eliminar la materia %s?", strings.ToUpper(name))
	m := util.NewMenu(functions, options, title)
	m.Run()

	return err
}

func (gg *Grades) Get(courseName string) (float64, error) {
	note, exists := (*gg)[courseName]
	if !exists {
		return -1, fmt.Errorf("No existe la materia \"%s\"", courseName)
	}
	return note, nil
}

func (gg *Grades) Add(courseName string, note float64) error {
	if gg.alreadyExists(courseName) {
		return fmt.Errorf("No se puede agregar la materia \"%s\" porque ya existe.", courseName)
	}
	if note < 0 || note > 10 {
		return fmt.Errorf("La nota debe estar entre 0 y 10.")
	}
	(*gg)[courseName] = note
	return nil
}

func (gg *Grades) Update(courseName string, note float64) error {
	if !gg.alreadyExists(courseName) {
		return fmt.Errorf("No se puede actualizar la materia \"%s\" porque no existe.", courseName)
	}
	(*gg)[courseName] = note
	return nil
}

func (gg *Grades) Delete(courseName string) error {
	if !gg.alreadyExists(courseName) {
		return fmt.Errorf("No se puede eliminar la materia \"%s\" porque no existe.", courseName)
	}
	delete(*gg, courseName)
	return nil
}

// Student
type Student struct {
	Name string
	Id   int
	Grades
}

var lastId int

func generateId() int {
	lastId++
	return lastId
}

func (s *Student) gpa() float64 { // grade point average
	return s.Grades.avg()
}

func (s *Student) toString() string {
	var str string
	title := "DATOS DEL ESTUDIANTE # " + fmt.Sprint(s.Id)
	underline := strings.Repeat("-", len(title))
	str += fmt.Sprintf("%-30s\t\n", title)
	str += fmt.Sprintf("%-30s\t\n", underline)

	str += fmt.Sprintf("Nombre: %s\t\n", s.Name)
	str += fmt.Sprintf("\n%s", s.Grades.toString())

	return str
}

func (s *Student) scan() error {
	// Scan student
	fmt.Printf("Ingreso de datos del alumno # %d\n\n", s.Id)
	fmt.Print("Ingrese nombre: ")
	s.Name = util.ScanSentence()

	// Scan grades
	fmt.Printf("\n\n")
	confirmAction := util.ConfirmAction("¿Quiere cargarle las notas al alumno ahora?")

	var err error
	if confirmAction {
		fmt.Println()
		err = s.Grades.scan()
		if err != nil {
			return err
		}
	}
	fmt.Printf("\n\nSe ha cargado correctamente el siguiente alumno:\n\n%s\n", s.toString())
	return nil
}

func (s *Student) update() error {
	fmt.Print("Ingrese el nombre del alumno: ")
	name := util.ScanSentence()

	s.Name = name
	fmt.Printf("\n\nSe ha modificado correctamente el alumno y quedó así:\n\n%s\n", s.toString())

	return nil
}

func (s *Student) GetGrade(name string) (float64, error) {
	return s.Grades.Get(name)
}

func (s *Student) AddGrade(name string, note float64) error {
	return s.Grades.Add(name, note)
}

func (s *Student) UpdateGrade(name string, note float64) error {
	return s.Grades.Update(name, note)
}

func (s *Student) DeleteGrade(name string) error {
	return s.Grades.Delete(name)
}

func (s *Student) AddGrades(grades *Grades) error {
	for name, note := range *grades {
		err := s.AddGrade(name, note)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewStudent(name string) *Student {
	return &Student{
		Name:   name,
		Id:     generateId(),
		Grades: Grades{},
	}
}

// Students
type Students map[int]Student

func (ss *Students) alreadyExists(id int) bool {
	_, exists := (*ss)[id]
	return exists
}

func (ss *Students) getGpaOfStudent(id int) float64 {
	student := (*ss)[id]
	return student.gpa()
}

func (ss *Students) toString() string {
	var str string
	title := util.TextWithFiller("-", 15, "LISTADO DE ALUMNOS")
	char := "/"
	underline := strings.Repeat(char, len(title))
	str += char + underline + char + "\n"
	str += char + title + char + "\n"
	str += char + underline + char + "\n\n"

	for _, s := range *ss {
		str += s.toString() + "\n\n"
	}

	return str
}

func (ss *Students) Get(id int) (*Student, error) {
	if !ss.alreadyExists(id) {
		return nil, fmt.Errorf("No existe el alumno con ID=%d.", id)
	}
	s := (*ss)[id]
	return &s, nil
}

func (ss *Students) Add(s *Student) error {
	if ss.alreadyExists(s.Id) {
		return fmt.Errorf("No se puede agregar al alumno con ID=%d porque ya existe.", s.Id)
	}
	(*ss)[s.Id] = *s
	return nil
}

func (ss *Students) Update(id int, name string) error {
	if !ss.alreadyExists(id) {
		return fmt.Errorf("No se puede actualizar al alumno con ID=%d porque no existe.", id)
	}
	s, err := ss.Get(id)
	if err != nil {
		return err
	}

	s.Name = name
	return nil
}

func (ss *Students) Delete(id int) error {
	if ss.alreadyExists(id) {
		return fmt.Errorf("No se puede eliminar al alumno con ID=%d porque no existe.", id)
	}
	delete(*ss, id)
	return nil
}

func (ss *Students) GetGradeOfStudent(id int, courseName string) (float64, error) {
	if !ss.alreadyExists(id) {
		return -1, fmt.Errorf("No se pueden mostrar las notas del alumno con ID=%d porque no existe.", id)
	}
	student := (*ss)[id]
	return student.GetGrade(courseName)
}

func (ss *Students) AddGradeToStudent(id int, courseName string, note float64) error {
	if !ss.alreadyExists(id) {
		return fmt.Errorf("No se pueden actualizar las notas del alumno con ID=%d porque no existe.", id)
	}
	student := (*ss)[id]
	student.AddGrade(courseName, note)
	return nil
}

func (ss *Students) UpdateGradeToStudent(id int, courseName string, note float64) error {
	if !ss.alreadyExists(id) {
		return fmt.Errorf("No se pueden actualizar las notas del alumno con ID=%d porque no existe.", id)
	}
	student := (*ss)[id]
	student.UpdateGrade(courseName, note)
	return nil
}

func (ss *Students) DeleteGradeToStudent(id int, courseName string) error {
	if !ss.alreadyExists(id) {
		return fmt.Errorf("No se pueden eliminar las notas del alumno con ID=%d porque no existe.", id)
	}
	student := (*ss)[id]
	student.DeleteGrade(courseName)
	return nil
}

func (ss *Students) AddGradesToStudent(id int, grades *Grades) error {
	if !ss.alreadyExists(id) {
		return fmt.Errorf("No se pueden actualizar las notas del alumno con ID=%d porque no existe.", id)
	}
	student := (*ss)[id]
	student.AddGrades(grades)
	return nil
}

// Exercise functions
func viewAllStudents(ss *Students) {
	fmt.Println(ss.toString())
}

func addStudent(ss *Students) {
	util.ClearConsole()
	s := NewStudent("")
	err := s.scan()
	if err != nil {
		lastId-- // If it wasn't created decrement the id created
		panic(err)
	}
	err = ss.Add(s)
	panic(err)
}

func updateStudent(ss *Students) {
	fmt.Print("Ingrese el ID del alumno a modificar: ") // Buscar por nombre
	id, err := strToInt(util.ScanSentence())
	if err != nil {
		panic(fmt.Errorf("Debe ingresar un ID numérico"))
	}

	s, err := ss.Get(id)
	if err != nil {
		panic(err)
	}

	functions := []func(){
		func() {
			err := s.update()
			if err != nil {
				panic(err)
			}
			(*ss)[id] = *s
		},
		func() {
			err := s.Grades.update()
			if err != nil {
				panic(err)
			}
			(*ss)[id] = *s
		},
	}

	options := []string{
		"Datos del Alumno",
		"Actualizar/Eliminar Materias",
	}

	title := fmt.Sprintf("Ingrese qué desea modificar del alumno de ID=%d", id)
	m := util.NewMenu(functions, options, title)
	m.Run()
}

func run() {
	students := initialMockDate()
	functions := []func(){
		func() { viewAllStudents(students) },
		func() { addStudent(students) },
		func() { updateStudent(students) },
	}

	options := []string{
		"VER TODOS LOS ALUMNOS.",
		"AGREGAR UN ALUMNO.",
		"MODIFICAR UN ALUMNO.",
		// "ELIMINAR UN ALUMNO.",
		//"ORDENAR ALUMNOS.",
		//"BUSCAR UN ALUMNO.",
	}

	menu := util.NewMenu(functions, options, "SISTEMA DE ALUMNOS")
	menu.Run()
}

func initialMockDate() *Students {
	grades1 := Grades{
		"Lengua":                 2,
		"Matemática":             5,
		"Algoritmos del flexman": 2,
	}
	student1 := NewStudent("Juan")
	student2 := NewStudent("Ana")

	students := Students{}
	students.Add(student1)
	students.Add(student2)

	students.AddGradesToStudent(student1.Id, &grades1)
	students.AddGradesToStudent(student2.Id, &grades1)

	return &students
}

func Ej5() {
	// Objetivo: Crear un programa en Go que gestione las calificaciones de los estudiantes utilizando mapas. y que maneje errores comunes y excepciones.

	// Descripción:

	// 1. Estructura del Estudiante:
	//     - Define una estructura `Estudiante` que tenga al menos dos campos: `Nombre` (string) y `ID` (int) y `calificaciones` (mapa)
	// 			 con llave siendo la materia (string) y el valor siendo la nota (float)
	// 2. Funciones para Manejar Calificaciones:
	//     - Escribe funciones para añadir un nuevo estudiante y sus calificaciones.
	//     - Implementa una función para actualizar las calificaciones de un estudiante existente.
	//     - Añade una función para eliminar las calificaciones de un estudiante.
	//     - Crea una función para calcular el promedio de calificaciones de un estudiante.
	// 3. Interacción con el Usuario:
	//     - Permite que el usuario agregue, actualice, elimine y vea las calificaciones de los estudiantes a través de la entrada del terminal.
	// 4. Manejo de Errores en la Entrada de Datos:
	//     - Al añadir o actualizar las calificaciones, asegúrate de que son números flotantes válidos. Si no lo son, muestra un mensaje de error.
	//     - Al seleccionar un ID de estudiante, verifica que sea un número entero válido.
	//     - Asegúrate de manejar posibles errores, como intentar actualizar las calificaciones de un estudiante que no existe.
	// 5. Funciones con Manejo de Errores:
	//     - Modifica las funciones para añadir, actualizar, y eliminar estudiantes para que manejen situaciones como intentar actualizar un estudiante que no existe.
	//     - Cada función debe retornar un error si ocurre una situación excepcional.
	// 6. Implementación de la Interfaz de Usuario:
	//     - Asegúrate de que el usuario reciba retroalimentación clara si algo va mal (por ejemplo, si intenta añadir un estudiante con un ID duplicado).

	// Instrucciones Adicionales:

	// - Puedes mejorar la interacción con el usuario implementando un menú simple en la terminal.
	// - Utiliza ciclos y declaraciones condicionales para validar la entrada del usuario.
	// - Practica escribir mensajes de error claros y útiles que ayuden al usuario a entender qué salió mal.
	// - Considera todos los casos posibles donde algo podría ir mal y cómo tu programa debería responder en esos casos.
	fmt.Print("EJERCICIO 5\n\n")

	defer func(){
		r := recover()
		if r != nil {
			errorText := fmt.Sprintf("\nHubo un error: %s\n", r)
			errorText += strings.Repeat("-", len(errorText)) + "\n\n"
			fmt.Print(errorText)
			util.WaitForEnter()
			run()
		}
	}()

	run()
}
