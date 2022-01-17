-- Adminer 4.8.1 MySQL 5.5.5-10.6.5-MariaDB dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

USE `config_storage`;

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
                       `id` int(11) NOT NULL AUTO_INCREMENT,
                       `op` text NOT NULL,
                       `namespace` int(11) NOT NULL,
                       `user` int(11) NOT NULL,
                       `status` int(1) NOT NULL COMMENT '0-ok 1-commied -1- restore',
                       PRIMARY KEY (`id`),
                       KEY `namespace` (`namespace`),
                       KEY `user` (`user`),
                       CONSTRAINT `log_ibfk_1` FOREIGN KEY (`namespace`) REFERENCES `namespace` (`id`),
                       CONSTRAINT `log_ibfk_2` FOREIGN KEY (`user`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `namespace`;
CREATE TABLE `namespace` (
                             `id` int(11) NOT NULL AUTO_INCREMENT,
                             `name` text NOT NULL,
                             `owner` int(11) NOT NULL,
                             PRIMARY KEY (`id`),
                             KEY `owner` (`owner`),
                             KEY `name` (`name`(768)),
                             CONSTRAINT `namespace_ibfk_1` FOREIGN KEY (`owner`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` int(11) NOT NULL AUTO_INCREMENT,
                        `username` varchar(20) NOT NULL,
                        `password` varchar(32) NOT NULL,
                        `is_admin` tinyint(1) NOT NULL,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- 2022-01-17 07:44:55
-- Adminer 4.8.1 MySQL 5.5.5-10.6.5-MariaDB dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

USE `config_storage`;

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
                       `id` int(11) NOT NULL AUTO_INCREMENT,
                       `op` text NOT NULL,
                       `namespace` int(11) NOT NULL,
                       `user` int(11) NOT NULL,
                       `status` int(1) NOT NULL COMMENT '0-ok 1-commied -1- restore',
                       PRIMARY KEY (`id`),
                       KEY `namespace` (`namespace`),
                       KEY `user` (`user`),
                       CONSTRAINT `log_ibfk_1` FOREIGN KEY (`namespace`) REFERENCES `namespace` (`id`),
                       CONSTRAINT `log_ibfk_2` FOREIGN KEY (`user`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `namespace`;
CREATE TABLE `namespace` (
                             `id` int(11) NOT NULL AUTO_INCREMENT,
                             `name` text NOT NULL,
                             `owner` int(11) NOT NULL,
                             PRIMARY KEY (`id`),
                             KEY `owner` (`owner`),
                             KEY `name` (`name`(768)),
                             CONSTRAINT `namespace_ibfk_1` FOREIGN KEY (`owner`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` int(11) NOT NULL AUTO_INCREMENT,
                        `username` varchar(20) NOT NULL,
                        `password` varchar(32) NOT NULL,
                        `is_admin` tinyint(1) NOT NULL,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- 2022-01-17 07:44:55
