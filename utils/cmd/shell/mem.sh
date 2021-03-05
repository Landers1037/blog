#!/bin/bash
free | awk '{print $3}'|head -2 |tail -1
exit 0