#!/bin/bash

# Verificar argumentos
if [ $# -ne 3 ]; then
    echo "Uso: $0 <número_costuras> <archivo_entrada> <directorio_resultados>"
    exit 1
fi

# Crear directorio de resultados si no existe
mkdir -p "$3"

# Compilar el programa
go build -o costuras *.go

# Ejecutar el programa
./costuras "$1" "$2" "$3"

# Medir tiempo de ejecución
time ./costuras "$1" "$2" "$3"

# Verificar resultado
if [ $? -eq 0 ]; then
    echo "Proceso completado exitosamente"
    echo "Resultados guardados en: $3"
else
    echo "Error durante la ejecución"
    exit 1
fi