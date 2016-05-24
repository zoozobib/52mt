#最近一次收益金额和积分
SELECT
	c_id,
	mobile_phone_number,
	reg_time,
	first_buy_time,
	first_buy_amount,
	income,
	integral
FROM
	cus_account
WHERE
	mobile_phone_number = '189****3704'
ORDER BY
	first_buy_time DESC
LIMIT 1;

#今天收益金额和积分
SELECT
	c_id,
	mobile_phone_number,
	reg_time,
	first_buy_time,
	first_buy_amount,
	income,
	integral
FROM
	cus_account
WHERE
	mobile_phone_number = '189****3704'
AND to_days(first_buy_time) = to_days(now());

#近七天收益金额和积分
SELECT
	c_id,
	mobile_phone_number,
	reg_time,
	first_buy_time,
	first_buy_amount,
	income,
	integral
FROM
	cus_account
WHERE
	mobile_phone_number = '189****3704'
AND DATE_SUB(CURDATE(), INTERVAL 7 DAY) <= date(first_buy_time);

#近30天收益金额和积分
SELECT
	c_id,
	mobile_phone_number,
	reg_time,
	first_buy_time,
	first_buy_amount,
	income,
	integral
FROM
	cus_account
WHERE
	mobile_phone_number = '1315***8575'
AND DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= date(first_buy_time);

#近30天收益金额和积分
SELECT
	c_id,
	mobile_phone_number,
	reg_time,
	first_buy_time,
	first_buy_amount,
	income,
	integral
FROM
	cus_account
WHERE
	mobile_phone_number = '189****3704'
AND DATE_FORMAT(first_buy_time, '%Y%m') = DATE_FORMAT(CURDATE(), '%Y%m');
