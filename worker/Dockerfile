FROM python:3.10-slim-buster

WORKDIR /app

COPY requirements.txt .

RUN  pip install --upgrade pip
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 6066

CMD ["faust", "-A", "worker", "worker", "-l", "info"]