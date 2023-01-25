FROM ubuntu:18.04

ENV DEBCONF_NOWARNINGS=yes

RUN apt-get update && apt-get install -y openssh-server vim sudo iputils-ping net-tools wget software-properties-common default-jdk build-essential git  cmake clang libssl-dev

RUN mkdir /var/run/sshd

RUN mkdir /root/python_test/
RUN mkdir /root/c_test/
RUN mkdir /root/java_test/

ARG ROOT_PASSWORD
RUN echo root:${ROOT_PASSWORD} | chpasswd

RUN sed -i 's/#\?PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
RUN sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd

ENV NOTVISIBLE "in users profile"
RUN echo "export VISIBLE=now" >> /etc/profile

EXPOSE 22

CMD ["/usr/sbin/sshd", "-D"]