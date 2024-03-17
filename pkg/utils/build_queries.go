package utils

import (
	"reflect"
	"strings"

	"github.com/bit-wiz/data-store-a/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	pageSize int64 = 3
)

// will search on all fields of the Citizen struct as req.
func SearchQ(s string) primitive.M {
	filter := bson.M{}

	if s == "" {
		return filter
	}

	fields := reflect.ValueOf(models.Citizen{})

	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Tag.Get("bson")

		fieldFilter := bson.M{
			fieldName: bson.M{
				"$regex": primitive.Regex{
					Pattern: s,
					Options: "i",
				},
			},
		}

		if _, ok := filter["$or"]; !ok {
			filter["$or"] = []bson.M{}
		}

		filter["$or"] = append(filter["$or"].([]bson.M), fieldFilter)
	}

	return filter
}

func FindOpts(sort string, page int) options.FindOptions {
	var opts options.FindOptions
	// [0]: sort type (asc, desc) | [1]: sort type (by field)
	sortdata := strings.Split(sort, "-")

	if sortdata[0] == "a" {
		opts.SetSort(bson.D{{sortdata[1], 1}})
	} else {
		opts.SetSort(bson.D{{sortdata[1], -1}})
	}

	opts.SetSkip((int64(page) - 1) * pageSize)
	opts.SetLimit(pageSize)

	return opts
}
