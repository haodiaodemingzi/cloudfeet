#!/bin/sh

export PATH='/opt/sbin:/opt/bin:/usr/sbin:/usr/bin:/sbin:/bin'
### Custom user script
### Called on Internet status changed
### $1 - Internet status (0/1)
### $2 - elapsed time (s) from previous state

logger -t "di" "Internet state: $1, elapsed time: $2s."

TOKEN_URL="https://cloudfeet.tk/api/v1/auth/token"
SCRIPT_URL="https://cloudfeet.tk/api/v1/pac/script"

parse_json_token()
{
  local json=$1
  local key=$2
  if [[ -z "$3" ]]; then
    local num=1
  else
    local num=$3
  fi
  echo "${json}" | awk -F"[,:}]" '{for(i=1;i<=NF;i++){if($i~/'${key}'\042/){print $(i+1)}}}' | tr -d '"' | sed -n ${num}p
}

get_cf_token()
{
  local AUTH_JSON=""
  if [ -z "$1" ]; then
    echo "no auth data. exit!" && exit 1;
  else
    AUTH_JSON=$1
  fi

  wget --no-check-certificate --post-data="$AUTH_JSON" "$TOKEN_URL" -O /tmp/cf_token
  token_json=`cat /tmp/cf_token` && rm /tmp/cf_token
  nvram set CLOUDFEET_TOKEN=`parse_json_token $token_json token`
}

########################################################################################

if [ -f "/bin/scutclient.sh" ]; then
	scutclient.sh restart
fi

if [ "x$1" == "x1" ]; then
	dev_imei=`cat /proc/dev_uid`
	signed_pwd=`echo -n "$dev_imei|jhp1!T!C!#okkzBoTK8!" | md5sum | awk '{print $1}'`
	post_data="{\"username\":\"$dev_imei\",\"password\":\"$signed_pwd\"}"
	#echo $post_data

	get_cf_token $post_data
	CLOUDFEET_TOKEN=`nvram get CLOUDFEET_TOKEN`

	#update ss config
	wget --no-check-certificate --header="Token: ${CLOUDFEET_TOKEN}" \
					-O - ${SCRIPT_URL} | bash
	nvram commit
	mtd_storage.sh save
fi

