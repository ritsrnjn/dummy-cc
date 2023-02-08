CREATE TABLE `accounts` (
	  `account_id` bigint(20) NOT NULL AUTO_INCREMENT,
	  `customer_id` varchar(255) NOT NULL,
	  `account_limit` int(11) NOT NULL,
	  `per_transaction_limit` int(11) NOT NULL,
	  `last_account_limit` int(11) NOT NULL,
	  `last_per_transaction_limit` int(11) NOT NULL,
	  `account_limit_update_time` bigint(20) NOT NULL,
	  `per_transaction_limit_update_time` bigint(20) NOT NULL,
	  PRIMARY KEY (`account_id`)	
	) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;