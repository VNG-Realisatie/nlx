
# nlx-directory chart

This chart depends on a central nginx-ingress with ssl-passthrough enabled.

```sh
helm install --namespace=nginx-ingress --values helm/nginx-ingress-values.yaml stable/nginx-ingress
```

## New environment

To deploy a new environment, run the first time with `dbResetDirectory=true`.
