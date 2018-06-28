DROP TABLE IF EXISTS `sln_assign`;
CREATE TABLE `sln_assign` (
  `id`                  int(11)       NOT NULL AUTO_INCREMENT,
  `sln_no`              varchar(40)   NOT NULL,
  `supplier_id`         int(11)       NOT NULL,
  `add_time`            int(11)       NOT NULL,
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;