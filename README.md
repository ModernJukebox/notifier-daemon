# NotifierDaemon

The Notifier daemon is, of course, a daemon that sends data to a configured server.
You can configure the interval at which the daemon sends data to the server and where the data is fetched from.
It uses http post requests to send the data to the server.

Basically, it is a simple http client that sends continuously data to the server.
