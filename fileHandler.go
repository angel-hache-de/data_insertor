package main

import (
	"data_inserter/models"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func GetRegistries(multiplicator int8) []models.Row {
	rows, err := GetFileRows()
	if err != nil {
		fmt.Println("Error reading rows:", err)
		return nil
	}

	data := ReadRows(int(multiplicator), rows)
	return data
}

func InsertRegistriesOneByOne() error {
	headers := GetHeaders()
	rows, err := GetFileRows()
	if err != nil {
		fmt.Println("Error reading rows:", err)
		return nil
	}

	input := make([]models.Row, 1)
	// Iterate over the rows and print each cell value
	rowToSave := models.Row{
		IdSensor: 1,
		Data:     make([]models.Medida, len(headers)-1),
	}
	input[0] = rowToSave
	medida := models.Medida{}
	f64 := 0.0
	for _, row := range rows[7:] {
		for index, cell := range row {
			if index == 0 {
				rowToSave.FechaToma = cell
				continue
			}

			medida.Atributo = headers[index]
			f64, _ = strconv.ParseFloat(cell, 32)
			medida.Medida = float32(f64)
			rowToSave.Data[index-1] = medida
		}
		if err := actualStore.SaveData(input); err != nil {
			return err
		}
	}

	return nil
}

func ReadRows(multiplicator int, rows [][]string) []models.Row {
	headers := GetHeaders()
	// Iterate over the rows and print each cell value
	data := make([]models.Row, 0)
	rowToSave := models.Row{
		IdSensor: 1,
		Data:     make([]models.Medida, len(headers)-1),
	}
	medida := models.Medida{}
	f64 := 0.0
	for _, row := range rows[7:] {
		for index, cell := range row {
			if index == 0 {
				rowToSave.FechaToma = cell
				continue
			}

			medida.Atributo = headers[index]
			f64, _ = strconv.ParseFloat(cell, 32)
			medida.Medida = float32(f64)
			rowToSave.Data[index-1] = medida
		}
		for i := 0; i < multiplicator; i++ {
			data = append(data, rowToSave)
		}
	}

	return data
}

func GetFileRows() ([][]string, error) {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer f.Close()

	// Get all the rows from the sheet
	return f.GetRows("data")
}

// como vrga normalizar el nombre weon
func GetHeaders() []string {
	return []string{"datetime", "aqi", "aqi_alto", "pm1", "pm1_alto", "pm2.5", "pm2.5_alto", "pm10", "pm10_alto", "temperatura", "temperatura_alta", "temperatura_baja", "humedad", "humedad_alta", "humedad_baja", "punto_rocio", "maxima_punto_rocio", "minima_punto_rocio", "bulbo_humedo", "maxima_bulbo_humedo", "minima_bulbo_humedo", "indice de calor", "maxima_indice_calor", "minima_indice_calor"}
}

func GetRows() {

}
