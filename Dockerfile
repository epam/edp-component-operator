FROM alpine:3.13.5

ENV USER_UID=1001

RUN apk add --no-cache ca-certificates==20191127-r2 \
                       openssh-client==8.1_p1-r0

USER ${USER_UID}
