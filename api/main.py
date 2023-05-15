from fastapi import FastAPI
from pydantic import BaseModel
from aiokafka import AIOKafkaProducer
import json

app = FastAPI(openapi_url="/api/openapi.json", 
              docs_url="/api/docs", 
              redoc_url="/api/redoc")

class BaseRecord(BaseModel):
    request_id: str
    request_type: str
    pipeline: dict
    data: dict

@app.on_event("startup")
async def startup_event():
    # Create an async Kafka producer
    producer = AIOKafkaProducer(bootstrap_servers='kafka:9092')
    await producer.start()
    app.state.producer = producer

@app.on_event("shutdown")
async def shutdown_event():
    # Stop the Kafka producer on shutdown
    await app.state.producer.stop()

@app.post("/publish")
async def publish_message(record: BaseRecord):
    # Get the Kafka producer from the app state
    producer = app.state.producer

    # Convert the record to a dictionary
    record_dict = record.dict()

    # Serialize the record to JSON bytes
    value_bytes = json.dumps(record_dict).encode('utf-8')

    # Publish the record to the 'section_gpu' topic
    await producer.send_and_wait('section_gpu', value=value_bytes)

    return {"message": "Message published successfully"}