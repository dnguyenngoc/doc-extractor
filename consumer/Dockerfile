# FROM wallies/python-cuda:3.11-cuda11.7-runtime
FROM python:3.11.3-slim-bullseye

WORKDIR /app

COPY requirements.txt .

RUN  pip install --upgrade pip

RUN pip install --no-cache-dir -r requirements.txt

COPY ./app /app

ENTRYPOINT [ "python3", "main.py" ]
