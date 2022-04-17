-- Adminer 4.8.1 MySQL 5.5.5-10.7.3-MariaDB dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `cluster`;
CREATE TABLE `cluster` (
  `raft_id` varchar(64) NOT NULL,
  `address` text NOT NULL,
  UNIQUE KEY `raftid` (`raft_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` int(11) NOT NULL COMMENT '0 - set 1 - del',
  `key` varchar(255) NOT NULL,
  `value` varchar(255) NOT NULL,
  `namespace_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `status` int(1) NOT NULL DEFAULT 0 COMMENT '0-ok 1-commied -1- restore',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `namespace_id` (`namespace_id`),
  CONSTRAINT `log_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `log_ibfk_3` FOREIGN KEY (`namespace_id`) REFERENCES `namespace` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `namespace`;
CREATE TABLE `namespace` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `user_id` int(11) NOT NULL,
  `raft_id` varchar(64) NOT NULL,
  `private_key` varchar(64) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING HASH,
  KEY `user_id` (`user_id`),
  KEY `raft_id` (`raft_id`),
  CONSTRAINT `namespace_ibfk_3` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `namespace_ibfk_4` FOREIGN KEY (`raft_id`) REFERENCES `cluster` (`raft_id`) ON DELETE CASCADE ON UPDATE CASCADE
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

INSERT INTO `user` (`id`, `username`, `password`, `is_admin`) VALUES
(2,	'lxh001',	'87d9bb400c0634691f0e3baaf1e2fd0d',	1);

DROP TABLE IF EXISTS `user_cluster`;
CREATE TABLE `user_cluster` (
  `user_id` int(11) NOT NULL,
  `raft_id` varchar(64) NOT NULL,
  UNIQUE KEY `user_id_raft_id` (`user_id`,`raft_id`),
  KEY `raft_id` (`raft_id`),
  CONSTRAINT `user_cluster_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `user_cluster_ibfk_3` FOREIGN KEY (`raft_id`) REFERENCES `cluster` (`raft_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `user_namespace`;
CREATE TABLE `user_namespace` (
  `user_id` int(11) NOT NULL,
  `namespace_id` int(11) NOT NULL,
  `type` int(11) NOT NULL COMMENT '0 - Owner 1 - Normal 2 - RO',
  UNIQUE KEY `user_id_namespace_id` (`user_id`,`namespace_id`),
  KEY `namespace_id` (`namespace_id`),
  CONSTRAINT `user_namespace_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `user_namespace_ibfk_3` FOREIGN KEY (`namespace_id`) REFERENCES `namespace` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- 2022-04-17 07:10:52
