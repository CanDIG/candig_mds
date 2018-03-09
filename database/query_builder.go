package database

import (
	"github.com/CanDIG/candig_mds/models"
	"gopkg.in/mgo.v2/bson"
)

//BuildQuery makes a query
func BuildQuery(query models.Query) bson.M {
	pipeline := bson.M{}
	for _, condition := range query.SelectedCondition {
		pipeline = AddConditionToQuery(condition, pipeline)
	}
	return pipeline
}

//AddConditionToQuery adds each condition to the map
func AddConditionToQuery(condition []string, p bson.M) bson.M {
	var logic string
	if condition[0] == "AND" {
		logic = "$and"
	} else if condition[0] == "OR" {
		logic = "$or"
	}
	switch condition[2] {
	case "Equal to":
		p = bson.M{
			logic: []interface{}{
				bson.M{condition[1]: bson.M{"$eq": condition[3]}},
				p,
			},
		}
		break
	case "Not equal to":
		p = bson.M{
			logic: []interface{}{
				bson.M{condition[1]: bson.M{"$ne": condition[3]}},
				p,
			},
		}
		break
	case "Greater than":
		p = bson.M{
			logic: []interface{}{
				bson.M{condition[1]: bson.M{"$gt": condition[3]}},
				p,
			},
		}
		break
	case "Less than":
		p = bson.M{
			logic: []interface{}{
				bson.M{condition[1]: bson.M{"$lt": condition[3]}},
				p,
			},
		}
		break
	case "Greater or equal to":
		p = bson.M{
			logic: []interface{}{
				bson.M{condition[1]: bson.M{"$gte": condition[3]}},
				p,
			},
		}
		break
	case "Less or equal to":
		p = bson.M{
			logic: []interface{}{
				bson.M{condition[1]: bson.M{"$lte": condition[3]}},
				p,
			},
		}
		break

	default:
		return p
	}
	return p
}
