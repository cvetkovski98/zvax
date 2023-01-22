helm install redis \
    --set auth.password=$REDIS_PASSWORD \
    --set securityContext.enabled=true \
    --set securityContext.fsGroup=2000 \
    --set securityContext.runAsUser=0 \
    --set volumePermissions.enabled=true \
    --set replica.replicaCount=0 \
    --set master.persistence.enabled=true \
    --set master.persistence.path=/data \
    --set master.persistence.size=1Gi \
    bitnami/redis

# create secrets
kubectl create secret generic auth-secret --from-file=./secrets/auth.gcp.config.yaml
kubectl create secret generic keys-secret --from-file=./secrets/keys.gcp.config.yaml
kubectl create secret generic qrcode-secret --from-file=./secrets/qrcode.gcp.config.yaml
kubectl create secret generic slots-secret --from-file=./secrets/slots.gcp.config.yaml
kubectl create secret generic sa-secret --from-file=./secrets/zvax-project-access-key.json

# run deployments
kubectl apply -f ./auth.deployment.yaml
kubectl apply -f ./keys.deployment.yaml
kubectl apply -f ./qrcode.deployment.yaml
kubectl apply -f ./slots.deployment.yaml
kubectl apply -f ./ingress.yaml

# set healthcheck confgi
kubectl apply -f ./healthz.config.yaml
