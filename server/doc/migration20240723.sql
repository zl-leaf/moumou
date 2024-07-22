alter table routers add column is_system tinyint(1) not null default 0 comment '是否系统菜单' after is_menu;

update routers set is_system=1 where name like '%router%';
update routers set is_system=1 where name like '%user%';
update routers set is_system=1 where name like '%role%';