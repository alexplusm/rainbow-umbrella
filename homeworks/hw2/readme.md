* brew install wrk
* https://github.com/wg/wrk
* https://medium.com/@felipedutratine/intelligent-benchmark-with-wrk-163986c1587f
* https://www.digitalocean.com/community/tutorials/how-to-benchmark-http-latency-with-wrk-on-ubuntu-14-04
* https://www.youtube.com/watch?v=eH4Tm4ASinE

### examples
```shell
    wrk -t12 -c1 -d30s --latency http://127.0.0.1:8080/api/v1/users

    wrk -t1 -c1 -d1s -s ./script.lua -H 'X-SessionId: 4814972b-138d-4d8d-b70b-35438b1b0feb' --latency http://127.0.0.1:8080/api/v1/users
```