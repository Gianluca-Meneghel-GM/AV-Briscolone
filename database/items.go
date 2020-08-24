package database

import "antaresvision.com/corso-go-2/models"

type ItemsReader interface {
	FetchItems() ([]models.Serial, error)
	FetchItem(ntind int, serial string) (*models.Serial, error)
}

type ItemsWriter interface {
	InsertItem(item *models.Serial) error
	UpdateItem(item models.Serial) error
}

type ItemsReaderWriter interface {
	ItemsReader
	ItemsWriter
}
