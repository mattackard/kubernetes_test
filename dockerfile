FROM scratch
ADD /bin/ip-echo /
EXPOSE 8080
CMD ["/ip-echo"]
