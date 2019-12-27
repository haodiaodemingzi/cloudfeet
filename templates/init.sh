#/bin/bash
token=''
get_json_value()
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

wget --no-check-certificate --post-data='{"username":"james","password":"1234"}'  http://localhost:8082/api/v1/auth/token -O /tmp/xxxtoken
token_json=`cat /tmp/xxxtoken`
get_json_value $token_json token

echo "token====${token}"
rm -f /tmp/xxxtoken

echo "wget --no-check-certificate --header="Token: ${token}" -O /tmp/gfwlist.confg http://localhost:8082/api/v1/pac/script"
wget --no-check-certificate --header="Token: ${token}" -O /tmp/start.sh http://localhost:8082/api/v1/pac/script
chmod +x /tmp/start.sh
sed -i 's/\&lt;/</g' /tmp/start.sh
dos2unix /tmp/start.sh
/bin/bash /tmp/start.sh