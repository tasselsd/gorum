-- MariaDB dump 10.19  Distrib 10.7.3-MariaDB, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: gorum
-- ------------------------------------------------------
-- Server version	10.7.1-MariaDB-1:10.7.1+maria~focal

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `t_comment`
--

DROP TABLE IF EXISTS `t_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `content` text DEFAULT NULL,
  `initiator_uid` bigint(20) DEFAULT NULL,
  `initiator` varchar(32) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `likes` bigint(20) DEFAULT NULL,
  `discuss_did` varchar(45) DEFAULT NULL,
  `sha1` char(40) DEFAULT NULL,
  `sha1_prefix` char(8) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_discuss`
--

DROP TABLE IF EXISTS `t_discuss`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_discuss` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `d_name` varchar(512) DEFAULT NULL,
  `content` text DEFAULT NULL,
  `initiator_uid` bigint(20) DEFAULT NULL,
  `initiator` varchar(32) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `likes` bigint(20) DEFAULT NULL COMMENT 'Number of likes',
  `division_rid` bigint(20) DEFAULT NULL,
  `division` varchar(32) DEFAULT NULL COMMENT 'Jurisdiction of region',
  `sha1` char(40) DEFAULT NULL,
  `sha1_prefix` char(8) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_label`
--

DROP TABLE IF EXISTS `t_label`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_label` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `l_name` varchar(32) DEFAULT NULL,
  `resource_id` bigint(20) DEFAULT NULL COMMENT 'which (discuss, region, user, comment etc.) resource to rel',
  `create_time` datetime DEFAULT NULL,
  `sha1` char(20) DEFAULT NULL,
  `sha1_prefix` char(8) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_region`
--

DROP TABLE IF EXISTS `t_region`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_region` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `r_name` varchar(32) DEFAULT NULL,
  `about` varchar(45) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `sha1` char(40) DEFAULT NULL,
  `sha1_prefix` char(8) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user`
--

DROP TABLE IF EXISTS `t_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `u_name` varchar(32) DEFAULT NULL,
  `avatar` varchar(512) DEFAULT NULL,
  `likes` bigint(20) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `passwd` char(40) DEFAULT NULL,
  `phone` varchar(15) DEFAULT NULL,
  `country_code` int(11) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `sha1` char(40) DEFAULT NULL,
  `sha1_prefix` char(8) DEFAULT NULL,
  `valid` int(11) DEFAULT 0,
  `once_token` char(80) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `u_name_UNIQUE` (`u_name`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-10 14:50:01
