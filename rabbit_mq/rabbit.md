# RabbitMQ

RabbitMq é um message broker(intermediário)

Implementa AMQP, MQTT, STOMP e HTTP (protocolos)

AMQP é baseado no TCP, então é very fast.

Desenvolvido em Erlang

Desacoplamento entre serviços, conseguimos ter plugins e extensões.

Rápido e Poderoso.

É usado a memoria ram, então por isso é muito rápido.

Padrão de mercado.

## Por baixo dos panos

Quando um serviço usa o rabbit, ele abre uma conexão com ele via TCP, e essa conexão é mantida, dentro dessa conexão tem channels (sub-conexões)

Isso evita o processo de criar conexão toda hora.

É chamado de multiplexing connections.

É usado 1 thread por channel (ou seja, gasta uma certa memoria)

## Funcionamento Básico

Publisher -> Exchange -> Queue -> Consumer

Publisher publica algo e o consumer consome algo. Est

Uma mensagem enviada vai ser colocada em uma fila, e quando o consumer lê dessa fila, a mensagem vai desaparecer.

O Exchange pode rotear mensagens para várias filas, basicamente ela pega uma mensagem, processa e descobre pra qual fila essa mensagem vai ser enviada.

## Tipo de exchange

- Direct: Mensagem vai pra exchange e o exchange envia pra uma fila
- Fanout: Mensagem vai pro exchange e esse exchange vai mandar pra todas as filas que estão relacionadas com essa exchange
- Topic: Mensagem vai pro exchange, essa exchange tem algumas regras, de acordo com a regra a mensagem é enviada pra uma fila
- Header: No header da mensagem tem a informação de qual fila que eu quero que a mensagem seja entregue

## Direct

                     /    Queue -> Consumer
Produces -> Exchange --   Queue -> Consumer
                     \    Queue -> Consumer

Cada traço é uma bind, ou seja, eu faço uma bind da minha Exchange pra uma ou mais Queues

No bind eu coloco uma routing key, que é uma key que relaciona uma Exchange a uma Queue

A mensagem enviada por um producer vai ter a routing key, e quando a exchange verificar isso, ela vai rotear pra fila correta.

## Fanout

                     /    Queue -> Consumer
Produces -> Exchange --   Queue -> Consumer
                     \    Queue -> Consumer

O Exchange fanout vai cair em uma exchange, e a exchange manda a mensagem pra todas as filas que são bindadas a essa exchange.

## Topic

                     /    Queue -> Consumer
Produces -> Exchange --   Queue -> Consumer
                     \    Queue -> Consumer

Também vai ter routing key, mas essas routing keys contem regras que parecem com expressões regulares.

Routing key vai usar uma formato com * ou #.

Exemplo, se tenho uma routing key `x.*`, toda mensagem que eu enviar que tiver a routing key começando em x.{qualquercoisa} vai ser enviada para a fila relacionada ao routing key `x.*`.

A expressão # indica uma correspondência de zero ou mais palavras.

O padrão de roteamento de "agreements.eu.berlin.#" corresponde a qualquer routing key que comece com "agreements.eu.berlin"

## Queues

São FIFO.

Existe um esquema de prioridade, mas tem que saber muito pra justificar e essa prioridade.

- Propriedades.
  - Durable: Se ela deve ser salva mesmo depois do restart do broker
  - Auto-delete: Removida automaticamente quando o consumer se desconecta
  - Expiry: Define o tempo que não há mensagens ou clientes consumindo (Quando a fila expira, ela é removida)
  - MessageTTL: Tempo de vida da mensagem (TTL = Time to live)
  - Overflow: Quando a fila transborda
    - Drop Head (remove a última mensagem)
    - Reject Publish (rejeita a mensagem, publisher recebe um erro)
  - Exclusive: Somente channel que criou pode acessar
  - Max length ou bytes: Quantidade máxima de mensagens ou tamanho das mensagens em bytes permitidas. Caso aconteça, termos um overflow e podemos excolher em remover as mensagens mais antigas ou rejeitar as mensagens novas.

Filas são declaradas.

## Dead Letter Queues (DLQ)

Algumas mensagens não conseguem ser entregues por algum motivo

No Rabbit, pode-se ser configurado uma exchange que roteia as mensagens para uma DLQ

Tais mensagens podem ser consumidas e averiguadas posteriormente, por exemplo, pode ser uma aplicação que pega isso e fica lendo.

## Lazy Queues

Mensagens são armazenadas em disco

As vezes o fluxo de mensagem é muito grande e os consumidores não estão tankando, quando o limite de memoria chega para o rabbit, a lazy queue vai guardar as mensagens em disco.

Exige alto I/O

Tem que ter caso especifico pra justificar o uso, pois é mais lento

Quando há milhões de mensagens em uma fila, por qualquer motivo, há possibilidade de liberar a memória. jogando especificamente as mensagens da fila no disco
