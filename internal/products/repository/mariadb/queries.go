package repository

const (
	sqlGetAll = `SELECT id,name,type,count,price FROM products`
	sqlGetOne = `SELECT id,name,type,count,price* FROM products WHERE id=?`
	sqlStore = `INSERT INTO products (name, type, count, price) VALUES (?, ?, ?, ?)`
	sqlLastID = `SELECT MAX(id) as last_id FROM products`
	sqlUpdate = `UPDATE products SET name=?, type=?, count=?, price=? WHERE id=?`
	sqlUpdateName = `UPDATE products SET name=? WHERE id=?`
	sqlDelete = `DELETE FROM products WHERE id=?`
)