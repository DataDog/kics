---
title: "Using Kubernetes native secret management"
group_id: "rules/k8s"
meta:
  name: "k8s/using_kubernetes_native_secret_management"
  id: "b9c83569-459b-4110-8f79-6305aa33cb37"
  display_name: "Using Kubernetes native secret management"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "INFO"
  category: "Secret Management"
---
## Metadata

**Id:** `b9c83569-459b-4110-8f79-6305aa33cb37`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Info

**Category:** Secret Management

### Description

 Use of an external secret storage and management system should be considered for complex secret management needs, rather than using Kubernetes Secrets directly. Additionally, ensure that access to secrets is carefully limited.


## Compliant Code Examples
```yaml
apiVersion: secrets-store.csi.x-k8s.io/v1
kind: SecretProviderClass
metadata:
  name: azure-kvname
  namespace: myNameSpace
spec:
  provider: azure
  parameters:
    usePodIdentity: "true"              
    keyvaultName: "<key Vault Name>"               
    objects:  |
      array:
        - |
          objectName: secret1          
          objectType: secret                                      
        - |
          objectName: key1               
          objectType: key
    tenantId: "<tenant ID which the Key Vault sits under"            
  secretObjects:                             
  - secretName: appsecrets   
    data:
    - key: secret1                          
      objectName: secret1                                        
    type: Opaque  

```

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: cluster-secrets
data:
  # Fill with your encoded base64 CA
  certificate-authority-data: Cg==
stringData:
  # Fill with your string Token
  bearerToken: "my-token"
---
apiVersion: secrets-store.csi.x-k8s.io/v1
kind: SecretProviderClass
metadata:
  name: azure-kvname
  namespace: myNameSpace
spec:
  provider: azure
  parameters:
    usePodIdentity: "true"              
    keyvaultName: "<key Vault Name>"               
    objects:  
      array:
        - {objectName: secret1, objectType: secret}
        - {objectName: key1 , objectType: key}                                      
    tenantId: "<tenant ID which the Key Vault sits under"            
  secretObjects:                             
  - secretName: appsecrets   
    data:
    - key: secret1                          
      objectName: secret1                                        
    type: Opaque  

```
## Non-Compliant Code Examples
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: cluster-secrets
data:
  # Fill with your encoded base64 CA
  certificate-authority-data: Cg==
stringData:
  # Fill with your string Token
  bearerToken: "my-token"

```