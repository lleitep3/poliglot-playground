FROM python:3.11-rc-bullseye

WORKDIR /app

RUN /usr/local/bin/python -m pip install --upgrade pip

COPY requirements.txt /app/requirements.txt

RUN pip install -r requirements.txt

COPY . /app


CMD ["pytest"]
