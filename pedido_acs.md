# Pedidos das ACs

Abaixo estão indicados os pedidos para cada AC do curso. Consulte os prazos em http://estudante.ibmec.br, e as orientações gerais para entrega [aqui](https://victor0machado.github.io/ed/orientacao_entregas.html).

- [Pedidos das ACs](#pedidos-das-acs)
  - [AC1](#ac1)
  - [AC2](#ac2)
  - [AC3 (em sala)](#ac3-em-sala)
  - [AC4](#ac4)
  - [AC5](#ac5)
  - [AC6 (em sala)](#ac6-em-sala)
  - [AC7](#ac7)
  - [AC8](#ac8)
  - [AC9](#ac9)
  - [AC10](#ac10)
  - [AC10](#ac10-1)

## AC1

1. Crie uma função `calculaMedia` que receba um número variável de parâmetros de tipo ponto flutuante e retorne a média aritmética desses valores.
2. Construa uma função `verificaParidade` que receba um inteiro e retorne se o número é par ou ímpar.
3. Desenvolva uma função `minhaPotencia` que receba dois inteiros, `base` e `expoente`, e retorne o resultado de base elevado ao expoente, sem usar funções prontas da linguagem.
4. Implemente uma função `converteCelsiusParaFahrenheit` que receba uma temperatura em Celsius e retorne a temperatura convertida em Fahrenheit.

## AC2

1. Faça um programa que cria um vetor de inteiros com 10 elementos. Então inicialize este vetor com números quaisquer e imprima na tela todos os seus elementos, um abaixo do outro.
2. Faça uma função/método que receba uma string como parâmetro e que retorne uma nova string, onde a sequência dos caracteres foi invertida. Dentro da parte principal (main), leia uma string digitada pelo usuário e passe para a função/método criada, imprimindo em seguida a string devolvida. Para esse exercício, pesquise sobre o comportamento e a interação entre dados do tipo `rune` e do tipo `string`.
3. Crie um tipo chamado `Ponto` que represente um ponto no plano cartesiano, com coordenadas `X` e `Y`. Em seguida, implemente um método chamado `DistanciaOrigem()` que calcule a distância desse ponto até a origem `(0,0)`.
4. Crie um pacote chamado `geometria` que contenha funções para calcular a área e o perímetro de um retângulo. Em seguida, use esse pacote em um programa principal para calcular a área e o perímetro de um retângulo com dimensões fornecidas pelo usuário.

## AC3 (em sala)

1. Elabore um struct `Contato`, que deve conter os campos nome e e-mail. O struct deve conter um método para alterar o e-mail.
2. Elabore uma função `adicionarContato`, que recebe um contato e um array com até 5 elementos e inclui o contato no primeiro elemento vazio do array.
3. Elabore uma função `excluirContato`, que recebe um array de 5 elementos e elimina o último contato válido.
4. Elabore uma interface por linha de comando na função `main`, que cria um array de 5 elementos e permite a inclusão ou exclusão de contatos.

## AC4

1. Implemente em Go a solução recursiva para o problema da Torre de Hanói.
2. Implemente em Go um algoritmo que resolva o seguinte problema: dado um array de números inteiros positivos, considerado ordenado crescentemente, e um valor alvo, encontre um par de números no array cuja soma seja igual ao valor alvo. Se nenhum par for encontrado, retorne um valor `(-1, -1)` indicando que nenhum par foi encontrado. Resolva esse problema com um algoritmo cuja complexidade é O(n).

## AC5

Acessar o convite para o site **Beecrowd** recebido no seu e-mail do Ibmec (matricula@alunos.ibmec.edu.br). Caso já tenha uma conta no Beecrowd com seu e-mail pessoal, me envie um e-mail (silva.victor@ibmec.edu.br) ou entre em contato pelo WhatsApp para que eu possa enviar um novo convite.

Na disciplina "Estruturas de Dados", realizar a lista com o nome "AC5". Suba o código no GitHub como normalmente e envie pelo https://estudante.ibmec.br o link para os exercícios no GitHub.

## AC6 (em sala)

Considere duas listas ordenadas, `a` e `b`, com `m` e `n` elementos cada uma, respectivamente. Todos os elementos das listas são distintos, ou seja, se o número `2` está na lista `a`, ele não aparece na lista `b`. Escreva um algoritmo em pseudocódigo que retorne uma única lista ordenada `v`, com `m + n` elementos, contendo todos os elementos de `a` e `b`.

## AC7

Acessar o convite para o site **Beecrowd** recebido no seu e-mail do Ibmec (matricula@alunos.ibmec.edu.br). Caso já tenha uma conta no Beecrowd com seu e-mail pessoal, me envie um e-mail (silva.victor@ibmec.edu.br) ou entre em contato pelo WhatsApp para que eu possa enviar um novo convite.

Na disciplina "Estruturas de Dados", realizar a lista com o nome "AC7". Suba o código no GitHub como normalmente e envie pelo https://estudante.ibmec.br o link para os exercícios no GitHub.

## AC8

Teste realizado em sala.

## AC9

Na disciplina "Estruturas de Dados" do **Beecrowd**, realizar a lista com o nome "AC9" e resolver na ferramenta. Em seguida, suba o código no GitHub como normalmente e envie pelo https://estudante.ibmec.br o link para os exercícios no GitHub.

## AC10

Na disciplina "Estruturas de Dados" do **Beecrowd**, realizar a lista com o nome "AC10" e resolver na ferramenta. Em seguida, suba o código no GitHub como normalmente e envie pelo https://estudante.ibmec.br o link para os exercícios no GitHub.

**Dicas:**

- No exercício 1161, você deverá rodar o código até chegar no EOF (_end of file_). Isso pode ser analisado capturando o erro no `fmt.Scanln`, como no caso abaixo:

  ``` go
  func leInfo():
    var x int
    _, err := fmt.Scanln(&x)

    if err != nil { return }
  ```

- No exercício 1221, para ser amis eficiente, você pode usar o conceito que diz que todo número não primo `n` possui pelo menos um divisor menor ou igual à raiz quadrada de `n`.

## AC10

Na disciplina "Estruturas de Dados" do **Beecrowd**, realizar a lista com o nome "AC11" e resolver na ferramenta. Em seguida, suba o código no GitHub como normalmente e envie pelo https://estudante.ibmec.br o link para os exercícios no GitHub.
