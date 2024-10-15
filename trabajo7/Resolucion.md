 Desarrollo:
#### Prerequisitos:

#### 4.1 Agregar Code Coverage a nuestras pruebas unitarias de backend y front-end e integrarlas junto con sus resultados en nuestro pipeline de build.
   ![Texto alternativo](imagenes/uno.png)
   ![Texto alternativo](imagenes/tres.png)
   ![Texto alternativo](imagenes/dos.png)
  ![Texto alternativo](imagenes/cinco.png)
     ![Texto alternativo](imagenes/cuatro.png)


#### Desarrollo del punto 4.2: Demostración de cómo integrar SonarCloud en un pipeline de CI/CD y cómo leer los reportes de análisis estático.

 #### 4.2.1 Integraremos SonarCloud para analizar el código fuente. Configurar SonarCloud en nuestro pipeline siguiendo instructivo 5.1
 Creo mi cuenta:
      ![Texto alternativo](imagenes/1.png)
Y desarrollo el instructivo 
      ![Texto alternativo](imagenes/2.png)
      ![Texto alternativo](imagenes/3.png)
      ![Texto alternativo](imagenes/4.png)
      ![Texto alternativo](imagenes/6.png)
      ![Texto alternativo](imagenes/7.png)
Modifico el pipeline:
![Texto alternativo](imagenes/8.png)
![Texto alternativo](imagenes/9.png)
![Texto alternativo](imagenes/10.png)
Voy al link: 
![Texto alternativo](imagenes/12.png)
![Texto alternativo](imagenes/13.png)
Podemos ver: los problemas detectados por Sonar Cloud, ya sean issues relacionados a la calidad del software y otros a atributos de codigo limpio.

Clean Code Atributtes:

Intencionalidad: Se refiere a si el código expresa claramente su propósito.
Consistencia: Señala si el código sigue un estilo coherente y consistente. La consistencia es crucial para que el código sea más legible y mantenible.
Adaptabilidad y Responsabilidad: 
Adaptabilidad tiene 0 problemas. La Adaptabilidad indica qué tan fácil es modificar el código para adaptarse a cambios
Responsabilidad tiene 3, si cada clase o método tiene una sola responsabilidad.
Software Quality Attributes

Seguridad: Muestra problemas relacionados con la seguridad del código.
Confiabilidad: Indica la capacidad del software para funcionar correctamente bajo condiciones esperadas.
Mantenibilidad: Evalúa cuán fácil es mantener el código.
Luego, se presenta la gravedad de los problemas, clasificados como alta, media o baja, lo que indica su nivel de criticidad. La mayoría de los problemas se categorizan como de severidad alta o media, lo que señala áreas críticas que necesitan atención urgente. Además, se especifica el tipo de problemas y el esfuerzo estimado para cada uno. SonarCloud ofrece un tiempo aproximado para resolver cada incidencia, lo que facilita la priorización de tareas y la planificación de las acciones necesarias para mejorar el código. Finalmente, se muestra el impacto de cada problema en el proyecto. 

4.3 Pruebas de Integración con Cypress:
![Texto alternativo](imagenes/11.png)

Pruebo el ejemplo: 
![Texto alternativo](imagenes/16.png)
Hago mi primer test: 
![Texto alternativo](imagenes/14.png)
Remplazo con mi url: 
![Texto alternativo](imagenes/15.png)
Vemos que aparece:
![Texto alternativo](imagenes/14.png)
Remplazo con mi url: 
![Texto alternativo](imagenes/20.png)
Ejecutamos en modo headless:
![Texto alternativo](imagenes/19.png)
Modifico prueba para que falle:
![Texto alternativo](imagenes/21.png)
Vemos que falla:
![Texto alternativo](imagenes/22.png)
![Texto alternativo](imagenes/23.png)
4.3.6 Grabar nuestras pruebas para que Cypress genere código automático y genere reportes:
![Texto alternativo](imagenes/25.png)
![Texto alternativo](imagenes/26.png)
4.3.7 Hacemos prueba de editar un empleado
![Texto alternativo](imagenes/27.png)

Desafios:
Integrar en el pipeline SonarCloud para nuestro proyecto Angular, mostrar el resultado obtenido en SonarCloud
![Texto alternativo](imagenes/28.png)
![Texto alternativo](imagenes/29.png)
Usando el link:
![Texto alternativo](imagenes/30.png)
Implementar en Cypress pruebas de integración que incluya los casos desarrollados como pruebas unitarias del front en el TP06.
![Texto alternativo](imagenes/31.png)
![Texto alternativo](imagenes/32.png)
![Texto alternativo](imagenes/33.png)
![Texto alternativo](imagenes/34.png)
![Texto alternativo](imagenes/35.png)
![Texto alternativo](imagenes/36.png)





















