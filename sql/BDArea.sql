CREATE TABLE `BDArea` (
`AreaId` int(32) unsigned NOT NULL AUTO_INCREMENT COMMENT '地区Id',
`CountryId` int(32)   COMMENT '国家Id',
`CorpId` int(32)   COMMENT '公司Id',
`Code` varchar(100) NOT NULL  COMMENT '地区编码',
`Name` varchar(100) NOT NULL  COMMENT '地区名称',
`ParentId` int(32)   COMMENT '上级Id',
`Level` tinyint NOT NULL  COMMENT '层级',
`IsSys` tinyint  DEFAULT '1' COMMENT '是否系统内置',
`OriginId` int(32)   COMMENT '来源地区Id',
`SealFlag` tinyint  DEFAULT '0' COMMENT '停用',
`Creator` int(32)   COMMENT '创建人',
`CreatTime` timeStamp   COMMENT '创建时间',
`Modifier` int(32)   COMMENT '修改人',
`Ts` TimeStamp   COMMENT '最后修改时间',
`DR` tinyint  DEFAULT '0' COMMENT '是否删除',
PRIMARY KEY (`AreaId`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='地区表';
