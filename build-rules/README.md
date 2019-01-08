# [csi-build-rules](https://github.com/kubernetes-csi/csi-build-rules)

These build and test rules can be shared between different Go projects
without modifications. Customization for the different projects happen
in the top-level Makefile.

The rules include support for building and pushing Docker images, with
the following features:
 - one or more command and image per project
 - push canary and/or tagged release images
 - automatically derive the image tag(s) from repo tags
 - never overwrite an existing release image

Sharing and updating
--------------------

`[git subtree](https://github.com/git/git/blob/master/contrib/subtree/git-subtree.txt)`
is the recommended way of maintaining a copy of the rules inside the
`build-rules` directory of a project. This way, it is possible to make
changes also locally, test them and then push them back to the shared
repository at a later time.

Cheat sheet:

- `git subtree pull --prefix=build-rules https://github.com/kubernetes-csi/csi-build-rules.git master` - update local copy to latest upstream
- edit, `git commit`, `git subtree push --prefix=build-rules git@github.com:<user>/csi-build-rules.git <my-new-or-existing-branch>` - push to a new branch before submitting a PR
