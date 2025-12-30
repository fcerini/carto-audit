#!/bin/bash

sh/build.sh

echo "copiando al server..."
sshpass -f ./sh/pass.txt scp carto-audit soflex@172.24.134.60:/home/soflex/carto-audit/

