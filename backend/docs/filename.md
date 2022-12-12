## Arquivos & Pastas
<br/>



### */cmd*
> **Objetivo**
> <br/>
> Todos os executáveis, porém devem ser separados por pastas

> **Arquivos**
> <br/>
> *nome_service/main.go*
# <br/>



### */common*
> **Objetivo**
> <br/>
> Coisas que são compartilhadas por todos ou quase todos os arquivos

> **Arquivos**
> <br/>
> *sem padrão*
# <br/>



### */docs*
> **Objetivo**
> <br/>
> Documentação e arquivos talvez importantes porém sem código em si

> **Arquivos**
> <br/>
> *sem padrão*
# <br/>



### */ext*
> **Objetivo**
> <br/>
> Quando for feita alguma integração, os arquivos devem ser colocados nessa pasta, por exemplo, conexões a outras APIs

> **Arquivos**
> <br/>
> *nome_do_terceiro/funcionalidade_do_arquivo_ext.go*
# <br/>



### */internal/db*
> **Objetivo**
> <br/>
> Todas as operações realizadas no banco de dados

> **Arquivos**
> <br/>
> *nomeDoProto_db.go*
# <br/>



### */internal/db/migration*
> **Objetivo**
> <br/>
> Arquivos SQL para inicializar o banco de dados

> **Arquivos**
> <br/>
> *0001_nome.up.sql* <br/>
> *0001_nome.down.sql*
# <br/>


 
### */internal/services/api*
> **Objetivo**
> <br/>
> Arquivos voltados as rotas da API, tanto a declaração dela quanto os handlers

> **Arquivos**
> <br/>
> *routerGroup_handler.go*
# <br/>



### */internal/services/grpc*
> **Objetivo**
> <br/>
> Arquivos voltados ao servico do grpc todos os arquivos que sejam um método do `.proto`

> **Arquivos**
> <br/>
> *nomeDoProto_service.go*
# <br/>



### */internal/services/rabbit*
> **Objetivo**
> <br/>
> Arquivos voltados ao serviço amqp 

> **Arquivos**
> <br/>
> *nomeDoProto_consumer.go*
> *nomeDoProto_producer.go*
# <br/>



### */internal/tests*
> **Objetivo**
> <br/>
> Testes unitários para serem executados no CI

> **Arquivos**
> <br/>
> *nomeDoProto_test.go*
# <br/>



### */internal/utils*
> **Objetivo**
> <br/>
> Arquivos que são funções que evitam duplicação de código e/ou funcionalidades que deveriam ser compartilhadas

> **Arquivos**
> <br/>
> *sem padrão*
# <br/>



### */protodefs/proto*
> **Objetivo**
> <br/>
> Arquivos .proto que são utilizados pelo GRPC

> **Arquivos**
> <br/>
> *nome.proto*
# <br/>



### */protodefs/gen*
> **Objetivo**
> <br/>
> Arquivos que são gerados pelo proto

> **Arquivos**
> <br/>
> *nome_grpc.pb.go*
> *nome.pb.go*
# <br/>