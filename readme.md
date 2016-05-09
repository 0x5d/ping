# Ping

A simple async ping pong app built for my Component-Based Software Development course, using Golang
and Kafka on Docker.

## Run it

- Make sure [Golang](https://golang.org/dl/) and
[Docker](https://docs.docker.com/engine/installation/) are properly installed.

- Run the [RabbitMQ container](https://hub.docker.com/_/rabbitmq/).

- Clone this repo.

  ```sh
  git clone https://github.com/castillobg/ping.git
  ```

- Build the `ping` executable.

  ```sh
  # Inside the cloned ping repo:
  go build
  ```

- Run `ping`.
