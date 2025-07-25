CREATE TABLE `user` (
                        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '用户唯一标识',
                        `username` VARCHAR(255) NOT NULL UNIQUE COMMENT '用户名',
                        `email` VARCHAR(255) NOT NULL COMMENT '邮箱地址',
                        `gender` VARCHAR(16) DEFAULT NULL COMMENT '性别（如 male/female/other）',
                        `age` INT DEFAULT NULL COMMENT '年龄',
                        `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';