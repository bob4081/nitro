package cmd

const CloudConfig = `#cloud-config
packages:
  - redis
  - jq
  - apt-transport-https
  - ca-certificates
  - curl
  - gnupg-agent
  - software-properties-common
  - sshfs
  - pv
  - httpie
  - unzip
  - mysql-client
  - postgresql-client
write_files:
  - path: /home/ubuntu/.nitro/databases/mysql/conf.d/5/mysql.cnf
    content: |
      [mysqld]
      max_allowed_packet=256M
      wait_timeout=86400
      default-authentication-plugin=mysql_native_password
  - path: /home/ubuntu/.nitro/databases/mysql/conf.d/8/mysql.cnf
    content: |
      [mysqld]
      max_allowed_packet=256M
      wait_timeout=86400
      default-authentication-plugin=mysql_native_password
      [mysqldump]
      column-statistics=0
  - path: /home/ubuntu/.nitro/databases/mysql/setup.sql
    content: |
      CREATE USER IF NOT EXISTS 'nitro'@'localhost' IDENTIFIED BY 'nitro';
      CREATE USER IF NOT EXISTS 'nitro'@'%' IDENTIFIED BY 'nitro';
      GRANT ALL PRIVILEGES ON *.* TO 'nitro'@'localhost' WITH GRANT OPTION;
      GRANT ALL PRIVILEGES ON *.* TO 'nitro'@'%' WITH GRANT OPTION;
      FLUSH PRIVILEGES;
  - path: /home/ubuntu/.nitro/databases/postgres/setup.sql
    content: |
      ALTER USER nitro WITH SUPERUSER;
  - path: /opt/nitro/nginx/template.conf
    content: |
      server {
          listen 80;
          listen [::]:80;

          root CHANGEWEBROOTDIR;

          index index.php;
          gzip_static  on;
          error_page 404 /index.php?$query_string;
          ssi on;
          server_name CHANGESERVERNAME;
          client_max_body_size 100M;

          location / {
              try_files $uri $uri/ /index.php$is_args$args;
          }

          location ~ \.php$ {
             include snippets/fastcgi-php.conf;
             fastcgi_pass unix:/var/run/php/phpCHANGEPHPVERSION-fpm.sock;
             fastcgi_read_timeout 240;
             fastcgi_param CRAFT_NITRO 1;
             fastcgi_param DB_USER nitro;
             fastcgi_param DB_PASSWORD nitro;
          }
      }
  - path: /opt/nitro/php-xdebug.ini
    content: |
      zend_extension=xdebug.so
      xdebug.remote_enable=1
      xdebug.remote_connect_back=0
      xdebug.remote_host=192.168.64.1
      xdebug.remote_port=9000
      xdebug.remote_autostart=1
      xdebug.idekey=PHPSTORM
runcmd:
  - sed -i 's|127.0.0.53|1.1.1.1|g' /etc/resolv.conf
  - add-apt-repository --no-update -y ppa:ondrej/php
  - echo "CRAFT_NITRO=1" >> /etc/environment
  - echo "DB_USER=nitro" >> /etc/environment
  - echo "DB_PASSWORD=nitro" >> /etc/environment
  - curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
  - sudo add-apt-repository --no-update -y "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
  - wget -q -O - https://packages.blackfire.io/gpg.key | sudo apt-key add -
  - echo "deb http://packages.blackfire.io/debian any main" | sudo tee /etc/apt/sources.list.d/blackfire.list
  - apt-get update -y
  - apt-get install -y nginx docker-ce docker-ce-cli containerd.io
  - usermod -aG docker ubuntu
  - mkdir -p /home/ubuntu/sites
  - mkdir -p /home/ubuntu/.nitro/databases/imports
  - mkdir -p /home/ubuntu/.nitro/databases/mysql/conf.d
  - mkdir -p /home/ubuntu/.nitro/databases/mysql/backups
  - mkdir -p /home/ubuntu/.nitro/databases/postgres/conf.d
  - mkdir -p /home/ubuntu/.nitro/databases/mysql/conf.d
  - mkdir -p /home/ubuntu/.nitro/databases/postgres/backups
  - cp /etc/skel/.bashrc /home/ubuntu/.bashrc
  - cp /etc/skel/.profile /home/ubuntu/.profile
  - cp /etc/skel/.bash_logout /home/ubuntu/.bash_logout
  - chown -R ubuntu:ubuntu /home/ubuntu/
`
