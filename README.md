# Troubleshooting

After adding the handlerWrapper service returns error 404:

// Todo: English this part
This might be the case when you set the name before you set the
HandlerWrapper. I assume that the server is completely overwritten
when you set the Handlerwrapper and hence the name you have specified is lost.

```
// Right
service = micro.NewService(
	micro.Server(server.NewServer(server.WrapHandler(auth.HandlerWrapper(authHandler)))),
        micro.Name(ServiceName),
)

// Wrong
service = micro.NewService(
        micro.Name(ServiceName),
	micro.Server(server.NewServer(server.WrapHandler(auth.HandlerWrapper(authHandler)))),
)
```