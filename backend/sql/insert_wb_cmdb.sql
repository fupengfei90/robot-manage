-- WB CMDB数据导入
INSERT INTO `wb_cmdb_info` (`system_name`, `environment`, `ip_address`, `port`, `status`, `owner`, `remark`) VALUES
('Wecube系统', '梁智隆
张文广', 'WeCube平台软件服务', 'auth-server
gateway
portal
core
S3', '4C8G+250G+1T（GFS）', 'TKO-MGMT
ST-MGMT', '10.111.2.11'),
('Wecube系统', '梁智隆
张文广', 'WeCube插件服务', 'Plugin  instance-Docker', '4C8G+250G
4C16G +250', 'TKO-MGMT
ST-MGMT', '10.111.2.15'),
('Wecube系统', '梁智隆
张文广', 'WeCube插件服务-监控', 'Plugin  instance-Docker', '4C16G+250G
4C8G+250G', 'TKO-MGMT
ST-MGMT', '10.111.2.10'),
('Wecube系统', '梁智隆
张文广', 'WeCube Nexus集群服务', 'Plugin  instance-Docker', '4C8G+250G+1T（GFS）', 'TKO-MGMT
ST-MGMT', '10.111.2.13'),
('Wecube系统', '梁智隆
张文广', 'WECMDB', 'WECMDB-ADAPTOR', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.47'),
('中间件系统', '包健', 'ELK', 'ES', '8C32G 系统(Linux) 60G+数据3T', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.109
10.110.8.56'),
('中间件系统', '包健', 'ELK', 'kibana', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.115'),
('中间件系统', '张哲楠
梁智隆', 'UM', 'UM-LADPPROXY(UM-SERV)、UM-BATCH', '2C4G 系统(Linux) 60G+数据100G
2C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.18'),
('中间件系统', '张哲楠
梁智隆', 'UM', 'UM-SSO', '2C4G 系统(Linux) 60G+数据100G
2C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.39'),
('中间件系统', '张哲楠
梁智隆', 'UM', 'UM-Core', '2C4G 系统(Linux) 60G+数据100G
2C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.44'),
('中间件系统', '张文广
梁智隆', 'MASA', 'MASA_CONF', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.12'),
('中间件系统', '张文广
梁智隆', 'MASA', 'MASA_REGISTRY', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.14'),
('中间件系统', '张文广
梁智隆', 'MASA', 'MASA_GateKeeper', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.46'),
('中间件系统', '张文广
梁智隆', 'MASA', 'MASA_EventMesh', 'NA', 'NA', 'NA'),
('中间件系统', '张文广
梁智隆', 'WEMQ', 'WEMQ_CC', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.31
10.110.8.65'),
('中间件系统', '张文广
梁智隆', 'WEMQ', 'WEMQ_Broker(Slave)', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.34
10.110.8.86'),
('中间件系统', '张文广
梁智隆', 'WEMQ', 'WEMQ_Broker(Async )', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.28
10.110.8.94'),
('中间件系统', '张文广
梁智隆', 'WEMQ', 'WEMQ_Broker(Sync)', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.36
10.110.8.113'),
('中间件系统', '张文广
梁智隆', 'WEMQ', 'WEMQ_Admin', '2C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.40
10.110.8.126'),
('中间件系统', '张文广
梁智隆', 'WEMQ', 'WEMQ_NameSRV', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.25
10.110.8.75'),
('中间件系统', '张文广
梁智隆', 'WEMQ', 'WEMQ_ACL', '2C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.49
10.110.8.134'),
('中间件系统', '张文广
梁智隆', 'RMB', 'RMB_GSL', '2C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.41
10.110.8.123'),
('中间件系统', '张文广
梁智隆', 'RMB', 'RMB_SGS', '2C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.27
10.110.8.57'),
('中间件系统', '张文广
梁智隆', 'RMB', 'RMB_OSS', '2C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.48
10.110.8.104'),
('中间件系统', '马宏斌
但宝平', 'GNS', 'GNS_MGT', '2C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.20'),
('中间件系统', '马宏斌
但宝平', 'GNS', 'GNS_QUERY', '2C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.38'),
('中间件系统', '马宏斌
但宝平', 'GNS', 'GNS_UPDATE', '2C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.37'),
('中间件系统', '马宏斌
但宝平', 'GNS', 'GNS_RECNCLN', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.9'),
('中间件系统', '马宏斌
但宝平', 'GNS', 'GNS_VERIFY', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.4'),
('中间件系统', '马宏斌
但宝平', 'GNS', 'GNS_TOOLKIT', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.13'),
('中间件系统', '刘宇泽
张哲楠', 'WEDOH', 'WEDOH-ADMIN', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.3'),
('中间件系统', '刘宇泽
张哲楠', 'WEDOH', 'WEDOH-DNS', '2C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.35'),
('中间件系统', '刘宇泽
张哲楠', 'WEDOH', 'WEDOH-MONITOR', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.5'),
('中间件系统', '刘宇泽
张哲楠', 'WEDOH', 'WEDOH-OBSERVER', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.15'),
('中间件系统', '包健
张哲楠', 'WEAPM', 'WEAPM-ADMSERVER', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.8'),
('中间件系统', '包健
张哲楠', 'WEAPM', 'WEAPM-SERVER', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.22'),
('中间件系统', '包健
张哲楠', 'WEAPM', 'WEAPM-VIEW', '2C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.21'),
('中间件系统', '包健
刘宇泽', 'FPS', 'FPS-TP', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.26
10.110.8.42'),
('中间件系统', '包健
刘宇泽', 'FPS', 'FPS-AP', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.29'),
('中间件系统', '马宏斌
但宝平', 'UMS', 'UMS-INFOS', '2C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.32'),
('大数据系统', '但宝平
邱国银', 'BDP', 'BDP-HADOOP(计算节点)', '4台
8C64G 系统(Linux) 200G+数据2T
4台
8C64G 系统(Linux) 60G+数据500G', 'TKO-WB-DH
ST-WB-DH', '10.110.17.39
10.110.17.25
10.110.17.35
10.110.17.41

10.110.17.44
10.110.17.26
10.110.17.34
10.110.17.43'),
('大数据系统', '但宝平
邱国银', 'BDP', 'BDP-HADOOP(管理节点)', '4C32G 系统(Linux) 200G+数据2T', 'TKO-WB-DH
ST-WB-DH', '10.110.17.17
10.110.17.15'),
('大数据系统', '但宝平
邱国银', 'BDP', 'BDP-HIVE', '8C32G 系统(Linux) 200G+数据500G', 'TKO-WB-DH
ST-WB-DH', '10.110.17.30
10.110.17.16'),
('大数据系统', '但宝平
邱国银', 'BDP', 'BDP-SPARK', '4C16G 系统(Linux) 200G+数据500G', 'TKO-WB-DH
ST-WB-DH', '10.110.17.12
10.110.17.11'),
('大数据系统', '包健
蹇晓况', 'BDP', 'BDP-UJES（Linkis）', '8C64G 系统(Linux) 200G+数据500G', 'TKO-WB-DH
ST-WB-DH', '10.110.17.27
10.110.17.21
10.110.17.48
10.110.17.19
10.110.17.23'),
('大数据系统', '包健
蹇晓况', 'BDP', 'BDP-GATEWAY', '8C64G 系统(Linux) 200G+数据500G', 'TKO-WB-DH
ST-WB-DH', '10.110.17.29
10.110.17.18
10.110.17.22
10.110.17.32
10.110.17.47'),
('大数据系统', '包健
蹇晓况', 'BDP', 'DSS-IDE', '4C16G 系统(Linux) 200G+数据500G', 'TKO-WB-DH
ST-WB-DH', '10.110.17.13
10.110.17.7'),
('大数据系统', '包健
蹇晓况', 'BDP', 'BDAP-WTSS', '8C64G 系统(Linux) 200G+数据500G', 'TKO-WB-DH
ST-WB-DH', '10.110.17.33
10.110.17.36
10.110.17.20
10.110.17.45
10.110.17.37'),
('大数据系统', '包健
蹇晓况', 'BDP', 'BDP-OPS（Managis）', '4C16G 系统(Linux) 200G+数据500G', 'TKO-WB-DH
ST-WB-DH', '10.110.17.8
10.110.17.6'),
('大数据系统', '包健
蹇晓况', 'BDP', 'KMS', '2C8G 系统(Linux) 60G+数据100G', 'TKO-WB-DH
ST-WB-DH', '10.110.17.31
10.110.17.46
10.110.17.49'),
('大数据系统', '包健
蹇晓况', 'BDP', 'BDP-HBASE(管理节点)', '4C32G 系统(Linux) 200G+数据2T', 'TKO-WB-DH
ST-WB-DH', '10.110.17.5
10.110.17.10
10.110.17.4'),
('大数据系统', '包健
蹇晓况', 'BDP', 'BDP-HBASE(计算节点)', '8C32G 系统(Linux) 200G+数据2T', 'TKO-WB-DH
ST-WB-DH', '10.110.17.3
10.110.17.2
10.110.17.9
10.110.17.14'),
('零售贷款', '梁智隆
郭庆计', 'ICNCBDP', 'ICNCBDP-ACCTADSDWH', '大数据抽数子系统，部署到jobserver，无需单独申请', 'N/A', 'N/A'),
('零售贷款', '梁智隆
郭庆计', 'ICNCBDP', 'ICNCBDP-ACCTDWH', '大数据抽数子系统，部署到jobserver，无需单独申请', 'N/A', 'N/A'),
('零售贷款', '梁智隆
郭庆计', 'ICNCBDP', 'ICNCBDP-AUTHADSDWH', '大数据抽数子系统，部署到jobserver，无需单独申请', 'N/A', 'N/A'),
('零售贷款', '梁智隆
郭庆计', 'ICNCBDP', 'ICNCBDP-AUTHDWH', '大数据抽数子系统，部署到jobserver，无需单独申请', 'N/A', 'N/A'),
('零售贷款', '梁智隆
郭庆计', 'ICNCBDP', 'ICNCBDP-FONDWH', '大数据抽数子系统，部署到jobserver，无需单独申请', 'N/A', 'N/A'),
('零售贷款', '马宏斌
但宝平', 'ECIF', 'ECIF-INDEX', '2C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.30'),
('零售贷款', '马宏斌
但宝平', 'ECIF', 'ECIF-CORE', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-零售
ST-WB-零售', '10.110.6.17'),
('零售贷款', '但宝平
邱国银', 'RDPBDP', 'RDPBDP', '大数据抽数子系统，部署到jobserver，无需单独申请', 'N/A', 'N/A'),
('零售贷款', '梁智隆
郭庆计', 'INBS', 'INBS-MIDFSDCN', '2C8G 系统(Linux) 60G+数据100G', 'TKO-WB-零售
ST-WB-零售', '10.110.6.15'),
('零售贷款', '梁智隆
郭庆计', 'INBS', 'INBS-MIDBATCHDCN', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-零售
ST-WB-零售', '10.110.6.7
10.110.6.11'),
('零售贷款', '梁智隆
郭庆计', 'INBS', 'INBS-SERVICEDCN', '2C8G 系统(Linux) 60G+数据100G', 'TKO-WB-零售
ST-WB-零售', '10.110.6.13'),
('零售贷款', '梁智隆
郭庆计', 'INBS', 'INBS-BATCHDCN', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-零售
ST-WB-零售', '10.110.6.11
10.110.6.7'),
('零售贷款', '梁智隆
郭庆计', 'ICTSP', 'ICTSP-SCHEDULE', '2C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.127'),
('零售贷款', '梁智隆
郭庆计', 'ICTSP', 'ICTSP-EXECUTOR', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.50'),
('零售贷款', '梁智隆
郭庆计', 'ICTSP', 'ICTSP-WEB', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.24'),
('零售贷款', '梁智隆
郭庆计', 'ICPS', 'ICPS-PROCESS', '2C8G 系统(Linux) 60G+数据100G', 'TKO-WB-零售
ST-WB-零售', '10.110.6.3
10.110.6.9'),
('零售贷款', '梁智隆
郭庆计', 'ICPS', 'ICPS-BATCHDCN', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-零售
ST-WB-零售', '10.110.6.4
10.110.6.12'),
('零售贷款', '梁智隆
郭庆计', 'IGLP', 'IGLP-BATCHADM', '2C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.67'),
('零售贷款', '梁智隆
郭庆计', 'IMPS', 'IMPS-SERVICE', '2C8G 系统(Linux) 60G+数据100G', 'TKO-WB-零售
ST-WB-零售', '10.110.6.14'),
('零售贷款', '梁智隆
郭庆计', 'IMPS', 'IMPS-CAMSERVICE', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.97'),
('零售贷款', '梁智隆
郭庆计', 'IMPS', 'IMPS-CAMBATCH', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.76'),
('零售贷款', '梁智隆
郭庆计', 'ICOPS', 'ICOPS_ADMNGINX', '1C4G 系统(Linux) 60G+数据50G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.144'),
('零售贷款', '梁智隆
郭庆计', 'IFON', 'IFON-WEBSR', '与ICOPS_ADMNGINX混合部署，无需单独申请主机资源', 'N/A', 'N/A'),
('零售贷款', '梁智隆
郭庆计', 'IFON', 'IFON-BATCH', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.99'),
('零售贷款', '梁智隆
郭庆计', 'IFON', 'IFON-SERVICE', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.131'),
('零售存款', '吴伟
江国路', 'IDP', 'IDP-CORE(零售核心)', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-零售
ST-WB-零售', '10.110.4.12
10.110.6.5
10.110.6.6
10.110.6.10'),
('零售存款', '吴伟
江国路', 'IDP', 'IDP-CORE(对公核心)', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-企业业务
ST-WB-企业业务', '10.110.6.5
10.110.4.12
10.110.4.44'),
('零售存款', '吴伟
江国路', 'IDP', 'IDP-BATCH（零售批量）', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-零售
ST-WB-零售', '10.110.6.8
10.110.6.16'),
('零售存款', '吴伟
江国路', 'IDP', 'IDP-BATCH（对公批量）', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-企业业务
ST-WB-企业业务', '10.110.4.17
10.110.4.34'),
('零售存款', '吴伟
江国路', 'IDP', 'IDP-ADMCORE', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.141
10.110.8.72'),
('零售存款', '吴伟
江国路', 'IDP', 'IDP-ADMBATCH', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.103
10.110.8.133'),
('零售存款', '吴伟
江国路', 'IDP', 'IDP-CONSOLE', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.108
10.110.8.128'),
('零售存款', '吴伟
江国路', 'IDP', 'IDP-GL', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.58
10.110.8.111'),
('零售存款', '吴伟
江国路', 'IDP', 'IDP-FILEBATCH', '1C4G 系统(Linux) 60G+数据100G
1C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.68
10.110.8.77'),
('零售存款', '吴伟
江国路', 'IDP', 'IDP-ETL', '大数据抽数子系统，无需单独申请', 'N/A', 'N/A'),
('零售存款', '吴伟
江国路', 'IDP', 'IDP-COUNTER', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.81
10.110.8.130'),
('零售存款', '吴伟
江国路', 'IDP', 'IDP-COUNTERBATCH', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.112
10.110.8.145'),
('对公离线数仓', '韩瑞波
邱国银', 'CODS', 'CODS-CUAM', '大数据抽数子系统，无需单独申请', 'N/A', 'N/A'),
('对公存款', '周智君
张哲楠', 'MEATS', 'MEATS-CORE', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-企业业务
ST-WB-企业业务', '10.110.4.11'),
('对公贷款', '韩瑞波
刘宇泽', 'MEFS', 'MEFS-CORE', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-企业业务
ST-WB-企业业务', '10.110.4.10'),
('对公贷款', '韩瑞波
刘宇泽', 'MEFS', 'MEFS-ADM', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.114'),
('对公存款', '周智君
张哲楠', 'CFAOS', 'CFAOS-WEAC', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-企业业务
ST-WB-企业业务', '10.110.4.3'),
('对公存款', '周智君
张哲楠', 'CFAOS', 'CFAOS-WEUAC', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.88'),
('对公贷款', '韩瑞波
刘宇泽', 'CUAMS', 'CUAMS-CORE', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-企业业务
ST-WB-企业业务', '10.110.4.7'),
('对公贷款', '韩瑞波
刘宇泽', 'CUAMS', 'CUAMS-ADMBATCH', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.118'),
('对公贷款', '韩瑞波
刘宇泽', 'CUAMS', 'CUAMS-ADM', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.106'),
('对公贷款', '韩瑞波
刘宇泽', 'CUAMS', 'CUAMS-TASK', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-企业业务
ST-WB-企业业务', '10.110.4.14'),
('对公贷款', '韩瑞波
刘宇泽', 'MARM', 'MARM-CORE', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-企业业务
ST-WB-企业业务', '10.110.4.4'),
('对公贷款', '韩瑞波
刘宇泽', 'MARM', 'MARM-ADM', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.69'),
('对公贷款', '韩瑞波
刘宇泽', 'MCFCM', 'MCFCM-CORE', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-企业业务
ST-WB-企业业务', '10.110.4.8
10.110.4.16'),
('对公贷款', '韩瑞波
刘宇泽', 'MCFCM', 'MCFCM-TASK', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-企业业务
ST-WB-企业业务', '10.110.4.6
10.110.4.13'),
('对公贷款', '韩瑞波
刘宇泽', 'MCFCM', 'MCFCM-ADM', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.122
10.110.8.105'),
('对公贷款', '韩瑞波
刘宇泽', 'MCFCM', 'MCFCM-ADMBATCH', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.120
10.110.8.136'),
('对公贷款', '韩瑞波
刘宇泽', 'MCFCM', 'MCFCM-CFSCORE', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-企业业务
ST-WB-企业业务', '10.110.4.5
10.110.4.9'),
('对公存款', '周智君
张哲楠', 'CCIF', 'CCIF-ADM', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.85'),
('对公存款', '周智君
张哲楠', 'CCIF', 'CCIF-CORE', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-企业业务
ST-WB-企业业务', '10.110.4.15'),
('对公存款', '周智君
张哲楠', 'CCIF', 'CCIF-INDEX', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.70'),
('对公贷款', '韩瑞波
刘宇泽', 'NRRS', 'NRRS-ADM', '4C16G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.55'),
('外汇系统', '马宏斌
邱国银', 'FXTS', 'FXTS-ADM', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.89'),
('外汇系统', '马宏斌
邱国银', 'FXTS', 'FXTS-PT', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-企业法人业务
ST-WB-企业法人业务', '10.110.2.9'),
('外汇系统', '马宏斌
邱国银', 'FXTS', 'FXTS-CORE', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-企业法人业务
ST-WB-企业法人业务', '10.110.2.15'),
('外汇系统', '马宏斌
邱国银', 'FXTS', 'FXTS-ACCT', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-企业法人业务
ST-WB-企业法人业务', '10.110.2.4'),
('对公存款', '周智君
张哲楠', 'CBM', 'CBM-CORE', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.93'),
('对公存款', '周智君
张哲楠', 'CBM', 'CBM-ADM', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.102'),
('财务系统', '马宏斌
江国路', 'ECL', 'ECL-CORE', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.66'),
('财务系统', '马宏斌
江国路', 'ECL', 'ECL-WEB', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.110'),
('财务系统', '马宏斌
江国路', 'ECL', 'ECL-BDP', '大数据抽数子系统，无需单独申请', 'N/A', 'N/A'),
('财务系统', '马宏斌
江国路', 'DAS', 'DAS-BATCHADPT', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.66'),
('财务系统', '马宏斌
江国路', 'DAS', 'DAS-ALMBDP', '大数据抽数子系统，无需单独申请', 'N/A', 'N/A'),
('财务系统', '马宏斌
江国路', 'DAS', 'DAS-GLMBDP', '大数据抽数子系统，无需单独申请', 'N/A', 'N/A'),
('外汇系统', '马宏斌
邱国银', 'FIPQES', 'FIPQES-FXFRONT', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-专线接入
ST-WB-专线接入', '10.114.1.78'),
('外汇系统', '马宏斌
邱国银', 'FIPQES', 'FIPQES-FXADM', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-管理与PaaS', '10.110.8.121'),
('外汇系统', '马宏斌
邱国银', 'FIPQES', 'FIPQES-FXCORE', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-企业法人业务
ST-WB-企业法人业务', '10.110.2.12
10.110.4.2'),
('TiDB', '包健
刘宇泽', 'TiDB', 'DB-TIKVSERVER', '4C16G 系统(Linux) 60G+数据500G', 'TKO-DB数据库
ST-DB数据库', '10.110.32.43
10.110.32.54
10.110.32.39'),
('TiDB', '包健
刘宇泽', 'TiDB', 'DB-TIDBSERVER', '2C8G 系统(Linux) 60G+数据250G', 'TKO-DB数据库
ST-DB数据库', '10.110.32.59
10.110.32.50'),
('TiDB', '包健
刘宇泽', 'TiDB', 'DB-TIDBPD', '2C8G 系统(Linux) 60G+数据250G', 'TKO-DB数据库
ST-DB数据库', '10.110.32.10
10.110.32.70'),
('TiDB', '包健
刘宇泽', 'TiDB', 'DB-TIDBMONITOR', '2C8G 系统(Linux) 60G+数据250G', 'TKO-DB数据库
ST-DB数据库', '10.110.32.67'),
('TiDB', '包健
刘宇泽', 'TiDB', 'DM-MASTER', '2C8G 系统(Linux) 60G+数据250G', 'TKO-DB数据库
ST-DB数据库', '10.110.32.26
10.110.32.48'),
('TiDB', '包健
刘宇泽', 'TiDB', 'DM-WORKER', '2C8G 系统(Linux) 60G+数据250G', 'TKO-DB数据库
ST-DB数据库', '10.110.32.49
10.110.32.30
10.110.32.51
10.110.32.52'),
('TiDB', '包健
刘宇泽', 'TiDB', 'DM-MONITOR', '2C8G 系统(Linux) 60G+数据200G', 'TKO-DB数据库
ST-DB数据库', '10.110.32.35'),
('WEREDIS', '吴伟
但宝平', 'WEREDIS', 'WEREDIS_OBSERVER', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.2'),
('WEREDIS', '吴伟
但宝平', 'WEREDIS', 'WEREDIS_PROXY', '4C8G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.23'),
('WEREDIS', '吴伟
但宝平', 'WEREDIS', 'WEREDIS_CLUSTER（业务集群）', '8C16G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.83
10.110.8.138
10.110.8.45
10.110.8.43'),
('WEREDIS', '吴伟
但宝平', 'WEREDIS', 'WEREDIS_CLUSTER（心跳集群）', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.16
10.110.8.73'),
('WEREDIS', '吴伟
但宝平', 'WEREDIS', 'WEREDIS_ADMIN', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.7'),
('WEREDIS', '吴伟
但宝平', 'WEREDIS', 'WEREDIS_EXPORTER', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.17'),
('WEREDIS', '吴伟
但宝平', 'WEREDIS', 'WEREDIS_DASHBOARD', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.6'),
('WEREDIS', '吴伟
但宝平', 'WEREDIS', 'WEREDIS_STAT', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.10'),
('WEREDIS', '吴伟
但宝平', 'WEREDIS', 'WEREDIS-MIGRATION', '1C4G 系统(Linux) 60G+数据100G', 'TKO-WB-管理与PaaS
ST-WB-PaaS', '10.110.8.11');
