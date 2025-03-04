FROM nginx:latest

# 필수 빌드 도구 및 종속성 설치
RUN apt-get update && apt-get install -y \
    git \
    gcc \
    make \
    libpcre3-dev \
    zlib1g-dev \
    wget \
    libssl-dev \
    openssl \
    build-essential

# NGINX 소스 다운로드 및 RTMP 모듈 빌드
RUN git clone https://github.com/arut/nginx-rtmp-module.git \
    && wget http://nginx.org/download/nginx-1.23.4.tar.gz \
    && tar -xvzf nginx-1.23.4.tar.gz \
    && cd nginx-1.23.4 \
    && ./configure --add-module=../nginx-rtmp-module \
    && make && make install

COPY nginx.conf /usr/local/nginx/conf/nginx.conf

COPY stat.xsl /usr/local/nginx/html/stat.xsl


EXPOSE 8080 1935
# NGINX 실행 명령
CMD ["/usr/local/nginx/sbin/nginx", "-g", "daemon off;"]