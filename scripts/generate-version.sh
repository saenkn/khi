#!/bin/bash

echo "export const VERSION=\"$(cat VERSION)\"" >> ./web/src/environments/version.prod.ts
echo "export const VERSION=\"$(cat VERSION)@$(git log -1 --pretty=format:%h )\"" >> ./web/src/environments/version.dev.ts