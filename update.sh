#!/bin/sh

DATE=$(date '+%Y-%m-%d %H:%M:%S')
git add -A
git commit -m "$DATE"
git push origin main
