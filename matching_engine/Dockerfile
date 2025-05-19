FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        build-essential \
        cmake \
        git \
        g++ \
        make \
        libboost-all-dev \
        libssl-dev \
        pkg-config \
        ca-certificates \
        uuid-dev \
        libasio-dev \
        libjsoncpp-dev && \
    update-ca-certificates && \
    apt-get clean

WORKDIR /workspace
RUN git clone --recursive https://github.com/drogonframework/drogon  && \
    cd drogon && \
    git submodule update --init --recursive && \
    mkdir build && \
    cd build && \
    cmake .. -DBUILD_SHARED_LIBS=ON -DCMAKE_INSTALL_PREFIX=/usr/local && \
    make -j$(nproc) && \
    make install && \
    # Проверка наличия конфигурационных файлов
    ls -la /usr/local/lib/cmake/Drogon/ || \
    (echo "Файлы конфигурации Drogon не найдены!" && exit 1)

WORKDIR /app
COPY . .

RUN mkdir build && \
    cd build && \
    cmake .. -DCMAKE_PREFIX_PATH=/usr/local -Ddrogon_DIR=/usr/local/lib/cmake/Drogon && \
    make -j$(nproc)

EXPOSE 18888

CMD ["./build/matching_engine"]