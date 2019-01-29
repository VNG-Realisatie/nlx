#!/bin/bash

echo "$(minikube ip)                              traefik.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                         docs.dev.nlx.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                   certportal.dev.nlx.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                    directory.dev.nlx.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)     directory-inspection-api.dev.nlx.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)   directory-registration-api.dev.nlx.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                      insight.dev.nlx.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                         demo.dev.nlx.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                        txlog.dev.rdw.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                     irma-api.dev.rdw.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                  insight-api.dev.rdw.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                        txlog.dev.brp.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                     irma-api.dev.brp.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                  insight-api.dev.brp.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                    txlog.dev.haarlem.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                 irma-api.dev.haarlem.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)              insight-api.dev.haarlem.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)                   outway.dev.haarlem.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)              application.dev.haarlem.minikube" | sudo tee -a /etc/hosts
