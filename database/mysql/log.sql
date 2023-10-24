CREATE TABLE `performance` (
 `id` INT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '用户ID',
 `name` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '姓名',
 `nickname` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '昵称',
 `avatar` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '头像',
 `background` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '个人中心背景图',
 `mobile` CHAR(11) NOT NULL DEFAULT '' COMMENT '手机号',
 `password` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '密码',
 `mail` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '邮箱地址',
 `identity` VARCHAR(18) NOT NULL DEFAULT '' COMMENT '身份证号',
 `gender` TINYINT UNSIGNED NOT NULL DEFAULT 4 COMMENT '性别: 0=未指定、1=男、2=女、3=第三性别、4=保密: 默认=4',
 `nation` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '民族',
 `birthday` INT NOT NULL DEFAULT 0 COMMENT '出生日期',
 `address` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '收货地址',
 `audience` JSON NOT NULL COMMENT '观影人信息,格式：[{"name":"xxx","identity_type": 1,"identity_number"}],identity 可选值：1=身份证、2=港澳台居民居住证、3=港澳居民来往内地通行证、4=台湾居民来往大陆通行证、5=护照、6=外国人永久居留身份证',
 `status` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否禁用: 0=未指定、1=开启、2=禁用: 默认=1',
 `login_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登录时间',
 `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
 `is_del` TINYINT NOT NULL DEFAULT 2 COMMENT '是否删除: 0=未指定、1=是、2=否: 默认=2',
 PRIMARY KEY (`id`) USING BTREE,
 INDEX idx_mobile(mobile)
) AUTO_INCREMENT=1 ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='演出活动表';