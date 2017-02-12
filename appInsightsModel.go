package main

type QueryResult struct {
	Tables []struct {
		TableName string `json:"TableName"`
		Columns   []struct {
			ColumnName string `json:"ColumnName"`
			DataType   string `json:"DataType"`
		} `json:"Columns"`
		Rows [][]string `json:"Rows"`
	} `json:"Tables"`
}
