alter table routers add column component varchar(255) not null default '' comment '自定义页面' after sort;

insert into routers(name,title,path,is_menu,pid,component) values('sys_user_manage', '账号管理','sys_user_manage',1,0,'');
insert into routers(name,title,path,is_menu,pid,component) values('sys_user_list', '账号列表','list',1,4,'sys/user/list');
insert into routers(name,title,path,is_menu,pid,component) values('sys_user_info', '账号详情','info',0,4,'sys/user/info');

insert into routers(name,title,path,is_menu,pid,component) values('sys_router_manage', '路由管理','sys_router_manage',1,0,'');
insert into routers(name,title,path,is_menu,pid,component) values('sys_router_list', '路由列表','list',1,7,'sys/router/list');
insert into routers(name,title,path,is_menu,pid,component) values('sys_router_info', '路由详情','info',0,7,'sys/router/info');
insert into routers(name,title,path,is_menu,pid,component) values('sys_router_add', '添加路由','add',0,7,'sys/router/add');
