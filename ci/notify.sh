#!/usr/bin/env bash

SLACK_USERNAME="cwc"
SLACK_EMOJI=":cwc:"
SLACK_CHANNEL="#cloud"


slack_notif() {
    token="${1}"
    if [[ $token ]]; then
        message="Version ${CI_COMMIT_TAG} of ${SLACK_USERNAME} has been released! ${VERSION}"
        endpoint="https://hooks.slack.com/services/${token}"
        payload="{\"text\": \"${message}\", \"username\": \"${SLACK_USERNAME}\", \"channel\": \"${SLACK_CHANNEL}\", \"icon_emoji\": \"${SLACK_EMOJI}\"}"
        curl -X POST "${endpoint}" -H "Accept: application/json" -d "${payload}"
    fi
}

slack_notif "${SLACK_TOKEN}"
slack_notif "${SLACK_TOKEN_PUBLIC}"
