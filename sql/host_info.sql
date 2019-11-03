DROP TABLE IF EXISTS `host_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `host_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `hostname` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '主机名',
  `address` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'IP地址',
  `db_kind` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '数据库种类，比如mysql、redis、tidb',
  `product` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '所属产品线',
  `cluster_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '集群名称',
  `role` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '在集群中的角色',
  `port` int(11) NOT NULL DEFAULT '0' COMMENT '端口',
  `add_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '添加时间',
  PRIMARY KEY (`id`),
  KEY `idx_hostname` (`hostname`),
  KEY `idx_db_kind` (`db_kind`),
  KEY `idx_cluster_name` (`cluster_name`),
  KEY `idx_role` (`role`),
  UNIQUE KEY `unq_address_port` (`address`,`port`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci AUTO_INCREMENT=180002 COMMENT='数据库、主机对应信息';
