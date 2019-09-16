FROM geertjohan/pgmodeler-cli:latest

RUN mkdir /go
ENV GOPATH /go
ENV QT_QPA_PLATFORM=

# Install postgres client
RUN wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - && \
    echo "deb http://apt.postgresql.org/pub/repos/apt/ bionic-pgdg main" > /etc/apt/sources.list.d/postgres.list && \
    apt-get update && \
    apt-get -y install --no-install-recommends postgresql-client-10 && \
    rm -rf /var/lib/apt/lists/*

# Install migrate
RUN wget --quiet -O - https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - && \
    echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ bionic main" > /etc/apt/sources.list.d/migrate.list && \
    apt-get update && \
    apt-get install -y migrate && \
    rm -rf /var/lib/apt/lists/*

# Install java
RUN apt-get update && apt-get -y install --no-install-recommends \
    git \
    ttf-dejavu \
    xvfb \
    maven openjdk-8-jdk \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Install go
ENV GOLANG_VERSION 1.13
RUN wget -O go.tgz "https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz" && \
    tar -C /usr/local -xzf go.tgz && \
    rm go.tgz
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Install modd
RUN go get github.com/cortesi/modd/cmd/modd

# Install docker so developers can run the diff script in modd
RUN curl -fsSLO https://download.docker.com/linux/static/stable/x86_64/docker-17.09.0-ce.tgz && \
    tar --strip-components=1 -xvzf docker-17.09.0-ce.tgz -C /usr/local/bin && \
    rm docker-17.09.0-ce.tgz

# Download/compile apgdiff
RUN git clone https://github.com/fordfrog/apgdiff.git /opt/apgdiff && \
    (cd /opt/apgdiff && mvn package) && \
    cp /opt/apgdiff/target/apgdiff-2.6.0-SNAPSHOT.jar /opt/apgdiff.jar &&\
    rm -rf /opt/apgdiff && \
    echo "#!/bin/bash\njava -jar /opt/apgdiff.jar \$@" > /usr/local/bin/apgdiff &&\
    chmod +x /usr/local/bin/apgdiff

WORKDIR /go/src/go.nlx.io/nlx/directory-db
CMD ["/bin/sh", "-c", "xvfb-run /go/bin/modd -f diff/modd.conf"]
