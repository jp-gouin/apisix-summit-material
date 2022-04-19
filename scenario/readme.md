Scenario of APISIX

STEP 1

kubectl  apply -f Step\ 1a-front-route-insecure.yaml
kubectl  apply -f Step\ 1b-back-route-insecure.yaml

Navigate to https://book-app.demo.jpgouin.pro 
API and UI is not protected

STEP 2

kubectl  delete -f Step\ 1a-front-route-secure.yaml && \
kubectl  apply -f Step\ 2a-front-route-secure.yaml
kubectl  delete -f Step\ 1b-back-route-secure.yaml && \
kubectl  apply -f Step\ 2b-back-route-secure.yaml

Refresh http://book-app.demo.jpgouin.pro -> Authenticate against keycloak using john
John cannot post book or author

Open new private nav login with Franck
Franck can post author and book

Go to keycloak
Explain permission

Show front and back config files

STEP 3 

kubectl  apply -f Step\ 3-back-route-apikey.yaml

Create consumer with UI

Can list authors
curl http://authkey-go-api.demo.jpgouin.pro/api/authors -H 'apikey: john-awesome-secure-key' -i

Cannot post author
curl -i -X POST -H 'apikey: john-awesome-secure-key' http://authkey-go-api.demo.jpgouin.pro/api/author

Create consumer with API : 
curl http://apisix-dashboard.demo.jpgouin.pro/apisix/admin/consumers -H 'X-API-KEY: edd1c9f034335f136f87ad84b625c8f1' -X PUT -d '
{
    "username": "john",
    "plugins": {
        "key-auth": {
            "key": "john-awesome-secure-key"
        }
    }
}'
