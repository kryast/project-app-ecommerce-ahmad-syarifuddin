-- Memasukkan ratings_id = 5 untuk transaction_items_id = 1
INSERT INTO reviews (transaction_items_id, ratings_id)
SELECT ti.id, 5
FROM transaction_items ti
WHERE ti.id = 1  -- transaction_items_id yang dipilih
AND ti.transaction_id = 2;  -- Memastikan hanya jika transaction_id = 2


insert into product_details (product_id, category_id, photo_id) values
(1,1,1),
(2,1,1),
(3,2,1),
(4,2,1),
(5,3,1),
(6,3,1),
(7,4,1),
(8,4,1),
(9,5,1),
(10,5,1);

insert into transaction_items (transaction_id, product_details_id, quantity) values
(1,1,5),
(2,1,5),
(2,2,3),
(2,3,4),
(2,1,2);

 SELECT
    pd.id,
    p.name AS product_name,
    p.detail AS product_detail,
    p.price AS product_price,
    c.name AS category_name,
    ph.photo_url
  FROM
    product_details pd
  JOIN
    products p ON p.id = pd.product_id
  JOIN
    categories c ON c.id = pd.category_id
  JOIN
    photos ph ON ph.id = pd.photo_id
    WHERE
    pd.id = 1;

 SELECT 
			ti.id,
            p.name AS product_name,
            SUM(ti.quantity) AS total_quantity
        FROM 
            transaction_items ti
        JOIN 
            product_details pd ON pd.id = ti.product_details_id
        JOIN 
            products p ON p.id = pd.product_id
        JOIN 
            transactions t ON t.id = ti.transaction_id
        WHERE 
            t.id = 1 
        GROUP BY 
			ti.id ,p.name
        ORDER BY 
            total_quantity DESC
        LIMIT 5;