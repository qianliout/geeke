#!/bin/bash

wget https://storage.googleapis.com/harbor-releases/release-2.8.0/harbor-online-installer-v2.8.0.tgz

tar zxvf harbor-offline-installer-v2.8.2.tgz -C /usr/src/
cd /usr/src/harbor