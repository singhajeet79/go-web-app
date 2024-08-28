# Install Argo CD
```
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "NodePort"}}'
kubectl patch svc argocd-server -n argocd -p '{"spec": {"ports": [{"port": 443, "nodePort": 32000}]}}' 
```
# Retrieve the Initial Admin Password
```
kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name 
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d | awk '{sub(/[-%]+$/, ""); print}'
```

# Access ArgoCD
http://<Node_IP>:<NodePort>

# Configure Argo CD CLI (Optional)
```
wget https://github.com/argoproj/argo-cd/releases/download/v2.4.0/argocd-linux-amd64
chmod +x argocd-linux-amd64
sudo mv argocd-linux-amd64 /usr/local/bin/argocd 
```

# Log in using CLI
```
argocd login <Node_IP>:<NodePort>
```
