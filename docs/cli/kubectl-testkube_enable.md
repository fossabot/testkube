## kubectl-testkube enable

Enable feature

```
kubectl-testkube enable <feature> [flags]
```

### Options

```
  -h, --help   help for enable
```

### Options inherited from Parent Commands

```
      --analytics-enabled   enable analytics
  -a, --api-uri string      api uri, default value read from config if set
  -c, --client string       client used for connecting to Testkube API one of proxy|direct (default "proxy")
      --namespace string    Kubernetes namespace, default value read from config if set (default "testkube")
      --oauth-enabled       enable oauth
      --verbose             show additional debug messages
```

### SEE ALSO

* [kubectl-testkube](kubectl-testkube.md)	 - Testkube entrypoint for kubectl plugin
* [kubectl-testkube enable analytics](kubectl-testkube_enable_analytics.md)	 - Enable collecting of anonymous analytics
* [kubectl-testkube enable oauth](kubectl-testkube_enable_oauth.md)	 - enable oauth authentication for direct api

