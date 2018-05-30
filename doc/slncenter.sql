SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sln_basic_info
-- ----------------------------
DROP TABLE IF EXISTS `sln_basic_info`;
CREATE TABLE `sln_basic_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sln_no` varchar(20) NOT NULL,
  `customer_id` int(11) NOT NULL DEFAULT '0',
  `supplier_id` int(11) NOT NULL DEFAULT '0',
  `sln_name` varchar(20) DEFAULT NULL,
  `sln_num` int(11) DEFAULT NULL,
  `sln_date` datetime(6) NOT NULL,
  `customer_price` decimal(12,2) DEFAULT NULL,
  `supplier_price` decimal(12,2) DEFAULT NULL,
  `sln_status` varchar(5) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sln_no` (`sln_no`),
  KEY `customer_id` (`customer_id`,`supplier_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for sln_user_info
-- ----------------------------
DROP TABLE IF EXISTS `sln_user_info`;
CREATE TABLE `sln_user_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sln_no` varchar(20) NOT NULL,
  `pay_ratio` int(11) NOT NULL,
  `welding_name` varchar(60) NOT NULL,
  `welding_note` varchar(300) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sln_no` (`sln_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for welding_device
-- ----------------------------
DROP TABLE IF EXISTS `welding_device`;
CREATE TABLE `welding_device` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sln_no` varchar(20) NOT NULL,
  `sln_role` varchar(2) DEFAULT 'C',
  `device_id` int(11) DEFAULT '0',
  `device_type` varchar(20) DEFAULT NULL,
  `device_model` varchar(255) DEFAULT NULL,
  `device_price` decimal(10,2) DEFAULT NULL,
  `device_num` int(11) DEFAULT NULL,
  `brand_name` varchar(255) DEFAULT NULL,
  `device_note` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sln_no` (`sln_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for welding_file
-- ----------------------------
DROP TABLE IF EXISTS `welding_file`;
CREATE TABLE `welding_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sln_no` varchar(20) NOT NULL,
  `sln_role` varchar(2) DEFAULT 'C',
  `file_name` varchar(100) NOT NULL,
  `file_type` varchar(20) NOT NULL,
  `file_url` varchar(200) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `sln_no` (`sln_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for welding_info
-- ----------------------------
DROP TABLE IF EXISTS `welding_info`;
CREATE TABLE `welding_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sln_no` varchar(20) NOT NULL,
  `welding_business` varchar(20) NOT NULL,
  `welding_scenario` varchar(20) NOT NULL,
  `welding_metal` varchar(20) NOT NULL,
  `welding_efficiency` varchar(20) NOT NULL,
  `welding_splash` varchar(20) NOT NULL,
  `welding_model` varchar(20) NOT NULL,
  `welding_method` varchar(20) NOT NULL,
  `welding_gas` varchar(20) NOT NULL,
  `gas_cost` varchar(20) NOT NULL,
  `max_height` decimal(6,1) NOT NULL,
  `max_radius` decimal(5,2) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sln_no` (`sln_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
