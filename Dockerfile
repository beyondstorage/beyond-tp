FROM ubuntu:20.04

ENV UID=1000
ENV GID=1000
ENV USER="developer"
ENV JAVA_VERSION="11"
ENV GO_VERSION="1.16.2"
ENV GO_URL="https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz"
ENV GO_ROOT="/usr/local/go"
ENV FLUTTER_CHANNEL="stable"
ENV FLUTTER_VERSION="2.0.1"
ENV FLUTTER_URL="https://storage.googleapis.com/flutter_infra/releases/$FLUTTER_CHANNEL/linux/flutter_linux_$FLUTTER_VERSION-$FLUTTER_CHANNEL.tar.xz"
ENV FLUTTER_HOME="/home/$USER/flutter"
ENV FLUTTER_WEB_PORT="8090"
ENV FLUTTER_DEBUG_PORT="42000"
ENV PATH="$GO_ROOT/:$FLUTTER_HOME/bin:$PATH"

# install all dependencies
ENV DEBIAN_FRONTEND="noninteractive"
RUN apt-get update \
  && apt-get install --yes --no-install-recommends openjdk-$JAVA_VERSION-jdk curl unzip sed git bash xz-utils libglvnd0 ssh xauth x11-xserver-utils libpulse0 libxcomposite1 libgl1-mesa-glx \
  && rm -rf /var/lib/{apt,dpkg,cache,log}

# go
RUN curl -sSL -o go.tar.gz $GO_URL \
  && tar -C /usr/local -xzf go.tar.gz

# create user
RUN groupadd --gid $GID $USER \
  && useradd -s /bin/bash --uid $UID --gid $GID -m $USER

USER $USER
WORKDIR /home/$USER

# flutter
RUN curl -sSL -o flutter.tar.xz $FLUTTER_URL \
  && mkdir -p $FLUTTER_HOME \
  && tar xf flutter.tar.xz -C /home/$USER \
  && rm flutter.tar.xz \
  && flutter config --no-analytics

# Build beyond-tp
RUN make build
