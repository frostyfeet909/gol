# gol
My GO repo that I love.

I got bored so I made the simplest API in Go with:

- ~~Gin~~ Echo framework
- Auto generated Swagger docs
- Fly.io deployment
- Neon PG database

It is deployed at [https://gol-app.fly.dev/swagger/index.html] - give it a few seconds to startup after the first request it's all on zero-scaling.

## Getting Started

1. `make install`
2. `make db`
3. `make run`

## Swagger Docs

Once the server is running, navigate to `http://localhost:8080/swagger/index.html` to view the Swagger UI and explore the API endpoints.
