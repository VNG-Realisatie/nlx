#!/bin/bash

echo "$(minikube ip)                              traefik.minikube"           | sudo tee -a /etc/hosts
echo "$(minikube ip)                         docs.nlx-dev-directory.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                   certportal.nlx-dev-directory.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                    directory.nlx-dev-directory.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)     directory-inspection-api.nlx-dev-directory.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)   directory-registration-api.nlx-dev-directory.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                      insight.nlx-dev-directory.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                         demo.nlx-dev-directory.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                        txlog.nlx-dev-rdw.minikube"       | sudo tee -a /etc/hosts
echo "$(minikube ip)                     irma-api.nlx-dev-rdw.minikube"       | sudo tee -a /etc/hosts
echo "$(minikube ip)                  insight-api.nlx-dev-rdw.minikube"       | sudo tee -a /etc/hosts
echo "$(minikube ip)                        txlog.nlx-dev-brp.minikube"       | sudo tee -a /etc/hosts
echo "$(minikube ip)                     irma-api.nlx-dev-brp.minikube"       | sudo tee -a /etc/hosts
echo "$(minikube ip)                  insight-api.nlx-dev-brp.minikube"       | sudo tee -a /etc/hosts
echo "$(minikube ip)                    txlog.nlx-dev-haarlem.minikube"       | sudo tee -a /etc/hosts
echo "$(minikube ip)                 irma-api.nlx-dev-haarlem.minikube"       | sudo tee -a /etc/hosts
echo "$(minikube ip)              insight-api.nlx-dev-haarlem.minikube"       | sudo tee -a /etc/hosts
echo "$(minikube ip)                   outway.nlx-dev-haarlem.minikube"       | sudo tee -a /etc/hosts
echo "$(minikube ip)              application.nlx-dev-haarlem.minikube"       | sudo tee -a /etc/hosts
