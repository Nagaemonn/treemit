# show help message
@default: help

App :=  'treemit'
Version := `grep '^const VERSION = ' cmd/main/version.go | sed "s/^VERSION = \"\(.*\)\"/\1/g"`

# show help message
@help:
    echo "Build tool for {{ App }} {{ Version }} with Just"
    echo "Usage: just <recipe>"
    echo ""
    just --list

# build the application with running tests
build: test
    go build -o treemit .cmd/main/

# run tests and generate the coverage report
test:
    go test -covermode=count -coverprofile=coverage.out ./...

# clean up build artifacts
clean:
    go clean
    rm -f coverage.out treemit build

# update the version if the new version is provided
update_version new_version = "":
    if [ "{{ new_version }}" != "" ]; then \
        sed 's/$VERSION/{{ new_version }}/g' .template/README.md > README.md; \
        sed 's/$VERSION/{{ new_version }}/g' .template/version.go > cmd/main/version.go; \
    fi

# build treemit for all platforms
make_distribution_files:
    for os in "linux" "windows" "darwin"; do \
        for arch in "amd64" "arm64"; do \
            mkdir -p dist/treemit-$os-$arch; \
            env GOOS=$os GOARCH=$arch go build -o dist/treemit-$os-$arch/treemit cmd/main/treemit.go; \
            cp README.md LICENSE dist/treemit-$os-$arch/; \
            tar cvfz dist/treemit-$os-$arch.tar.gz -C dist treemit-$os-$arch; \
        done; \
    done

upload_assets tag:
    # リリースが存在するまで最大10回リトライ
    for i in $(seq 1 10); do \
        gh release view {{ tag }} --repo Nagaemonn/treemit && break; \
        echo 'Waiting for GitHub release to be available...'; \
        sleep 3; \
    done
    gh release upload --repo Nagaemonn/treemit {{ tag }} dist/*.tar.gz