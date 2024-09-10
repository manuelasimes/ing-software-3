Desarrollo:
4.1. Crear una cuenta en Azure
Ya tenia creada una. 
4.2. Crear un recurso Web App en Azure Portal y navegar a la url provista
   ![Texto alternativo](imagenes/1.png)
   ![Texto alternativo](imagenes/2.png)
   ![Texto alternativo](imagenes/3.png)
   Navegamos a la URL provista:
   ![Texto alternativo](imagenes/4.png)
4.3. Actualizar Pipeline de Build para que use tareas de DotNetCoreCLI@2 como en el pipeline clásico, luego crear un Pipeline de Release en Azure DevOps con CD habilitada
   ![Texto alternativo](imagenes/7.png)
   ![Texto alternativo](imagenes/8.png)
   ![Texto alternativo](imagenes/9.png)
   Creamos pipeline de release:
  ![Texto alternativo](imagenes/5.png)
  ![Texto alternativo](imagenes/6.png)
  ![Texto alternativo](imagenes/10.png)
  Habilitamos CD:
  ![Texto alternativo](imagenes/11.png)
  ![Texto alternativo](imagenes/12.png)
  ![Texto alternativo](imagenes/13.png)
  ![Texto alternativo](imagenes/14.png)
4.4. Optimizar Pipeline de Build
![Texto alternativo](imagenes/47.png)

4.5. Verificar el deploy en la url de la WebApp /weatherforecast
![Texto alternativo](imagenes/46.png)
4.6. Realizar un cambio al código del controlador para que devuelva 7 pronósticos, realizar commit, evaluar ejecución de pipelines de build y release, navegar a la url de la webapp/weatherforecast y corroborar cambio
Hacemos cambio para que devuelva 7 pronosticos:
![Texto alternativo](imagenes/48.png)
![Texto alternativo](imagenes/49.png)
Vemos que efectivamente nos muestra 7 pronosticos:
![Texto alternativo](imagenes/50.png)

4.7. Clonar la Web App de QA para que contar con una WebApp de PROD a partir de un Template Deployment en Azure Portal y navegar a la url provista para la WebApp de PROD.
![Texto alternativo](imagenes/17.png)
![Texto alternativo](imagenes/18.png)
Vemos que se ha clonado la web app correctamente: 
![Texto alternativo](imagenes/19.png)
4.8. Agregar una etapa de Deploy a Prod en Azure Release Pipelines
![Texto alternativo](imagenes/20.png)
![Texto alternativo](imagenes/21.png)
![Texto alternativo](imagenes/22.png)
4.9. Realizar un cambio al código del controlador para que devuelva 10 pronósticos, realizar commit, evaluar ejecución de pipelines de build y release, navegar a la url de la webapp/weatherforecast y corroborar cambio, verificar que en la url de la webapp_prod/weatherforecast se muestra lo mismo.
![Texto alternativo](imagenes/25.png)
![Texto alternativo](imagenes/26.png)
![Texto alternativo](imagenes/27.png)
![Texto alternativo](imagenes/28.png)
![Texto alternativo](imagenes/33.png)
4.10. Modificar pipeline de release para colocar una aprobación manual para el paso a Producción.
![Texto alternativo](imagenes/24.png)

4.11. Realizar un cambio al código del controlador para que devuelva 5 pronósticos, realizar commit, evaluar ejecución de pipelines de build y release, navegar a la url de la webapp/weatherforecast y corroborar cambio, verificar que en la url de la webapp_prod/weatherforecast aun se muestra la versión anterior.
Hacemos cambio para que devuelva 5 pronosticos:
![Texto alternativo](imagenes/51.png)
Vemos que la webapp si nos muestra el cambio: 
![Texto alternativo](imagenes/52.png)
EN cambio, en webapp-prod, como no aceptamos el cambio, sigue mostrando 7: 
![Texto alternativo](imagenes/53.png)

4.12. Aprobar el pase ya sea desde el release o desde el mail recibido. 
4.12.1. Notar que se puede dar la aprobación pero posponer su aplicación hasta una determinada fecha
4.13. Esperar a la finalización de la etapa de Pase a Prod y luego corroborar que en la url de la webapp_prod/weatherforecast se muestra la nueva versión coinicidente con la de QA. image
![Texto alternativo](imagenes/28.png)
![Texto alternativo](imagenes/29.png)
![Texto alternativo](imagenes/30.png)
![Texto alternativo](imagenes/31.png)
![Texto alternativo](imagenes/32.png)
Y vemos como muestra los 10 pronosticos:
![Texto alternativo](imagenes/33.png)

4.14. Realizar un pipeline (no release) que incluya el deploy a QA y a PROD con una aprobación manual. El pipeline debe estar construido en YAML sin utilizar el editor clásico de pipelines ni el editor clásico de pipelines de release.
![Texto alternativo](imagenes/34.png)
![Texto alternativo](imagenes/35.png)
![Texto alternativo](imagenes/36.png)
![Texto alternativo](imagenes/38.png)
![Texto alternativo](imagenes/39.png)
![Texto alternativo](imagenes/40.png)
![Texto alternativo](imagenes/41.png)
![Texto alternativo](imagenes/42.png)
![Texto alternativo](imagenes/43.png)
![Texto alternativo](imagenes/45.png)


