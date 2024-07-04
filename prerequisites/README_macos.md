# Workshop para Iniciantes - Microsoft Reactor - Golang SP

## Instalação de Dependências

Para garantir uma dinâmica melhor no dia, aqui está um guia de preparação com todas as aplicações que deverão ser instaladas antes de participar do evento.

## Descobrindo a Arquitetura do Mac

1. Clique no ícone da Apple no canto superior esquerdo da tela.
2. No menu suspenso, selecione a opção `Sobre Este Mac`.
3. Na janela que se abre, você verá informações sobre seu Mac, incluindo o tipo de processador.
    - Se você vir algo como `Intel Core`, a arquitetura do seu processador é baseada em x86.
    - Se você vir algo como `Chip Apple M1/2/3`, a arquitetura do seu processador é baseada em ARM ou Apple Silicon.

## Instalando o Go

O passo mais importante para participar do workshop é garantir que você possui o Go, que inclui ferramentas e componentes essenciais para o desenvolvimento da nossa aplicação.

[Acesse o site oficial do Go](https://go.dev/dl/) e baixe a versão mais recente.

1. Baixe o pacote de instalação (arquivo .pkg) adequado para a arquitetura do seu Mac (ARM/Apple Silicon ou x86).
2. Execute o instalador e siga as instruções na tela.
3. Após a instalação, abra o Terminal e verifique a instalação com:

```bash
   go version
```

## Instalando o Git

O Git geralmente já está instalado no macOS. Verifique abrindo o Terminal e digitando:

```bash
git --version
```

Se o Git não estiver instalado, você verá uma mensagem solicitando a instalação das ferramentas de desenvolvedor da Apple (Xcode Command Line Tools). Siga as instruções na tela para instalar.

## Instalando o Visual Studio Code (Opcional)
1. Acesse o [site oficial do VSCode](https://code.visualstudio.com/Download) e baixe a versão mais recente.
2. Execute o instalador (arquivo .dmg) e arraste o ícone do VSCode para a pasta Aplicativos.
3. Após a instalação, abra o VSCode e verifique a versão indo em Code -> About.

## Verificação Final

Para garantir que tudo está instalado corretamente, execute os seguintes comandos no Terminal:

``` bash
    go version
    git --version
    code --version
```

Se todas as versões forem exibidas corretamente, você está pronto para o workshop!
