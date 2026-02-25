#!/bin/bash

IMAGE_PREFIX="chinese-api"
ARCH_SUFFIX="amd64"
CONTAINER_NAME="chinese-api-app"

echo "🔄 开始重启 ${CONTAINER_NAME} 容器..."

# 1. 查找当前运行的容器使用的镜像
CURRENT_IMAGE=$(docker inspect --format='{{.Config.Image}}' "$CONTAINER_NAME" 2>/dev/null)

if [[ -z "$CURRENT_IMAGE" ]]; then
    echo "❌ 容器 ${CONTAINER_NAME} 不存在或未运行"
    echo "📝 正在尝试启动新容器..."
    
    # 调用启动脚本（假设启动脚本名为 start.sh）
    if [[ -f "./start.sh" ]]; then
        bash ./start.sh
        exit $?
    else
        echo "❌ 找不到启动脚本 start.sh，请手动启动"
        exit 1
    fi
fi

echo "📦 当前容器使用的镜像: ${CURRENT_IMAGE}"

# 2. 获取本地可用的最大版本号
LATEST_TAG=$(docker images --format "{{.Repository}}:{{.Tag}}" \
  | grep "^${IMAGE_PREFIX}:" \
  | grep "${ARCH_SUFFIX}" \
  | sed -E "s/^${IMAGE_PREFIX}:([0-9]+\.[0-9]+\.[0-9]+)-${ARCH_SUFFIX}$/\1/" \
  | sort -Vr \
  | head -n 1)

if [[ -z "$LATEST_TAG" ]]; then
  echo "❌ 未找到符合格式的镜像版本（chinese-api:x.y.z-amd64）"
  exit 1
fi

LATEST_FULL_IMAGE="${IMAGE_PREFIX}:${LATEST_TAG}-${ARCH_SUFFIX}"

echo "🆕 本地最新镜像: ${LATEST_FULL_IMAGE}"

# 3. 检查是否需要更新镜像
if [[ "$CURRENT_IMAGE" != "$LATEST_FULL_IMAGE" ]]; then
    echo "⚠️  检测到新版本镜像，将使用最新版本重启"
    echo "📥 从 ${CURRENT_IMAGE} 更新到 ${LATEST_FULL_IMAGE}"
    
    # 停止并删除旧容器
    docker rm -f "$CONTAINER_NAME" 2>/dev/null
    
    # 启动新容器
    docker run -d \
      --name "$CONTAINER_NAME" \
      --restart=always \
      -p 8080:8080 \
      -v /data/project/chinese/config/app.yml:/app/config/app.yml \
      --network common-app-net \
      "$LATEST_FULL_IMAGE"
      
    if [[ $? -eq 0 ]]; then
        echo "✅ 容器已使用新镜像重启成功"
    else
        echo "❌ 容器重启失败"
        exit 1
    fi
else
    echo "✅ 当前已是最新版本，执行热重启"
    
    # 4. 重启现有容器（保持相同配置）
    docker restart "$CONTAINER_NAME"
    
    if [[ $? -eq 0 ]]; then
        echo "✅ 容器热重启成功"
    else
        echo "❌ 容器重启失败，尝试重新创建"
        
        # 获取容器的完整配置并重新创建
        OLD_CONTAINER_ID=$(docker inspect --format='{{.Id}}' "$CONTAINER_NAME" 2>/dev/null)
        
        # 停止并删除旧容器
        docker rm -f "$CONTAINER_NAME"
        
        # 使用相同配置重新创建
        docker run -d \
          --name "$CONTAINER_NAME" \
          --restart=always \
          -p 8080:8080 \
          -v /data/project/chinese/config/app.yml:/app/config/app.yml \
          --network common-app-net \
          "$CURRENT_IMAGE"
          
        if [[ $? -eq 0 ]]; then
            echo "✅ 容器重新创建成功"
        else
            echo "❌ 容器重新创建失败"
            exit 1
        fi
    fi
fi

# 5. 等待容器启动并检查状态
echo "⏳ 等待容器启动..."
sleep 3

# 检查容器状态
CONTAINER_STATUS=$(docker inspect --format='{{.State.Status}}' "$CONTAINER_NAME" 2>/dev/null)

if [[ "$CONTAINER_STATUS" == "running" ]]; then
    echo "✅ 容器状态: 运行中"
    
    # 可选：查看最近日志
    echo "📋 最近日志:"
    docker logs --tail 10 "$CONTAINER_NAME"
else
    echo "❌ 容器状态异常: ${CONTAINER_STATUS}"
    echo "📋 错误日志:"
    docker logs "$CONTAINER_NAME"
    exit 1
fi

echo "🎉 容器重启完成"