package provide

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type NameRow struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getNameTable(ctx context.Context, connection *pgxpool.Pool, query string) (*NameRow, error) {
	var result NameRow
	err := pgxscan.Get(context.Background(), connection, &result, query)
	return &result, err
}

func getNameById(connection *pgxpool.Pool, table string, id int) (*NameRow, error) {
	return getNameTable(
		context.Background(),
		connection,
		fmt.Sprintf(
			`SELECT %v_id as id, name FROM %v WHERE %v_id = %v`,
			table, table, table, id,
		),
	)
}

func getIdByName(connection *pgxpool.Pool, table, name string) (*NameRow, error) {
	return getNameTable(
		context.Background(),
		connection,
		fmt.Sprintf(
			`SELECT %v_id as id, name FROM %v WHERE name = '%v'`,
			table, table, name,
		),
	)
}

func listNameTable(connection *pgxpool.Pool, table string) (*[]*NameRow, error) {
	result := &[]*NameRow{}
	err := pgxscan.Select(
		context.Background(),
		connection,
		result,
		fmt.Sprintf(`SELECT %v_id as id, name FROM %v`, table, table),
	)
	return result, err
}
