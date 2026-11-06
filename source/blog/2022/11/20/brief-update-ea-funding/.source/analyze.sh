#!/bin/bash

cd /home/loki/Documents/core/ea/fresh/misc/openphil-funding

# wget https://www.openphilanthropy.org/wp-admin/admin-ajax.php?action=generate_grants&nonce=1920c9d172 -O grants.csv
# rm admin-ajax.php?action=generate_grants
Rscript analysis.R
