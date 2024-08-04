#!/bin/bash

if ! command -v brew &> /dev/null; then
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi

brew install go
brew install git

if ! command -v code &> /dev/null; then
    brew install --cask visual-studio-code
fi

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
source ~/.zshrc

go version
git --version
code --version
