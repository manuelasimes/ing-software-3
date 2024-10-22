Desarrollo del punto 4.1:
Creamos el App service: 
      ![Texto alternativo](imagenes/1.png)
Se creo correctamente: 
      ![Texto alternativo](imagenes/2.png)
Agregamos variables a nustro pipeline: 
      ![Texto alternativo](imagenes/3.png)

4.1.1 - Agregar a nuestro pipeline una nueva etapa que dependa de nuestra etapa de Construcción y Pruebas y de la etapa de Construcción de Imagenes Docker y subida a ACR realizada en el TP08
Agregar tareas para crear un recurso Azure Container Instances que levante un contenedor con nuestra imagen de back utilizando un AppServicePlan en Linux
Pegamos el deploy de app service:
      ![Texto alternativo](imagenes/5.png)
Se ejecuta correctamente:
      ![Texto alternativo](imagenes/31.png)
      ![Texto alternativo](imagenes/8.png)

Se crea correctamente la app en el App Service Plan:
      ![Texto alternativo](imagenes/6.png)
Y vemos que se conecta correctamente con la base de datos y funciona:
  ![Texto alternativo](imagenes/7.png)


4.2 Desafíos:
4.2.1 Agregar tareas para generar Front en Azure App Service con Soporte para Contenedores
Agreagamos las tareas:
  ![Texto alternativo](imagenes/32.png)
Se ejecuta corretctamente:
  ![Texto alternativo](imagenes/9.png)
Y se crea la app:
  ![Texto alternativo](imagenes/10.png)
  ![Texto alternativo](imagenes/12.png)

4.2.2 Agregar variables necesarias para el funcionamiento de la nueva etapa considerando que debe haber 2 entornos QA y PROD para Back y Front.
  ![Texto alternativo](imagenes/13.png)

4.2.3 Agregar tareas para correr pruebas de integración en el entorno de QA de Back y Front creado en Azure App Services con Soporte para Contenedores.
  ![Texto alternativo](imagenes/15.png)
  Funcionan: 
    ![Texto alternativo](imagenes/14.png)

4.2.4 Agregar etapa que dependa de la etapa de Deploy en QA que genere un entorno de PROD.
    ![Texto alternativo](imagenes/33.png)
Se crean las apps de prod:   
![Texto alternativo](imagenes/19.png)
Funcionan:
    ![Texto alternativo](imagenes/20.png)
    ![Texto alternativo](imagenes/34.png)


4.2.5 Entregar un pipeline que incluya:
A) Etapa Construcción y Pruebas Unitarias y Code Coverage Back y Front
Para hacerlo reutilice el tp 7:
    ![Texto alternativo](imagenes/38.png)

Vemos el code coverage:
    ![Texto alternativo](imagenes/35.png)

B) Construcción de Imágenes Docker y subida a ACR
    ![Texto alternativo](imagenes/39.png)
    ![Texto alternativo](imagenes/40.png)

C) Deploy Back y Front en QA con pruebas de integración para Azure Web Apps
    ![Texto alternativo](imagenes/41.png)
    Pruebas se ejecutan correctamente:
    ![Texto alternativo](imagenes/25.png)
TAnto el front como la api funcionan:
    ![Texto alternativo](imagenes/24.png)
    ![Texto alternativo](imagenes/23.png)

D) Deploy Back y Front en QA con pruebas de integración para ACI
    ![Texto alternativo](imagenes/41.png)
    ![Texto alternativo](imagenes/42.png)
    ![Texto alternativo](imagenes/43.png)
    ![Texto alternativo](imagenes/44.png)

E) Deploy Back y Front en QA con pruebas de integración para Azure Web Apps con Soporte para contenedores
Este punto esta desarrollado más especificamente en la primer parte del trabajo. 
    ![Texto alternativo](imagenes/45.png)
    ![Texto alternativo](imagenes/46.png)
    ![Texto alternativo](imagenes/47.png)

F) Aprobación manual de QA para los puntos C,D,E
    ![Texto alternativo](imagenes/37.png)
Para el deploy a prod de Web Apps:     
![Texto alternativo](imagenes/26.png)
Y para los demas deploy a Prod:
    ![Texto alternativo](imagenes/49.png)

G) Deploy Back y Front en PROD para Azure Web Apps

H) Deploy Back y Front en PROD para ACI
I) Deploy Back y Front en PROD para Azure Web Apps con Soporte para contenedores
Deploy a Prod hechos: 
    ![Texto alternativo](imagenes/28.png)
Vemos como funcionan: 
    ![Texto alternativo](imagenes/50.png)
    ![Texto alternativo](imagenes/51.png)
    ![Texto alternativo](imagenes/52.png)
    ![Texto alternativo](imagenes/53.png)
    Web Apps:
   ![Texto alternativo](imagenes/56.png)
   ![Texto alternativo](imagenes/58.png)


