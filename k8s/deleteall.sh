kubectl delete deployment auth-deployment keys-deployment qrcode-deployment slots-deployment
kubectl delete service auth-service keys-service qrcode-service slots-service
kubectl delete secret auth-secret keys-secret qrcode-secret slots-secret sa-secret
helm uninstall redis
