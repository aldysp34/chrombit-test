-- Soal nomor 2
SELECT
    customer_id,
    customer_name,
    COUNT(order_id) AS total_orders
FROM
    orders
WHERE
    STR_TO_DATE(order_date, '%d-%m-%Y') >= DATE_SUB(NOW(), INTERVAL 6 MONTH)
GROUP BY
    customer_id, customer_name
ORDER BY
    total_orders DESC
LIMIT 5;


-- Soal nomor 3
SELECT *
FROM orders
WHERE STR_TO_DATE(order_date, '%d-%m-%Y') BETWEEN '2020-01-17' AND '2020-02-21';