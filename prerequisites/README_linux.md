# Workshop para Iniciantes - Microsoft Reactor - Golang SP

## Instalação de Dependências

Para garantir uma dinâmica melhor no dia, aqui está um guia de preparação com todas as aplicações que deverão ser instaladas antes de participar do evento.

## Instalando o Go

O passo mais importante para participar do workshop é garantir que você possui o Go, que inclui ferramentas e componentes essenciais para o desenvolvimento da nossa aplicação.

[Acesse o site oficial do Go](https://go.dev/dl/) e baixe a versão mais recente.

### Linux

1. Baixe o tarball (arquivo .tar.gz) adequado para a sua arquitetura.
2. Abra o Terminal e extraia o tarball no diretório `/usr/local`:
   ```bash
   tar -C /usr/local -xzf <arquivo.tar.gz>
    ```
3. Adicione o binário do Go ao `PATH` do sistema:
    - Se estiver utilizando Bash:
    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    source ~/.bashrc
    ```
    - Se estiver utilizando Zsh:
    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
    source ~/.zshrc
    ```
4. Verifique a instalação com:
```bash
go version
```

## Instalando o Git

1. Abra o Terminal e execute:
```bash
sudo apt update
sudo apt install git -y
```

2. Verifique a instalação com:
```bash
git --version
```

## Instalando o Visual Studio Code (Opcional)

1. Adicione o repositório o VSCode:
```bash
wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > packages.microsoft.gpg
sudo install -o root -g root -m 644 packages.microsoft.gpg /usr/share/keyrings/
sudo sh -c 'echo "deb [arch=amd64 signed-by=/usr/share/keyrings/packages.microsoft.gpg] https://packages.microsoft.com/repos/vscode stable main" > /etc/apt/sources.list.d/vscode.list'
sudo apt update
```

2. Instale o VSCode:
```bash
sudo apt install code -y
```

3. Verifique a instalação com:
```bash
code --version
```

## Verificação Final

Para garantir que tudo está instalado corretamente, execute os seguintes comandos no Terminal:

```bash
go version
git --version
code --version
```

Se todas as versões forem exibidas corretamente, você está pronto para o workshop!
