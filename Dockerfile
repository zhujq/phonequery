
FROM alpine:latest
WORKDIR /
ADD phonequery /phonequery 
ADD entrypoint.sh /entrypoint.sh
ADD phone.dat /phone.dat
RUN  chmod +x /phonequery  && chmod 777 /entrypoint.sh
ENTRYPOINT  /entrypoint.sh 

EXPOSE 8080
