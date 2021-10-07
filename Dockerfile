FROM alpine:3.13.6

ENV USER_UID=1001

RUN apk add --no-cache ca-certificates==20191127-r5 \
                       openssh-client==8.4_p1-r4

USER ${USER_UID}
