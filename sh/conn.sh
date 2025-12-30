#!/bin/bash

sshpass -f ./sh/pass.txt ssh -t 'soflex@172.24.134.60' -p 22 "cd /home/soflex/carto-audit/; bash"

