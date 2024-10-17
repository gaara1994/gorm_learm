

# 1.环境安装

## 1.postgres

下载镜像

```shell
docker pull postgres:17
```

安装数据库

```shell
docker run --name postgres-17 \
  -e POSTGRES_PASSWORD=MyNewPass4! \
  -d -p 5432:5432 \
  -v /var/lib/postgresql/data:/var/lib/postgresql/data \
  postgres:17
```

创建数据库

```sql
CREATE DATABASE "gorm_learn"
WITH
  TEMPLATE = "template1"
;

COMMENT ON DATABASE "gorm_learn" IS 'gorm_learn学习使用';
```



