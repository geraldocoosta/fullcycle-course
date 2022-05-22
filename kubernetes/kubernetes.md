# KUBERNETES

É um orquestrador de containers.

Conseguimos pegar os containers docker e pro k8s pra ele subir isso

Automatiza a implantação, o dimensionamento e o gerenciamento de aplicativos em contêineres.

## Conceitos importantes

É disponibilizado através de um conjunto de apis

normalmente acessamos a api usando o CLI kubectl

Tudo é baseado em estado.

O kubernetes é um cluster (conjunto de muquinas), com `1 * master` e `1 * nodes`.

No master tempos

- kube-apiserver
- kube-controller-manager
- kube-scheduler

Nos nodes:

- Kubelet
- Kubeproxy

Os nodes se comunicam com o master para receber comandos

## Dinâmica "Superficial"

Cluster: conjunto de máquinas (cada node é uma maquina (vm ou maquina mesmo))

Cada máquina possui uma quantidade de vCPU e Memória

Pods: Unidade que contém os containers provisionados

O pod representa os processos rodando no Cluster

Pode ter mais de container rodando por pod, mas é exceção, como por exemplo na utilização do Istio.

Deployment: É um outro tipo de objeto do kubernetes, tem objetivo de provisionar os Pods

Pra isso a gente tem o replicaSet, onde falamos quantas replicas de pods nós vamos querer.

Geralmente o replica set é settado dentro de um manifesto Deployment.

Se o limite de cpu e memoria do cluster for atingido, o kubernetes **não** cria mais pods. O pod fica pendente até que tenha mais recursos, com outro nó ou colocando mais recursos.

Se um nó cair, os pods que pertenciam a esse nó vão ser redistribuídos para outros.

O k8s ira verificar a saúde dos pods pra se der ruim, subir outro saudável e matar o que não está saudável.

## Comandos importantes

```bash
# Mapeando para, quando eu acessar o localhost 8080 da minha máquina, ir para a porta 8080 de um pod
kubectl port-forward pod/goserver 8080:8080

#Deletando po no namespace default
kubectl delete pod goserver

# Buscando pods no namespace default
kubectl get pod

# Buscando replicasets namespace default
kubectl get replicaSet

# Mostra informações sobre o pod
kubectl describe pod namePod

# Trás o historico de implantações de um objecto, nesse caso, um deployment com nome nomedeployment
kubectl rollout history deployment nomedeployment

# Volta pra ultima versão que estava rodando
kubectl rollout undu deployment nomedeployment

# voltando para uma versão especifica
kubectl rollout undo deployment nomedeployment --to-revision=2
```

O replica set é um conjunto de pods, que são replicas de um mesmo pod.

Qual o problema dele? Basicamente, ele verifica a quantidade de pods que estão no ar, e só, se eu mudar o manifesto e aplicar ele, ele não irá criar outros pods.

Para ele subir outra versão, é necessário que ele seja deletado e criado novamente.
