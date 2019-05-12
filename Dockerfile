FROM python:3-alpine
WORKDIR /home
RUN pip install speedtest-cli
COPY speedtest.linux.amd64 speedtest

CMD while true; do ./speedtest; sleep 3600; done
