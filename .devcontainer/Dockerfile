ARG GO_VERSION=1.21.4

FROM golang:$GO_VERSION-bookworm
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000
ARG GOPLS_VERSION=latest

RUN addgroup --gid $USER_GID $USERNAME && \
    adduser $USERNAME --shell /bin/sh --disabled-password -u $USER_UID --gid $USER_GID && \
    mkdir -p /etc/sudoers.d && \
    echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME

RUN go install golang.org/x/tools/gopls@${GOPLS_VERSION}
RUN for tool in tools/gopls tools/cmd/goimports lint/golint; \
    do go install golang.org/x/${tool}@latest; \
    done

# Detect shadowing bugs
RUN go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest

# Visual Studio Code tools
RUN go install github.com/cweill/gotests/gotests@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/josharian/impl@latest
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

# Setup shell
USER $USERNAME
#RUN sh -c "$(wget -O- https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" "" --unattended &> /dev/null
#ENV ENV="/home/$USERNAME/.ashrc" \
#    ZSH=/home/$USERNAME/.oh-my-zsh \
#    EDITOR=vi \
#    LANG=en_US.UTF-8 \
#    PATH=/home/vscode/.local/bin:$PATH
#RUN printf 'ZSH_THEME="agnoster"\nENABLE_CORRECTION="false"\nplugins=(git copyfile extract colorize dotenv encode64 golang)\nsource $ZSH/oh-my-zsh.sh' > "/home/$USERNAME/.zshrc"
#RUN echo "exec `which zsh`" > "/home/$USERNAME/.ashrc"

USER root

# Allow the developer to download third-party modules
RUN chown -R $USER_UID:$USER_GID /go/pkg

#WORKDIR /app
#COPY . . 
#RUN go mod download
#RUN go build -o /godocker

#EXPOSE 8080

#CMD [ "/godocker" ]