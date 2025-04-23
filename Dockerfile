FROM postgres:17

# 1️⃣ Installer les dev headers Postgres, outils de build, Git et CA certs
RUN apt-get update \
 && apt-get install -y --no-install-recommends \
      postgresql-server-dev-17 \
      build-essential \
      git \
      ca-certificates \
 && update-ca-certificates \
 && rm -rf /var/lib/apt/lists/*

# 2️⃣ Cloner & installer pg-ulid
RUN git clone https://github.com/andrielfn/pg-ulid.git /pg-ulid \
 && cd /pg-ulid \
 && make install \
 && cd / \
 && rm -rf /pg-ulid

# 3️⃣ Init script pour activer l’extension ULID
COPY init-ulid.sql /docker-entrypoint-initdb.d/
