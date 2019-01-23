# php环境搭建
之前用node做了一个接口转发的项目，了解了中间件，api网关，集群，nginx配置等服务端相关的知识。

现在想再深入的了解，公司的服务端用的是yii2框架，就先研究一下。

搭建php开发环境需要准备3个东西，一个是`web服务器`，一个是`php-fpm`，还有一个是`数据库`;

- web服务器是什么
> web服务器一般都是指静态服务器，启动了服务器就能在浏览器访问静态页面，比如nginx。
- php-fpm
> 只有web服务器还不够，有时页面中需要动态的处理数据，还需要访问数据库的能力，这时就要选择一种后端语言，这里选择php，这时就需要php-fpm了。

> 要解释php-fpm这个东西，还需要和php-cgi和fastcgi一起说。首先，如果一个请求进来，nginx会把请求转发给php-fpm，但是转发的请求是需要按照fastcgi规定的格式组装和传输的，而php-fpm实现了fastcgi协议的程序，能解析fastcgi传过来的数据。

> php是一门解释性语言，php-cgi就是php的解释器，可以一行行的去执行php脚本，但是php-cgi没有进程管理的能力，比如php-fpm会预先启动一个master和多个worker进程等待请求，php-fpm在收到fastcgi传来的请求后，会调度一个worker进程去处理请求，请求完成后返回给fastcgi，再由nginx返回给客户端。

# nginx

web服务器就选用nginx，下面是`docker-compose`的内容:
```
nginx:
  image: nginx:latest
  volumes:
    - /path/to/nginx/conf.d:/etc/nginx/conf.d
    - /path/to/yii/project/web:/var/www/html
  ports:
    - "8080:80"
  links:
    - php
```

在nginx的conf.d里面：
```
server {
    charset utf-8;
    client_max_body_size 128M;

    listen 80;

    server_name localhost;
    root        /var/www/html;
    index       index.php;
    
    location / {
        try_files $uri $uri/ /index.php$is_args$args;
    }

    location ~ ^/assets/.*\.php$ {
        deny all;
    }
    
    location ~ \.php$ {
        fastcgi_pass php:9000;
        fastcgi_index  index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        include fastcgi_params;
    }

    location ~* /\. {
        deny all;
    }
}
```

我们把默认的静态文件路径从`/etc/share/nginx/html`改为`/var/www/html`，目的是保持和下面的`php-fpm`路径一致，不然就要写多个root，`$document_root`的值就是`root`的值。

这样，web服务器就搭建好了，就可以访问静态页面了。

# 搭建php-fpm

`docker-compose`的配置：

```
php:
  build: ./config/php
  volumes:
    - /path/to/php/project:/var/www
    - /path/to/php/project/web:/var/www/html
  ports:
    - "9000:9000"
  links:
    - mysql
```

Dockerfile的配置

```
FROM php:7.2.8-fpm
RUN docker-php-ext-install pdo_mysql
```

因为数据库用到了mysql，需要安装一个`pdo_mysql`插件，所以单独抽取出来用Dockerfile。

# mysql


```
mysql:
  image: mysql:8
  ports:
    - "3306"
  expose:
    - "3306"
  environment:
    MYSQL_ROOT_PASSWORD: root
    MYSQL_DATABASE: web
    MYSQL_USER: user
    MYSQL_PASSWORD: pwd
  command: ['--default-authentication-plugin=mysql_native_password']
```

开发环境就搭建好了。