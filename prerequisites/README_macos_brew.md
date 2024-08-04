# Workshop para Iniciantes - Microsoft Reactor - Golang SP

## Instalação de Dependências

Para garantir uma dinâmica melhor no dia, aqui está um guia de preparação com todas as aplicações que deverão ser instaladas antes de participar do evento.

## Instalando o Homebrew

Homebrew é um gerenciador de pacotes que facilita a instalação de software no macOS. Siga os passos abaixo para instalar o Homebrew:

1. Abra o Terminal.
2. Execute o seguinte comando para instalar o Homebrew:
   ```bash
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
   ```
3. Após a instalação, verifique se o Homebrew foi instalado corretamente executando:
```bash
brew --version
```

## Instalando o Go

1. Com o Homebrew instalado, execute o seguinte comando no Terminal para instalar o Go:

```bash
brew install go
```

2. Abra o Terminal e execute o seguinte comando para adicionar o Go ao PATH do seu sistema operacional:

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

| Se você estiver usando o Zsh, substitua `~/.bashrc` por `~/.zshrc`.


3. Verifique se o Go foi instalado coretamente executando o comando:
```bash
go version
```

## Instalando o Git

1. Execute o seguinte comando no Terminal:

```bash
brew install git
```

2. Verifique se o Git foi instalado corretamente:

```bash
git --version
```

## Instalando o Visual Studio Code (Opcional)

1. Para instalar o VSCode, execute o seguinte comando:

```bash
brew install --cask visual-studio-code
```

2. Verifique a instalação:

```bash
code --version
```
