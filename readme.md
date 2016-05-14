# Ping

2/2 of a simple async ping pong app built for my Component-Based Software Development course, using Golang and RabbitMQ.

## Run it

- If you haven't already, head over to the [pong](https://github.com/castillobg/pong) repo, read the
readme, and get everything running. Then, come back.

- Welcome back!

- Clone this repo.

  ```sh
  git clone https://github.com/castillobg/ping.git
  ```

- Build the `ping` executable.

  ```sh
  # Inside the cloned ping repo:
  go build
  ```

- Run ping:
  ```
  ./ping -port 8081 -broker rabbit -address localhost:5672
  ```
  or just simply
  ```
  ./ping -port 8081
  ```

- You can then send a POST request to localhost:8080/api/pings and wait for a `pong`.
  ```
  curl -X POST localhost:8080/api/pings
  ```
