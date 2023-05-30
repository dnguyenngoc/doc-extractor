import json
import logging
from confluent_kafka import Consumer
from setting import config

logging.basicConfig(level=logging.INFO)

# function for start consumer
def start_consumer(server, config):
    consumer = Consumer(server)
    consumer.subscribe(config['in_topics'])
    
    logging.info(consumer)

    while True:
        msg = consumer.poll(1.0)

        if msg is None:
            continue

        if msg.error():
            logging.info("Consumer error: {}".format(msg.error()))
            continue

        key = msg.key()
        value = msg.value()
        
        logging.info("Group ID: %s, Topic: %s, Key: %s, Value: %s" %(config['group_id'], msg.topic(), key, value))

    consumer.close()


# Read consumer configurations from the JSON file
with open('consumer_config.json') as file:
    configs = json.load(file)

# Start consumers based on the configurations
config_general = {
    "bootstrap.servers": config.KAFKA_BOOTSTRAP_SERVERS,
    "group.id": "pipeline-demo"
}

for c_name, c_config in configs.items():
    logging.info("[START] start consumer: %s" %(c_name))
    start_consumer(server=config_general, config=c_config)

    