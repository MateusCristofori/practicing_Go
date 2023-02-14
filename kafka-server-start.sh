# # Inicialização do Zookeeper.
# zookeeper-server-start.sh config/zookeeper.properties
# # Inicialização do servidor Kafka.
# kafka-server-start.sh config/server.properties
# # Criação de um novo tópico
# kafka-topics.sh --create --topic readtest --bootstrap-server localhost:9092
# Criação de um produtor(PUB)
kafka-console-producer --bootstrap-server=localhost:9092 --topic=readtest
# Criação de um consumidor(SUB)
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=readtest