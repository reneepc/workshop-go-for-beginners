#!/bin/bash

if command -v go &> /dev/null; then
    sudo rm -rf /usr/local/go
fi

if command -v git &> /dev/null; then
    brew uninstall git
fi

if command -v code &> /dev/null; then
    brew uninstall --cask visual-studio-code
fi

sed -i '' '/\/usr\/local\/go\/bin/d' ~/.zshrc
source ~/.zshrc

if ! command -v go &> /dev/null; then echo "Go uninstalled successfully."; fi
if ! command -v git &> /dev/null; then echo "Git uninstalled successfully."; fi
if ! command -v code &> /dev/null; then echo "VSCode uninstalled successfully."; fi
