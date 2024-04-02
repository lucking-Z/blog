package mongo

import (
	"blogs/pkg/filter"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func transformBson(ex *filter.Expression) (bson.E, error) {
	switch ex.Op {
	case filter.OpEq:
		return bson.E{Key: ex.AttrName, Value: bson.M{"$eq": ex.Value}}, nil
	case filter.OpNotEq:
		return bson.E{Key: ex.AttrName, Value: bson.M{"$ne": ex.Value}}, nil
	case filter.OpGt:
		return bson.E{Key: ex.AttrName, Value: bson.M{"$gt": ex.Value}}, nil
	case filter.OpGte:
		return bson.E{Key: ex.AttrName, Value: bson.M{"$gte": ex.Value}}, nil
	case filter.OpLt:
		return bson.E{Key: ex.AttrName, Value: bson.M{"$lt": ex.Value}}, nil
	case filter.OpLte:
		return bson.E{Key: ex.AttrName, Value: bson.M{"$lte": ex.Value}}, nil
	case filter.OpIn:
		return bson.E{Key: ex.AttrName, Value: bson.D{{"$in", ex.Values}}}, nil
	default:
		return bson.E{}, fmt.Errorf("expression relational operator is unknown:%d", ex.Op)
	}
}
func BuildMongo(ex []filter.Expression) ([]bson.E, error) {
	tmp := make([]bson.E, 0)
	for k, v := range ex {
		if v.Op == filter.OpAnd {
			exps := ex[k].GetExps()
			if len(exps) < 1 {
				return nil, errors.New("expression empty")
			}
			res, err := BuildMongo(exps)
			if err != nil {
				return nil, err
			}
			tmp = append(tmp, bson.E{Key: "$and", Value: bson.A{res}})
			//tmp = append(tmp, res...)
			//
		} else if v.Op == filter.OpOr {
			exps := ex[k].GetExps()
			if len(exps) < 1 {
				return nil, errors.New("expression empty")
			}
			res, err := BuildMongo(exps)
			if err != nil {
				return nil, err
			}
			tmp = append(tmp, bson.E{Key: "$or", Value: bson.A{res}})
		} else {
			res, err := transformBson(&ex[k])
			if err != nil {
				return nil, err
			}
			tmp = append(tmp, res)
		}
	}
	return tmp, nil
}
