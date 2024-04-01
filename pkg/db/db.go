package db

import (
	"blogs/pkg/filter"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

type ConfigDB struct {
	Username string
	Password string
	Host     string
	Port     int
	DbName   string
}

var instance *gorm.DB

func GetInstance() *gorm.DB {
	return instance
}

func transformQuery(ex *filter.Expression) (clause.Expression, error) {
	switch ex.Op {
	case filter.OpEq:
		return clause.Eq{Column: ex.AttrName, Value: ex.Value}, nil
	case filter.OpNotEq:
		return clause.Neq{Column: ex.AttrName, Value: ex.Value}, nil
	case filter.OpGt:
		return clause.Gt{Column: ex.AttrName, Value: ex.Value}, nil
	case filter.OpGte:
		return clause.Gte{Column: ex.AttrName, Value: ex.Value}, nil
	case filter.OpLt:
		return clause.Lt{Column: ex.AttrName, Value: ex.Value}, nil
	case filter.OpLte:
		return clause.Lte{Column: ex.AttrName, Value: ex.Value}, nil
	case filter.OpIn:
		return clause.IN{Column: ex.AttrName, Values: ex.Values}, nil
	default:
		return nil, fmt.Errorf("expression relational operator is unknown:%d", ex.Op)
	}
}

func Build(ex []filter.Expression, wrapInParentheses bool) ([]clause.Expression, error) {
	tmp := make([]clause.Expression, 0)
	for k, v := range ex {
		if v.Op == filter.OpAnd || v.Op == filter.OpOr {
			exps := ex[k].GetExps()
			res, err := Build(exps, true)
			if err != nil {
				return nil, err
			}
			if !wrapInParentheses {
				tmp = append(tmp, res...)
			} else {
				if v.Op == filter.OpAnd {
					tmp = append(tmp, clause.And(res...))
				} else if v.Op == filter.OpOr {
					tmp = append(tmp, clause.Or(res...))
				}
			}
		} else {
			res, err := transformQuery(&ex[k])
			if err != nil {
				return nil, err
			}
			tmp = append(tmp, res)
		}
	}
	return tmp, nil
}

func InitDB(confObj ConfigDB) {
	sync.OnceFunc(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&&interpolateParams=true&loc=Local",
			confObj.Username,
			confObj.Password,
			confObj.Host,
			confObj.Port,
			confObj.DbName)
		loggerIns := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
			},
		)
		mysqlConfig := mysql.Config{DSN: dsn}
		gormConfig := &gorm.Config{
			Logger: loggerIns,
		}
		var err error
		instance, err = gorm.Open(mysql.New(mysqlConfig), gormConfig)
		if err != nil {
			panic(fmt.Sprintf("db init failed:%s", err.Error()))
		}

		sqlDB, err := instance.DB()
		if err != nil {
			panic(fmt.Sprintf("db init failed:%s", err.Error()))
		}
		//最大空闲连接数
		sqlDB.SetMaxIdleConns(10)
		//最大打开数量
		sqlDB.SetMaxOpenConns(100)
		//最大可服务时间
		sqlDB.SetConnMaxLifetime(time.Hour)
	})
}
