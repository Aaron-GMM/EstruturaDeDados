---
nomeAluno: "Seu nome"
matricula: "Digite sua matrícula aqui por obséquio"
---

# Bom trabalho

## TKO

```bash
## Criando seu novo repositório
tko init --remote [fup|ed|poo]

## Interagindo com seu repositório
tko open <pasta>

## Rodando um código diretamente
tko test <arquivo_codigo>  # usando interface default
tko run  <arquivo_codigo>  # usando interface curses

## Atualizando o tko
pipx upgrade tko
```

## Comandos do Git

- `git add <arquivo>`: Adiciona um arquivo ao stage
- `git add .`: Adiciona todos os arquivos ao stage
- `git commit -m "mensagem"`: Cria um commit com os arquivos no stage
- `git push`: Envia os commits para o repositório remoto
- `git pull`: Atualiza o repositório local com as mudanças do repositório remoto
- `git status`: Mostra o estado atual do repositório
- `git log`: Mostra o histórico de commits
- `git clone <url>`: Clona um repositório remoto para uma pasta local

## Resolvendo Conflitos

- Antes de trabalhar num respositório, lembre de fazer o pull.
- Se quando for fazer o push, der erro, tente as seguintes opções.
  - git pull --ff-only
  - git rebase
  - git pull --no-rebase
- Se não der certo, cole o erro no chatgpt e siga as instruções.
- Se não der certo, mande o erro no grupo do telegram.

## Lembre-se

- Sua consciência é a melhor ferramenta contra o plágio.
