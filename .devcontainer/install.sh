#!/bin/bash


# Instala pacotes Python
pipx install tko

# Instala pacotes Node.js globais
npm install -g typescript esbuild

# Instala pacotes locais
cd /workspaces
npm install --save-dev @types/node readline-sync

source .bashrc || true
