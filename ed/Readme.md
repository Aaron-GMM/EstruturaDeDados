# Repositório de exercícios para Estrutura de Dados

Esse é um repositório que integra o conteúdo das atividades de Estrutura de Dados (ES) para os cursos:

- Ciência da Computação
- Engenharia de Computação
- Sistemas de Informação
- Engenharia de Software

Esse curso apresenta os fundamentos teóricos e atividades práticas para a compreensão de como funcionam as estruturas de dados básicas e como utilizá-las para resolver problemas computacionais.

A linguagem escolhida para o curso é o C++, pois permite uma compreensão melhor das estruturas de baixo nível, como gestão da memória e um nível de controle que permite controlar e compreender a complexidade dos algoritmos.

O prérequisito para o curso é ter conhecimento básico de programação e lógica de programação, em linguagem C, C++, ou Python.

Cada grande bloco é composto de várias grupos, chamados de missões. Uma missão representa a aquisição de um conhecimento ou habilidade específica. Cada missão é composta de várias atividades, chamadas de tarefas. Uma tarefa pode ser uma leitura, vídeo, exercício teórico ou prática. Algumas atividades de código estão no modelo de maratona, com casos de testes que avaliam se o código resolve o problema proposto.

Para baixar as atividades, marcar as atividades feitas e executar os códigos de teste, você pode utilizar o aplicativo tko.

O próprio aplicativo vai orientar a sequência sugerida na qual as atividades devem ser realizadas e o mínimo a ser obtido para avançar para as próximas missões.

Bom trabalho.

## Introdução

Significado das tags nas questões:

| marcador  | significado
| --------- | -----------
| test      | Problema de código com testes
| open      | Problema de programação abertos sem testes
| info      | Conteúdo informativo de leitura ou vídeo
| task      | Outras atividades

### Aquecimento com FUP

- [ ] `@pintando :test`[Pintando a casa](https://github.com/qxcodefup/arcade/blob/master/base/pintando/Readme.md)
- [ ] `@mumia    :test`[Criança, jovem, adulto](https://github.com/qxcodefup/arcade/blob/master/base/mumia/Readme.md)
- [ ] `@fuga     :test`[Fuga de helicóptero](https://github.com/qxcodefup/arcade/blob/master/base/fuga/Readme.md)
- [ ] `@pedra    :test`[Pedra na lua](https://github.com/qxcodefup/arcade/blob/master/base/pedra/Readme.md)
- [ ] `@casais   :test`[Casais na arca](https://github.com/qxcodefup/arcade/blob/master/base/casais/Readme.md)
- [ ] `@baruel   :test`[Figurinhas repetidas](https://github.com/qxcodefup/arcade/blob/master/base/baruel/Readme.md)
- [ ] `@gomos    :test`[Gomos da cobra](https://github.com/qxcodefup/arcade/blob/master/base/gomos/Readme.md)

### Conceitos Básicos de C++

Nesse primeiro bloco, vamos aprender ou relembrar os conceitos básicos da linguagem C++.

- [ ] `#cpp_intro   :info`[Objetivos Introdução](https://github.com/qxcodeed/arcade/blob/master/wiki/video_intro.md)
- [ ] `#cpp_pointer :info`[Vídeo - Qual a diferença entre ponteiros e referências](https://youtu.be/uz_sTcNdguY?si=YznFcnV6sisot_Sc)
- [ ] `#cpp_oop     :info`[Vídeo - Introdução a OOP in C++](https://youtu.be/w7F587dNwqA?si=2UMhfS9_DcAh-gjF)
- [ ] `#cpp_acesso  :info`[Vídeo - Modificadores de Acesso com exemplos](https://youtu.be/n1RfuPbzG-M?si=kwaSjo9Ng1g4waOK)

### Aprendendo um pouco sobre OOP em C++ <!-- @tad r:intro -->

A implementação de Orientação a Objetos em C++ tem várias peculiaridades como a forma de fazer o "toString" ou a sobrecarga de operadores e métodos.

- [ ] `#cpp_cpref  :info`[Como fazer sobrecarga de cópia e referencia](https://github.com/qxcodeed/arcade/blob/master/wiki/tad/sobrecarga.md)
- [ ] `#cpp_oper   :info`[Sobrecarga de operadores de classes](https://github.com/qxcodeed/arcade/blob/master/wiki/tad/sobrecarga_operadores.md)
- [ ] `L2 @ponto   :test`[Implemente a classe ponto](https://github.com/qxcodeed/arcade/blob/master/base/ponto/Readme.md)
- [ ] `L2 @fracao  :test`[Implemente a classe fração reduzida](https://github.com/qxcodeed/arcade/blob/master/base/fracao/Readme.md)

### Análise de Algoritmos <!-- @analise r:intro -->

A análise de algoritmos é essencial para entender e avaliar a eficiência das soluções computacionais. Ela permite medir o desempenho dos algoritmos em termos de tempo e espaço, facilitando a comparação entre diferentes abordagens para a mesma tarefa. Compreender a complexidade computacional ajuda a tomar decisões informadas sobre qual algoritmo utilizar, otimizando recursos e melhorando a performance do sistema.

Estudar análise de algoritmos também revela os limites e aplicabilidades de cada solução, evitando o uso inadequado e promovendo a criação de algoritmos mais eficientes e inovadores. Esse conhecimento é crucial não apenas para a resolução de problemas reais, mas também para o desenvolvimento de software otimizado e para a preparação em entrevistas técnicas, onde a análise de algoritmos é frequentemente testada.

- [ ] `#analise    :info`[Análise de Algoritmos](https://github.com/qxcodeed/arcade/blob/master/wiki/analise/analise.md)

## Recursão

Nesse bloco, abordaremos o conceito de recursividade, uma técnica de programação que permite que uma função chame a si mesma para resolver problemas de forma mais elegante e eficiente. Entender recursividade pode ajudar a simplificar problemas complicados e tornar a programação mais fácil e eficaz.

### Recursão: Estudo <!-- @rec_intro r:intro -->

- [ ] `#cpp_rec_view :info`[Vídeo - Introdução à recursão](https://youtu.be/nEohgkZkm_c?si=XIDgq6jPKlAd9gbz)
- [ ] `#cpp_rec_read :info`[Recursividade](https://github.com/qxcodeed/arcade/blob/master/wiki/recursao/recursao.md)

### Recursão: Treinando conceitos I<!-- @con_rec r:rec_intro -->

- [ ] `#cpp_intro       :info` [L0 - @binomiais - Calculando coeficientes](https://github.com/qxcodeed/arcade/blob/master/base/binomiais/Readme.md)
- [ ] `@escada       L0 :test` [De quantas maneiras podemos subir](https://github.com/qxcodeed/arcade/blob/master/base/escada/Readme.md)
- [ ] `@diagonal     L1 :test` [Imprimindo um char por linha](https://github.com/qxcodeed/arcade/blob/master/base/diagonal/Readme.md)
- [ ] `@padroes1     L1 :test` [Descobrindo quantos bloquinhos](https://github.com/qxcodeed/arcade/blob/master/base/padroes1/Readme.md)
- [ ] `@sufixos      L1 :test` [L1 - @sufixos - Imprimindo palavra em pedaços](https://github.com/qxcodeed/arcade/blob/master/base/sufixos/Readme.md)
- [ ] `@empilhamento L1 :test` [Empilhamento de divisões](https://github.com/qxcodeed/arcade/blob/master/base/empilhamento/Readme.md)
- [ ] `@mdc          L2 :test` [Implemente o mdc recursivo](https://github.com/qxcodeed/arcade/blob/master/base/mdc/Readme.md)

### Recursão: Treinando conceitos II <!-- @rec_basic r:con_rec -->

- [ ] `@padroes2  L2 :test`[Descobrindo quantas pecinhas](https://github.com/qxcodeed/arcade/blob/master/base/padroes2/Readme.md)
- [ ] `@santo JM  L2 :test`[O milagre de ficar pobre](https://github.com/qxcodeed/arcade/blob/master/base/santo/Readme.md)
- [ ] `@fibonacci L2 :test`[Procriando coelhos](https://github.com/qxcodeed/arcade/blob/master/base/fibonacci/Readme.md)
- [ ] `@quadrado  L2 :test`[Aragão calculando n ao quadrado](https://github.com/qxcodeed/arcade/blob/master/base/quadrado/Readme.md)
- [ ] `@bitcoin   L2 :test`[Investindo em ativos digitais](https://github.com/qxcodeed/arcade/blob/master/base/bitcoin/Readme.md)

### Recursão: Desafios <!-- @rec_desafios r:rec_basic t:7 -->

- [ ] `@triangulo  L1 :test`[Imprimindo ao contrário](https://github.com/qxcodeed/arcade/blob/master/base/triangulo/Readme.md)
- [ ] `@fibonacci2 L3 :test`[Fazendo filhos e morrendo](https://github.com/qxcodeed/arcade/blob/master/base/fibonacci2/Readme.md)
- [ ] `@padroes3   L3 :test`[Contando pontos](https://github.com/qxcodeed/arcade/blob/master/base/padroes3/Readme.md)
- [ ] `@queimada   L2 :test`[Tocando fogo na floresta](https://github.com/qxcodeed/arcade/blob/master/base/queimada/Readme.md)
- [ ] `@treinando  L2 :test`[Treinando recursão com vetores](https://github.com/qxcodeed/arcade/blob/master/base/treinando/Readme.md)

### Recursão: Exercícios com matemática <!-- @rec_primos r:rec_basic -->

- [ ] `@eh_primo  L1 :test`[Retorna se é ou não é](https://github.com/qxcodeed/arcade/blob/master/base/eh_primo/Readme.md)
- [ ] `@enesimo   L2 :test`[Gere o enesimo primo](https://github.com/qxcodeed/arcade/blob/master/base/enesimo/Readme.md)
- [ ] `@priminhos L2 :test`[Carrega a lista de primos](https://github.com/qxcodeed/arcade/blob/master/base/priminhos/Readme.md)

### Recursão: Backtracking <!-- @backtracking r:rec_desafios -->

- [ ] `@labirinto L2  :test`[Escapando do labirinto](https://github.com/qxcodeed/arcade/blob/master/base/labirinto/Readme.md)
- [ ] `@distancia L2  :test`[Colocando os números no lugar certo](https://github.com/qxcodeed/arcade/blob/master/base/distancia/Readme.md)
- [ ] `@sudoku    L3  :test`[L3 - @sudoku](https://github.com/qxcodeed/arcade/blob/master/base/sudoku/Readme.md)

## Ordenação

Nesse bloco, abordaremos os algoritmos de busca e ordenação, duas técnicas essenciais na programação. Os algoritmos de busca são usados para encontrar elementos específicos dentro de uma estrutura de dados (localizar um nome específico em uma lista de contatos, por exemplo), enquanto os algoritmos de ordenação organizam esses elementos em uma sequência desejada (como organizar uma lista de livros por título em ordem alfabética). Compreender essas técnicas é crucial para manipular e acessar dados de forma eficiente em seus programas.

### Ordenação: Estudo <!-- @sort_intro r:analise -->

- [ ] `#alg_ord_video   :info` [Vídeo - implementação de todos os algoritmos de ordenação](https://youtu.be/k6nODikH_D8?si=tpbDRg139HvdzSIS)
- [ ] `#alg_ord_leitura :info` [Algoritmos de Busca](https://github.com/qxcodeed/arcade/blob/master/wiki/busca_ordenacao/busca.md)
- [ ] `#alg_ord_simples :info`[Algoritmos de Ordenação](https://github.com/qxcodeed/arcade/blob/master/wiki/busca_ordenacao/ordenacao.md)
- [ ] `#alg_ord_recurs  :info`[Algoritmos de Ordenação Recursiva: Merge Sort e Quick Sort](https://github.com/qxcodeed/arcade/blob/master/wiki/ordenacao_recursiva/ordenacao_rec.md)

### Ordenação: Construindo algoritmos <!-- @sort_build r:sort_intro -->

- [ ] `@tapioca  L2 :test`[Tapioca Sort](https://github.com/qxcodeed/arcade/blob/master/base/tapioca/Readme.md)
- [ ] `@mergevet L2 :test`[Criando um vetor simples e fazendo um merge](https://github.com/qxcodeed/arcade/blob/master/base/mergevet/Readme.md)

### Ordenação: Resolvendo exercícios simples <!-- @sort_basic r:sort_intro -->

- [ ] `@vacina  L0 :test`[Otimizando vacinas](https://github.com/qxcodeed/arcade/blob/master/base/vacina/Readme.md)
- [ ] `@indices L1 :test`[Índices originais](https://github.com/qxcodeed/arcade/blob/master/base/indices/Readme.md)

## Vetores

Os vetores são estruturas de dados fundamentais que permitem acesso direto e rápido a qualquer elemento, tornando operações de leitura e escrita extremamente eficientes. Compreender como os vetores gerenciam a memória e os custos associados ao redimensionamento é essencial para tomar decisões informadas sobre seu uso. Além disso, os vetores são amplamente utilizados em diversas aplicações e são nativamente suportados por muitas linguagens de programação, o que facilita sua integração com outras partes do sistema.

### Vetores: Construindo <!-- @array_intro r:intro q:4 -->

- [ ] `@ponteiros L2 :test`[Implementando vector com ponteiros](https://github.com/qxcodeed/arcade/blob/master/base/ponteiros/Readme.md)
- [ ] `@classes   L2 :test`[Implementando vector com @classes](https://github.com/qxcodeed/arcade/blob/master/base/classes/Readme.md)
- [ ] `@inteiros  L0 :test`[Lista Sequencial de Inteiros](https://github.com/qxcodeed/arcade/blob/master/base/inteiros/Readme.md)
- [ ] `@ordenada  L0 :test`[Lista Sequencial Ordenada](https://github.com/qxcodeed/arcade/blob/master/base/ordenada/Readme.md)
- [ ] `@circular  L0 :test`[Lista Sequencial Circular](https://github.com/qxcodeed/arcade/blob/master/base/circular/Readme.md)

### Vetores: Resolvendo Problemas Simples <!-- @array_basic r:array_intro -->

Treinando o uso da classe vector e seus métodos na resolução de problemas básicos.

- [ ] `@minmax     L0 :test`[MinMax](https://github.com/qxcodeed/arcade/blob/master/base/minmax/Readme.md)
- [ ] `@palindromo L0 :test`[Palíndromo](https://github.com/qxcodeed/arcade/blob/master/base/palindromo/Readme.md)
- [ ] `@repetido   L0 :test`[Verifique se algum elemento do vetor se repete](https://github.com/qxcodeed/arcade/blob/master/base/repetido/Readme.md)
- [ ] `@buscas     L1 :test`[Buscas de strings](https://github.com/qxcodeed/arcade/blob/master/base/buscas/Readme.md)
- [ ] `@rotacoes   L1 :test`[Rotacoes em uma lista](https://github.com/qxcodeed/arcade/blob/master/base/rotacoes/Readme.md)
- [ ] `@princesa   L2 :test`[Princesa v1](https://github.com/qxcodeed/arcade/blob/master/base/princesa/Readme.md)
- [ ] `@alternando L2 :test`[Princesa v2](https://github.com/qxcodeed/arcade/blob/master/base/alternando/Readme.md)

### Vetores: Otimizando problemas com maps e sets <!-- @array_struct r:array_basic -->

Os exercícios desse bloco podem ser resolvidos com vectors, porém a maioria das atividades precisa de atenção para a manipulação eficiente dos elementos. Estruturas como map, set, list e stack podem ser mais eficientes para resolver os problemas propostos.

- [ ] `@apagando    L2 :test`[Apagando elementos de forma eficiente](https://github.com/qxcodeed/arcade/blob/master/base/apagando/Readme.md)
- [ ] `@manipulacao L2 :test`[Utilizando a biblioteca algorithm](https://github.com/qxcodepoo/arcade/blob/master/base/manipulacao/Readme.md)
- [ ] `@especiais   L2 :test`[Utilizando map e set](https://github.com/qxcodepoo/arcade/blob/master/base/especiais/Readme.md)

## Listas Encadeadas

Agora estudaremos as listas encadeadas, estruturas de dados flexíveis e eficientes que podem crescer e diminuir de tamanho facilmente, adaptando-se às necessidades de armazenamento dos elementos. Diferentemente dos vetores, que têm um tamanho fixo ou precisam ser redimensionados de forma custosa, as listas encadeadas não utilizam um espaço de armazenamento contínuo na memória. Elas são representadas por uma sequência de elementos ligados entre si, podendo ser simplesmente ou duplamente encadeadas, lineares ou circulares. Além disso, as listas encadeadas servem como base para a implementação de diversas outras estruturas de dados, como pilhas, filas e grafos.

### Lista: Construindo <!-- @list_build r:array_intro -->

- [ ] `#lista_video   :info`[Vídeo - Implementando Lista Ligada Dupla](https://youtu.be/gA4YAtXnnKw?si=cC2aea_5N4cFUxJY)
- [ ] `@ligada L2     :test`[Lista simples sem cabeça e rabo](https://github.com/qxcodeed/arcade/blob/master/base/ligada/Readme.md)
- [ ] `@lista_simples :test`[Lista Simplesmente Encadeada - CPP](https://github.com/qxcodeed/arcade/blob/master/base/lista_simples/Readme.md)

### Lista: Resolvendo Problemas Simples <!-- @list_basic r:list_build -->

- [ ] `@princesa-lista   :test`[Princesa com lista ligada](https://github.com/qxcodeed/arcade/blob/master/base/princesa-lista/Readme.md)
- [ ] `@princesa-lista-2 :test`[Princesa2 com lista ligada](https://github.com/qxcodeed/arcade/blob/master/base/princesa-lista-2/Readme.md)

### Lista: Resolvendo Problemas Avançados <!-- @list_advanced r:list_basic -->

- [ ] `@banco L3    :test`[Gerenciando a fila do banco](https://github.com/qxcodeed/arcade/blob/master/base/banco/Readme.md)
- [ ] `@editor L2   :test`[Editor de texto](https://github.com/qxcodeed/arcade/blob/master/base/editor/Readme.md)
- [ ] `@controlz L3 :test`[Editor com Control Z](https://github.com/qxcodeed/arcade/blob/master/base/controlz/Readme.md)

### Lista: Variações de Listas Encadeadas

- [ ] `@cauda       :test`[Lista Simplesmente Encadeada com Cauda](https://github.com/qxcodeed/arcade/blob/master/base/cauda/Readme.md)
- [ ] `@duplamente  :test`[Lista Duplamente Encadeada](https://github.com/qxcodeed/arcade/blob/master/base/duplamente/Readme.md)
- [ ] `@invertendo  :test`[Inversão usando next e previous](https://github.com/qxcodeed/arcade/blob/master/base/invertendo/Readme.md)
- [ ] `@invertendo2 :test`[Pilha - invertendo com pilhas auxiliares](https://github.com/qxcodeed/arcade/blob/master/base/invertendo2/Readme.md)
- [ ] `@concatenar  :test`[Concatenar duas listas duplamente encadeadas](https://github.com/qxcodeed/arcade/blob/master/base/concatenar/Readme.md)
- [ ] `@converte    :test`[Converte simples em dupla](https://github.com/qxcodeed/arcade/blob/master/base/converte/Readme.md)

## Pilha e Fila

A próxima estrutura estudada é a Pilha, uma estrutura de dados eficiente e simples, baseada na política LIFO (Last-In First-Out), onde o último elemento inserido é o primeiro a ser removido. As operações básicas dessa estrutura incluem "empilha", que adiciona um elemento no topo, e "desempilha", que remove o elemento do topo. Essa estrutura é vantajosa em situações que exigem um controle de acesso restrito, como em editores de texto para desfazer ações ou no gerenciamento de chamadas de função em programas. Diferente de vetores, que permitem acesso direto a qualquer elemento, a Pilha limita as operações ao topo, o que a torna mais simples e eficiente em certos contextos. Comparada às listas encadeadas, a Pilha é mais restritiva, mas sua simplicidade pode ser benéfica para o desempenho de operações específicas.

### Pilha: Construindo <!-- @pilha_build r:list_build -->

- [ ] `@pilha-encadeada  :test`[Construindo Pilha com Lista Encadeada - CPP](https://github.com/qxcodeed/arcade/blob/master/base/pilha-encadeada/Readme.md)
- [ ] `@pilha-sequencial :test`[Construindo uma Pilha com Lista Sequencial CPP](https://github.com/qxcodeed/arcade/blob/master/base/pilha-sequencial/Readme.md)
- [ ] `@pilha-dupla      :test`[Pilha Dupla - CPP](https://github.com/qxcodeed/arcade/blob/master/base/pilha-dupla/Readme.md)

### Pilha - Resolvendo problemas <!-- @pilha r:pilha_build -->

- [ ] `@queimando L2  :test`[Queimando árvores usando pilha](https://github.com/qxcodeed/arcade/blob/master/base/queimando/Readme.md)
- [ ] `@pilhado   L2  :test`[Fugindo do labirinto](https://github.com/qxcodeed/arcade/blob/master/base/pilhado/Readme.md)
- [ ] `@polonesa  L1  :test`[Transforme da notação padrão para polonesa reversa](https://github.com/qxcodeed/arcade/blob/master/base/polonesa/Readme.md)
- [ ] `@minmax-pilha  :test`[MinMax Pilha CPP](https://github.com/qxcodeed/arcade/blob/master/base/minmax-pilha/Readme.md)
- [ ] `@balanceada L2 :test`[Fórmula balanceada](https://github.com/qxcodeed/arcade/blob/master/base/balanceada/Readme.md)

## Fila

Agora estudaremos a estrutura de dados Fila, uma lista linear que adota a política FIFO (First-In First-Out) para a manipulação de elementos. Isso significa que o primeiro elemento inserido é o primeiro a ser removido, como uma fila de atendimento em um banco, onde as pessoas são atendidas na ordem em que chegaram. As operações principais dessa estrutura são "enqueue", que adiciona um elemento no final da fila, e "dequeue", que remove o elemento do início. Isso é essencial em aplicações como filas de impressão ou sistemas que gerenciam tarefas em sequência, garantindo que a ordem de chegada seja mantida. Em comparação com vetores, que oferecem acesso aleatório, a Fila é mais limitada, mas sua estrutura é ideal para o processamento sequencial. Embora listas encadeadas possam ser usadas para implementar Filas, a estrutura da Fila se destaca por sua organização que prioriza a ordem cronológica, essencial em sistemas de tempo real.

### Fila: Construindo <!-- @fila_build r:list_build -->

- [ ] `@fila            :test`[Fila como duas pilhas](https://github.com/qxcodeed/arcade/blob/master/base/fila/Readme.md)
- [ ] `@fila-encadeada  :test`[Fila encadeada CPP](https://github.com/qxcodeed/arcade/blob/master/base/fila-encadeada/Readme.md)
- [ ] `@fila-sequencial :test`[Fila sequencial CPP](https://github.com/qxcodeed/arcade/blob/master/base/fila-sequencial/Readme.md)
- [ ] `@fila-empilhada  :test`[Utilizando Pilhas para a Criação de uma Fila CPP](https://github.com/qxcodeed/arcade/blob/master/base/fila-empilhada/Readme.md)

### Fila - Resolvendo problemas <!-- @fila_solver r:fila_build -->

- [ ] `@chaves   L1 :test`[Chaves de disputa](https://github.com/qxcodeed/arcade/blob/master/base/chaves/Readme.md)
- [ ] `@caminhao L2 :test`[Caminhao abastecendo](https://github.com/qxcodeed/arcade/blob/master/base/caminhao/Readme.md)
- [ ] `@caminho  L3 :test`[Menor caminho entre dois pontos](https://github.com/qxcodeed/arcade/blob/master/base/caminho/Readme.md)

## Árvores <!-- group -->

Agora vamos estudar as árvores, estruturas de dados que são especialmente adequadas para a representação de hierarquias. Diferente de listas e vetores, que são lineares, uma árvore é uma estrutura de dados bidimensional não-linear, permitindo organizar informações de maneira mais flexível e eficiente para cenários que exigem relações hierárquicas. Um tipo comum de árvore é a árvore binária, na qual cada nó pode ter no máximo dois filhos. As árvores são fundamentais em várias aplicações de computação, como na organização de expressões matemáticas, nos sistemas de arquivos de computadores, e em algoritmos de busca e ordenação. Além disso, árvores servem como base para outras estruturas de dados importantes. Para entender o funcionamento dessa estrutura de dados, exploraremos sua:

- Construção

### Árvore - Construindo <!-- @tree r:list_build -->

- [ ] `@serial      L2 :test`[Lendo uma árvore a partir de um serial](https://github.com/qxcodeed/arcade/blob/master/base/serial/Readme.md)
- [ ] `@formatada   L2 :test`[Mostrando a árvore formatada](https://github.com/qxcodeed/arcade/blob/master/base/formatada/Readme.md)
- [ ] `@soma        L2 :test`[Soma dos nós da árvore](https://github.com/qxcodeed/arcade/blob/master/base/soma/Readme.md)
- [ ] `@altura      L2 :test`[Altura e profundidade da árvore](https://github.com/qxcodeed/arcade/blob/master/base/altura/Readme.md)
- [ ] `@arvore-path L2 :test`[Caminho até o nó](https://github.com/qxcodeed/arcade/blob/master/base/arvore-path/Readme.md)
- [ ] `@dividido    L2 :test`[Percorrer in-order](https://github.com/qxcodeed/arcade/blob/master/base/dividido/Readme.md)
- [ ] `@arvore      L2 :test`[Construir árvore binaria CPP](https://github.com/qxcodeed/arcade/blob/master/base/arvore/Readme.md)
- [ ] `@incompleta  L3 :test`[Arvore incompleta](https://github.com/qxcodeed/arcade/blob/master/base/incompleta/Readme.md)
- [ ] `@abb         L3 :test`[Árvore Binária de Busca](https://github.com/qxcodeed/arcade/blob/master/base/abb/Readme.md)