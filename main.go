package main

import (
	"encoding/csv"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("Corriendo Algoritmo")

	file, err := os.Open("Nacidos_2019.csv")
	if err != nil {
		log.Fatal("Error al abrir el archivo Nacidos_2019")
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Error al leer contenido de Nacidos_2019")
	}

	start := time.Now()
	log.Println("Se inicia limpieza de datos")

	newFile, err := os.Create("cleaned_data.csv")
	if err != nil {
		log.Fatal("No se logró crear el archivo para guardar los nuevos datos")
	}

	defer newFile.Close()

	csvWriter := csv.NewWriter(newFile)

	for i, row := range data {

		// Solo se extraen columnas de interes
		newRow := []string{row[0], row[8], row[12], row[20], row[21], row[22], row[23], row[28], row[33], row[34], row[35], row[9], row[29], row[30]}

		// Se guarda nombre de las columnas
		if i == 0 {
			_ = csvWriter.Write(newRow)

			// Solo se guarda datos que pertenezcan a Antioquia
		} else if row[0] == "05" {
			newRow[0] = "Antioquia"
			newRow[2] = numeroConsultas(row[12])
			newRow[3] = edadMadre(row[20])
			newRow[4] = estadoConyugalMadre(row[21])
			newRow[5] = nivelEducativo(row[22])
			newRow[6] = sinInfo(row[23])
			newRow[7] = sinInfo(row[28])
			newRow[8] = edadPadre(row[33])
			newRow[9] = nivelEducativo(row[34])
			newRow[10] = sinInfo(row[35])
			newRow[11] = mesOcurrencia(row[9])
			newRow[13] = sinInfo(row[30])

			_ = csvWriter.Write(newRow)
		}

	}

	csvWriter.Flush()

	elapsed := time.Since(start)
	log.Printf("Se finaliza limpieza de datos en %s\n", elapsed/time.Millisecond)
}

func numeroConsultas(val string) string {
	switch val {
	case "00":
		return "Ninguna"
	case "99":
		return "Sin informacion"
	default:
		return val
	}
}

func edadMadre(val string) string {
	switch val {
	case "1":
		return "De 10-14 Anios"
	case "2":
		return "De 15-19 Anios"
	case "3":
		return "De 20-24 Anios"
	case "4":
		return "De 25-29 Anios"
	case "5":
		return "De 30-34 Anios"
	case "6":
		return "De 35-39 Anios"
	case "7":
		return "De 40-44 Anios"
	case "8":
		return "De 45-49 Anios"
	case "9":
		return "De 50-54 Anios"
	case "99":
		return "Sin informacion"
	default:
		return val
	}
}

func estadoConyugalMadre(val string) string {
	switch val {
	case "1":
		return "No esta casada y lleva dos o mas anios viviendo con su pareja"
	case "2":
		return "No esta casada y lleva menos de dos anios viviendo con su pareja"
	case "3":
		return "Esta separada, divorciada"
	case "4":
		return "Esta viuda"
	case "5":
		return "Esta soltera"
	case "6":
		return "Esta casada"
	case "9":
		return "Sin informacion"
	default:
		return val
	}
}

func nivelEducativo(val string) string {
	switch val {
	case "01":
		return "Preescolar"
	case "02":
		return "Basica primaria"
	case "03":
		return "Basica secundaria"
	case "04":
		return "Media academica o clasica"
	case "05":
		return "Media tecnica"
	case "06":
		return "Normalista"
	case "07":
		return "Tecnica profesional"
	case "08":
		return "Tecnologica"
	case "09":
		return "Profesional"
	case "10":
		return "Especializacion"
	case "11":
		return "Maestria"
	case "12":
		return "Doctorado"
	case "13":
		return "Ninguno"
	case "99":
		return "Sin informacion"
	default:
		return val
	}
}

func sinInfo(val string) string {
	switch val {
	case "99":
		return "Sin informacion"
	default:
		return val
	}
}

func edadPadre(val string) string {
	switch val {
	case "999":
		return "Sin informacion"
	default:
		return val
	}
}

func mesOcurrencia(val string) string {
	switch val {
	case "01":
		return "Enero"
	case "02":
		return "Febrero"
	case "03":
		return "Marzo"
	case "04":
		return "Abril"
	case "05":
		return "Mayo"
	case "06":
		return "Junio"
	case "07":
		return "Julio"
	case "08":
		return "Agosto"
	case "09":
		return "Septiembre"
	case "10":
		return "Octubre"
	case "11":
		return "Noviembre"
	case "12":
		return "Diciembre"
	default:
		return val
	}
}
