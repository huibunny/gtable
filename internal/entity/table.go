// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// User -.
// ValueList contains values of each table
// value type:
// 1. kv
// 		map[string]interface{}
// 2. list
//      []interface{}
type TableValue struct {
	ValueList []interface{} `json:"value_list" example:""`
}
