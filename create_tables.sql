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
      FOREIGN KEY (`customer_id`) REFERENCES `customers`(`customer_id`)
	) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


CREATE TABLE `offers` (
	  `offer_id` bigint(20) NOT NULL AUTO_INCREMENT,
	  `account_id` bigint(20) NOT NULL,
	  `limit_type` varchar(255) NOT NULL,
	  `new_limit` int(11) NOT NULL,
	  `activation_time` bigint(20) NOT NULL,
	  `expiration_time` bigint(20) NOT NULL,
	  `status` varchar(255) NOT NULL,
	  PRIMARY KEY (`offer_id`),
	  FOREIGN KEY (`account_id`) REFERENCES `accounts`(`account_id`)
	) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;