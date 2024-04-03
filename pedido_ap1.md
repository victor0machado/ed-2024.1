# Pedido da AP1

## Contexto

O McRonald's é um food truck em franco crescimento no último ano. Foi fundado por dois amigos apaixonados por boa comida e interessados em trazer uma nova opção de gastronomia no bairro em que vivem.

Atualmente, cada pedido do McRonald's é registrado manualmente e é passado para a cozinha. Após o preparo do pedido, o chef faz a expedição e distribui para os garçons entregarem na mesa ou para os motoboys fazerem as entregas. Não há um cardápio digital, todo o processo é feito no papel. Pedidos de delivery possuem uma taxa adicional fixa igual a R$10,00.

Esse processo está gerando diversos problemas. O primeiro é que diversos pedidos estão sendo registrados incorretamente, já que os atendentes confundem os nomes dos produtos com frequência. Em segundo lugar, os pedidos não estão sendo produzidos em ordem, então alguns clientes ficam esperando um tempo demasiado longo. Por último, os sócios não estão conseguindo levantar qual o faturamento real do empreendimento.

## Requisitos

Identificando essas falhas no processo, sua equipe foi contratada para elaborar um MVP para o back-end de uma solução de pedidos eletrônicos para o McRonald's. O sistema deve ser construído como uma API REST em Golang.

### Entidades do projeto

As seguintes estruturas são necessárias para o desenvolvimento do projeto:

* Produto: Cada produto deve ser armazenado com um identificador único, representado por um número inteiro, começando no `1`. Caso um produto seja removido, o seu identificador não deve ser substituído. Ou seja, se o usuário cadastrou 20 produtos (`1`, `2`, ... `20`), e excluiu o produto `20`, o próximo produto a ser cadastrado terá o identificador `21`. Cada produto deve conter as seguintes informações (além do identificador):
    * Nome: nome curto;
    * Descrição: descrição completa do item;
    * Valor: o preço do produto, em reais;
* Pedido: O pedido deve incluir os seguintes campos:
    * Identificador: número único, iniciado com `1`;
    * Delivery: campo booleano, valendo `false` caso o pedido seja para consumo ou retirada no local, e `true` caso seja para entrega;
    * Produtos: lista com produtos disponíveis no cardápio. Cada produto pode ter uma quantidade de itens diferente;
    * Valor total: valor do pedido, incluindo os produtos e taxa de entrega (se houver);
* Lista de produtos: Uma lista com alocação sequencial (`slice`) contendo todos os produtos cadastrados no sistema;
* Fila de pedidos: Uma fila com alocação sequencial (`slice`) contendo os pedidos cadastrados e ativos;
* Métricas do sistema: Uma estrutura que armazena o número total de produtos cadastrados, o número de pedidos já encerrados, o número de pedidos em andamento e o faturamento total até o momento.

### Operações

Para a lista de produtos, sua equipe deve desenvolver algoritmos para:

* Adicionar um novo produto, gerando um identificador único;
* Remover um produto existente, a partir do seu identificador;
* Buscar um produto pelo identificador;
* Listar todos os produtos da lista.

Para a fila de pedidos, sua equipe deve desenvolver algoritmos para:

* Incluir um novo pedido;
* Exibir os pedidos em aberto;
* Expedir um pedido.

### Rotas da API

O MVP da API deve conter as seguintes rotas:

| **Rota**  | **Método** | **Descrição**                                                                                                 |
|-----------|------------|---------------------------------------------------------------------------------------------------------------|
| /produto  |    POST    | Inclui um novo produto na lista, a partir de um JSON contendo as informações necessárias.                     |
| /produto  |     GET    | Obtém informações de um produto, em formato JSON, a partir da informação do ID único, passado via formulário. |
| /produto  |   DELETE   | Exclui um produto da lista, a partir da informação do ID único, passado via formulário.                       |
| /produtos |     GET    | Obtém todos os produtos cadastrados, em formato JSON.                                                         |
| /pedido   |    POST    | Inclui um novo pedido na fila, a partir de um JSON contendo as informações necessárias.                       |
| /pedidos  |     GET    | Obtém todos os pedidos ativos no momento, em formato JSON.                                                    |
| /abrir    |    POST    | Abre a loja para iniciar a expedir os pedidos (ver mais informações abaixo).                                  |
| /fechar   |    POST    | Fecha a loja, parando de expedir os pedidos (ver mais informações abaixo).                                    |

As rotas `/abrir` e `/fechar` correspondem a um processo automatizado. A rota `/abrir` muda uma _flag_ booleana para `true`, enquanto que a rota `/fechar` muda para `false`.

Enquanto a loja estiver aberta, o sistema deve realizar a expedição de pedidos a cada 30 segundos, enquanto a fila estiver com pedidos ativos. Esse processo deve rodar paralelamente ao servidor sendo executado. Se a loja estiver fechada, nenhum pedido é expedido. Considere o seguinte fluxo hipotético, com as horas listadas ao lado no formato `HH:MM:SS`:

* 10:20:10 - Sistema recebe pedido 1;
* 10:21:52 - Sistema recebe pedido 2;
* 10:24:04 - Sistema recebe pedido 3;
* 10:30:00 - Loja é aberta;
* 10:30:30 - Pedido 1 é expedido;
* 10:30:43 - Sistema recebe pedido 4;
* 10:31:00 - Pedido 2 é expedido;
* 10:31:30 - Pedido 3 é expedido;
* 10:31:33 - Sistema recebe pedido 5;
* 10:31:40 - Loja é fechada.

Ao final da execução dessas etapas, o sistema terá expedido 3 pedidos, a cada 30 segundos. 2 novos pedidos (`4` e `5`) terão sido incluídos na fila de pedidos, mas como a loja foi fechada às `10:31:40`, esses pedidos não serão expedidos até que a loja seja reaberta.

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
* Avaliação estática do código: o código é eficiente e bem escrito, está organizado, não apresenta redundâncias e atende às boas práticas de programação.

Após a entrega e avaliação, as equipes receberão um barema contendo o descritivo da avaliação, bem como pesos e detalhamento de cada critério de avaliação.

## Demais informações

* Prazo do projeto, até 28/04/2024, sem possibilidade de adiamento do prazo;
* Submissão do link do GitHub atráves do http://estudante.ibmec.br;
* Apenas uma pessoa de cada equipe precisa enviar o link;
* As equipes devem ter no mínimo 3 e no máximo 5 pessoas.
