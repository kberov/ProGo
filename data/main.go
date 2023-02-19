package main

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	Printfln("Working with Databases%s",
		"\n Installing a Database Driver")
	Printfln("Available drivers:")
	listDrivers()
	dbh, err := openDatabase()
	if err == nil {
		// 697 Chapter 26 ■ Working with Databases
		Printfln("Opened database…")
		Printfln("\n Querying for Multiple Rows")
		queryDatabase(dbh)
		//Printfln("…closing it")
		//_ = dbh.Close()
	} else {
		panic("Error: " + err.Error())
	}

	Printfln("\n Scanning Values into a Struct")
	products := queryProducts(dbh)
	Printfln("All queried products:\n%#v", products)
	products = queryProductsWithCategories(dbh, QPOpts{})
	Printfln("All queried products:\n%#v", products)

	Printfln("\n Executing Statements with Placeholders")

	for _, cat := range []string{"Soccer", "Watersports"} {
		Printfln("--- %v Results ---", cat)
		products := queryProductsWithCategories(dbh, QPOpts{CatName: cat})
		for i, p := range products {
			Printfln("#%v: %v %v %v", i, p.Name, p.Category.Name, p.Price)
		}
	}

	Printfln("\n Executing Queries for Single Rows")
	for _, id := range []int{1, 3, 10} {
		p := queryProductRow(dbh, id)
		Printfln("Product: %v", p)
	}

	Printfln("\n Executing Other Queries")
	/* The Exec method is used for executing statements that don’t produce rows.
	 * The result from the Exec method is a Result value, which defines the methods
	 * described in Table 26-8, and an error that indicates problems executing the
	 * statement. */
	newProduct := Product{Name: "Stadium", Category: Category{Id: 2}, Price: 79500}
	newID := insertRow(dbh, &newProduct)

	p := queryProductRow(dbh, int(newID))
	Printfln("New Product: %v", p)

	Printfln("\n Using Prepared Statements")
	insertAndUseCategory("Misc Products", 2)
	Printfln("Changed product: %v", queryProductRow(dbh, 2))

	Printfln("\n Using Transactions")
	if err := insertAndUseCategoryTx(dbh, "Лека Атлетика", 6, 7, 8); err != nil {
		Printfln("Transaction failed: %v", err)
	}

	Printfln("Changed product: %v", queryProductRow(dbh, 6))
	if err := insertAndUseCategoryTx(dbh, "Olimpics", 6, 100); err != nil {
		Printfln("Transaction failed: %v", err)
	}

	Printfln("\n Using Reflection to Scan Data into a Struct")
	products, _ = queryDatabaseReflect(dbh)
	for _, p := range products {
		Printfln("Product: %v", p)
	}
}

func queryDatabaseReflect(db *sql.DB) (products []Product, err error) {
	rows, err := db.Query(`SELECT Products.Id, Products.Name, Products.Price,
Categories.Id as "Category.Id", Categories.Name as "Category.Name"
FROM Products, Categories
WHERE Products.Category = Categories.Id`)
	if err != nil {
		return
	} else {
		results, err := scanIntoStruct(rows, &Product{})
		if err == nil {
			products = (results).([]Product)
		} else {
			Printfln("Scanning error: %v", err)
		}
	}
	return
}

func scanIntoStruct(rows *sql.Rows, target any) (results any, err error) {
	targetVal := reflect.ValueOf(target)
	if targetVal.Kind() == reflect.Ptr {
		targetVal = targetVal.Elem()
	}
	if targetVal.Kind() != reflect.Struct {
		return
	}
	colNames, _ := rows.Columns()
	colTypes, _ := rows.ColumnTypes()
	references := []any{}
	fieldVal := reflect.Value{}
	var placeholder any
	Printfln("colTypes: %#v", colTypes)
	for i, colName := range colNames {
		colNameParts := strings.Split(colName, ".")
		fieldVal = targetVal.FieldByName(colNameParts[0])
		if fieldVal.IsValid() && fieldVal.Kind() == reflect.Struct &&
			len(colNameParts) > 1 {
			var namePart string
			// 718 Chapter 26 ■ Working with Databases
			for _, namePart = range colNameParts[1:] {
				compFunction := matchColName(namePart)
				fieldVal = fieldVal.FieldByNameFunc(compFunction)
			}
		}
		if !fieldVal.IsValid() ||
			!colTypes[i].ScanType().ConvertibleTo(fieldVal.Type()) {
			references = append(references, &placeholder)
		} else if fieldVal.Kind() != reflect.Ptr && fieldVal.CanAddr() {
			fieldVal = fieldVal.Addr()
			references = append(references, fieldVal.Interface())
		}
	}
	resultSlice := reflect.MakeSlice(reflect.SliceOf(targetVal.Type()), 0, 10)
	for rows.Next() {
		err = rows.Scan(references...)
		if err == nil {
			resultSlice = reflect.Append(resultSlice, targetVal)
		} else {
			break
		}
	}
	results = resultSlice.Interface()
	return
}

func matchColName(colName string) func(string) bool {
	return func(fieldName string) bool {
		return strings.EqualFold(colName, fieldName)
	}
}

func insertAndUseCategoryTx(db *sql.DB, catName string, productIDs ...int) (err error) {
	tx, err := db.Begin()
	updatedFailed := false
	var uerr error
	if err == nil {
		catResult, err := tx.Stmt(insertNewCategory).Exec(catName)
		if err == nil {
			newID, _ := catResult.LastInsertId()
			preparedStatement := tx.Stmt(changeProductCategory)
			for _, id := range productIDs {
				changeResult, err := preparedStatement.Exec(newID, id)
				if err == nil {
					changes, _ := changeResult.RowsAffected()
					if changes == 0 {
						uerr = errors.New(fmt.Sprintf(
							"Updating category to %v  with newID %d for field with Id %d failed!", catName, newID, id))
						updatedFailed = true
						break
					}
				}
			}
		}
	}
	if err != nil || uerr != nil || updatedFailed {
		Printfln("Aborting transaction: %v; Update failed: %t", uerr, updatedFailed)
		if err = tx.Rollback(); err != nil {

			Printfln("Rollback failed %v", err)
		}

	} else {
		if err = tx.Commit(); err != nil {
			Printfln("Commit failed %v", err)
		}
	}
	return
}

func insertAndUseCategory(name string, productIDs ...int) {
	result, err := insertNewCategory.Exec(name)
	if err == nil {
		newID, _ := result.LastInsertId()
		for _, id := range productIDs {
			_, _ = changeProductCategory.Exec(int(newID), id)
		}
	} else {
		Printfln("Prepared statement error: %v", err)
	}
}
func insertRow(db *sql.DB, p *Product) (id int64) {
	res, err := db.Exec(`
INSERT INTO Products (Name, Category, Price)
VALUES (?, ?, ?)`, p.Name, p.Category.Id, p.Price)
	if err == nil {
		id, err = res.LastInsertId()
		if err != nil {
			Printfln("Result error: %v", err.Error())
		}
	} else {
		Printfln("Exec error: %v", err.Error())
	}
	return
}

func queryProductRow(db *sql.DB, id int) (p Product) {
	row := db.QueryRow(`
		SELECT p.id, p.name, p.price,c.id AS c_id,c.name AS c_name
		FROM Products p, Categories c
		WHERE p.category = c.id AND p.id = ?`, id)
	if row.Err() == nil {
		scanErr := row.Scan(&p.Id, &p.Name, &p.Price,
			&p.Category.Id, &p.Category.Name)
		if scanErr != nil {
			Printfln("Scan error: %v", scanErr)
		}
	} else {
		Printfln("Row error: %v", row.Err().Error())
	}
	return
}

type Category struct {
	Id   int
	Name string
}

type Product struct {
	Id   int
	Name string
	Category
	Price float64
}

// We use this struct as ever changing parameters for queryProductsWithCategories(…)
type QPOpts struct {
	CatName string
}

func queryProductsWithCategories(db *sql.DB, opts QPOpts) (products []Product) {
	var andCatName string
	if opts.CatName != "" {
		andCatName = " AND c.Name=?"
	}
	rows, err := db.Query(`
		SELECT p.id, p.name, p.price,c.id AS c_id,c.name AS c_name
		FROM Products p, Categories c
		WHERE p.category = c.id `+andCatName, opts.CatName)
	Printfln("\n Understanding the rows.Scan Method")
	if err == nil {
		for rows.Next() {
			p := Product{}
			scanErr := rows.Scan(&p.Id, &p.Name, &p.Price,
				&p.Category.Id, &p.Category.Name)
			if scanErr == nil {
				Printfln("Product Row: %v ", p)
				products = append(products, p)
			} else {
				Printfln("scanErr: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error: %v", err)
	}
	return
}

func queryProducts(db *sql.DB) (products []Product) {
	rows, err := db.Query("SELECT * from Products")
	Printfln("\n Understanding the rows.Scan Method")
	if err == nil {
		for rows.Next() {
			p := Product{}
			scanErr := rows.Scan(&p.Id, &p.Name, &p.Category.Id, &p.Price)
			if scanErr == nil {
				Printfln("Product Row: %v ", p)
				products = append(products, p)
			} else {
				Printfln("scanErr: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error: %v", err)
	}
	return
}

func queryDatabase(db *sql.DB) {
	rows, err := db.Query("SELECT * from Products")
	Printfln("\n Understanding the rows.Scan Method")
	if err == nil {
		for rows.Next() {
			var id, category int
			var name string
			var price float64
			scanErr := rows.Scan(&id, &name, &category, &price)
			if scanErr == nil {
				Printfln("Row: %v %v %v %v", id, name, category, price)
			} else {
				Printfln("scanErr: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error: %v", err)
	}
}
