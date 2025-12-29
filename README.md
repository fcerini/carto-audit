# carto-audit
Auditoria de la carga de direcciones y la cartografia

# Resumen

Auditar todas las direcciones que no pudieron ser georreferenciadas en el SAE, inicialmente para la provincia de Cordoba.

## Objetivos

- Mejorar la cartografia
- Auditar la georreferenciacion por parte de los operadores

## Datos de entrada

- Base de datos de vectores de calles de la zona o provincia a auditar.
- Llamadas no georreferenciadas:
    - Relato del operador
    - Direccion cargada por el operador
    - Audio de la llamada, para poder transcribir
    - Ubicacion proporcionada por el sistema ELS de Google.

## Temas

- Importar la base de vectores de calles (cuadras) en un servicio que permita buscar por direcciones o por Lat,Lon.
Se va a usar PostgreSQL con la extensión PostGIS.
- Armar una base de datos asociando las llamadas, los audios y la mejor posicion del ELS.
- Transcribir todos los audios y actualizar la llamada.
- Procesar los relatos y las transcripciones con un modelo que permita extraer la direccion. DireccionIA.
- Calcular la direccion mediante la posicion del ELS (georreferenciacion inversa). DireccionEls
- Georreferenciar
    - Direccion cargada por el operador DireccionOperador
    - Direccion generada por el modelo desde la transcripcion + relato. DireccionIA

Al final se obtienen los siguientes datos para el analisis

- DireccionOperador, LatLon Soflex, LatLon Google
- DireccionIA, LatLon Soflex, LatLon Google
- DireccionEls,  LatLon ELS

Luego de esto generar un proceso que analice cada llamada para determinar los casos posibles

- La direccion existe y deberia agregarse a la cartografia de Soflex.
- La direccion ya existe en la cartografia de Soflex pero el operador no pudo encontrarla.
- La direccion no existe o no puede calcularse desde el relato.
- Otros casos, por ej: LatLon georreferenciada es muy diferente a la del ELS.

### Información Adicional sobre sistema ELS de Google.
Cuando se llama al 911, el celular emite varias posiciones LONGITUD y LATITUD. Se debe quedar con la última porque es la más precisa (el sistema con el tiempo va perfeccionando la ubicación)
