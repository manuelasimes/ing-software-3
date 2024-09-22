Prerequisitos:
En primer lugar, instalé lo necesraio para desarrollar el trabajo. 
    ![Texto alternativo](imagenes/1.jpg)
    ![Texto alternativo](imagenes/2.jpg)
    ![Texto alternativo](imagenes/3.jpg)

4.1 Creación de una BD SQL Server para nuestra App
Opte por la opción 2: 
    ![Texto alternativo](imagenes/4.jpg)
    ![Texto alternativo](imagenes/5.jpg)
    ![Texto alternativo](imagenes/6.jpg)
Clonar el repo https://github.com/ingsoft3ucc/Angular_WebAPINetCore8_CRUD_Sample.git
    ![Texto alternativo](imagenes/clonar.jpg)

B. Seguir las instrucciones del README.md del repo clonado prestando atención a la modificación de la cadena de conexión en el appSettings.json para que apunte a la BD creada en 4.1
    ![Texto alternativo](imagenes/8.jpg)
C. Navegar a http://localhost:7150/swagger/index.html y probar uno de los controladores para verificar el correcto funcionamiento de la API.
    ![Texto alternativo](imagenes/7.jpg)
D. Navegar a http://localhost:4200 y verificar el correcto funcionamiento de nuestro front-end Angular
    ![Texto alternativo](imagenes/9.jpg)
E. Una vez verificado el correcto funcionamiento de la Aplicación procederemos a crear un proyecto de pruebas unitarias para nuestra API.

4.3 Crear Pruebas Unitarias para nuestra API
A. En el directorio raiz de nuestro repo crear un nuevo proyecto de pruebas unitarias para nuestra API
    ![Texto alternativo](imagenes/10.jpg)
B. Instalar dependencias necesarias
    ![Texto alternativo](imagenes/11.jpg)
C. Editar archivo UnitTest1.cs reemplazando su contenido por
    ![Texto alternativo](imagenes/12.jpg)
D. Renombrar archivo UnitTest1.cs por EmployeeControllerUnitTests.cs
    ![Texto alternativo](imagenes/13.jpg)
E. Editar el archivo EmployeeCrudApi.Tests/EmployeeCrudApi.Tests.csproj para agregar una referencia a nuestro proyecto de EmployeeCrudApi reemplazando su contenido por
Lo edité
F. Ejecutar los siguientes comandos para ejecutar nuestras pruebas
G. Verificar que se hayan ejecutado correctamente las pruebas
    ![Texto alternativo](imagenes/14.jpg)
    ![Texto alternativo](imagenes/15.jpg)
H. Verificar que no estamos usando una dependencia externa como la base de datos.

I. Modificar la cadena de conexión en el archivo appsettings.json para que use un usuario o password incorrecto y recompilar el proyecto EmployeeCrudApi
    ![Texto alternativo](imagenes/16.jpg)
    ![Texto alternativo](imagenes/17.jpg)
J. Verificar que nuestro proyecto ya no tiene acceso a la BD navegando a http://localhost:7150/swagger/index.html y probando uno de los controladores:
    ![Texto alternativo](imagenes/18.jpg)
K. En la carpeta de nuestro proyecto EmployeeCrudApi.Tests volver a correr las pruebas

L. Verificar que se hayan ejecutado correctamente las pruebas inclusive sin tener acceso a la BD, lo que confirma que es efectivamente un conjunto de pruebas unitarias que no requieren de una dependencia externa para funcionar.

    ![Texto alternativo](imagenes/19.jpg)
M. Modificar la cadena de conexión en el archivo appsettings.json para que use el usuario y password correcto y recompilar el proyecto EmployeeCrudApi
N. Verificar que nuestro proyecto vuelve a tener acceso a la BD navegando a http://localhost:7150/swagger/index.html y probando uno de los controladores:
    ![Texto alternativo](imagenes/20.jpg)

4.4 Creamos pruebas unitarias para nuestro front de Angular:
A. Nos posicionamos en nuestro proyecto de front, en el directorio EmployeeCrudAngular/src/app

B. Editamos el archivo app.component.spec.ts reemplazando su contenido por:
    ![Texto alternativo](imagenes/21.jpg)
C. Creamos el archivo employee.service.spec.ts reemplazando su contenido por:
    ![Texto alternativo](imagenes/22.jpg)
D. Editamos el archivo employee.component.spec.ts ubicado en la carpeta employee reemplazando su contenido por:
Lo edité
E. Editamos el archivo addemployee.component.spec.ts ubicado en la carpeta addemployee reemplazando su contenido por:
    ![Texto alternativo](imagenes/23.jpg)
F. En el directorio raiz de nuestro proyecto EmployeeCrudAngular ejecutamos el comando

ng test
G. Vemos que se abre una ventana de Karma con Jasmine en la que nos indica que los tests se ejecutaron correctamente
Efectivamente se abió: 
    ![Texto alternativo](imagenes/24.jpg)
Las pruebas se ejecutaron correctamente. 
I. Verificamos que no esté corriendo nuestra API navegando a http://localhost:7150/swagger/index.html y recibiendo esta salida:
    ![Texto alternativo](imagenes/25.jpg)
4.5 Agregamos generación de reporte XML de nuestras pruebas de front.
A. Instalamos dependencia karma-junit-reporter
    ![Texto alternativo](imagenes/26.jpg)

B. En el directorio raiz de nuestro proyecto (al mismo nivel que el archivo angular.json) creamos un archivo karma.conf.js con el siguiente contenido
    ![Texto alternativo](imagenes/27.jpg)

C. Ejecutamos nuestros test de la siguiente manera:

ng test --karma-config=karma.conf.js --watch=false --browsers ChromeHeadless
D. Verificamos que se creo un archivo test-result.xml en el directorio test-results que está al mismo nivel que el directorio src
    ![Texto alternativo](imagenes/28.jpg)

4.6 Modificamos el código de nuestra API y creamos nuevas pruebas unitarias:
Realice la ssiguientes 5 modificaciones: 
Al agregar y al editar un empleado, controlar que el nombre del empleado no esté repetido.
    ![Texto alternativo](imagenes/29.jpg)
Vemos que funciona: 
    ![Texto alternativo](imagenes/30.jpg)
La longitud máxima del nombre y apellido del empleado debe ser de 100 caracteres.
Almacenar el nombre en la BD siempre con la primera letra de los nombres en Mayuscula y todo el apellido en Mayusculas. Ejemplo, si recibo juan carlos chamizo, se debe almacenar como Juan Carlos CHAMIZO.
Asegurar que el nombre del empleado no contenga caracteres especiales o números, a menos que sea necesario (por ejemplo, caracteres especiales en apellidos como "O'Connor" o "García").
    ![Texto alternativo](imagenes/32.jpg)
Validar que el nombre tenga un número mínimo de caracteres, por ejemplo, al menos dos caracteres para evitar entradas inválidas como "A".
    ![Texto alternativo](imagenes/31.jpg)
B. Crear las pruebas unitarias necesarias para validar las modificaciones realizadas en el código
    ![Texto alternativo](imagenes/33.jpg)
    ![Texto alternativo](imagenes/34.jpg)
    ![Texto alternativo](imagenes/35.jpg)
    ![Texto alternativo](imagenes/36.jpg)
    ![Texto alternativo](imagenes/37.jpg)

4.7 Modificamos el código de nuestro Front y creamos nuevas pruebas unitarias:
A. Realizar en el código del front las mismas modificaciones hechas a la API.

Vemos como responde el front en estos dos ejemplos:
    ![Texto alternativo](imagenes/38.jpg)
    ![Texto alternativo](imagenes/39.jpg)
B. Las validaciones deben ser realizadas en el front sin llegar a la API, y deben ser mostradas en un toast como por ejemplo https://stackblitz.com/edit/angular12-toastr?file=src%2Fapp%2Fapp.component.ts o https://stackblitz.com/edit/angular-error-toast?file=src%2Fapp%2Fcore%2Frxjsops.ts
Vemos como se usa Toastr:
    ![Texto alternativo](imagenes/40.jpg)
    ![Texto alternativo](imagenes/42.jpg)
Para este tercer caso, se ve como se modifica el apellido por todas mayusculas, y como se crea una lista local sin llegar al back y no deja repetir nombres:
    ![Texto alternativo](imagenes/43.jpg)
C. Crear las pruebas unitarias necesarias en el front para validar las modificaciones realizadas en el código del front.
    ![Texto alternativo](imagenes/44.jpg)
    ![Texto alternativo](imagenes/45.jpg)
    ![Texto alternativo](imagenes/46.jpg)






