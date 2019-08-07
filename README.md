# Contenedores desde cero

Este repositorio demuestra a generación de un contenedor en pocas líneas de código. El contenedor esta realizado en Go y está basado en el código y las platicas de [Liz Rice](https://github.com/lizrice) "Building a container from scratch in Go".

## Requisitos

Para ejecutar los ejemplos es necesario instalar [Vagrant](https://www.vagrantup.com/downloads.html).

## ¿Cómo ejecutar?

El ambiente Vagrant y este repo incluye todo lo necesario para ejecutar los ejemplos. Desde la raíz del repositorio ejecutar:

```
vagrant up

vagrant ssh
```
para acceder al ambiente.


Dentro de la carpeta `/vagrant/demo` se encuentra el archivo **containers.go** el cual contiene el demo completo.

Dentro de las carpetas `/vagrant/demo/slides` se encuentran cada uno de los ejemplos correspondientes con la presentación.

Para ejecutar cualquier ejemplo basta con acceder a alguna de las carpetas, compilar el ejemplo y ejecutarlo:

```
go build -o demo && sudo ./demo run echo "hola mundo"
```

Recordando que la sintaxis del la demostración es:

**`./demo run <comando> <args>`**

## Contenido de las carpetas

Dentro de la carpeta `demo` se encuentra las siguientes carpetas, en orden de las slides:

* `/demo/containers.go`:  Demo completo con los namespaces UTS, PID, MNT y USER.

* `/demo/slides/base1`: Contiene wrapper simple de comandos.

* `/demo/slides/base2`: Contiene el wrapper padre (host) que ejecuta al (hijo). El padre permite modificar el ambiente del hijo y generar el contenedor

* `/demo/slides/utsnamespace`: Contiene el namespace UTS el cual permite modificar el hostname del guest sin modificar el del host.

* `/demo/slides/pidnamespace`: Contiene el namespace PID el cual permite aislar procesos.

* `/demo/slides/chroot`: Permite cargar un nuevo filesystem (Alpine en el código)

* `/demo/slides/procfs`: Contiene el namespace MNT el cual permite montar el filesystem `proc`.

* `/demo/slides/cgroups`: Contiene el acceso a los cgroups generando un control group para limitar la memoria.

# ToDo:

Agregar el resto de los namespaces:
* NET
* IPC