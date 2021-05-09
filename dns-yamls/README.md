# Usage

* Deploy internal zone file
    *  `gcloud dns managed-zones create internal-seneca --description="Zone for internal Seneca services." --visibility=private --dns-name=internal.seneca.com. --networks=default`
    * `gcloud dns record-sets import seneca-internal.yaml --zone="internal-seneca" --delete-all-existing` 