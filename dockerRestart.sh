#!/bin/bash
APP_NAME="public-tv-server"
id=$(sudo docker restart $APP_NAME)
sudo docker logs -f $id