#!/bin/bash


nvram set ss_enable=1
nvram set ss_mtu=1492
nvram set ss_multiport=22,80,443,1935
nvram set ss_lower_port_only=2
#nvram set ss_usage=" -O auth_sha1_v4 -o http_simple"
nvram set ss_server={{.Server}}
nvram set ss_server_port={{.Port}}
nvram set ss_key={{.Password}}
nvram set ss_method={{.Method}}
nvram set ss_obfs=plain
nvram set ss_obfs_param=""
nvram set ss_protocol=origin
nvram set ss_proto_param=""



logger "更新出国服务器配置..更新出国服务器配置..更新出国服务器配置.."

token=`nvram get CLOUDFEET_TOKEN`
wget --no-check-certificate --header="Token: ${token}" -O /tmp/gfwlist.conf {{.GfwListURL}}

if [ -f /tmp/gfwlist.conf ];then
	mv /tmp/gfwlist.conf /etc/storage/dnsmasq/dnsmasq.d/gfwlist.dnsmasq.conf
	echo "" > /etc/storage/dnsmasq/gfwhosts
fi

report_script="/tmp/reportdns.sh"
script_data='#!/bin/bash
logger "域名缓存分析.."
kill -USR1 `cat /var/run/dnsmasq.pid`
sleep 5
CLOUDFEET_TOKEN=`nvram get CLOUDFEET_TOKEN`
wget --no-check-certificate --header="Token: ${CLOUDFEET_TOKEN}" \
         --post-data="`cat /tmp/dns.cache|uniq`" {{.DomainsUploadURL}}
'
printf  "$script_data" > $report_script

chmod 755 ${report_script}

ipset_gfw=ss_spec_dst_fw

restart_dnsproxy
sleep 1
restart_dns
sleep 1

/usr/bin/shadowsocks.sh restart
