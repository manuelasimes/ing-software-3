4.1 Modificar nuestro pipeline para construir imágenes Docker de back y front y subirlas a ACR
      ![Texto alternativo](imagenes/2.png)
      ![Texto alternativo](imagenes/3.png)
      ![Texto alternativo](imagenes/4.png)
  4.1.2 Crear un recurso ACR en Azure Portal siguiendo el instructivo 5.1
        ![Texto alternativo](imagenes/1.png)
              ![Texto alternativo](imagenes/5.png)
  4.1.3 Modificar nuestro pipeline en la etapa de Build y Test
Aca tenemos para la publicacion del dockerfile de back: 
![Texto alternativo](imagenes/6.png)
De forma analoga, como en el instructivo, hice lo mismo para el front. 
4.1.4 En caso de no contar en nuestro proyecto con una ServiceConnection a Azure Portal para el manejo de recursos, agregar una service connection a Azure Resource Manager como se indica en instructivo 5.2
Si contaba con una:
![Texto alternativo](imagenes/8.png)
![Texto alternativo](imagenes/9.png)

4.1.5 Agregar a nuestro pipeline variables
![Texto alternativo](imagenes/10.png)
4.1.6 Agregar a nuestro pipeline una nueva etapa que dependa de nuestra etapa de Build y Test
![Texto alternativo](imagenes/11.png)
4.1.7 - Ejecutar el pipeline y en Azure Portal acceder a la opción Repositorios de nuestro recurso Azure Container Registry. Verificar que exista una imagen con el nombre especificado en la variable backImageName asignada en nuestro pipeline
![Texto alternativo](imagenes/12.png)
![Texto alternativo](imagenes/13.png)
![Texto alternativo](imagenes/14.png)
4.1.8 - Agregar tareas para generar imagen Docker de Front (DESAFIO)
![Texto alternativo](imagenes/17.png)
![Texto alternativo](imagenes/16.png)
![Texto alternativo](imagenes/19.png)
4.1.9 - Agregar a nuestro pipeline una nueva etapa que dependa de nuestra etapa de Construcción de Imagenes Docker y subida a ACR
Agregar variables a nuestro pipeline:
![Texto alternativo](imagenes/20.png)
Agregar variable secreta cnn-string-qa desde la GUI de ADO que apunte a nuestra BD de SQL Server de QA como se indica en el instructivo 5.3
![Texto alternativo](imagenes/21.png)

Agregar tareas para crear un recurso Azure Container Instances que levante un contenedor con nuestra imagen de back
![Texto alternativo](imagenes/22.png)
![Texto alternativo](imagenes/23.png)
Para que funcione, debemos hacer lo siguiente de forma local:
![Texto alternativo](imagenes/24.png)
![Texto alternativo](imagenes/25.png)
4.1.10 - Ejecutar el pipeline y en Azure Portal acceder al recurso de Azure Container Instances creado. Copiar la url del contenedor y navegarlo desde browser. Verificar que traiga datos.
   ![Texto alternativo](imagenes/27.png)
![Texto alternativo](imagenes/28.png)
4.1.11 - Agregar tareas para generar un recurso Azure Container Instances que levante un contenedor con nuestra imagen de front (DESAFIO)
ANtes debemos modificar archivos de EmployeeCrudAngular: 
![Texto alternativo](imagenes/29.png)
![Texto alternativo](imagenes/30.png)
![Texto alternativo](imagenes/31.png)
![Texto alternativo](imagenes/32.png)
![Texto alternativo](imagenes/36.png)
Subimos los cambios:
![Texto alternativo](imagenes/33.png)
Luego:
![Texto alternativo](imagenes/34.png)
![Texto alternativo](imagenes/35.png)
![Texto alternativo](imagenes/37.png)
![Texto alternativo](imagenes/38.png)
4.1.12 - Agregar tareas para correr pruebas de integración en el entorno de QA de Back y Front creado en ACI.
![Texto alternativo](imagenes/41.png)
Debemos hacer cambios en las pruebas:
![Texto alternativo](imagenes/39.png)
Se corren correctamente:
![Texto alternativo](imagenes/40.png)
![Texto alternativo](imagenes/42.png)

4.2.4 Agregar etapa que dependa de la etapa de Deploy en ACI QA y genere contenedores en ACI para entorno de PROD.
Creamos un environment:
![Texto alternativo](imagenes/43.png)
![Texto alternativo](imagenes/44.png)
![Texto alternativo](imagenes/45.png)
Agregamos nuevas variables:
![Texto alternativo](imagenes/46.png)
![Texto alternativo](imagenes/52.png)
Agregamos las tareas al pipeline:
![Texto alternativo](imagenes/53.png)
Corremos el pipeline y aceptamso:
![Texto alternativo](imagenes/47.png)
Se ejecuta correctamente:
![Texto alternativo](imagenes/48.png)
![Texto alternativo](imagenes/49.png)
![Texto alternativo](imagenes/50.png)
![Texto alternativo](imagenes/51.png)








