package goveem

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GetData(query string) ([]map[string]interface{}, int) {
	rows, err := DB.Query(query)
	if err != nil {
		return nil, 500
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, 500
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
			return nil, 500
		}
		for i, colName := range columns {
			rowMap[colName] = *(values[i].(*interface{}))
		}
		result = append(result, rowMap)
	}

	if err := rows.Err(); err != nil {
		return nil, 500
	}

	defer DB.Close()

	if err != nil {
		fmt.Println("Error", err)
		return nil, 500
	}
	return result, 200
}

func UpdateData(query string, updateFields map[string]interface{}) int {
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
		return 500
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatal("Error retrieving affected rows:", err)
		return 400
	}

	if rowsAffected >= 0 {
		return 200
	}

	return 400
}

func UpdateDataQ(query string) int {
	if strings.Contains(query, "update") && !strings.Contains(query, "delete") {
		res, err := DB.Exec(query)

		if err != nil {
			return 500
		}
		rowsAffected, err := res.RowsAffected()
		if rowsAffected >= 0 {
			return 200
		}
		return 400
	}
	return 500
}

func QueryModify(query string) int {
	res, err := DB.Exec(query)
	if err != nil {
		return 500
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 500
	}
	if rowsAffected >= 0 {
		return 200
	}
	return 400
}

func DeleteData(query string) int {
	if strings.Contains(query, "delete") && !strings.Contains(query, "update") {
		res, err := DB.Exec(query)
		if err != nil {
			return 500
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return 500
		}
		if rowsAffected >= 0 {
			return 200
		}
		return 400
	}
	return 500
}
