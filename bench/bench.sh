#!/bin/bash

go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out