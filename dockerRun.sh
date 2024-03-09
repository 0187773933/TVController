#!/bin/bash
APP_NAME="public-tv-controller"
sudo docker rm -f $APP_NAME || echo ""
id=$(sudo docker run -dit \
--name $APP_NAME \
--restart="always" \
$APP_NAME)
sudo docker logs -f $id