#!/bin/bash

# Verificar argumentos
if [ $# -ne 3 ]; then
    echo "Uso: $0 <número_costuras> <archivo_entrada> <directorio_resultados>"
    exit 1
fi

# Crear directorio de resultados si no existe
mkdir -p "$3"

# Compilar el programa
go build -o costuras *.go || exit

# Captura el tiempo inicial en milisegundos
start=$(date +%s%3N) 

# Ejecutar el programa
./costuras "$1" "$2" "$3"

# Captura el tiempo final en milisegundos
end=$(date +%s%3N)   

# Calcula el tiempo medido en ms
exe_time="$((end - start)) ms"

echo Ejecutado en "$exe_time"

# Verificar resultado
if [ $? -eq 0 ]; then
    echo "Proceso completado exitosamente"
    echo "Resultados guardados en: $3"
else
    echo "Error durante la ejecución"
    exit 1
fi