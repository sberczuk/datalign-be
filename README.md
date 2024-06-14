# Project Info

this project consists of 2 repos:

- [https://github.com/sberczuk/datalign-be](https://github.com/sberczuk/datalign-be): the backend service (this one)
- [https://github.com/sberczuk/dainterview-fe](https://github.com/sberczuk/dainterview-fe): the front end service

To build and run the docker containers, see the readmes in each project.

To once he services are running:

The front end is available at  http://localhost:8080
and connects to the backend on port 3000

The front end could have been a static page too

## Missing things

There are some TODOs ih the code that highlight a few issues. This (pverlapping) list covers the more significant ones.

- Ports are hard coded. in real life they would be configurable
- The Front End sets a no-cors attribute. We might consider using a proxy in a real case
- The React Tests are a bit thin. As I mentioned. My Front end testing experience is evolving. I'm using vitetest. But I'd add tests for
  - The Validation in the component
  - The Post and parsing of the response

# Running

 ```bash
docker build -t datalign-backend .
docker run  -p 3000:3000 datalign-backend
 ```

# Choices

Web framework
- [Fiber](https://github.com/gofiber/fiber)