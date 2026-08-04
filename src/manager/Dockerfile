FROM python:3-alpine

WORKDIR /app

COPY ./requirements.txt .

RUN pip install -r requirements.txt

COPY . .

ENV ATTOCI_API_HOST=0.0.0.0
ENV ATTOCI_API_PORT=8089
ENV ATTOCI_API_KEY=12345678
ENV ATTOCI_API_LOG_LEVEL=CRITICAL
ENV ATTOCI_DB_MIGRATIONS_PATH=./data/migrations
ENV ATTOCI_DB_DATABASES_PATH=./data/databases
ENV ATTOCI_DB_JSON_PATH=./data/json

CMD ["python", "main.py"]