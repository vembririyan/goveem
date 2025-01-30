package goveem

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func SELECT(query string) ([]map[string]interface{}, int, error) {
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println("Error", err)

		return nil, 500, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, 500, err
	}

	var result []map[string]interface{}

	values := make([]interface{}, len(columns))

	for rows.Next() {
		rowMap := make(map[string]interface{})
		for i := range columns {
			var value interface{}
			values[i] = &value
		}

		if err := rows.Scan(values...); err != nil {
			return nil, 500, err
		}
		for i, colName := range columns {
			rowMap[colName] = *(values[i].(*interface{}))
		}
		result = append(result, rowMap)
	}

	if err := rows.Err(); err != nil {
		return nil, 500, err
	}

	defer DB.Close()

	if err != nil {
		fmt.Println("Error", err)
		return nil, 500, err
	}
	return result, 200, err
}

func ExecQuery(query string) (int, error) {
	res, err := DB.Exec(query)

	if err != nil {
		fmt.Println("Error", err)
		return 500, err
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatal("Error retrieving affected rows:", err)
		return 400, err
	}

	if rowsAffected >= 0 {
		return 200, err
	}

	return 500, err
}

func UpdateData(query string, updateFields map[string]interface{}) (int, error) {
	var setClauses []string
	var values []interface{}

	var i_column int = 1

	for column, value := range updateFields {
		switch DB_VERSION {
		case "mysql":
			setClauses = append(setClauses, fmt.Sprintf("%s= ? ", column))
		case "postgre":
			i := "$" + strconv.Itoa(i_column)
			setClauses = append(setClauses, fmt.Sprintf("%s= "+i, column))
		default:
			setClauses = append(setClauses, fmt.Sprintf("%s= ? ", column))
		}

		values = append(values, value)
		i_column++
	}

	setClause := strings.Join(setClauses, ", ")
	updateQuery := fmt.Sprintf(query, setClause)

	defer DB.Close()
	res, err := DB.Exec(updateQuery, values...)

	if err != nil {
		return 500, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatal("Error retrieving affected rows:", err)
		return 400, err
	}

	if rowsAffected >= 0 {
		return 200, err
	}

	return 400, err
}
