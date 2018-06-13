### Usage
```go
`k8s:"a,b,c"`
```
> Support 2 type configmaps
>  1) configmap from file
>      a) configmap name (required)
>      b) file name (required)
>      c) override namespace or empty string
>
>  2) configmap from literal
>      a) configmap name (required)
>      b) empty string
>      c) override namespace or empty string
