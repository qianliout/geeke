#!/bin/bash

sudo apt update
sudo apt -y install git
sudo apt -y install zsh
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
