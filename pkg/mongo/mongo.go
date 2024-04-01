package mongo

import (
	"blogs/pkg/filter"
	"go.mongodb.org/mongo-driver/mongo"
)

func Build(ex []filter.Expression) {
	mongo.Client{}.Database().Collection().Find()
}
