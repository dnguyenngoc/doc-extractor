# Base image
FROM python:3.10-slim-buster

# Set the working directory in the container
WORKDIR /app

# Copy the requirements.txt file to the container
COPY requirements.txt .

# Install the dependencies
RUN pip install --no-cache-dir -r requirements.txt

RUN pip install uvicorn
RUN pip install aiokafka

# Copy the application code to the container
COPY . .

# Expose the port on which the FastAPI application will run
EXPOSE 8081

# Command to run the FastAPI application
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8081"]