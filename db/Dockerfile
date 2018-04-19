FROM geertjohan/pgmodeler-cli as build

RUN mkdir /go
ENV GOPATH /go
ENV QT_QPA_PLATFORM=

RUN wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - && \
    echo "deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main" > /etc/apt/sources.list.d/postgres.list && \
    apt update && \
    apt -y install --no-install-recommends postgresql-client-10 && \
    rm -rf /var/lib/apt/lists/*

RUN apt-get update && apt-get -y install --no-install-recommends \
    golang-go \
    git \
    ttf-dejavu \
    xvfb \
    maven openjdk-8-jdk \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Install modd and migrate
RUN go get github.com/cortesi/modd/cmd/modd && \
    go get -u -d github.com/mattes/migrate/cli github.com/lib/pq && \
    go build -tags 'postgres' -o /usr/local/bin/migrate github.com/mattes/migrate/cli

# Install docker so developers can run the diff script in modd
RUN curl -fsSLO https://download.docker.com/linux/static/stable/x86_64/docker-17.09.0-ce.tgz && \
    tar --strip-components=1 -xvzf docker-17.09.0-ce.tgz -C /usr/local/bin && \
    rm docker-17.09.0-ce.tgz

# Download/compile apgdiff
RUN git clone https://github.com/GeertJohan/apgdiff.git /opt/apgdiff && \
    (cd /opt/apgdiff && mvn package) && \
    cp /opt/apgdiff/target/apgdiff-2.5.0-SNAPSHOT.jar /opt/apgdiff.jar &&\
    rm -rf /opt/apgdiff && \
    echo "#!/bin/bash\njava -jar /opt/apgdiff.jar \$@" > /usr/local/bin/apgdiff &&\
    chmod +x /usr/local/bin/apgdiff

ADD db/wait-for-postgres.sh /usr/local/bin/wait-for-postgres.sh
ADD db/pgmodeler-cli-config /root/.config/pgmodeler/

WORKDIR /go/src/github.com/VNG-Realisatie/nlx/db
CMD ["/bin/sh", "-c", "/usr/local/bin/wait-for-postgres.sh xvfb-run /go/bin/modd"]


FROM ubuntu:artful

# RUN apt-get update && \
#     apt-get -y install --no-install-recommends wget ca-certificates && \
#     wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - && \
#     echo "deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main" > /etc/apt/sources.list.d/postgres.list && \
#     apt update && \
#     apt -y install --no-install-recommends postgresql-client-10 && \
#     rm -rf /var/lib/apt/lists/*

COPY --from=build /usr/local/bin/migrate /usr/local/bin/migrate
ADD db/migrations /opt/db-migrations

ENV PGHOST postgres
ENV PGUSER postgres
ENV PGPASS postgres
ENV PGDB nlx

RUN echo $'#!/bin/sh\n\
\n\
/usr/local/bin/migrate -path=/opt/db-migrations -database=postgres://\${PGUSER}:\${PGPASS}@\${PGHOST}/\${PGDB}?sslmode=disable up\n'\
> /opt/migrate-up.sh && chmod +x /opt/migrate-up.sh

CMD ["/opt/migrate-up.sh"]
