package utils

import "time"

const (
	TIMEFORMAT      = "2006-01-02 15:04:05"
	DATEFORMAT      = "2006-01-02"
	UPDATE_PRODUCT  = `update m_product set id=?, product_code=?,product_name=?,category_id=?,updated_at=? where id=?`
	ALL_TRANSACTION = `select td.oi as order_id, 
			td.pname as product_name,
			td.cname as category_name,
			td.pcode as product_code,
			td.price as price,
			td.qty as quantity,
			td.od as order_date,
			td.oun as outlet_name, 
			r.regional_name as regional_name
	from(
			select 	tc.oi as oi,
					tc.prodname as pname,
					tc.catename as cname,
					tc.prodcod as pcode,
					tc.price as price,
					tc.qty as qty,
					tc.ord_date as od,
					o.outlet_name as oun,
					o.regional as region 
			from(
					select 
						tb.order_id as oi,
						tb.product_name as prodname,
						c.category_name as catename,
						tb.product_code as prodcod,
						tb.price as price,
						tb.qty as qty,
						tb.order_date as ord_date,
						tb.outlet_code as o_code,
						tb.is_active 
					from(
						select 
							ta.order_id as order_id,
							p.product_name as product_name,
							p.product_code as product_code,
							p.category_id as category_id,
							ta.qty as qty,
							ta.price as price,
							ta.order_date as order_date,
							ta.outlet_code as outlet_code,
							ta.is_active as is_active
						from(
							select 
								tbl_po.order_id as order_id,
								pp.product_price_id as product_price_id,
								tbl_po.order_date as order_date,
								tbl_po.outlet_code as outlet_code,
								tbl_po.qty as qty,
								pp.product_id as product_id,
								pp.product_price as price,
								pp.is_active as is_active
							from(
								select 
									poi.order_id as order_id,
									poi.product_id as product_id,
									po.order_date as order_date,
									po.outlet_code as outlet_code,
									po.id as poid,
									poi.qty as qty
								from purchase_order_item poi 
								join purchase_order po on poi.order_id = po.id
							)tbl_po
							join m_product_price pp on tbl_po.product_id=pp.product_price_id
						)ta 
						join m_product p on ta.product_id = p.id
					)tb
					join m_category c on tb.category_id = c.id
			)tc
			join m_outlet o on o.outlet_code = tc.o_code
	)td
	join m_region r on td.region = r.regional_code`

	THIS_DAY   = `where date(td.od)=?`
	THIS_MONTH = `where month(td.od)=? or monthname(td.od)=?`

	ADD_PRODUCT       = `insert into m_product(id,product_code,product_name,category_id,created_at,updated_at) values(?,?,?,?,?,?)`
	ADD_CATEGORY      = `insert into m_category(id,category_name,created_at,updated_at) values(?,?,?,?)`
	DELETE_PRODUCT    = `delete from m_product where id=?`
	GET_ALL_PRODUCTS  = `select p.id,p.product_code,p.product_name,c.id, c.category_name from m_product p join m_category c on p.category_id = c.id`
	GET_PRODUCT_BY_ID = `select p.id,p.product_code,p.product_name,c.id, c.category_name from m_product p join m_category c on p.category_id = c.id where p.id=?`
)

var CurDate = time.Now().Format(TIMEFORMAT)
var Date = time.Now().Format(DATEFORMAT)
