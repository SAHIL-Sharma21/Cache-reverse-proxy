# Caching Reverse Proxy Server

- A Simple caching reverse proxy server built in ```GOLang``` that caches the response from  an origin server and serves caches responses to improve performance.
If the requested content is cached, the server returns the cached response; if not, it fetches the response from the origin server, caches it, and forwards it to the client.

## Features
- **Caching**: Caches responses from the origin server to reduce the number of requests to the origin.
- **Proxy**: Forwards client requests to the origin server if the requested resource is not cached.
- **CLI SUpport**: Users can specify the origin server and the port for the proxy server through command-line flags.

## Usage

### Pre-requisites
- Go version 1.19+ or higher
- A running origin server from which the proxy can fetch responses.

### Installation

1. clone this repo
```bash
    git clone https://github.com/SAHIL-Sharma21/caching-proxy.git
```
2. Navigate to the root directory
```bash
    cd caching-proxy
```

3. Build the project
```bash
    go build -o cmd/main.go
``` 

4. Run the server
```bash
    ./cmd/main.go --port <PORT> --origin <ORIGIN_SERVER_URL>
```
- ```--port```: The port on which the proxy server will listen for requests.(default: 8080)
- ```--origin```: The origin server's base URL to which requests will be forwarded when the response is not cached.

### Example
```bash
    ./cmd/main.go --port 8080 --origin https://jsonplaceholder.typicode.com
```
### Testing the proxy server

```bash
    curl http://localhost:8080/posts/1
```
- On the first request, the proxy fetches the response from the origin server and caches it.

```bash
    curl http://localhost:8080/posts/1
```
- On the second request, the proxy returns the cached response.

## CONTRIBUTING

- [GitHub](https://github.com/SAHIL-Sharma21/Cache-reverse-proxy) 

Feel free to fork this project and submit pull requests for improvements, bug fixes, or additional features.


