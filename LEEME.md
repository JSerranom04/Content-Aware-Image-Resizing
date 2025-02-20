# Práctica 2: Programación Dinámica - Seam Carving
Asignatura: Algoritmia Básica
Curso 2023/24

## Estructura del Directorio practica2_NIA1_NIA2

### Archivos Principales
- `main.go`: Programa principal que implementa el algoritmo de seam carving
- `pixel.go`: Funciones para el cálculo de energía y brillo de píxeles
- `recurrency.go`: Implementación de la programación dinámica
- `seam.go`: Funciones para encontrar y eliminar costuras
- `image.go`: Funciones de lectura/escritura de imágenes
- `types.go`: Definición de tipos de datos
- `seam_test.go`: Tests unitarios
- `ejecutar.sh`: Script para automatizar la ejecución
- `LEEME.md`: Este archivo

### Requisitos
- Go 1.20 o superior
- Biblioteca estándar de Go (no requiere dependencias externas)

### Compilación

bash
go build -o costuras .go

### Ejecución

bash
./costuras <número_costuras> <archivo_entrada> <directorio_resultados>

Ejemplo:

bash
./costuras 50 ./pruebas/imagen.png ./resultados


### Descripción del Algoritmo
El programa implementa el algoritmo de seam carving para reducir el ancho de una imagen mediante la eliminación de costuras verticales de mínima energía. El proceso consta de tres pasos principales:

1. Cálculo de energía: Para cada píxel se calcula su energía usando el gradiente en X e Y
2. Programación dinámica: Se encuentra la costura de mínima energía usando la ecuación de recurrencia
3. Eliminación: Se remueve la costura encontrada y se repite el proceso

### Estructura de Datos
- `MatrixComponent`: Estructura que almacena los valores RGBA y brillo de cada píxel
- Matrices bidimensionales para representar la imagen y las energías

### Tests
Para ejecutar los tests:

bash
go test -v




### Imágenes de Prueba
El directorio `pruebas/` contiene imágenes de ejemplo:
- boat.png: Imagen de un barco
- water.png: Imagen de agua con áreas homogéneas
- rocket.png: Imagen de un cohete
- imagen.png: Imagen de prueba general

### Notas Adicionales
- El programa acepta imágenes en formato PNG
- Los resultados se guardan en el directorio especificado
- Se muestra el progreso durante la ejecución
- El tiempo de ejecución depende del tamaño de la imagen y número de costuras
- La implementación incluye optimizaciones para recalcular solo las áreas afectadas por cada costura

### Ejemplo de Rendimiento
En una imagen de 1466x1220 píxeles, reduciendo 350 costuras:
- Dimensiones iniciales: 1466x1220
- Dimensiones finales: 1466x870
- Tiempo aproximado: 2-3 minutos

Autores: Victor Orrios Baron y Juan José Serrano Mora