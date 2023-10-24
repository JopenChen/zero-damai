CREATE TABLE `performance_session` (
 `id` INT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT 'ID',
 `performance_id` INT UNSIGNED NOT NULL COMMENT '关联的演出活动ID',
 `tickets` JSON COMMENT '规格票数，格式为：[{"seat": "vip","price": 2023,"quantity": 2000}]。seat: 席位、price: 票价、quantity: 数量。',
 `show_at` INT NOT NULL DEFAULT 0 COMMENT '演出时间',
 `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
 `is_del` TINYINT NOT NULL DEFAULT 2 COMMENT '是否删除: 0=未指定、1=是、2=否: 默认=2',
 PRIMARY KEY (`id`) USING BTREE
) AUTO_INCREMENT=1 ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='演出活动场次表';