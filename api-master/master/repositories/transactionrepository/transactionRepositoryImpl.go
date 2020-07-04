package transactionrepository

import (
	"database/sql"
	"fmt"
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/utils"
	"strconv"

	guuid "github.com/google/uuid"
)

type TransactionRepoImpl struct {
	db *sql.DB
}

func InitTransactionRepoImpl(db *sql.DB) TransactionRepository {
	return &TransactionRepoImpl{db: db}
}

func (t *TransactionRepoImpl) GetAll() ([]*models.Transaction, error) {
	rows, err := t.db.Query(utils.ALL_TRANSACTION)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var listTransaction []*models.Transaction
	for rows.Next() {
		transaction := models.Transaction{}
		err := rows.Scan(&transaction.OrderID, &transaction.ProductName, &transaction.CategoryName, &transaction.ProductCode, &transaction.Price, &transaction.Quantity, &transaction.OrderDate, &transaction.OutletName, &transaction.RegionName)
		if err != nil {
			return nil, err
		}
		price, _ := strconv.Atoi(transaction.Price)
		qty, _ := strconv.Atoi(transaction.Quantity)
		transaction.TotalPrice = price * qty
		listTransaction = append(listTransaction, &transaction)
	}
	return listTransaction, nil
}

func (t *TransactionRepoImpl) GetThisDay() ([]*models.Transaction, error) {
	// thisday := fmt.Sprintf(utils.TesDate)
	rows, err := t.db.Query(utils.ALL_TRANSACTION+" "+utils.THIS_DAY, utils.Date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var listTransaction []*models.Transaction
	for rows.Next() {
		transaction := models.Transaction{}
		err := rows.Scan(&transaction.OrderID, &transaction.ProductName, &transaction.CategoryName, &transaction.ProductCode, &transaction.Price, &transaction.Quantity, &transaction.OrderDate, &transaction.OutletName, &transaction.RegionName)
		if err != nil {
			return nil, err
		}
		price, _ := strconv.Atoi(transaction.Price)
		qty, _ := strconv.Atoi(transaction.Quantity)
		transaction.TotalPrice = price * qty
		listTransaction = append(listTransaction, &transaction)
	}
	return listTransaction, nil
}

func (t *TransactionRepoImpl) GetByMonth(monthdata string) ([]*models.Transaction, error) {
	month := fmt.Sprintf(utils.ALL_TRANSACTION + " " + utils.THIS_MONTH)
	rows, err := t.db.Query(month, monthdata, monthdata)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var listTransaction []*models.Transaction
	for rows.Next() {
		transaction := models.Transaction{}
		err := rows.Scan(&transaction.OrderID, &transaction.ProductName, &transaction.CategoryName, &transaction.ProductCode, &transaction.Price, &transaction.Quantity, &transaction.OrderDate, &transaction.OutletName, &transaction.RegionName)
		if err != nil {
			return nil, err
		}
		price, _ := strconv.Atoi(transaction.Price)
		qty, _ := strconv.Atoi(transaction.Quantity)
		transaction.TotalPrice = price * qty
		listTransaction = append(listTransaction, &transaction)
	}

	return listTransaction, nil
}

func (t *TransactionRepoImpl) Add(data []*models.Transaction) (list []*models.Transaction, err error) {
	for _, trans := range data {
		pName := trans.ProductName
		outletName := trans.OutletName
		row := t.db.QueryRow(`select id,product_code, product_name,category_id from m_product where product_name =?`, pName)
		var prod = models.Product{} //param
		err := row.Scan(&prod.ID, &prod.ProductCode, &prod.ProductName, &prod.ProductCategory.CategoryId)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		row = t.db.QueryRow(`select * from m_outlet where outlet_name=?`, outletName)
		var trOutlet = models.Outlet{} //param
		err = row.Scan(&trOutlet.OutletCode, &trOutlet.OutletName, &trOutlet.Region)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		row = t.db.QueryRow(`select product_price_id,product_price from m_product_price where product_id=? and is_active=1`, prod.ID)
		var trProdPrice = models.ProductPrice{}
		err = row.Scan(&trProdPrice.ProductPriceID, &trProdPrice.ProductPrice)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		tx, err := t.db.Begin()
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		orderID := guuid.New()

		_, err = tx.Exec(`insert into purchase_order values(?,?,?,?,?)`, orderID, utils.CurDate, trOutlet.OutletCode, utils.CurDate, utils.CurDate)
		if err != nil {
			fmt.Println(err.Error())
			tx.Rollback()
			return nil, err
		}
		_, err = tx.Exec(`insert into purchase_order_item(qty,created_at,updated_at,order_id,product_id) values(?,?,?,?,?)`, trans.Quantity, utils.CurDate, utils.CurDate, orderID, trProdPrice.ProductPriceID)
		if err != nil {
			fmt.Println(err.Error())
			tx.Rollback()
			return nil, err
		}

		list = append(list, trans)
		err = tx.Commit()
		return list, err
	}
	return list, nil
}
