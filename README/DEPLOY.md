# Robot Manage 部署文档（非Docker方式）

## 一、本地构建

### 1. 执行构建脚本
```bash
chmod +x build.sh
./build.sh
```

构建完成后会生成 `deploy_YYYYMMDD_HHMMSS.tar.gz` 部署包。

## 二、服务器部署

### 1. 上传部署包
```bash
scp deploy_*.tar.gz root@9.135.10.13:/data/robot/robot-manage/
```

### 2. 解压部署包
```bash
tar -xzf deploy_*.tar.gz
cd deploy_*
```

### 3. 配置数据库
编辑配置文件：
```bash
vim configs/config.yaml
```

修改数据库连接信息：
```yaml
database:
  host: localhost
  port: 3306
  user: root
  password: your_password
  dbname: robot_manage
```

### 4. 初始化数据库
```bash
# 登录MySQL
#mysql -u root -p

# 创建数据库
#CREATE DATABASE IF NOT EXISTS robot_manage DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# 导入SQL文件
#USE robot_manage;
#source sql/init.sql;
```

### 5. 启动服务
```bash
./start.sh
```

### 6. 查看日志
```bash
tail -f robot-manage.log
```

### 7. 停止服务
```bash
./stop.sh
```

## 三、Nginx配置（可选）

如果需要使用域名访问，配置Nginx反向代理：

```nginx
server {
    listen 80;
    server_name your-domain.com;

    # 前端静态文件
    location / {
        root /opt/deploy_*/dist;
        try_files $uri $uri/ /index.html;
    }

    # 后端API代理
    location /api/ {
        proxy_pass http://localhost:8088;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

重启Nginx：
```bash
sudo nginx -t
sudo systemctl restart nginx
```

## 四、systemd服务配置（推荐）

创建systemd服务文件：
```bash
sudo vim /etc/systemd/system/robot-manage.service
```

内容：
```ini
[Unit]
Description=Robot Manage Service
After=network.target mysql.service

[Service]
Type=simple
User=root
WorkingDirectory=/opt/deploy_YYYYMMDD_HHMMSS
ExecStart=/opt/deploy_YYYYMMDD_HHMMSS/robot-manage
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

启用并启动服务：
```bash
sudo systemctl daemon-reload
sudo systemctl enable robot-manage
sudo systemctl start robot-manage
sudo systemctl status robot-manage
```

查看日志：
```bash
sudo journalctl -u robot-manage -f
```

## 五、更新部署

### 1. 停止旧服务
```bash
cd /opt/deploy_old
./stop.sh
# 或
sudo systemctl stop robot-manage
```

### 2. 备份数据库
```bash
mysqldump -u root -p robot_manage > backup_$(date +%Y%m%d).sql
```

### 3. 部署新版本
```bash
cd /opt
tar -xzf deploy_new.tar.gz
cd deploy_new
# 复制旧配置
cp ../deploy_old/configs/config.yaml configs/
```

### 4. 更新systemd服务路径
```bash
sudo vim /etc/systemd/system/robot-manage.service
# 修改 WorkingDirectory 和 ExecStart 路径
sudo systemctl daemon-reload
```

### 5. 启动新服务
```bash
sudo systemctl start robot-manage
```

## 六、常见问题

### 1. 端口被占用
```bash
# 查找占用端口的进程
lsof -ti:8088
# 停止进程
kill -9 <PID>
```

### 2. 数据库连接失败
- 检查MySQL是否启动
- 检查配置文件中的数据库信息
- 检查防火墙设置

### 3. 前端访问404
- 检查Nginx配置
- 确认dist目录存在
- 检查文件权限

### 4. 查看服务状态
```bash
# systemd方式
sudo systemctl status robot-manage

# 手动启动方式
ps aux | grep robot-manage
```

## 七、监控和维护

### 1. 日志轮转
创建 `/etc/logrotate.d/robot-manage`：
```
/opt/deploy_*/robot-manage.log {
    daily
    rotate 7
    compress
    delaycompress
    missingok
    notifempty
}
```

### 2. 定期备份数据库
添加crontab任务：
```bash
crontab -e
```

添加：
```
0 2 * * * mysqldump -u root -p'password' robot_manage > /backup/robot_manage_$(date +\%Y\%m\%d).sql
```

### 3. 磁盘空间监控
```bash
df -h
du -sh /opt/deploy_*
```
