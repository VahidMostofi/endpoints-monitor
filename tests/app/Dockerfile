FROM alpine

WORKDIR /usr/src/app

FROM alpine:latest
RUN echo "http://dl-8.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories \
     && apk --no-cache --update-cache add postgresql-dev g++ linux-headers gfortran build-base bash libpng-dev openblas-dev wget freetype-dev python3-dev py3-pip \
     && ln -s /usr/include/locale.h /usr/include/xlocale.h \
     && pip3 install setuptools wheel \
     && pip3 install numpy 

COPY requirements.txt .
RUN pip3 install --no-cache-dir -r requirements.txt

COPY app.py .

CMD [ "python3", "./app.py" ]