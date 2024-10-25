## smoke test

As described by https://protohackers.com/problem/0

This implements an echo service as described in rfc862:

>TCP Based Echo Service
>
>   One echo service is defined as a connection based application on TCP.
>   A server listens for TCP connections on TCP port 7.  Once a
>   connection is established any data received is sent back.  This
>   continues until the calling user terminates the connection.

