#!/bin/bash

if command -v go &> /dev/null; then
    sudo rm -rf /usr/local/go
fi

# Uninstall Git
if command -v git &> /dev/null; then
    sudo apt remove --purge -y git
fi

if command -v code &> /dev/null; then
    sudo apt remove --purge -y code
fi

sed -i '/\/usr\/local\/go\/bin/d' ~/.bashrc
source ~/.bashrc

if ! command -v go &> /dev/null; then echo "Go uninstalled successfully."; fi
if ! command -v git &> /dev/null; then echo "Git uninstalled successfully."; fi
if ! command -v code &> /dev/null; then echo "VSCode uninstalled successfully."; fi

