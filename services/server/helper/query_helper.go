package helper

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type QueryExecuter interface {
	DoQuery(query string, args ...interface{}) ([]map[string]interface{}, error)
	DoQueryRow(query string, args ...interface{}) (map[string]interface{}, error)
	DoQuerySlice(query string, args ...interface{}) ([]interface{}, error)
}

type QueryHelper struct {
	DB *sqlx.DB
}

func parseRows(rows *sqlx.Rows) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	for rows.Next() {
		datum := make(map[string]interface{})
		err := rows.MapScan(datum)
		if err != nil {
			return nil, err
		}
		data = append(data, datum)
	}
	return data, nil
}

func (q *QueryHelper) DoQuery(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := q.DB.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	data, err := parseRows(rows)
	log.Println("Query %v %v return %v", query, args, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (q *QueryHelper) DoQueryRow(query string, args ...interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	err := q.DB.QueryRowx(query, args...).MapScan(data)
	log.Printf("Query %v %v return %v", query, args, data)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}

	}
	return data, nil
}

func (q *QueryHelper) DoQuerySlice(query string, args ...interface{}) ([]interface{}, error) {
	rows, err := q.DB.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	res := []interface{}{}
	for rows.Next() {
		var ret interface{}
		err := rows.Scan(&ret)
		if err != nil {
			return nil, err
		}
		res = append(res, ret)

	}
	log.Printf("Query Ret: %v", res)
	return res, err
}
