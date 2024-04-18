package models

type Row struct {
	IdSensor  int8     `json:"idSensor"`
	FechaToma string   `json:"fechaToma"`
	Ubicacion *string  `json:"ubicacion"`
	Data      []Medida `json:"data"`
}

type Medida struct {
	Atributo string  `json:"atributo"`
	Medida   float32 `json:"medida"`
}

type StoreProcedureResult struct {
	Errors *string `json:"errors"`
	Ok     *bool   `json:"ok"`
}
