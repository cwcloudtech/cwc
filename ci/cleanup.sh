i=0
for project in $(curl -X GET --header "PRIVATE-TOKEN:${GITLAB_TOKEN}"  ${GITLAB_HOST}/api/v4/projects/${GITLAB_PROJECT}/releases | jq -r '.[] | @base64'); do
    i=$(( i + 1 ))
    _jq() {
     echo ${project} | base64 --decode | jq -r ${1}
    }
    if [ $i -gt 5 ]
     then
      curl -X DELETE --header "PRIVATE-TOKEN:${GITLAB_TOKEN}" ${GITLAB_HOST}/api/v4/projects/${GITLAB_PROJECT}/releases/$(_jq '.tag_name') >/dev/null
    fi
done
