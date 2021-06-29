FROM alpine:3.13.5

ENV USER_UID=1001

RUN apk add --no-cache ca-certificates==20191127-r5 \
                       openssh-client==8.4_p1-r3

USER ${USER_UID}
