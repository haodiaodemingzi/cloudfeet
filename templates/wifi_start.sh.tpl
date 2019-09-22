#/bin/bash
token=''
function get_json_value()
{
  local json=$1
  local key=$2
  if [[ -z "$3" ]]; then
    local num=1
  else
    local num=$3
  fi

  token=$(echo "${json}" | awk -F"[,:}]" '{for(i=1;i<=NF;i++){if($i~/'${key}'\042/){print $(i+1)}}}' | tr -d '"' | sed -n ${num}p)
}

nvram set ss_enable=1
nvram set ss_mtu=1492
nvram set ss_multiport=22,80,443,1935
nvram set ss_lower_port_only=2
#nvram set ss_usage=" -O auth_sha1_v4 -o http_simple"
nvram set ss_server={{.Server}}
nvram set ss_server_port={{.Port}}
nvram set ss_key={{.Password}}
nvram set ss_method={{.Method}}
nvram set ss_obfs=http_simple
nvram set ss_obfs_param=""
nvram set ss_protocol=auth_sha1_v4
nvram set ss_proto_param=""



logger "更新出国服务器配置..更新出国服务器配置..更新出国服务器配置.."

wget --no-check-certificate --post-data='{"username":"james","password":"1234"}'  {{.AuthURL}} -O /tmp/xxxtoken
token_json=`cat /tmp/xxxtoken`
get_json_value $token_json token

echo "token====${token}"

rm -f /tmp/xxxtoken
wget --no-check-certificate --header="Token: ${token}" -O /tmp/gfwlist.confg {{.GfwListURL}}

if [ -f /tmp/gfwlist.conf ];then
	mv /tmp/gfwlist.conf /etc/storage/dnsmasq/dnsmasq.d/gfwlist.dnsmasq.conf
fi
ipset_gfw=ss_spec_dst_fw

restart_dnsproxy
sleep 1
restart_dns
sleep 1

/usr/bin/shadowsocks.sh restart