# didactic-container
📚 📦 **Repository for learning how containers works in Go.** 
⚠️ `In progress`

<img src="https://github.com/realnfcs/didactic-container/blob/main/public/dcont_image.png" alt="container_img" width=400em height=300em align="left" hspace="10" vspace="6">
The objective of the project is to create an environment/infrastructure like a Docker to understand concepts and methods such as the creation of containers and their configuration, management and download of images, among others. This repository is learning purpose only, any contribution is welcome.


## Credits for:

Thanks for all knowledge you two have shared with us

* Liz Rice / [@lizrice](https://github.com/lizrice) 
* Fabio Akita / [@akitaonrails](https://github.com/akitaonrails)

## Documentation

`This small documentation is on PT-BR (Portuguese, Brazil) idiom, in the future, the documentation will have a english idiom.`

**[PT-BR]** ⚠️ `Em progresso`
- [O que é um container?](#o-que-é-um-container)
	- [Cgroup](#cgroup)
	- [Namespaces](#namespaces)
- [Project Organization](#project-organization)
- [Image Engine](#image-engine)
	- [Como vai funcionar?](#como-vai-funcionar)
	- [Por que o SQLite?](#por-que-o-sqlite)
	- [TODO](#todo)


### O que é um container?

**Containers são patchs na Kernel do SO que possibilita restringir os recursos do sistema aos processos**.
Com um conjunto de funcionalidade como Cgroups e Namespaces, podemos restringir o que o(s) processo(s) dentro 
do container conseguem **usar e ver**.

	Fazemos a Kernel "mentir" para os processos. Essa "mentira" chamados de container. - Akita
 
Junto com as imagens, que são "cópias" ou snapshots de sistema operacional em um grande arquivo, podemos empacotar um filesystem 
inteiro para podermos usar-lo várias vezes, o que é comparável com as imagens do Docker. Usando a imagens do Ubuntu, por exemplo,
conseguimos usar suas funcionalodades básicas atribuidas pelo seu filesystem, como o gerenciador de pacotes apt, uso da Kernel 
"containerizado" etc.

Com a imagem em mãos, podemos subir um container para realizar os procedimentos necessários, como rodar um servidor HTTP, realizar
processos etc.

### Cgroup

Como mencionado antes, o gerenciador de um container usa da funcionalidade do Cgroup para limitar o que um container consegue usar
da máquina, como CPU e RAM, e da Kernel, como pids (Process ID).

### Namespaces

Namespaces limita o que o contaiener consegue ver, como Process ID externos, processos da Rede, entre outros. Funcionam também como 
uma forma de etiquetar processos que, quando criado e implementado em um container, cria um PID 1 onde seus "filhos" seguirão essa
sequência.

## Project Organization

<div align="center">
	<img src="https://github.com/realnfcs/didactic-container/blob/main/public/project_infrastructure.png" alt="project_organization" align="center">
</div>

## Image Engine

	go run main.go image <cmd> <params>

Na image Engine, teremos a CLI que ficará responsável pela instalação, configuração e gerenciamento das imagens/filesystems que serão usados para criar um container.

### Como vai funcionar?

Tendo o exemplo com o filesystem do Alpine. O usuário terá que digitar o comando ` go run main.go image create `  para iniciar a criação e gerenciamento das imagens.

Ao digitar o comando, algo similar terá que aparecer:


	$ go run main.go image create
	
	$ Com qual das imagens abaixo você que criar seu container?
	
	$ > Ubuntu
	
	$ > Alpine
	
	$ > Personalizado


**Se ele escolher o personalizado, apenas terá que informar o 
local onde o filesystem está localizado.**

Escolhendo o nosso exemplo usando o Alpine, o CLI terá que: 

1. Primeiramente verificar se já não está instalado.
2. Se não estiver, instalar o filesystem.
3. Guardar em um local apropriado

**Esse será o próposito desta parte -> Instalar a image a ser utilizada!**

Para fins de facilidade, a Image Engine também terá um comando `info` para observar as imagens instaladas e personalizadas.

### Por que o SQLite?

Para melhor gerenciamento, usarei o SQLite para guardar informações das imagens e de outros futuros dados, então, para adiantar, já o deixarei configurado.

## TODO

- [ ] Preparar o ambiente para criar a Image Engine
- [ ] Criar a conexão SQLite
- [ ] Criar as funcionalidades para instalar e configurar as Imagens/Filesystems a serem usadas
- [ ] Criar a funcionalidade para ver as imagens instaladas
