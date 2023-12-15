package practica

import (
	"fmt"
	"math"
)

// File
type File interface {
	getName() string
	getDetails() string
}

type Files map[string]File

func (files *Files) Add(file File) {
	(*files)[file.getName()] = file
}

func (files *Files) Get(fileName string) string {
	file := (*files)[fileName]
	fileInfo := fmt.Sprintf("Información del documento '%s': %s.", file.getName(), file.getDetails())
	return fileInfo
}

// BaseFile
type BaseFile struct {
	name      string
	extension string
	size      int // Bytes
	fileType  string
}

func newBaseFile(name string, ext string, size int, fileType string) *BaseFile {
	return &BaseFile{name: name, extension: ext, size: size, fileType: fileType}
}

func (bf *BaseFile) getName() string {
	return bf.name
}

func (bf *BaseFile) getType() string {
	return bf.fileType
}

// Files
type TextFile struct {
	BaseFile
	Content string
}

func (textFile *TextFile) getDetails() string {
	details := fmt.Sprintf("%s: Este es un documento de texto con extensión .%s", textFile.fileType, textFile.extension)
	return details
}

func newTextFile(name string, ext string, size int, content string) *TextFile {
	return &TextFile{
		BaseFile: *newBaseFile(name, ext, size, "Texto"),
		Content:  content,
	}
}

type ImageFile struct {
	BaseFile
	Width  int // pixels
	Height int // pixels
}

func (imageFile *ImageFile) getDetails() string {
	details := fmt.Sprintf("%s con resolución: %dx%d y extensión .%s", imageFile.getType(), imageFile.Width, imageFile.Height, imageFile.extension)
	return details
}

func newImageFile(name string, ext string, size int, width int, height int) *ImageFile {
	return &ImageFile{
		BaseFile: *newBaseFile(name, ext, size, "Imagen"),
		Width:    width,
		Height:   height,
	}
}

type AudioFile struct {
	BaseFile
	Duration float64 // secs
}

func (audioFile *AudioFile) getDetails() string {
	minutes, seconds := math.Modf(audioFile.Duration / 60)
	seconds *= 60
	durationText := fmt.Sprintf("%02.0f:%02.0f", minutes, seconds)
	details := fmt.Sprintf("%s con duración: %s minutos y extensión .%s", audioFile.getType(), durationText, audioFile.extension)
	return details
}

func newAudioFile(name string, ext string, size int, duration float64) *AudioFile {
	return &AudioFile{
		BaseFile: *newBaseFile(name, ext, size, "Audio"),
		Duration: duration,
	}
}

func EjPairProgramming() {
	// Pair programming: Sistema de Almacenamiento de Documentos
	// Descripción del problema:
	// Crea un sistema de almacenamiento de documentos que admita diferentes tipos de documentos (por ejemplo, texto, imagen, audio). Utiliza una interfaz común llamada
	// Document y un mapa para almacenar varios documentos, donde la clave es el nombre del documento.
	// Cada tipo de documento debe tener al menos un atributo de ese tipo. Por ejemplo:
	// txt: texto.
	// jpg: size.
	// Implementa las siguientes operaciones:
	// Agregar Documento:
	// Crea una función que pueda agregar un nuevo documento al mapa.
	// Recuperar Documento:
	// Crea una función que pueda recuperar y mostrar la información de un documento específico por su nombre.
	// Output de ejemplo
	// Información del documento 'DocumentoTexto': Texto: Este es un documento de texto.
	// Información del documento 'DocumentoImagen': Imagen con resolución: 1080x720
	fmt.Print("EJERCICIO 3\n\n")

	files := []File{
		newAudioFile("ArchivoAudio", "mp3", 100, 140),
		newImageFile("ArchivoImagen", "jpg", 100, 1080, 720),
		newTextFile("ArchivoTexto", "txt", 100, "Este es un documento de texto"),
	}

	docs := make(Files)
	for _, f := range files {
		docs.Add(f)
	}
	
	fmt.Println("ARCHIVOS: ")
	for k:= range docs {
		fmt.Printf("%+v\n", docs.Get(k))
	}
}
