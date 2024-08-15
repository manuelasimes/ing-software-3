# ing-software-3
## Trabajo Práctico 1 - Git Básico

### 1- Objetivos de Aprendizaje
 - Utilizar herramientas de control de configuración de software
 - Familiarizarse con los comandos más utilizados
 - Configurar el repositorio principal de cada alumno para la materia

### 2- Unidad temática que incluye este trabajo práctico
Este trabajo práctico corresponde a la unidad Nº: 1

### 3- Consignas a desarrollar en el trabajo práctico:
  - Los ejercicios representan casos concretos y rutinarios en uso de este tipo de herramientas
  - En los puntos donde corresponda, proveer los comandos de git necesarios para llevar a cabo el punto.
  - Cuando se especifique alguna descripción, realizarlo de la manera más clara posible y con ejemplos cuando sea necesario.

### 4- Desarrollo:

#### 1- Instalar Git
Los pasos y referencias asumen el uso del sistema operativo Windows, en caso otros SO seguir recomendaciones específicas.

  - Bajar e instalar el cliente git. Por ejemplo, https://git-for-windows.github.io/
  - Bajar e instalar un cliente visual.
 Por ejemplo, TortoiseGit para Windows o SourceTree para Windows/MAC:
    - https://tortoisegit.org/
    - https://www.sourcetreeapp.com/
    - Lista completa: https://git-scm.com/downloads/guis/

#### 2- Crear un repositorio local y agregar archivos
  - Crear un repositorio local en un nuevo directorio.
   ![Texto alternativo](tp01/1.jpg)

  - Agregar un archivo Readme.md, agregar algunas líneas con texto a dicho archivo.
    ![Texto alternativo](tp01/2.jpg)
    ![Texto alternativo](tp01/3.jpg)

  - Crear un commit y proveer un mensaje descriptivo.
    ![Texto alternativo](tp01/4.jpg)


#### 3- Configuración del Editor Predeterminado
 - Instalar Notepad ++ para Windows o TextMate para Mac OS, colocarle un alias y configurarlo como editor predeterminado
   
#### 4- Creación de Repos 01 -> Crearlo en GitHub, clonarlo localmente y subir cambios
  - Crear una cuenta en https://github.com
  - Crear un nuevo repositorio en dicha página con el Readme.md por defecto
    ![Texto alternativo](tp01/5.jpg)

  - Clonar el repo remoto en un nuevo directorio local
    ![Texto alternativo](tp01/6.jpg)

  - Editar archivo Readme.md agregando algunas lineas de texto
    ![Texto alternativo](tp01/7.jpg)
    ![Texto alternativo](tp01/8.jpg)

  - Editar (o crear si no existe) el archivo .gitignore agregando los archivos *.bak
    ![Texto alternativo](tp01/11.jpg)
![Texto alternativo](tp01/10.jpg)
![Texto alternativo](tp01/12.jpg)

  - Crear un commit y porveer un mensaje descriptivo
  - Intentar un push al repo remoto
    ![Texto alternativo](tp01/13.jpg)
![Texto alternativo](tp01/14.jpg)

  - En caso de ser necesario configurar las claves SSH requeridas y reintentar el push.

#### 5- Creación de Repos 02-> Crearlo localmente y subirlo a GitHub
  - Crear un repo local
    ![Texto alternativo](tp01/15.jpg)

  - Agregar archivo Readme.md con algunas lineas de texto
    ![Texto alternativo](tp01/16.jpg)

  - Crear repo remoto en GitHub
  - Asociar repo local con remoto
  - Crear archivo .gitignore
  - Crear un commit y proveer un mensaje descriptivo
  - Subir cambios.
    ![Texto alternativo](tp01/17.jpg)
    ![Texto alternativo](tp01/18.jpg)
![Texto alternativo](tp01/19.jpg)
![Texto alternativo](tp01/20.jpg)
![Texto alternativo](tp01/21.jpg)

![Texto alternativo](tp01/22.jpg)



#### 6- Ramas
  - Crear una nueva rama
    
  - Cambiarse a esa rama
  - Hacer un cambio en el archivo Readme.md y hacer commit
    ![Texto alternativo](tp01/23.jpg)
![Texto alternativo](tp01/24.jpg)

  - Revisar la diferencia entre ramas
![Texto alternativo](tp01/25.jpg)

#### 7- Merges
  - Hacer un merge FF
  - Borrar la rama creada
  - Ver el log de commits
        ![Texto alternativo](tp01/26.jpg)
        ![Texto alternativo](tp01/27.jpg)


  - Repetir el ejercicio 6 para poder hacer un merge con No-FF
    ![Texto alternativo](tp01/28.jpg)
    ![Texto alternativo](tp01/29.jpg)

#### 8- Resolución de Conflictos
  - Instalar alguna herramienta de comparación. Idealmente una 3-Way:
    - P4Merge https://www.perforce.com/downloads/helix-visual-client-p4v:
![alt text](p4merge.png)
    - Se puede omitir registración. Instalar solo opción Merge and DiffTool.
 - ByondCompare trial version https://www.scootersoftware.com/download.php
    - Configurar Tortoise/SourceTree para soportar esta herramienta.
    - https://www.scootersoftware.com/support.php?zz=kb_vcs
    - https://medium.com/@robinvanderknaap/using-p4merge-with-tortoisegit-87c1714eb5e2
  - Crear una nueva rama conflictBranch
        ![Texto alternativo](tp01/30.jpg)

  - Realizar una modificación en la linea 1 del Readme.md desde main y commitear
  - En la conflictBranch modificar la misma línea del Readme.md y commitear
     ![Texto alternativo](tp01/31.jpg)

  - Ver las diferencias con git difftool main conflictBranch
            ![Texto alternativo](tp01/32.jpg)

  - Cambiarse a la rama main e intentar mergear con la rama conflictBranch
            ![Texto alternativo](tp01/33.jpg)

  - Resolver el conflicto con git mergetool
            ![Texto alternativo](tp01/34.jpg)

  - Agregar .orig al .gitignore
            ![Texto alternativo](tp01/35.jpg)

  - Hacer commit y push
        ![Texto alternativo](tp01/36.jpg)

#### 9- Familiarizarse con el concepto de Pull Request

  - Explicar que es un pull request.
    Un pull request (PR) es una solicitud para que los cambios realizados en una rama de un repositorio se integren en otra rama, es decir, una petición para que los cambios en una rama (generalmente una rama de características o una rama de desarrollo) sean revisados e integrados en otra rama (como main o master) del mismo repositorio o en un repositorio diferente.
  - Crear un branch local y agregar cambios a dicho branch.

  - Subir el cambio a dicho branch y crear un pull request.
            ![Texto alternativo](tp01/37.jpg)

  - Completar el proceso de revisión en github y mergear el PR al branch master.
        ![Texto alternativo](tp01/38.jpg)
        ![Texto alternativo](tp01/39.jpg)
            ![Texto alternativo](tp01/40.jpg)
            ![Texto alternativo](tp01/41.jpg)
            ![Texto alternativo](tp01/42.jpg)





#### 10- Algunos ejercicios online
  - Entrar a la página https://learngitbranching.js.org/
  - Completar los ejercicios **Introduction Sequence**
  - Opcional - Completar el resto de los ejercicios para ser un experto en Git!!!
        ![Texto alternativo](tp01/43.jpg)

#### 11- Crear Repositorio de la materia
  - Crear un repositorio para la materia en github. Por ejemplo **ing-software-3**
  - Subir archivo(s) .md con los resultados e imágenes de este trabajo práctico. Puede ser en una subcarpeta **trabajo-practico-01**

### Referencias

- https://try.github.io/
- https://github.github.com/training-kit/downloads/es_ES/github-git-cheat-sheet.pdf
- https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet
