-- MySQL dump 10.14  Distrib 5.5.64-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: ticket_system
-- ------------------------------------------------------
-- Server version	5.5.64-MariaDB

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
-- Current Database: `ticket_system`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `ticket_system` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `ticket_system`;

--
-- Table structure for table `t_announcement`
--

DROP TABLE IF EXISTS `t_announcement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_announcement` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `t_number` varchar(32) DEFAULT NULL COMMENT '公告编号',
  `t_title` varchar(32) DEFAULT NULL COMMENT '公告标题',
  `t_center` text COMMENT '公告内容',
  `t_initiatorid` varchar(32) DEFAULT NULL COMMENT '公告发起人id',
  `t_create` varchar(64) NOT NULL COMMENT '公告创建时间',
  `t_update` varchar(64) NOT NULL COMMENT '公告更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_announcement`
--

LOCK TABLES `t_announcement` WRITE;
/*!40000 ALTER TABLE `t_announcement` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_announcement` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_dperson`
--

DROP TABLE IF EXISTS `t_dperson`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_dperson` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `t_dusername` varchar(32) DEFAULT NULL COMMENT '接收人姓名',
  `t_dphone` varchar(14) NOT NULL COMMENT '接收人联系方式',
  `t_number` varchar(64) NOT NULL COMMENT '接收人编号',
  `t_function` varchar(32) DEFAULT NULL COMMENT '接收人的职位',
  `t_lock` int(1) DEFAULT NULL,
  `t_passwd` varchar(64) NOT NULL,
  `t_org` varchar(32) DEFAULT NULL,
  `t_type` varchar(32) NOT NULL,
  `is_dis` smallint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQUE` (`t_dphone`)
) ENGINE=InnoDB AUTO_INCREMENT=172 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_dperson`
--

LOCK TABLES `t_dperson` WRITE;
/*!40000 ALTER TABLE `t_dperson` DISABLE KEYS */;
INSERT INTO `t_dperson` VALUES (152,'高斌','2147483647','c371138c-9614-4ce9-a91c-3c030794','运维专员',0,'123123123','信息部','普通员工',1),(153,'霍建华','4294967294','c371138c-9614-4ce9-a91c-3c030710','运维专员',0,'123123','信息部','管理员',1),(166,'小豆','15847449876','5ae4c8e2-5288-48a7-8a04-bbba8c9f','运维工程师',0,'123123','信息部','普通员工',1),(167,'绿豆','15978946513','5ab32146-7342-45ce-be6b-fd47b93b','影片制作',0,'123123','信息部','普通员工',1),(168,'华建年','1587964655','e244848d-acbb-4e87-8c02-c5a0722c','运维工程师',0,'3123123','信息部','普通员工',0),(169,'小白菜专属','18877776666','9566c74d-1003-7c4d-7bbb-0407d1e2c649','私人顾问',0,'123123','信息部','普通员工',0),(170,'周润发','16547812345','81855ad8-681d-0d86-d1e9-1e00167939cb','大明星',0,'123123','信息部','普通员工',0),(171,'小宝','9874561233','95af5a25-3679-51ba-a2ff-6cd471c483f1','大宝专属',0,'123123','信息部','普通员工',0);
/*!40000 ALTER TABLE `t_dperson` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_info`
--

DROP TABLE IF EXISTS `t_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `t_number` varchar(32) DEFAULT NULL COMMENT '工单编号',
  `t_type` varchar(16) NOT NULL DEFAULT '0' COMMENT '工单类型',
  `t_name` varchar(64) DEFAULT NULL COMMENT '工单名称',
  `t_org` varchar(16) NOT NULL DEFAULT '0' COMMENT '工单组织',
  `t_disnameid` varchar(64) DEFAULT NULL COMMENT '工单处理人员id',
  `t_srcnameid` varchar(64) DEFAULT NULL COMMENT '工单发起人员id',
  `t_status` varchar(16) DEFAULT '0' COMMENT '工单状态',
  `t_extend` varchar(128) DEFAULT NULL COMMENT '工单备注表',
  `t_create` varchar(64) DEFAULT NULL,
  `t_endtime` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `t_number` (`t_number`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_info`
--

LOCK TABLES `t_info` WRITE;
/*!40000 ALTER TABLE `t_info` DISABLE KEYS */;
INSERT INTO `t_info` VALUES (1,'5065517a-3123-4c9f-ae67-0c84cb1f','咨询','0: zabbix agent 挂了','系统部','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:27','2020-04-25 18:24:26',''),(2,'524897e5-b85a-4561-8cd1-56b5af75','网络','1: zabbix agent 挂了','系统部','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:28','2020-04-25 18:24:27',''),(3,'dc1fd99d-c0a7-4bee-821c-66b87665','网络','2: zabbix agent 挂了','系统部','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:29','2020-04-25 18:24:28',''),(4,'ad44bec0-8be8-4a16-a97a-447feb93','网络','3: zabbix agent 挂了','系统部','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:30','2020-04-25 18:24:29',''),(5,'1be8ac1f-e528-4468-b687-4a43dea3','网络','4: zabbix agent 挂了','系统部','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:31','2020-04-25 18:24:30',''),(6,'1edc5ac3-8a22-42bf-ba6b-10e4f97c','网络','5: zabbix agent 挂了','系统部','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:32','2020-04-25 18:24:31',''),(7,'a0713e55-5d88-40ff-b712-bafb4ef5','系统','6: zabbix agent 挂了','系统部','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:33','2020-04-25 18:24:32',''),(8,'e3cf5e6c-12b6-41a6-a9f0-17ce8634','系统','7: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:34','2020-04-25 18:24:33',''),(9,'a9b982dd-5167-4da4-81d0-b42d117e','系统','8: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:35','2020-04-25 18:24:34',''),(10,'3410ce67-645d-4436-ba21-07551091','系统','9: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:36','2020-04-25 18:24:35',''),(11,'c358d102-ec81-4208-afc2-aea55ccf','系统','10: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:37','2020-04-25 18:24:36',''),(12,'ce5071cd-b418-4f2d-90d6-c06befac','系统','11: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:38','2020-04-25 18:24:37',''),(13,'78664404-b1a4-4c94-b9cf-6d5cd7fd','系统','12: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:39','2020-04-25 18:24:38',''),(14,'197360c8-7ec7-4989-9762-ded66109','系统','13: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:40','2020-04-25 18:24:39',''),(15,'829da9d4-2572-48f5-be51-8638b64d','服务','14: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','完成','2020-04-25 18:24:41','2020-04-25 18:24:40','2020-04-25 19:24:40'),(16,'80bf60e5-4ff8-4143-aaf8-54450970','服务','15: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','完成','2020-04-25 18:24:42','2020-04-25 18:24:41','2020-04-25 20:24:40'),(17,'dd16f43b-830c-4b2f-8f92-d8884311','服务','16: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:43','2020-04-25 18:24:42',''),(18,'29e43af1-3fe2-42e6-88fa-eec16985','服务','17: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:44','2020-04-25 18:24:43',''),(19,'9c83b7c0-c287-4eb1-ac81-306d769d','耗材','18: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:45','2020-04-25 18:24:44',''),(20,'73a528b8-b194-42a6-a826-33604d6e','耗材','19: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:46','2020-04-25 18:24:45',''),(21,'be3ce67f-53eb-4978-b1e6-774e07eb','耗材','20: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:47','2020-04-25 18:24:46',''),(22,'4101c78c-e6c4-49c7-ba9d-bdf3b54e','耗材','21: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:48','2020-04-25 18:24:47',''),(23,'7a102f3f-c26d-43e4-97ca-4254abc2','耗材','22: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:49','2020-04-25 18:24:48',''),(24,'c001627d-2975-4619-9fd0-56cea4dd','耗材','23: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:50','2020-04-25 18:24:49',''),(25,'1f3dee7e-7e1a-40ab-83d4-51bed63d','耗材','24: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','处理中','2020-04-25 18:24:51','2020-04-25 18:24:50',''),(26,'3f03c5a8-d6c9-4a93-b6e5-da15498a','耗材','25: zabbix agent 挂了','物业管理','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','已挂起','2020-04-25 18:24:52','2020-04-25 18:24:51',''),(27,'d5d940f1-0deb-4d5c-b877-5446dff7','耗材','26: zabbix agent 挂了','出料中心','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','已挂起','2020-04-25 18:24:53','2020-04-25 18:24:52',''),(28,'0617b6b9-a4f1-4416-b706-24731096','耗材','27: zabbix agent 挂了','出料中心','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','已挂起','2020-04-25 18:24:54','2020-04-25 18:24:53',''),(29,'b959b310-c7d1-4a7f-b028-49e628b8','其他','28: zabbix agent 挂了','出料中心','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','已挂起','2020-04-25 18:24:55','2020-04-25 18:24:54',''),(30,'adf524d8-58a7-41e9-bc8a-090226a7','其他','29: zabbix agent 挂了','出料中心','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','已挂起','2020-04-25 18:24:56','2020-04-25 18:24:55',''),(31,'1e200c6b-e353-42e7-92aa-27bdb171','其他','30: zabbix agent 挂了','出料中心','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','已挂起','2020-04-25 18:24:57','2020-04-25 18:24:56',''),(32,'c4155918-ef64-473c-a318-8fe0c065','其他','31: zabbix agent 挂了','出料中心','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','完成','2020-04-25 18:24:58','2020-04-25 18:24:57','2020-04-25 18:24:58'),(33,'54ec540a-6618-4b7c-a773-bdb6689d','其他','32: zabbix agent 挂了','出料中心','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','完成','2020-04-25 18:24:59','2020-04-25 18:24:58','2020-04-25 18:24:58'),(34,'3b63ee1c-f845-449d-a7bf-c3c47678','网络','33: zabbix agent 挂了','出料中心','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','完成','miss','2020-04-25 18:24:59','2020-05-03 10:52:25'),(35,'004a6342-27d9-4e77-9670-b4d2f4b4','网络','34: zabbix agent 挂了','出料中心','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','完成','2020-04-25 18:25:01','2020-04-25 18:25:00','2020-04-25 19:24:58'),(36,'9b606c32-fff4-463b-b8c2-8fbccd8b','网络','35: zabbix agent 挂了','出料中心','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','完成','2020-04-25 18:25:02','2020-04-25 18:25:01','2020-04-25 18:24:58'),(37,'5fbafce5-af06-4e7d-bde4-585b747b','网络','36: zabbix agent 挂了','出料中心','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','完成','早就弄忘了， 一直忘记关单','2020-04-25 18:25:02','2020-05-02 15:55:47'),(38,'b4fb79fa-62bc-4dce-9a09-09799f39','网络','37: zabbix agent 挂了','出料中心','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','已挂起','2020-04-25 18:25:04','2020-04-25 18:25:03',''),(39,'aef4ab75-827d-4743-a716-d70c98a9','系统','38: zabbix agent 挂了','出料中心','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','已挂起','2020-04-25 18:25:05','2020-04-25 18:25:04',''),(40,'89b100be-768d-49ec-9aa6-bdc91b48','系统','39: zabbix agent 挂了','出料中心','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024c9e','完成','2020-04-25 18:25:06','2020-04-25 18:25:05','2020-05-03 10:48:43'),(42,'46702332-07ab-420c-8c6e-899a26b0','耗材','B210维修打印机','物业管理','c371138c-9614-4ce9-a91c-3c030710','5d03367c-6b56-4567-9463-e6024sdf','已挂起','由于问题复杂，暂时挂起','2020-05-01 22:31:40',''),(43,'7a2eea9f-d8f0-4e95-bca2-19841491','网络','210办公室网络不通','物业管理','','5d03367c-6b56-4567-9463-e6024sCG','待接手','仓储中心 2区 A楼 A210办公室 不能上网','2020-05-02 07:55:43',''),(44,'075e9fde-088f-4e87-887b-6b5f359f','服务','510爱奇艺视频软件打不开了','信息部','','5d03367c-6b56-4567-9463-e6024sCG','已挂起','玩什么爱奇艺  ，一边去','2020-05-02 07:57:23',''),(45,'e09ca656-7900-4fdb-8c63-f6d56677','服务','笔记本电脑安装操作系统','信息部','c371138c-9614-4ce9-a91c-3c030794','5d03367c-6b56-4567-9463-e6024sCG','完成',' 不用修了， 哥换新电脑了','2020-05-02 13:35:18','2020-05-02 15:57:43'),(52,'cca91788-285d-4083-b94b-26f3f6c9','服务','远程桌面连不上','信息部','','5d03367c-6b56-4567-9463-e6024sCG','已挂起','远程vpn连接不上，提示远程连接404','2020-05-02 18:35:04',''),(53,'dbcdfff6-4fb5-40d6-af54-0bb367cf','服务','测试发布工单','运维部','c371138c-9614-4ce9-a91c-3c030710','5d03367c-6b56-4567-9463-e6024sCG','完成','处理完成','2020-05-02 18:39:34','2020-05-03 11:16:30'),(54,'433a1f22-341e-40b3-9075-55b4b26b','系统','远程桌面无法连接','运维部','5ae4c8e2-5288-48a7-8a04-bbba8c9f','5d03367c-6b56-4567-9463-e6024sCG','完成','小绿豆，圆又圆。 解决事情，转一圈','2020-05-03 10:49:58','2020-05-03 10:55:29'),(55,'52fdfc07-2182-654f-163f-5f0f9a62','服务','电脑无法上网','信息部','9566c74d-1003-7c4d-7bbb-0407d1e2c649','9566c74d-1003-7c4d-7bbb-0407d1e2','已挂起','先过来唠会','2020-05-03 11:30:40',''),(57,'6694d2c4-22ac-d208-a007-2939487f','服务','协助视察工作','信息部','81855ad8-681d-0d86-d1e9-1e00167939cb','a1f093d3-d3aa-4f97-8a82-6cbd35d7','完成','来个搓搓','2020-05-03 11:38:24','2020-05-03 11:39:25');
/*!40000 ALTER TABLE `t_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_org`
--

DROP TABLE IF EXISTS `t_org`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_org` (
  `id` int(2) NOT NULL AUTO_INCREMENT,
  `t_name` varchar(32) DEFAULT NULL,
  KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_org`
--

LOCK TABLES `t_org` WRITE;
/*!40000 ALTER TABLE `t_org` DISABLE KEYS */;
INSERT INTO `t_org` VALUES (0,'信息部'),(1,'运维部'),(2,'物业管理'),(3,'出料中心'),(4,'系统部');
/*!40000 ALTER TABLE `t_org` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_sperson`
--

DROP TABLE IF EXISTS `t_sperson`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_sperson` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `t_username` varchar(32) DEFAULT NULL COMMENT '发单人姓名',
  `t_phone` varchar(14) NOT NULL COMMENT '发单人联系手机',
  `t_addr` varchar(64) DEFAULT NULL COMMENT '发单人联系地址',
  `t_number` varchar(32) NOT NULL COMMENT '发单人编号',
  `t_lock` int(1) DEFAULT NULL,
  `t_passwd` varchar(64) NOT NULL,
  `t_org` varchar(32) DEFAULT NULL,
  `t_type` varchar(32) NOT NULL,
  `is_src` smallint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `t_phone` (`t_phone`)
) ENGINE=InnoDB AUTO_INCREMENT=117 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_sperson`
--

LOCK TABLES `t_sperson` WRITE;
/*!40000 ALTER TABLE `t_sperson` DISABLE KEYS */;
INSERT INTO `t_sperson` VALUES (1,'太叔凤','2147483648','中煤鄂能化工','5d03367c-6b56-4567-9463-e6024c9e',0,'123123','信息部','实习员工',1),(105,'柳林','4294967295','中煤鄂能化工','5d03367c-6b56-4567-9463-e6024sdf',0,'123123','信息部','普通员工',1),(110,'master','15847445046','中煤鄂能化工六号楼B210','5d03367c-6b56-4567-9463-e6024sCG',0,'123123','信息部','管理员',1),(111,'豆豆','15847445478','办公B楼B210','a1f093d3-d3aa-4f97-8a82-6cbd35d7',0,'123123','系统部','普通员工',1),(112,'黑头盔','16758914563','办公八楼B801','428ef4b1-b3c9-444b-90e3-19a8ba33',0,'123123','系统部','实习员工',1),(113,'艾小豆','18867587946','办公区7楼A712','150da71b-0b62-45af-99be-d5038e23',0,'123123','物业管理','实习员工',1),(114,'小白菜','17788779966','厂区2楼物业办公室','52fdfc07-2182-654f-163f-5f0f9a62',0,'123123','物业管理','普通员工',1),(115,'小圆菜','16655487455','厂区1楼a111','9566c74d-1003-7c4d-7bbb-0407d1e2',0,'123123','物业管理','普通员工',1),(116,'大宝','1234566789','A201','eb9d18a4-4784-045d-87f3-c67cf227',0,'123123','信息部','普通员工',1);
/*!40000 ALTER TABLE `t_sperson` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_sptype`
--

DROP TABLE IF EXISTS `t_sptype`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_sptype` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `t_name` varchar(32) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_sptype`
--

LOCK TABLES `t_sptype` WRITE;
/*!40000 ALTER TABLE `t_sptype` DISABLE KEYS */;
INSERT INTO `t_sptype` VALUES (0,'普通员工'),(1,'实习员工'),(2,'管理员'),(3,'超级管理员');
/*!40000 ALTER TABLE `t_sptype` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_status`
--

DROP TABLE IF EXISTS `t_status`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_status` (
  `id` int(4) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_status`
--

LOCK TABLES `t_status` WRITE;
/*!40000 ALTER TABLE `t_status` DISABLE KEYS */;
INSERT INTO `t_status` VALUES (1,'待接手'),(2,'处理中'),(3,'已挂起'),(4,'完成');
/*!40000 ALTER TABLE `t_status` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_type`
--

DROP TABLE IF EXISTS `t_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t_type` (
  `id` int(2) NOT NULL AUTO_INCREMENT,
  `t_name` varchar(32) DEFAULT NULL COMMENT '工单类型',
  KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_type`
--

LOCK TABLES `t_type` WRITE;
/*!40000 ALTER TABLE `t_type` DISABLE KEYS */;
INSERT INTO `t_type` VALUES (0,'耗材'),(1,'系统'),(2,'服务'),(3,'网络'),(4,'咨询'),(5,'其他');
/*!40000 ALTER TABLE `t_type` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-05-03 12:47:33
