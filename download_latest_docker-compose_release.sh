#!/bin/bash
set -x
url="https://github.com/docker/compose/releases/latest"
latest_version=$(curl $url -s -L -I -o /dev/null -w '%{url_effective}' | { read redirected_url; echo "${redirected_url##*/}"; })
curl -L https://github.com/docker/compose/releases/download/$latest_version/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
