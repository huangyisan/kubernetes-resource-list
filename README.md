# kubernetes-resource-list

This app can show a list of Kubernetes API resources and subresources

## Why

When I was writing my RBAC configuration, I couldn't find resources, especially subresources, and the official kubernetes documentation doesn't list them anywhere.

## How To Use
```shell
$ ./kube-resource --help
Usage of ./kube-resource:
  -c, --kubeconfig string   (optional) absolute path to the kubeconfig file
  -p, --prefer              (optional) only display the supported resources with the version preferred by the server.
  -s, --search string       (optional) only display the supported resources for a group and version.
  -v, --version             show app version
```

### Example
```shell
$ ./kube-resource
APIGROUPS                	GROUPVERSION        	NAME                               	VERBS
core                     	v1                  	bindings                           	[create]
core                     	v1                  	componentstatuses                  	[get list]
core                     	v1                  	configmaps                         	[create delete deletecollection get list patch update watch]
core                     	v1                  	endpoints                          	[create delete deletecollection get list patch update watch]
core                     	v1                  	events                             	[create delete deletecollection get list patch update watch]
core                     	v1                  	limitranges                        	[create delete deletec

...
..
.
```