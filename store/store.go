package store

import "data_inserter/models"

type IStore interface {
	SaveData([]models.Row) error
}
