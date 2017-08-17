CREATE DATABASE `mybanana` CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `mybanana`;

CREATE TABLE `banana` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
