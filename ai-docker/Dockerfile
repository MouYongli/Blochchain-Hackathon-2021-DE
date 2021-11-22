FROM python:3.8

WORKDIR /

COPY * /
RUN pip install --upgrade pip \ 
    && pip install --no-cache-dir -r requirements.txt \
    && cd /

EXPOSE 5000

CMD [ "python", "/main.py" ]
