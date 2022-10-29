#!/bin/bash
output_dir="bin"
app_base_name="excel2csv"

build_list=(
    "windows/amd64"
    "linux/amd64"
    "darwin/amd64"
    "darwin/arm64"
)

for build in ${build_list[@]}; do
    IFS='/' read -r -a build_info <<<"$build"
    GOOS=${build_info[0]}
    GOARCH=${build_info[1]}

    output_path="${output_dir}/${app_base_name}_${GOOS}_${GOARCH}"
    if [ "$GOOS" = "windows" ]; then
        output_path="${output_path}.exe"
    fi
    GOOS=$GOOS GOARCH=$GOARCH \
        go build -ldflags='-w -s' -o $output_path

    # command -v upx && upx $output_path &
done
