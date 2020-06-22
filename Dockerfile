FROM scratch
COPY ./aws_eks_hello /aws_eks_hello
ENTRYPOINT ["/aws_eks_hello"]
