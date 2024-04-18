package store

import (
	"data_inserter/models"
	"encoding/json"

	"errors"

	_ "github.com/microsoft/go-mssqldb"
)

func (sqlServer *SqlServer) SaveData(data []models.Row) error {
	databaseConnection, err := sqlServer.openDatabaseConnection()

	if err != nil {
		return errors.New("(origin -- SaveData) No se pudo crear la conexion a MSSQL " + err.Error())
	}
	defer sqlServer.closeDatabaseConnection(databaseConnection)

	stringData, err := json.Marshal(data)
	if err != nil {
		return errors.New("(origin -- SaveData) error marshaling the rows" + err.Error())
	}

	result, err := databaseConnection.Query(SAVE_DATA_SP,
		stringData,
	)

	if err != nil {
		return errors.New("(origin -- SaveData) error executing the sp: " + err.Error())
	}

	if !result.Next() {
		return errors.New("(origin -- SaveData) no se pudo guardar el error")
	}

	var mssqlResult models.StoreProcedureResult
	if err = result.Scan(&mssqlResult.Errors, &mssqlResult.Ok); err != nil {
		return errors.New("(origin -- SaveData) error scaning the sp result " + err.Error())
	}

	if !*mssqlResult.Ok {
		return errors.New("(origin -- SaveData) no se pudo guardar la data: " + *mssqlResult.Errors)
	}

	return nil
}
