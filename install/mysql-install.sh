#!/bin/bash

sudo  apt-get  -y update
sudo apt install -y  mysql-server
sudo systemctl start mysql.service
sudo systemctl enable mysql.service