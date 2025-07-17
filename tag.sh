#!/bin/bash

# 统一版本号
VERSION="1.0.5"

# 要打 tag 的 commit（可以是分支名、HEAD、hash）
TARGET_COMMIT="master"

for dir in */; do
  # 去掉最后的 /
  DIR_NAME="${dir%/}"

  # 构造 tag 名
  TAG_NAME="${DIR_NAME}/v${VERSION}"

  echo "Creating tag: $TAG_NAME -> $TARGET_COMMIT"
  git tag "$TAG_NAME" "$TARGET_COMMIT"
  git push origin "$TAG_NAME"
done