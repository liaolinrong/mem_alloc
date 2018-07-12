FROM registry.hundsun.com/library/centos:7.2.1511
COPY mem_alloc /mem_alloc
RUN chmod +x /mem_alloc
CMD ["mem_alloc", "-n", "1000", "-t", "1000"]
