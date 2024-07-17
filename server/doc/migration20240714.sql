alter table users change column created_at `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间';
alter table users change column updated_at `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间';
alter table users change column deleted_at `deleted_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间';

alter table routers change column created_at `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间';
alter table routers change column updated_at `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '更新时间';
alter table routers change column deleted_at `deleted_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间';
