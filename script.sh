#!/usr/bin/env bash




find ./groupcache | grep \.go | grep -v test |  xargs wc -l
