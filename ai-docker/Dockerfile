FROM python:3.8

WORKDIR /model

COPY * /model
RUN cd /model \
    && pip install --upgrade pip \ 
    && pip install --no-cache-dir -r requirements.txt

VOLUME [ "/model" ]

EXPOSE 5000

CMD [ "python", "/model/main.py" ]
