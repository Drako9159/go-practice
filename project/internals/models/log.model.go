package models

import  (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"GoBaby/internals/utils"
)

type Los struct {
	Date primitive.DateTime `bson:"date"`
	// bynary-encoded json-like document
	Duration int
}   

func (c Log) FormatDuration(duration int) string {
	return utils.FormatDuration(duration)
}

func (c Log) FormatPrimitiveDateTime(date primitive.DateTime) string {
	return utils.FormatPrimitiveDate(date)
}