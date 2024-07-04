# Workshop para iniciantes - Microsoft Reactor - Golang SP

## Instalação de dependências

Para garantir uma dinâmica melhor no dia, aqui está um guia de preparação com todas as aplicações que deverão ser instaladas antes de atender ao evento.

## Instalando o Go

O passo mais importante para participar do Workshop é garantir que você possui o Go, o que inclui ferramentas e componentes essenciais para o desenvolvimento da nossa aplicação.

[Acesse o site oficial do Go](https://go.dev/dl/) e baixe a versão mais recente, ou clique nos links abaixo para baixar o instalador.

### Windows

Baixe o instalador (arquivo .msi), execute-o e siga as instruções na tela.

### Mac

Baixe o pacote de instalação (arquivo .pkg) adequado, de acordo com a arquitetura do seu Mac ARM ou x86.

Caso não saiba a arquitetura do seu Mac, siga os seguintes passos:

1. Clique no ícone da Apple no canto superior esquerdo da tela.
2. No menu suspenso, selecione a opção `Sobre Este Mac`.
3. Na janela que se abre, você verá informações sobre seu Mac, incluindo o tipo de processador.
    - Se você vir algo como `Intel Core`, a arquitetura do seu processador é baseada em x86.
    - Se você vir algo como `Chip Apple M1/2/3`, a arquitetura do seu processador é baseada em ARM ou Apple Silicon.

### Linux

Baixe o tarball (arquivo .tar.gz) e abra seu terminal.

No terminal, extraia o tarball no diretório:

```bash
tar -C /usr/local -xzf <arquivo.tar.gz>
```

Para que consiga executar os comandos do Go é necessário adicionar o binário do Go ao `PATH` do seu sistema operacional. Você pode fazer isso temporáriamente com o comando `export PATH=$PATH:/usr/local/go/bin` ou definitivamente executando os comandos abaixo:

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

Se seu shell for o zsh:

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
source ~/.zshrc
```

## Instalando o Git

Acesse o [site oficial do Git](https://git-scm.com/downloads) e baixe a versão mais recente.

### Windows

Baixe o instalador (arquivo .exe), execute-o e siga as instruções.

### Mac

O Git geralmente já estará instalado no seu sistema operacional. Você pode conferir isso abrindo seu terminal e executando o comando:

```bash
git version
```

Caso já não esteja instalado, baixe o pacote 
