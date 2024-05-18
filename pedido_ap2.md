# Pedido da AP2

## Requisitos

Seu time foi solicitado a evoluir a aplicação do McRonald's, incluindo algumas novas funcionalidades para o sistema. O sistema deve ser construído como uma API REST em Golang.

## Novas funcionalidades

- Transforme a sua lista de produtos em uma árvore binária de busca. Essa árvore deve estar ordenada conforme o nome do produto. Todas as operações relativas ao produto devem refletir essa nova estrutura;
- Crie um novo parâmetro opcional para a requisição "GET" dos pedidos. Neste parâmetro, o usuário pode selecionar o retorno ordenado dos pedidos pelo seu valor total, de forma crescente. Ou seja, a lista retornada deve conter primeiro os pedidos de preço mais baixo. Caso este parâmetro não seja passado, a lista deve ser retornada conforme a sua ordem original (pelo id). Os dois valores possíveis para este parâmetro são "bubblesort" e "quicksort", em que a lista é ordenada conforme o algoritmo indicado;
- Modifique o cálculo do valor total do pedido, incluindo um desconto de 10% caso o valor total exceda R$100. Lembre-se de arredondar o valor para duas casas decimais;
- Atualize todas os "print" no terminal do programa, de forma a exibir as informações de dia e hora do sistema antes da mensagem. Utilize o formato `DD/MM/AAAA HH:MM:SS <mensagem>`;
- Crie um novo parâmetro para a rota `/abrir`, chamado `intervalo`, que substitui o intervalo padrão de 30 segundos entre a expedição de cada pedido;
- Desenvolva um arquivo `readme.md`, explicando o funcionamento da API.

## Bônus

Duas novas funcionalidades foram solicitadas, caso a equipe consiga desenvolver:

- Inclua um novo algoritmo de ordenação à sua escolha (selection sort, insertion sort ou merge sort) na requisição "GET" dos pedidos;
- Inclua duas novas métricas na lista de métricas do programa: ticket médio (total faturado até o momento, dividido pelo número de pedidos) e tempo de funcionamento do programa (número de segundos desde que o sistema entrou em operação).

## Critérios de aceitação

Para o sistema ser aceito pelos clientes, as seguintes condições devem ser satisfeitas:

* O sistema precisa ser desenvolvido na linguagem Go;
* O projeto deve ser entregue no seu formato de código fonte;
* O sistema deve ser entregue em um repositório **privado** no GitHub, adicionando um dos sócios (@victor0machado) como colaborador;
* O repositório possui um arquivo readme.md, descrevendo as pessoas da equipe e como utilizar o programa (comandos aceitos, formatos das estruturas, exemplos de respostas, etc.).

O não atendimento de algum dos critérios de aceitação pode acarretar em penalização na solução.

## Critérios de avaliação

O sistema será avaliado pelos clientes segundo critérios objetivos e subjetivos, quais sejam:

* Testes funcionais: os testes realizados pelos clientes não apresentam falhas. Ou seja, os requisitos funcionais detalhados acima estão sendo atendidos completamente;
* Avaliação estática do código: o código é eficiente e bem escrito, está organizado em pacotes com responsabilidades distintas, não apresenta redundâncias e atende às boas práticas de programação;
* Tratamento de erros: o projeto apresenta tratamento de erros adequado, analisando casos como produto com dados incompletos ou inválidos, pedidos incompletos, etc.

Após a entrega e avaliação, as equipes receberão um barema contendo o descritivo da avaliação, bem como pesos e detalhamento de cada critério de avaliação. Até 1 ponto adicional será concedido para cada funcionalidade bônus desenvolvida pela equipe, conforme a qualidade da entrega. A nota do time não pode passar de 10.

## Demais informações

* Prazo do projeto, até 16/06/2024, sem possibilidade de adiamento do prazo;
* Submissão do link do GitHub atráves do e-mail victor0machado@gmail.com (não enviar pelo Estudante ou pelo e-mail do Ibmec!);
* Apenas uma pessoa de cada equipe precisa enviar o link;
* As equipes devem ter no mínimo 3 e no máximo 5 pessoas.
