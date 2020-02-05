FROM centos:7
RUN useradd -m --uid 1987 otdd-test-runner && \
    echo "otdd-test-runner ALL=NOPASSWD: ALL" >> /etc/sudoers
RUN yum install -y sudo iptables
COPY ./otdd-test-runner /home/otdd-test-runner/
COPY ./run.sh /home/otdd-test-runner/
USER otdd-test-runner
WORKDIR /home/otdd-test-runner/
CMD sh ./run.sh
