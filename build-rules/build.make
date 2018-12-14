# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# The following make rules support building one or more command,
# packaging them in a containers and pushing to a registry. The expected
# repository layout is:
# - cmd/*/*.go - source code for each command
# - cmd/*/Dockerfile - docker file for each command
#
# The image name is the same as the command name. A new build
# gets pushed under different versions:
# - "canary" when building the master branch, overwriting older builds
# - "<branch name>-canary" for other branches, overwriting older builds
# - "vA.B.C[-something]" when build a revision tagged as that *and* when there is no such
#   image already, i.e. releases never get modified once pushed
#
# Using this approach it is not necessary to modify files to change
# versions and therefore it is also not necessary to create release
# branches in advance.
#
# Instead, the master branch can enter a freeze period where only
# release-critical commits are accepted, then a revision gets tagged
# and the next build will produce the release image. Once it becomes
# necessary to do a bugfix release, a branch can be created based on
# the tag and bug fixes can be merged there until that branch also
# gets tagged.

.PHONY: build-% build container-% container push-% push clean test

# A space-separated list of all commands in the repository, must be
# set in main Makefile of a repository.
# CMDS=

# Revision that gets built into each binary via the main.version string.
# Always includes the revision as a suffix.
REV=$(shell git describe --long --tags --match='v*' --dirty)
# Previous taggged revision (may or may not be the same)
TAGGED_REV=$(shell git describe --tags --match='v*' --abbrev=0)

REGISTRY_NAME=quay.io/k8scsi
IMAGE_TAGS=$(shell echo $$(git rev-parse --abbrev-ref HEAD)-canary | sed -e 's/^master-canary$$/canary/'; \
                   if [ "$$(git rev-list -n1 HEAD)" = "$$(git rev-list -n1 $(TAGGED_REV))" ]; then echo '$(TAGGED_REV)'; fi)
IMAGE_NAME=$(REGISTRY_NAME)/$*


ifdef V
TESTARGS = -v -args -alsologtostderr -v 5
else
TESTARGS =
endif

build-%:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-X main.version=$(REV) -extldflags "-static"' -o ./bin/$* ./cmd/$*

container-%: build-%
	docker build -t $*:latest -f ./cmd/$*/Dockerfile --label revision=$(REV) .

push-%: container-%
	set -ex; \
	push_image () { \
		docker tag $*:latest $(IMAGE_NAME):$$tag; \
		docker push $(IMAGE_NAME):$$tag; \
	}; \
	for tag in $(IMAGE_TAGS); do \
		if echo $$tag | grep -q -e '-canary$$'; then \
			: "creating or overwriting canary image"; \
			push_image; \
		elif docker pull $(IMAGE_NAME):$$tag 2>&1 | tee /dev/stderr | grep -q "manifest for $(IMAGE_NAME):$$tag not found"; then \
			: "creating release image"; \
			push_image; \
		else \
			: "release image $(IMAGE_NAME):$$tag already exists, skipping push"; \
		fi; \
	done

build: $(CMDS:%=build-%)
container: $(CMDS:%=container-%)
push: $(CMDS:%=push-%)

clean:
	-rm -rf bin

test:
	go test `go list ./... | grep -v 'vendor'` $(TESTARGS)
	go vet `go list ./... | grep -v vendor`
	files=$$(find . -name '*.go' | grep -v './vendor'); \
	if [ $$(gofmt -d $$files | wc -l) -ne 0 ]; then \
		echo "formatting errors:"; \
		gofmt -d $$files; \
		false; \
	fi
