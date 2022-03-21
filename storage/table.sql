create table `blog_tag`(
    `id` int(10) unsigned not null auto_increment,
    `name` varchar(100) default "" comment '标签名称',
    `state` tinyint(3) unsigned default '1' comment '状态 0 为禁用，1为启用',
    `created_on` int(10) unsigned default 0 comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default 0 comment '修改时间',
    `modified_by` varchar(100) default '' comment '修改人',
    `deleted_on` int(10) unsigned default 0 comment '删除时间',
    `is_del` tinyint(3) unsigned default 0 comment '是否删除 0 为未删除，1为已删除',
    Primary key (`id`)
) ENGINE=InnoDB default charset=utf8mb4 comment="标签管理";


create table `blog_article`(
    `id` int(10) unsigned not null auto_increment,
    `title` varchar(100) default "" comment '文章标题',
    `desc` varchar(255) default "" comment '文章简述',
    `cover_image_url` varchar(255) default '' comment '封面图片地址',
    `content` longtext comment '文章内容',
    `created_on` int(10) unsigned default 0 comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default 0 comment '修改时间',
    `modified_by` varchar(100) default '' comment '修改人',
    `deleted_on` int(10) unsigned default 0 comment '删除时间',
    `is_del` tinyint(3) unsigned default 0 comment '是否删除 0 为未删除，1为已删除',
    Primary key (`id`)
) ENGINE=InnoDB default charset=utf8mb4 comment="文章管理";

create table `blog_article_tag` (
    `id` int(10) unsigned not null auto_increment,
    `article_id` int(11) not null default 0 comment '文章ID',
    `tag_id` int(10) unsigned not null default 0 comment '标签ID',
    `created_on` int(10) unsigned default 0 comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default 0 comment '修改时间',
    `modified_by` varchar(100) default '' comment '修改人',
    `deleted_on` int(10) unsigned default 0 comment '删除时间',
    `is_del` tinyint(3) unsigned default 0 comment '是否删除 0 为未删除，1为已删除',
    Primary key(id)
) ENGINE=InnoDB default charset=utf8mb4 comment '标签文章关联表';