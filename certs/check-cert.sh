#!/bin/bash

echo ''
echo '------------------------------'
openssl x509 -in tls.crt -text -noout

echo ''
echo '------------------------------'
openssl x509 -in tls.crt -text -noout | grep DNS:

echo ''
echo '------------------------------'
openssl x509 -in tls.crt -noout -checkhost localhost


