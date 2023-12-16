package main

import (
	"fmt"
)

func main() {
	// Completar el ejercicio que está dado en la presentación
	// (https://docs.google.com/presentation/d/10CGVwkuIXYUsMajtsXOQlALbP8rW6Wc3ZDi9C6O9qXw/edit#slide=id.g1ed2b8452b7_0_34) sobre el algoritmo `addNote`.

	// 1. Se debe verificar que si solo se consideran las primeras 3 notas (las demás se ignoran)
	// 2. Las notas deben ser un valor entre 1 y 10
	//     a. Pregunta para pensar 🤔: ¿Es esta validación algo que se pueda extraer fuera del addNote? Preguntando de otro modo… la validación de que la nota sea un entero entre 1 y 10…¿ puede estar en un algoritmo aparte?
	// 3. Se pide verificar que el algoritmo que calcula los promedios funcione correctamente.
	// Podés guiarte con este código(https://go.dev/play/p/E3FLIICI7MJ)
	fmt.Print("EJERCICIO 2\n\n")

	var notesByName map[string][]int = make(map[string][]int)
	exit := false
	for !exit {
		displayMenu()
		operation, err := enterOperation()
		if err != nil {
			fmt.Println("Error introduciendo operación, el error fue:", err)
			continue
		}
		switch operation {
		case 1:
			name, note, err := enterNameAndNote()
			if err != nil {
				fmt.Println("Error introduciendo nombre y/o nota, el error fue:", err)
				continue
			}
			AddNote(notesByName, name, note)
		case 2:
			averageByName := calculateAverages(notesByName)
			fmt.Println(averageByName)
		case 3:
			exit = true
		}
		fmt.Println("algo")
	}
}

func AddNote(notesByName map[string][]int, name string, note int) error {
	// agregar código para verificar que no se estén violando los requisitos de agregar nota!
	// nota debe ser un valor entre 1 y 10
	// no debe haber más de tres notas por estudiante!
	if len(notesByName[name]) == 3 {
		return fmt.Errorf("No se pueden ingresar más de 3 notas por alumno.")
	}
	if !isInRange(note, 0, 10) {
		return fmt.Errorf("La nota debe estar entre 0 y 10 [nota=%d].", note)
	}
	if noteAlreadyExists(notesByName, note, name) {
		return fmt.Errorf("La nota %d ya está cargada para el alumno %s.", note, name)
	}
	notesByName[name] = append(notesByName[name], note)
	return nil
}

func noteAlreadyExists(notesByName map[string][]int, note int, name string) bool {
	notes := notesByName[name]
	return isNumInArray(notes, note)
}

func isNumInArray(arr []int, num int) bool {
	for _, n := range arr {
		if n == num {
			return true
		}
	}
	return false
}

func isInRange(num int, inf int, sup int) bool {
	return num >= inf && num <= sup
}

func displayMenu() {
	fmt.Println("Sistema alumnex 2.0")
	fmt.Println("====================")

	fmt.Printf("[%-30s]\n", "1. Ingreso de nota")
	fmt.Printf("[%-30s]\n", "2. Ver promedios")
	fmt.Printf("[%-30s]\n", "3. Salir")
}

func enterOperation() (int, error) {
	var operation int
	_, err := fmt.Scanf("%d", &operation)
	if err != nil {
		return 0, err
	}
	return operation, nil
}

func enterNameAndNote() (string, int, error) {
	fmt.Print("Nombre:")
	var name string
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		return "", 0, err
	}

	fmt.Print("Nota:")
	var note int
	_, err = fmt.Scanf("%d", &note)
	if err != nil {
		return "", 0, err
	}
	return name, note, nil
}

func calculateAverages(notesByName map[string][]int) map[string]float64 {
	var averageByName map[string]float64 = make(map[string]float64)
	for name, notes := range notesByName {
		notesSum := 0
		for _, note := range notes {
			notesSum += note
		}
		average := float64(notesSum) / float64(len(notes))
		averageByName[name] = average
	}
	return averageByName
}
