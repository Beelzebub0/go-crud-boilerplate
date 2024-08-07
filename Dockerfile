FROM scratch

WORKDIR /usr/src/app

COPY build/app ./

CMD [ "/usr/src/app/app", "-config=/usr/src/app/conf.yaml" ]
