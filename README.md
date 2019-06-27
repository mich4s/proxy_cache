# Proxy Cache
This is simple project of small one binary golang app to create proxy/cache for http web applications. This is good for static endpoints writen in frameworks with big bootstrap time(eg. php Laravel/Symfony). With this you can configure how often cache needs to be reloaded and all your stateless compute heavy endpoints will no longer be computed each request.


### TODO

- [x] Basic proxy with cache option
- [ ] Route params for better rest url matching
- [ ] Support for header states
- [ ] Cache reload worker
- [ ] Traffic monitor

Since this is free-time project I cannot focus on providing production ready solution here. However I use customized version of this library with my own PHP application and for now it works fine :)