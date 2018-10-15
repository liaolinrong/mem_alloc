FROM registry.hundsun.com/library/busybox:1.29.2
COPY mem_alloc /mem_alloc
RUN chmod +x /mem_alloc
CMD ["/mem_alloc", "-n", "500", "-t", "1000", "-d", "3600"]
