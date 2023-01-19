helm install redis --set auth.password=$REDIS_PASSWORD --set replica.replicaCount=1 bitnami/redis

# create secrets
kubectl create secret generic auth-secret --from-file=./secrets/gcp/auth.gcp.config.yaml
kubectl create secret generic keys-secret --from-file=./secrets/gcp/keys.gcp.config.yaml
kubectl create secret generic qrcode-secret --from-file=./secrets/gcp/qrcode.gcp.config.yaml
kubectl create secret generic slots-secret --from-file=./secrets/gcp/slots.gcp.config.yaml
kubectl create secret generic sa-secret --from-file=./secrets/zvax-project-access-key.json

# run deployments
kubectl apply -f ./auth.deployment.yaml
kubectl apply -f ./keys.deployment.yaml
kubectl apply -f ./qrcode.deployment.yaml
kubectl apply -f ./slots.deployment.yaml
kubectl apply -f ./ingress.yaml

# set healthcheck confgi
kubectl apply -f ./healthz.config.yaml
