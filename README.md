# Fibserver
Server to produce the requested number in the fibonacci sequence.

# Usage
\[server]:[port]/[fibIndex]

e.g. localhost:8080/10

55

# Installation (for now)
> docker build --tag fibserver .
>
> docker run --publish 8080:8080 fibserver