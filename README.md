# Telegram Queue Service

## Projeto

Esse projeto tem como objetivo começar a explorar o Message Broker **RabbitMQ**. Por hora, será um simples projeto, onde através da integração com serviços terceiros (meus serviços por exemplo) eu enviarei uma mensagem(contexto de notificação) para a fila, onde posteriormente o *microservice* irá consumir tal fila, assim, fazendo com que a mensagem enfim chegue ao **Telegram**.

## Organização deste repositório

Como mencionado acima, teremos o principal serviço consumindo a fila do rabbitmq. Também teremos uma pasta denominada *examples*, onde terá um simples exemplo de integração.

### Exemplo

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	queueService "github.com/douglaszuqueto/telegram/queue"
)

func main() {
	configQueue := queueService.Config{
		IP:       os.Getenv("RABBITMQ_IP"),
		Port:     os.Getenv("RABBITMQ_PORT"),
		Username: os.Getenv("RABBITMQ_USERNAME"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
	}

	queue, err := queueService.New(&configQueue)
	if err != nil {
		log.Panic(err.Error())
	}

	defer queue.Stop()

	var counter int

	for {
        msg := fmt.Sprintf("*Message*: %v", counter)

        err = queue.SendMessage(msg)
        if err != nil {
            log.Panic(err.Error())
        }

        log.Printf("Sending message: %v", msg)
        counter++

        time.Sleep(5 * time.Second)
    }
}
```

No exemplo, nada mais é que um laço infinito com um contador e um tempo para ficar enviando as mensagens ao Telegram à cada 5 segundos.


## Variáveis de ambiente - environment

As variáveis abaixo, para conhecimento serão de escopo global - serão utilizadas no serviço principal como também no exemplo de integração.

```
RABBITMQ_IP=0.0.0.0
RABBITMQ_PORT=5672
RABBITMQ_USERNAME=guest
RABBITMQ_PASSWORD=guest

TELEGRAM_TOKEN=9876543210
TELEGRAM_CHATID=123456789
```

## Como executar o projeto?

Ambos serviços - produtor/consumidor estarão dentro de uma pasta bin - foi compilado e esta pronto para serem executados. Por enquanto os serviços foram compilados para o distros baseadas em Linux e possui versões *amd64* e *ARM*, portanto é possível rodar todos os serviços na Raspberry.

Por questão de facilidade, eu recomendo que você faça o download desse repositório através do **GIT** ou clicando em *Download* aqui no Github mesmo :)

### Requisítos minimos

* Ter uma instáncia do RabbitMQ rodando. Seja no seu computador, Docker, na Raspberry ou até mesmo no [Cloud AMQP](https://www.cloudamqp.com/) caso não queira instalar o serviço
* Possuir um *bot* no telegram;

### Rodando o serviço principal
Como citado anteriormente, você precisa das variáveis de ambiente declaradas. Você pode criar um arquivo exportando as variáveis ou rodar cada serviço com as mesmas de forma inline.

#### Ex. 1:
Crie um arquivo .env(ou com o nome de sua escolha) e coloque o conteúdo abaixo - não esqueça de colocar os valores corretamente.

```
export RABBITMQ_IP=0.0.0.0
export RABBITMQ_PORT=5672
export RABBITMQ_USERNAME=guest
export RABBITMQ_PASSWORD=guest

export TELEGRAM_TOKEN=9876543210
export TELEGRAM_CHATID=123456789
```
Com o arquivo criado o que você precisa é exportar de fato as variáveis. Para isso basta rodar o comando **source .env**. Agora você pode rodar o binário:

```
./bin/telegram-service-amd64
```

Se quiser executar na raspberry, não esqueça de pegar a versão *arm* dos binários.

#### Ex. 2:
De forma inline...

```
RABBITMQ_IP=0.0.0.0 RABBITMQ_PORT=5672 RABBITMQ_USERNAME=guest RABBITMQ_PASSWORD=guest TELEGRAM_TOKEN=9876543210 TELEGRAM_CHATID=123456789 ./bin/telegram-service-amd64
```

### Rodando o exemplo

O procedimento para executar o exemplo serão as mesmas do passo anterior - muda apenas o diretório de onde o binário produtor se encontra.

```
export RABBITMQ_IP=0.0.0.0
export RABBITMQ_PORT=5672
export RABBITMQ_USERNAME=guest
export RABBITMQ_PASSWORD=guest
```

```
source .env

./examples/amqp/bin/telegram-amqp-amd64
```

ou...

```
RABBITMQ_IP=0.0.0.0 RABBITMQ_PORT=5672 RABBITMQ_USERNAME=guest RABBITMQ_PASSWORD=guest ./examples/amqp/bin/telegram-amqp-amd64
```

**Obs** No serviço produtor é utilizado apenas as variavéis referente ao RabbitMQ.


## Diagrama de funcionamento

![img](https://raw.githubusercontent.com/douglaszuqueto/telegram-queue-service/master/.github/diagram-example.png)

Como você pode observar, temos 3 etapas que ocorrem quando uma simples mensagem é enviada.

* 1º - Producer => Exchange: Aqui a exchange é a porta de entrada, toda mensagem vai para essa camada e depois é roteada para a(s) fila(s)
* 2º - Exchange => Queue: Nesta etapa, de fato a mensagem chega na fila, pronta para ser consumida por seu(s) consumer(s)
* 3º - Queue => Consumer: Na finaleira temos o consumer, aqui a mensagem chega e é aplicada a regra de negócio que for. No contexto atual - a mensagem que chega é enviada para o *Telegram* através de sua API

### Exemplo de diagrama um pouco mais complexo 

No diagrama abaixo, é um pouco mais completo do que realmente esse projeto faz. Mas serve de base para saber das possibilidades. 

Por exemplo: Além dos alertas enviados ao telegram, poderiamos ter mais 2 serviços. Um cadastrando as mensagens no banco de dados e outro salvando em um arquivo de log(texto).

![img](https://raw.githubusercontent.com/douglaszuqueto/telegram-queue-service/master/.github/diagram.png)

## Resultado final

### Serviço principal & Exemplo

![img](https://raw.githubusercontent.com/douglaszuqueto/telegram-queue-service/master/.github/screenshot_3.png)

### Mensagem recebida no telegram

![img](https://raw.githubusercontent.com/douglaszuqueto/telegram-queue-service/master/.github/screenshot.png)
