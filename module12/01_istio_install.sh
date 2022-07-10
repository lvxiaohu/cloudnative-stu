export all_proxy=socks5://192.168.31.251:7891
curl -L https://istio.io/downloadIstio | sh -
cd istio-1.14.1 
cp bin/istioctl /usr/local/bin
istioctl install --set profile=demo -y

k create ns httpmesh

k label ns httpmesh istio-injection=enabled
k -n httpmesh create deploy my-nginx --image=kuokou1/nginx:v1
