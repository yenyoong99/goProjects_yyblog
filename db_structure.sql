/*
MySQL - 5.7.26 : Database - yyblog
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`yyblog` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `yyblog`;

/*Table structure for table `blogs` */

DROP TABLE IF EXISTS `blogs`;

CREATE TABLE `blogs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'article id',
  `tags` text NOT NULL COMMENT 'label',
  `created_at` int(11) NOT NULL COMMENT 'created time',
  `published_at` int(11) NOT NULL COMMENT 'publish time',
  `updated_at` int(11) NOT NULL COMMENT 'update time',
  `title` varchar(255) NOT NULL COMMENT 'article title',
  `author` varchar(255) NOT NULL COMMENT 'author',
  `content` text NOT NULL COMMENT 'article content',
  `status` tinyint(4) NOT NULL COMMENT 'article status',
  `summary` varchar(255) NOT NULL COMMENT 'article summary info',
  `create_by` varchar(255) NOT NULL COMMENT 'article',
  `audit_at` int(11) NOT NULL COMMENT 'audit review time',
  `is_audit_pass` tinyint(4) NOT NULL COMMENT 'check audit review is pass',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_title` (`title`) COMMENT 'titile add unique'
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4;

/*Data for the table `blogs` */

/*Table structure for table `tokens` */

DROP TABLE IF EXISTS `tokens`;

CREATE TABLE `tokens` (
  `created_at` int(11) NOT NULL COMMENT 'created time',
  `updated_at` int(11) NOT NULL COMMENT 'update time',
  `user_id` int(11) NOT NULL COMMENT 'user id',
  `username` varchar(255) NOT NULL COMMENT 'username, not allow duplicate',
  `access_token` varchar(255) NOT NULL COMMENT 'user access token',
  `access_token_expired_at` int(11) NOT NULL COMMENT 'token expired time',
  `refresh_token` varchar(255) NOT NULL COMMENT 'refresh token',
  `refresh_token_expired_at` int(11) NOT NULL COMMENT 'refresh token expired time',
  PRIMARY KEY (`access_token`) USING BTREE,
  UNIQUE KEY `idx_token` (`access_token`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `tokens` */

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` int(11) NOT NULL COMMENT 'created time',
  `updated_at` int(11) NOT NULL COMMENT 'update time',
  `username` varchar(255) NOT NULL COMMENT 'username, not allow duplicate',
  `password` varchar(255) NOT NULL COMMENT 'encryption password',
  `label` varchar(255) NOT NULL COMMENT 'user tag',
  `role` tinyint(4) NOT NULL COMMENT 'user role',
  `email` varchar(255) NOT NULL COMMENT `user email`
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_user` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

/*Data for the table `users` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
