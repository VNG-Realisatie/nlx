#!/bin/bash

minikube_ip=$(minikube ip)
echo "${minikube_ip}                              traefik.minikube"           | sudo tee -a /etc/hosts
echo "${minikube_ip}                         docs.nlx-dev-directory.minikube" | sudo tee -a /etc/hosts
echo "${minikube_ip}                   certportal.nlx-dev-directory.minikube" | sudo tee -a /etc/hosts
echo "${minikube_ip}                    directory.nlx-dev-directory.minikube" | sudo tee -a /etc/hosts
echo "${minikube_ip}     directory-inspection-api.nlx-dev-directory.minikube" | sudo tee -a /etc/hosts
echo "${minikube_ip}   directory-registration-api.nlx-dev-directory.minikube" | sudo tee -a /etc/hosts
echo "${minikube_ip}                      insight.nlx-dev-directory.minikube" | sudo tee -a /etc/hosts
echo "${minikube_ip}                         demo.nlx-dev-directory.minikube" | sudo tee -a /etc/hosts
echo "${minikube_ip}                        txlog.nlx-dev-rdw.minikube"       | sudo tee -a /etc/hosts
echo "${minikube_ip}                     irma-api.nlx-dev-rdw.minikube"       | sudo tee -a /etc/hosts
echo "${minikube_ip}                  insight-api.nlx-dev-rdw.minikube"       | sudo tee -a /etc/hosts
echo "${minikube_ip}                        txlog.nlx-dev-brp.minikube"       | sudo tee -a /etc/hosts
echo "${minikube_ip}                     irma-api.nlx-dev-brp.minikube"       | sudo tee -a /etc/hosts
echo "${minikube_ip}                  insight-api.nlx-dev-brp.minikube"       | sudo tee -a /etc/hosts
echo "${minikube_ip}                        txlog.nlx-dev-haarlem.minikube"   | sudo tee -a /etc/hosts
echo "${minikube_ip}                     irma-api.nlx-dev-haarlem.minikube"   | sudo tee -a /etc/hosts
echo "${minikube_ip}                  insight-api.nlx-dev-haarlem.minikube"   | sudo tee -a /etc/hosts
echo "${minikube_ip}                       outway.nlx-dev-haarlem.minikube"   | sudo tee -a /etc/hosts
echo "${minikube_ip}                  application.nlx-dev-haarlem.minikube"   | sudo tee -a /etc/hosts
