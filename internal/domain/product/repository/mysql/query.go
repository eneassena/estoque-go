package mysql

const (
	QUERY_INSERT       = "INSERT INTO products (name, price, count) VALUES (?,?,?)"
	QUERY_SELECT       = "SELECT name,price,count FROM products"
	QUERY_UPDATE_COUNT = "UPDATE products SET count=? WHERE id=?"
	QUERY_DELETE       = "DELETE FROM products WHERE id=?"
	QUERY_SELECT_BY_ID = "SELECT name, price, count FROM products WHERE id=?"
)
