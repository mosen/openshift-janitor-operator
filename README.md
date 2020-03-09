# openshift-janitor-operator #

> Make sure your OpenShift Dev Cluster stays free of Junk!

## Summary ##

The OpenShift Janitor takes care of Projects that people have created and abandoned on an OpenShift cluster.

Normally this happens in a cluster which is not tightly controlled but which is multi-tenanted.

### Whitelisting Projects ###

The following system related namespaces are whitelisted by default and will never be deleted:

    default
    kube-*
    openshift
    openshift-*
    
## Usage ##

### Sweep ###

A **Sweep** object is a one time cluster scan for old Projects. 
To sweep your cluster every few days or weeks, use the **Janitor** resource.

    apiVersion: com.github.mosen.openshift-janitor/v1alpha1
    kind: Sweep
    metadata:
      name: example-sweep
    spec:
    
      # The age of a project, after which the owner will be warned about its removal.
      warnAgeDays: 83
    
      # The maximum number of days old a Project can be before it is removed
      deleteAgeDays: 90
    
      # Don't delete these projects/namespaces.
      ignoreProjects:
        - my-favourite-project
    
      # Don't delete projects/namespaces which have this annotation.
      ignoreAnnotation:
        com.github.mosen.openshift-janitor: "ignore"

### Janitor ###

(TODO)
