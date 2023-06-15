#!/usr/bin/env bash

curl -I -w '%{http_code}' "https://csbackend.fly.dev/"

fly proxy 5433 -a csbackend-ts