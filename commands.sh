# get container details
curl --unix-socket /var/run/docker.sock http://docker/containers/ba4de042fd9a/json| jq .

curl --unix-socket /var/run/docker.sock http://docker/containers/ba4de042fd9a/json | jq -r '.Name'

curl --unix-socket /var/run/docker.sock http://docker/containers/494cbfa9cbcd/json | jq -r '.State'
