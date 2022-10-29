#!/bin/bash
unalias -a
set -eu
output_dir="bin"
app_base_name="excel2csv"

build_list=(
    "windows/amd64"
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
)

mkdir -p $output_dir
for build in "${build_list[@]}"; do
    IFS='/' read -r -a build_info <<<"$build"
    GOOS=${build_info[0]}
    GOARCH=${build_info[1]}

    output_path="${output_dir}/${app_base_name}_${GOOS}_${GOARCH}"
    if [ "$GOOS" = "windows" ]; then
        output_path="${output_path}.exe"
    fi
    GOOS=$GOOS GOARCH=$GOARCH \
        go build -ldflags='-w -s' -o "$output_path"

    #Some versions of upx for macos may have problems and become unusable after compression
    if [[ "$GOOS" != "darwin" ]]; then
        command -v upx && upx $output_path &
    fi
done
wait
echo "done"
ls -lh $output_dir
