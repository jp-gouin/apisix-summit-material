# Setup the environment 

## Kind setup

``` 
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
  kubeadmConfigPatches:
  - |
    kind: JoinConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 30080
    hostPort: 80
    protocol: TCP
  - containerPort: 30443
    hostPort: 443
    protocol: TCP
- role: worker
- role: worker
``` 

With the following configuration file , use `./kind create cluster --name apisix --config=/home/ubuntu/apisix/cluster/kind-conf.yaml`  to create the kind cluster and expose `30080` and `30443` ports outside the cluster.

The `30080` port will be bound to Apisix to expose HTTP resources, the `30443` port will be bound to Apisix to expose HTTPS resources

## Remote connection


# Apisix deployment

```
helm repo add apisix https://charts.apiseven.com
helm repo update
```

Apisix installation : 
```
helm install -n apisix --create-namespace -f apisix-values.yaml apisix apisix/apisix
```

# keycloak deployment

```
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
``` 

# Certs generation

We use [omgwtfssl]() to generate our certificates for the example.

``` 
podman run  -e SSL_SUBJECT="*.demo.jpgouin.pro" --security-opt label=disable  paulczar/omgwtfssl
``` 

Let's only copy the certificate `cert.pem` and `key.pem` from the output of the command

``` 
-----BEGIN RSA PRIVATE KEY-----
  MIIEowIBAAKCAQEA3PbtZ0KzGi8v+CXxp+djyE0yg2lBmicPUSiPlKNeNVugrpd6
  HZThuptnY2liFKty0b+qASnvb7GZYi5USSeG0UPuTlJOeFOeCEpD4BidQ6hpMM8f
  24/J9ldRdzYQfwhhSlBUO1uTC7VOLMlEGczF7iCpEA7r0bsGJZ5pgvfqpUdKm19O
  nFSOyL523E0J9t1LJLQUp5kaNUtVj7IaB1FXN4JH9lL8OrFolde+YUw2kOnsukAL
  348cllQ1wNgrAENAw+ZVse4UI70oFz0tsm9TkgSMqfV26DlodLGvaLi8BgtKmzI0
  R8giFXQZn1JYpR/BVqPqtw+D2VWtuOKdLhzsFwIDAQABAoIBAD6JVB9JbU8X08ez
  VXVjs4XY81blzz9FlrwtcpVqtxr617sR++mrXUdm/BfSl4OqElDyoba6m51M8lAr
  NIePvEQ1qZhZCwRbBfrqqdMepEOlwj+iHy4Qv+QMSm4myvxo7sANhAlYnYFOytU6
  ESAMKGiArdhxoKQ2/kJfVxVN/QbRY49ZTpgetcrYWnaeLNF3O94kATW8rV3T9Cy9
  Wh33JXnzalbHKAb1WKMvwKCGmw7idUn0wbhMKWvSruEyM4bhbQLad12qraDk0FUo
  45ZZXwEXAXbjp158oc8Q0rSqYJud57+kzKFestgpdMQvx7ke7gWPqTufOtzL7eXa
  zpGVNqECgYEA/cd2fE5QWNE8E5ISPe7+5UfNA5ztICAZeBPML3EzQ5JQC82y4S1g
  7n0exGdkBSkxxRDHMaskGlU1MhgFuu9Aamvo2tFfATLqNiaA7DcwrdHUIBISFHM3
  angUajj4fu/4ZIHWNyTAVgbfAkU8HPS25bMWNHrNsPpUnDil26GLRFsCgYEA3uXz
  VuUaYMlvp3x10n67KB8LLn+v2pRgRH4j/G2TycL/DZaGLKJbeAvorckA0lDz40Dh
  Mz2qwB2Crs5nYRcSIaTKto4pAQEFf7ADvW8kEtz65GyIz+e672p/UrEkz2pXiADI
  vrjDnLRB/X3ozFJopiKss0OX6gDAxNjycKnxU/UCgYAZrl2nJqWq10GlGVsPOWhB
  +4obm33DinwJUreO9X2ikOPYzfAUKWtttuuuJGhSvWBz+MavmAoHQCgp8ZRi3mM+
  Yb0mp5ldnbFl2W+id9NNQ7abqDh8KUyqUYx/U1SSQ+Z3BUyQN//etMNMj4UNyRXZ
  GM5ecS7vgeWeCX1/RFG7kQKBgCNZAhsAoAtDRpKf+StVb6awLZzFIQUfzJhwDJLm
  aa5wXvQvr716Tdkewlp06s1viw255zyyBdcLlwLTtq898mElegQzTStclxultSIg
  cu8O3jbFQ3j8/bckA2cAsp619YWa4jrkoBjEuUZS5k0osHHQQ+T1ziyMzAUrhl75
  hEzVAoGBAMqO3dWOhS8WVd+erLlT2h1i1ftcRX620xSLpK/Ric7NnfP7DD1cbA09
  iJV9puiKN8aXJXInRXflPMbzBDgu+NZ1D/D+C7fDTcCS93x14ct+th5BuK8SgX0u
  /+qMeHYeDhjtAw2DCBhH2XUfKnfu5xVC50Im87SjhxrJdo0ITJLF
  -----END RSA PRIVATE KEY-----
  ``` 
  ```
  -----BEGIN CERTIFICATE-----
  MIIC9jCCAd6gAwIBAgIUJtW7drCTbF/774ZNwEyzcFRZGjMwDQYJKoZIhvcNAQEL
  BQAwEjEQMA4GA1UEAwwHdGVzdC1jYTAeFw0yMjAzMjkwOTQ4MDJaFw0yMjA1Mjgw
  OTQ4MDJaMB0xGzAZBgNVBAMMEiouZGVtby5qcGdvdWluLnBybzCCASIwDQYJKoZI
  hvcNAQEBBQADggEPADCCAQoCggEBANz27WdCsxovL/gl8afnY8hNMoNpQZonD1Eo
  j5SjXjVboK6Xeh2U4bqbZ2NpYhSrctG/qgEp72+xmWIuVEknhtFD7k5STnhTnghK
  Q+AYnUOoaTDPH9uPyfZXUXc2EH8IYUpQVDtbkwu1TizJRBnMxe4gqRAO69G7BiWe
  aYL36qVHSptfTpxUjsi+dtxNCfbdSyS0FKeZGjVLVY+yGgdRVzeCR/ZS/DqxaJXX
  vmFMNpDp7LpAC9+PHJZUNcDYKwBDQMPmVbHuFCO9KBc9LbJvU5IEjKn1dug5aHSx
  r2i4vAYLSpsyNEfIIhV0GZ9SWKUfwVaj6rcPg9lVrbjinS4c7BcCAwEAAaM5MDcw
  CQYDVR0TBAIwADALBgNVHQ8EBAMCBeAwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsG
  AQUFBwMBMA0GCSqGSIb3DQEBCwUAA4IBAQAO/XvTZESx5T5IMLJ0RxgNIwWat75u
  0pUqvBz7ASBmENEqerZMZ4HmBvLE7mmIKNOndUWrxteWCZqyOIQyNURGsFCMnUeH
  10ufS9DuA1D4Sx2W08NWJuYWE67+0Bd2jrQW3yWlFP90oZ0WtB0TehqD7RyEx4AS
  H5J1WdJ6+wPnMsOhgJFxJJjusxZz/7BcYLk5bgvEOTntJr7NSkRa+haNCtuKwqf3
  NUW/dG4Vy/Fj77lNL8cOGfYqyaMQhpAgs2eNvEO49vbrA3GRCynuTlqLgXvfKiSD
  FZSP6wECbwNBPsamzL9OOqTEkfaPXJh4ls3Y2a1ttr1FMIxBX1jFc2jF
  -----END CERTIFICATE-----
  ```


Upload keys to Apisix
``` 
kubectl exec -it -n apisix apisix-5b66d9f798-m785g -- curl -X PUT http://127.0.0.1:9180/apisix/admin/ssl/1 -H 'X-API-Key: edd1c9f034335f136f87ad84b625c8f1' -d '
{
    "cert": "-----BEGIN CERTIFICATE-----
MIIC9jCCAd6gAwIBAgIUO08VZYbn6JuFKasbiEEUkVMQnwcwDQYJKoZIhvcNAQEL
BQAwEjEQMA4GA1UEAwwHdGVzdC1jYTAeFw0yMjAzMjkxMDA0MDFaFw0yMjA1Mjgx
MDA0MDFaMB0xGzAZBgNVBAMMEiouZGVtby5qcGdvdWluLnBybzCCASIwDQYJKoZI
hvcNAQEBBQADggEPADCCAQoCggEBAL7rt3U871i0Qa/s4HVig8HcP1GR1eMyq/dI
jcCNRLToidOqE9eI4yQBS/hzgU1eKQJztyGj9L4VINFBA2assBYCV8HO24k/aU4Z
TWmNsiGgfnPwT610FXM65H6uFAgRuTyI9pbNGSULFpV/B1vP5A7kFoNPiRcJwoR8
/CiNtGdlO8f5/w5Vw+qgmGoLqW68rvUx/wtUDCtINyj/YxD1FqXXPfRE1+RmdBxc
KuvvEuqIOmtRCryAqtlqq6LC4Ekaz3f3DOLCGct+HPYIORJ0d+nf/O5PYyQbVtTE
9bsSjq8aPINQKG4OZxrwtS7DV/m7LphH/kU0V+JEltOIVgugeLkCAwEAAaM5MDcw
CQYDVR0TBAIwADALBgNVHQ8EBAMCBeAwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsG
AQUFBwMBMA0GCSqGSIb3DQEBCwUAA4IBAQCT9tUDp2DOkntvLOSchOto+DQdFL6D
ySDOFWxBoyDLtZGXAAOuXRPN2t9Lo8qmTLErMfgpDRNzuuuhVvzeu92+3NRyl6FQ
8Rcciz795E/StuisDje6Ov48IpLhAN3Q4OtiLGY/K23LfI2YWifglnFqb88JVzlS
kgmxJGsaZhTgKNtDwd7JCUfKHywBGqGBpMBppVAfYnSP1tZbjgTXlAFpHW7DHWrH
ez4NWtnWaDm3r0VNBzXl4aLnxlHWIos5jdj8C9xWUTHSlBNOazX/ri66HtRqW+Bd
vVHiFZfvHIOqZ0gUoM371v5ynWyubQuMggigBHx6c+rsMaIYQ1fPaI4z
-----END CERTIFICATE-----",
    "key": "-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAvuu3dTzvWLRBr+zgdWKDwdw/UZHV4zKr90iNwI1EtOiJ06oT
14jjJAFL+HOBTV4pAnO3IaP0vhUg0UEDZqywFgJXwc7biT9pThlNaY2yIaB+c/BP
rXQVczrkfq4UCBG5PIj2ls0ZJQsWlX8HW8/kDuQWg0+JFwnChHz8KI20Z2U7x/n/
DlXD6qCYagupbryu9TH/C1QMK0g3KP9jEPUWpdc99ETX5GZ0HFwq6+8S6og6a1EK
vICq2WqrosLgSRrPd/cM4sIZy34c9gg5EnR36d/87k9jJBtW1MT1uxKOrxo8g1Ao
bg5nGvC1LsNX+bsumEf+RTRX4kSW04hWC6B4uQIDAQABAoIBAF2KBWdWRHn0Tt6P
suUwMBeR/h/L0Lwwvlo4XOMDE6+C4swmXvRRp9+jFwKSLL6oLHV8FjRQLP/YQCzl
qnHql5cg7LgSzjKvM7vE382gU67wACRzj8YOwtei+gIq7buKW3QSMqZttC0XAqia
BRPxhe3ZyEHrfOIhsQDSb4S68LFa6jDtF4gAEu404nvDtnWGvFXJN2rd2U0Kc5G5
+uWrS3u25jm9r/77gG846yBHxWiMV8r4QeKrcJFo7C3rbtpkr6ZdsjeXvFdx5Lu7
sX6mvzwifZxqu8xjprHSy8CtYIMcRKJJKvH7EeOS2RSBAiF8MPZhedbKdtynIaJH
y1ND5BECgYEA5GQChQAMfAoqY6hLUbz4jfRF1tNMZCX+wJh6AGSjMC2kQ6apfKI3
9lTLmoM2vMSUumtGUea2IjfJzja6+ftMp962X2BQC7lUliQoISFVnXhxVt9gq5VM
atnVgKenp5DK0kzNaG42mXA888ooMlq3VBS1l2tMSqxd6urhasDX3K0CgYEA1gAg
5mfvoG/YyOcsnrTj0nN0uOjLLlUAXP6EmWCqAcbEuMb4XJ3VXA5/C4QhqxB/jFj8
jCXeJyogpBct0f2R/cPE4PjRlgnSHGjGIdNQ7HRmeoDgqsrXF+IELp/BKNOhXqnx
cW8OWL1srB5fYKHXxKXaAxfF9mY5eaUMAH/ZYb0CgYAq2NRcLUkqUCHptl4DBKfA
we4EQnnXZAVqDnD89+Rhmn5xrqenWSuUjA3ye+FeAqdAXfXAUb8jpkG27S+gzFNy
PucJ96CswmUrEbdxl7ZfJ4Rj3t5c8lJ2zU0vqMRcPF7Am2YBzINzv8m/ltJ5t+ki
3gu/T4Ltk26/LreBdpJItQKBgDdJ6TYd/EcTHospTtniGkoxEiMD4hqiU2mzSEbo
NoQm+oRSw7AKBym1hRVQmfI8XPfBtd3vmqm7tJswceIjBSju/1qwblW5S9OTLj1m
/y1YFXHpAiKeLVw6RKJXG+yUYMi4V2zbKHW82urNHg81QS3JO7440iiK2KaZAbTP
UzZxAoGBAMPdbImUsTanurFhlaq1yZdMsNaaD5d7kiO7Xb/nIOshlSGlbJuy9xwn
F8TE3Qc3iEntX0ttGaUw4F0ZiRZESYFT9A97CR2OWIuM1jliV1XKO9xG6KiB3q2X
60Fq8e9IMIDvgvwp0A3b9OCvttHceU1MP11QEQTyKOI5AB+6SNCK
-----END RSA PRIVATE KEY-----",
    "snis": ["*.demo.jpgouin.pro"]
}
'
``` 

# Keycloak configuration

## TL;DR 
Use the `realm-export.json` file to quickly setup Keycloak, **you still need to create the users**

## Setup Users , groups and roles

1. Create two users , `franck` an author and `jhon` an editor
2. Create a `author_group`  and a `editor_group` groups
3. Add `franck` to the `author_group`
4. Add `john` to `editor_group`
5. Create a Role `editor` and `author` roles
6. Map `author_group` to `author` role
7. Mao `editor_group` to editor role

## Setup Client

### Create a OIDC client : 
```
Name : poc-apisix
Access Type : Confidential
Authorization Enabled : On
Root URL : https://book-app.demo.jpgouin.pro
Valid Redirect URIs : https://book-app.demo.jpgouin.pro/*
``` 

### Setup Authorization

**Resources**
Resources are the list of URI api that you want to protect.

In our scenario we will define 4 Resources to protect : 
1. Author creation
2. Book creation
3. List authors
4. List books

Our demo application is using 4 apis to access and create resources

| Api            | Action        |
|----------------|---------------|
| `/api/author`  | Create author |
| `/api/book`    | Create book   |
| `/api/authors` | List authors  |
| `/api/books`   | List books    |

**Authorization Scopes**
Authorization scopes are HTTP method that will be allowed for users and roles.

We will define , for the purpose of the demo, 3 scopes : 
1. GET
2. POST
3. OPTIONS

As you can imagine for an application , you will probably ending up with all HTTP method :) 

**Policies**

Policies are the matching between roles and resources.
In our scenario we will have 4 Policies : 
1. Only defined group can list authors
2. Only defined group can list books
3. Only defined group can create author
4. Only defined group can create book 

Our example is pretty straightforward but here you can mixed roles and really authorized in depth each apis of your application !

**Permissions**

Finally the last thing to configure is the mapping between resources and policies. 

We are defining 4 mapping for our example for each policies.

# Apps deployment

## Front end 

The front end is composed of two kubernetes files 
| File            | Description        |
|----------------|---------------|
| `deployment.yaml `  | The deployment of the application|
| `service.yaml `  | The service to expose the application within the kubernetes cluster|

**Deployment** 
``` 
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml 
```
This will install the application in the `poc-front-app` namespace

```
kubectl get po,svc -n poc-front-app
NAME                                 READY   STATUS    RESTARTS   AGE
pod/poc-front-app-797c9c94bf-qjhnq   1/1     Running   0          22m

NAME                       TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)   AGE
service/poc-front-appsvc   ClusterIP   10.96.194.4   <none>        80/TCP    5d21h
```

## Go api 

The go api is composed of two kubernetes files 
| File            | Description        |
|----------------|---------------|
| `deployment.yaml `  | The deployment of the application|
| `service.yaml `  | The service to expose the application within the kubernetes cluster|

**Deployment** 
``` 
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml 
```
This will install the application in the `poc-go-api` namespace

```
kubectl get po,svc -n poc-go-api
NAME                              READY   STATUS    RESTARTS   AGE
pod/poc-go-api-854d54cd84-dz4fp   1/1     Running   0          22m

NAME                    TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
service/poc-go-apisvc   ClusterIP   10.96.118.46   <none>        80/TCP    5d5h
```

# DNS

In my case, i have a DNS account and i setup the dns accordingly , but you migth want to update your `/etc/hosts` file and add the following line : 
```
<public_ip_where_kind_is_running> book-app.demo.jpgouin.pro demo-api.demo.jpgouin.pro
```
# Scenario 1 : Expose the application

First thing first, we need to expose the application to outside our cluster.

We are using Apisix Ingress Controller to help us with that. 

Instead of having to curl Apisix and create the Upstream and the Route, we are going to use some CRD from Apisix Ingress Controller that will do the magic for us !

```
kubectl apply -f Step\ 1a-front-route-insecure.yaml
kubectl apply -f Step\ 1b-back-route-insecure.yaml
```

The route for the Go api is using a plugin : 
```
 plugins:
      - name: cors
        enable: true
        config:
          allow_credential: true
          allow_origins: "**"
          allow_origins_by_regex: [".*.demo.jpgouin.pro"]
          allow_methods: GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS
          expose_headers: X-PINGOTHER, Content-Type, Authorization
          allow_headers: X-PINGOTHER, Content-Type, Authorization
```

This plugin allows some Cross Origin Request Sharing, known as `cors`. indeed our front end is accessing a back-end that is not provided by the same server. 

Navigate to `https://book-app.demo.jpgouin.pro` , you should access the `front-app` \
You should also be able to create an Author and add a Book as well !

# Scenario 2 : Secure the application

It's time to secure the access to the application ! 

For that we are going to protect our `front-app` using the `openid-connect` plugin that will handle the generation of the token against Keycloak by redirecting us to Keycloak UI to login.

Then we are going to protect our `go-api` using the amazing `authz-keycloak` plugin. With the following config : 
``` 
config:
    token_endpoint: https://sso.demo.jpgouin.pro/auth/realms/apisix/protocol/openid-connect/token
    http_method_as_scope: true
    lazy_load_paths: true
    client_id: poc-apisix
    audience: poc-apisix
    client_secret: <redated>
    discovery: https://sso.demo.jpgouin.pro/auth/realms/apisix/.well-known/uma2-configuration
    realm: apisix
    ssl_verify: false
``` 
The key here is to use `http_method_as_scope` along `lazy_load_paths` to leverage the `Authorization` tab from Keycloak.

Try to connect to the `front-app` and let you flow to the redireciton to keycloak :)

Login using `franck` user. 

Create a new author , called `john` by filling the form, the author should be created as expected.

Now try to add a new book , you should have a `403` return from the server (and more precisely from Apisix) !\
It works as expected , yay, `franck` is not part of the `editor_group`  so he cannot create a new book !

Open an incognito tab on your browser and login using `john`.

`john` is able to see the author created by `franck`.

Let's go and try to add a book `myAwesomeBookvol1` and write `john` as the author.

`john` can add books and view them as he is part of the `editor` group :)

Refresh the page where `franck` is logged in , and voila the book from `john` is displayed !

Now let's play a little with the Authorization :)

Let's say that the Author list is sensitive and only auditor should be allowed to access it !

1. Navigate to Keycloak and in the `poc-apisix` client, click on the authorization tab and in policies.
2. Remove the `editor` group from the policies `Only user in group can list author` and save.
3. Refresh the page where `john` is logged in , there is now an error message `403` instead of the list of author !

Without having to code anything on our application or change the scope of our user we have successfully change the behaviour of our application and deeply secured the `api` and the access to the `ui`.