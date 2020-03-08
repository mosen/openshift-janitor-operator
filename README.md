# openshift-janitor-operator #

> Make sure your OpenShift Dev Cluster stays free of Junk!

## Summary ##

The OpenShift Janitor takes care of Projects that people have created and abandoned on an OpenShift cluster.

Normally this happens in a cluster which is not tightly controlled but which is multi-tenanted.

### Whitelisting Projects ###

The following system related namespaces are whitelisted by default:

    default
    kube-node-lease
    kube-public
    kube-system
    kube-service-catalog
    kube-*
    
    openshift
    openshift-*
    
