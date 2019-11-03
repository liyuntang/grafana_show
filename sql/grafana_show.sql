DROP TABLE IF EXISTS `grafana_show`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `grafana_show` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `endpoint` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '数据库实例名称',
  `metric` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '监控项',
  `value` int(11) NOT NULL DEFAULT '0' COMMENT '监控值',
  `monit_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '监控时间',
  PRIMARY KEY (`id`),
  KEY `idx_monit_time_idx_metric` (`monit_time`,`metric`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci AUTO_INCREMENT=466695540 COMMENT='grafana展示用表';
