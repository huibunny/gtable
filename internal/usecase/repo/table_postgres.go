package repo

import (
	"context"
	"fmt"

	"gtable/config"
	"gtable/internal/entity"
	"gtable/pkg/postgres"
)

const _defaultEntityCap = 64

// TableRepo -.
type TableRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *TableRepo {
	return &TableRepo{pg}
}

// Load -.
func (r *TableRepo) Load(ctx context.Context, tables []config.Table) (entity.TableValue, int, error) {
	errcode := 0
	var data entity.TableValue

	for _, table := range tables {
		switch table.Type {
		case "kv":
			kv, errcode, err := r.loadKV(ctx, table.Name, table.Column)
			if err != nil {
				return data, errcode, err
			} else {
				data.ValueList = append(data.ValueList, kv)
			}
		case "list":
			valuesList, errcode, err := r.loadList(ctx, table.Name, table.Column)
			if err != nil {
				return data, errcode, err
			} else {
				data.ValueList = append(data.ValueList, valuesList)
			}
		default:
			//
		}
	}

	return data, errcode, nil
}

func (r *TableRepo) loadKV(ctx context.Context, table string, columns []string) (map[string]interface{}, int, error) {
	kv := make(map[string]interface{})
	var key string
	var value map[string]interface{}
	sql, _, err := r.Builder.
		Select(columns...).
		From(table).ToSql()
	if err != nil {
		return kv, -1, fmt.Errorf("TableRepo - Load - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return kv, -2, fmt.Errorf("TableRepo - Load - r.Pool.Query: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&key, &value)
		if err != nil {
			return kv, -3, fmt.Errorf("TableRepo - Load - Scan: %w", err)
		}
		kv[key] = value
	}

	return kv, 0, nil
}

func (r *TableRepo) loadList(ctx context.Context, table string, columns []string) ([]map[string]interface{}, int, error) {
	var valuesList []map[string]interface{}
	sql, _, err := r.Builder.
		Select(columns...).
		From(table).ToSql()
	if err != nil {
		return valuesList, -1, fmt.Errorf("TableRepo - Load - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return valuesList, -2, fmt.Errorf("TableRepo - Load - r.Pool.Query: %w", err)
	}

	for rows.Next() {
		valueMap := make(map[string]interface{})
		values, err := rows.Values()
		if err != nil {
			return valuesList, -3, fmt.Errorf("TableRepo - Load - rows.Values: %w", err)
		}
		for index, value := range values {
			valueMap[columns[index]] = value
		}

		valuesList = append(valuesList, valueMap)
	}
	return valuesList, 0, nil
}
