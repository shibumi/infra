#!/bin/bash

hetzner_cloud_api_key=$(gopass api/hetzner.com/motoko)
echo "{ \"hetzner_cloud_api_key\": \"${hetzner_cloud_api_key}\" }"
