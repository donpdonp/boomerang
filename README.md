A simple URL shortener/redirect.

To create a new redirect, do an HTTP POST to the slug or short-name that 
you'd like to use in the future. Put the destination URL in the body of the POST.

```
POST /<slug>
<url>
```

Then use the slug normally, http://my_server/slug will redirect to the associated URL.

```
GET /<slug>
```
