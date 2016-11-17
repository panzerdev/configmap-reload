# Kubernetes ConfigMap Reload

[![license](https://img.shields.io/github/license/jimmidyson/configmap-reload.svg?maxAge=2592000)](https://github.com/jimmidyson/configmap-reload)

**configmap-reload** is a simple binary to trigger a reload when a Kubernetes ConfigMap is updated.
It watches the mounted volume dir and notifies the target process that the config map has been changed.
It executes a command with Kubectl exec to the name container

### Usage

```
Usage of /configmap-reload:
  -command string
    	Command beeing executed on trigger. Arguments need to be seperated by ',' like 'nginx,-s,reload'
  -container string
    	Container name in pod
  -folder string
    	Folder to watch for changes in it
```

### License

This project is [Apache Licensed](LICENSE.txt)
