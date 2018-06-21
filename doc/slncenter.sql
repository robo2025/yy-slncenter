SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sln_basic_info for sewage
-- --------------------------------------------------------------------------
DROP TABLE IF EXISTS `sewage_info`;
CREATE TABLE `sewage_info` (
  `id`                  int(11)       NOT NULL AUTO_INCREMENT,
  `sln_no`              varchar(20)   NOT NULL,
  `sewage_business`     varchar(20)   DEFAULT NULL,
  `sewage_scenario`     varchar(20)   DEFAULT NULL,
  `tech_method`         varchar(20)   DEFAULT NULL,
  `general_norm`        varchar(20)   DEFAULT NULL,
  `other_norm`          varchar(20)   DEFAULT NULL,
  `daily_capacity`      decimal(8, 2) DEFAULT NULL,
  `disinfector`         int(11)       DEFAULT NULL,
  `valve`               int(11)       DEFAULT NULL,
  `blower`              int(11)       DEFAULT NULL,
  `stirrer`             int(11)       DEFAULT NULL,
  `pump`                int(11)       DEFAULT NULL,
  `doser`               int(11)       DEFAULT NULL,
  `aux_equipment_nums`  int(11)       DEFAULT NULL,
  `total_equipment_nums` int(11)      DEFAULT NULL,
  `operating_size`      decimal(8,1)  DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sln_no` (`sln_no`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
-- ----------------------------
-- Table structure for sln_basic_info
-- ----------------------------
DROP TABLE IF EXISTS `sln_basic_info`;
CREATE TABLE `sln_basic_info` (
  `id`             int(11)     NOT NULL AUTO_INCREMENT,
  `sln_no`         varchar(20) NOT NULL,
  `sln_name`       varchar(20)          DEFAULT NULL,
  `sln_type`       varchar(20)          DEFAULT NULL,
  `sln_date`       int(11)              DEFAULT NULL,
  `sln_expired`    int(11)              DEFAULT NULL,
  `customer_id`    int(11)     NOT NULL,
  `customer_price` decimal(12, 2)       DEFAULT NULL,
  `supplier_id`    int(11)              DEFAULT NULL,
  `supplier_price` decimal(12, 2)       DEFAULT NULL,
  `sln_status`     varchar(5)           DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sln_no` (`sln_no`),
  KEY `customer_id` (`customer_id`) USING BTREE
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for sln_supplier_info
-- ----------------------------
DROP TABLE IF EXISTS `sln_supplier_info`;
CREATE TABLE `sln_supplier_info` (
  `id`            int(11)     NOT NULL AUTO_INCREMENT,
  `sln_no`        varchar(20) NOT NULL,
  `user_id`       int(11)     NOT NULL,
  `total_price`   decimal(10, 2)       DEFAULT NULL,
  `freight_price` decimal(11, 2)       DEFAULT NULL,
  `pay_ratio`     smallint(6)          DEFAULT NULL,
  `expired_date`  int(11)              DEFAULT NULL,
  `delivery_date` smallint(6)          DEFAULT NULL,
  `sln_desc`      text,
  `sln_note`      varchar(255)         DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sln_uniq` (`sln_no`, `user_id`) USING BTREE,
  KEY `sln_no` (`sln_no`) USING BTREE
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for sln_user_info
-- ----------------------------
DROP TABLE IF EXISTS `sln_user_info`;
CREATE TABLE `sln_user_info` (
  `id`           int(11)     NOT NULL AUTO_INCREMENT,
  `sln_no`       varchar(20) NOT NULL,
  `pay_ratio`    int(11)              DEFAULT NULL,
  `welding_name` varchar(60)          DEFAULT NULL,
  `sln_note` varchar(300)         DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sln_no` (`sln_no`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for welding_device
-- ----------------------------
DROP TABLE IF EXISTS `sln_device`;
CREATE TABLE `sln_device` (
  `id`               int(11)     NOT NULL AUTO_INCREMENT,
  `sln_no`           varchar(20) NOT NULL,
  `user_id`          int(11)     NOT NULL,
  `sln_role`         varchar(2)  NOT NULL DEFAULT 'C',
  `device_id`        varchar(20)          DEFAULT NULL,
  `device_type`      varchar(30)          DEFAULT NULL,
  `device_component` varchar(30)          DEFAULT NULL,
  `device_name`      varchar(50)          DEFAULT NULL,
  `device_model`     varchar(255)         DEFAULT NULL,
  `device_price`     decimal(10, 2)       DEFAULT NULL,
  `device_num`       int(11)              DEFAULT NULL,
  `brand_name`       varchar(255)         DEFAULT NULL,
  `device_note`      varchar(255)         DEFAULT NULL,
  `device_origin`    varchar(50)          DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sln_no` (`sln_no`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for welding_file
-- ----------------------------
DROP TABLE IF EXISTS `sln_file`;
CREATE TABLE `sln_file` (
  `id`        int(11)     NOT NULL AUTO_INCREMENT,
  `sln_no`    varchar(20) NOT NULL,
  `user_id`   int(11)     NOT NULL,
  `sln_role`  varchar(2)  NOT NULL DEFAULT 'C',
  `file_name` varchar(100)         DEFAULT NULL,
  `file_type` varchar(20)          DEFAULT NULL,
  `file_url`  varchar(200)         DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sln_no` (`sln_no`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for welding_info
-- ----------------------------
DROP TABLE IF EXISTS `welding_info`;
CREATE TABLE `welding_info` (
  `id`                 int(11)     NOT NULL AUTO_INCREMENT,
  `sln_no`             varchar(20) NOT NULL,
  `welding_business`   varchar(20)          DEFAULT NULL,
  `welding_scenario`   varchar(20)          DEFAULT NULL,
  `welding_metal`      varchar(20)          DEFAULT NULL,
  `welding_efficiency` varchar(20)          DEFAULT NULL,
  `welding_splash`     varchar(20)          DEFAULT NULL,
  `welding_model`      varchar(20)          DEFAULT NULL,
  `welding_method`     varchar(20)          DEFAULT NULL,
  `welding_gas`        varchar(20)          DEFAULT NULL,
  `gas_cost`           varchar(20)          DEFAULT NULL,
  `max_height`         decimal(8, 2)        DEFAULT NULL,
  `max_radius`         decimal(8, 2)        DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sln_no` (`sln_no`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for welding_support
-- ----------------------------
DROP TABLE IF EXISTS `sln_support`;
CREATE TABLE `sln_support` (
  `id`      int(11)     NOT NULL AUTO_INCREMENT,
  `sln_no`  varchar(20) NOT NULL,
  `user_id` int(11)     NOT NULL,
  `name`    varchar(50)          DEFAULT NULL,
  `price`   decimal(11, 2)       DEFAULT NULL,
  `note`    varchar(255)         DEFAULT NULL,
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for welding_tech_param
-- ----------------------------
DROP TABLE IF EXISTS `welding_tech_param`;
CREATE TABLE `welding_tech_param` (
  `id`        int(11)     NOT NULL AUTO_INCREMENT,
  `sln_no`    varchar(20) NOT NULL,
  `user_id`   int(11)     NOT NULL,
  `name`      varchar(30)          DEFAULT NULL,
  `value`     varchar(255)         DEFAULT NULL,
  `unit_name` varchar(10)          DEFAULT NULL,
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
