FROM python:3.6
RUN mkdir /app
WORKDIR /app
COPY proto /app/proto
COPY server.py /app/
COPY requirements.txt /app/
RUN pip install -r requirements.txt
ENTRYPOINT ["python"]
CMD ["server.py"]
