#!bin/sh

# 必要なものをDLする
sudo yum -y install squid
sudo yum -y install httpd
sudo yum -y install expect
sudo yum -y update openssl
sudo yum -y install firewalld

PW="test"

expect -c ""
set timeout 20

# Userの設定
# sudo htpasswd -c /etc/squid/.htpasswd sho
sudo spawn htpasswd -c /etc/squid/.htpasswd userid

# Passwordの設定
expect \"password:\"
send \"${PW}\n\"
expect \"password:\"
send \"${PW}\n\"

# squidの詳細設定
sudo sed -i -e "61,73d" /etc/squid/squid.conf
sudo sed -i -e '27i auth_param basic program /usr/lib64/squid/basic_ncsa_auth /etc/squid/.htpasswd\nauth_param basic children 5\nauth_param basic realm Squnewid Basic Authentication\nauth_param basic credentialsttl 24 hours\nacl password proxy_auth REQUIRED\nhttp_access allow password\n\nforwarded_for off\nrequest_header_access Referer deny all\nrequest_header_access X-Forwarded-For deny all\nrequest_header_access Via deny all\nrequest_header_access Cache-Control deny all\nvisible_hostname unknown\nno_cache deny all\n' /etc/squid/squid.conf
sudo sed -i 's/3128/userport/g' /etc/squid/squid.conf

# firewall関連の設定
sudo firewall-cmd --zone=public --add-port=3128/tcp --permanent
sudo systemctl start firewalld.service
sudo firewall-cmd --reload
#? sudo /etc/rc.d/init.d/iptables stop
sudo systemctl start squid
# sudo rm -f script.sh

# squidが立ち上がっているかのコマンド
# sudo systemctl status squid.service

# firewallが立ち上がっているかのコマンド
# sudo systemctl status firewalld -l