CREATE TABLE `users` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
    `password` varchar(200) NOT NULL DEFAULT '' COMMENT '密码',
    `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
    `deleted_at` int NOT NULL DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户';


CREATE TABLE `routers` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(50) NOT NULL COMMENT '标识名称',
    `title` VARCHAR(10) NOT NULL COMMENT '菜单名',
    `path` VARCHAR(50) NOT NULL COMMENT '路径',
    `is_menu` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否菜单',
    `pid` bigint NOT NULL DEFAULT 0 COMMENT '父节点',
    `sort` int NOT NULL DEFAULT 0 COMMENT '排序',
    `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
    `deleted_at` int NOT NULL DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '路由';