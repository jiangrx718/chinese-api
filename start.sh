#!/bin/bash

IMAGE_PREFIX="chinese-api"
ARCH_SUFFIX="amd64"
CONTAINER_NAME="chinese-api-app"

# 颜色定义（可选）
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 日志函数
log_success() { echo -e "${GREEN}✅ $1${NC}"; }
log_error() { echo -e "${RED}❌ $1${NC}"; }
log_warning() { echo -e "${YELLOW}⚠️  $1${NC}"; }
log_info() { echo -e "${BLUE}📝 $1${NC}"; }

echo "🚀 开始启动 ${CONTAINER_NAME} 容器..."

# 获取本地可用的最大版本号（假设版本号格式是 x.y.z）
LATEST_TAG=$(docker images --format "{{.Repository}}:{{.Tag}}" \
  | grep "^${IMAGE_PREFIX}:" \
  | grep "${ARCH_SUFFIX}" \
  | sed -E "s/^${IMAGE_PREFIX}:([0-9]+\.[0-9]+\.[0-9]+)-${ARCH_SUFFIX}$/\1/" \
  | sort -Vr \
  | head -n 1)

if [[ -z "$LATEST_TAG" ]]; then
  log_error "未找到符合格式的镜像版本（chinese-api:x.y.z-amd64）"
  exit 1
fi

FULL_IMAGE="${IMAGE_PREFIX}:${LATEST_TAG}-${ARCH_SUFFIX}"
log_info "找到最新镜像: ${FULL_IMAGE}"

# 停止并删除旧容器（如存在）
if docker ps -a --format "{{.Names}}" | grep -q "^${CONTAINER_NAME}$"; then
    log_warning "发现已存在的容器 ${CONTAINER_NAME}，正在停止并删除..."
    docker rm -f "$CONTAINER_NAME" 2>/dev/null
    if [[ $? -eq 0 ]]; then
        log_success "旧容器删除成功"
    else
        log_error "旧容器删除失败"
    fi
fi

# 启动容器
log_info "正在启动容器 ${CONTAINER_NAME}..."

CONTAINER_ID=$(docker run -d \
  --name "$CONTAINER_NAME" \
  --restart=always \
  -p 8080:8080 \
  -v /data/project/chinese/config/app.yml:/app/config/app.yml \
  --network common-app-net \
  "$FULL_IMAGE" 2>&1)

# 检查启动是否成功
if [[ $? -ne 0 ]] || [[ -z "$CONTAINER_ID" ]]; then
    log_error "容器启动失败！"
    echo "失败原因: ${CONTAINER_ID}"
    exit 1
fi

# 获取实际的容器ID（如果docker run命令成功返回的是容器ID）
if [[ "$CONTAINER_ID" =~ ^[a-f0-9]{12,}$ ]]; then
    log_success "容器已启动，容器ID: ${CONTAINER_ID:0:12}"
else
    # 如果返回的不是容器ID，可能是其他输出
    CONTAINER_ID=$(docker ps -lq --filter "name=${CONTAINER_NAME}" --format "{{.ID}}")
    if [[ -n "$CONTAINER_ID" ]]; then
        log_success "容器已启动，容器ID: ${CONTAINER_ID:0:12}"
    fi
fi

# 等待容器初始化的时间（根据应用调整）
log_info "等待容器初始化..."
sleep 3

# 检查容器状态
CONTAINER_STATUS=$(docker inspect --format='{{.State.Status}}' "$CONTAINER_NAME" 2>/dev/null)

if [[ "$CONTAINER_STATUS" == "running" ]]; then
    log_success "容器状态: 运行中"
    
    # 检查应用健康状态（如果应用有健康检查端点）
    log_info "检查应用健康状况..."
    
    # 示例：检查HTTP服务是否可访问（可选）
    if curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/health >/dev/null 2>&1; then
        log_success "应用健康检查通过"
    else
        log_warning "HTTP健康检查未通过，但容器在运行中"
    fi
    
    # 显示容器基本信息
    echo ""
    log_info "容器信息:"
    echo "----------------------------------------"
    docker ps --filter "name=$CONTAINER_NAME" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
    echo "----------------------------------------"
    
else
    log_error "容器启动失败！容器状态: ${CONTAINER_STATUS:-未找到}"
    
    echo ""
    log_error "========== 错误日志开始 =========="
    docker logs "$CONTAINER_NAME" 2>/dev/null || echo "无法获取容器日志"
    log_error "========== 错误日志结束 =========="
    
    echo ""
    log_error "========== 容器详情 =========="
    docker inspect "$CONTAINER_NAME" 2>/dev/null | grep -A 5 '"State"'
    
    # 清理失败的容器
    log_warning "正在清理失败的容器..."
    docker rm -f "$CONTAINER_NAME" 2>/dev/null
    
    exit 1
fi

# 显示最近启动日志（可选）
echo ""
log_info "最近启动日志:"
echo "----------------------------------------"
docker logs --tail 20 "$CONTAINER_NAME" 2>/dev/null || echo "暂无日志"
echo "----------------------------------------"

log_success "${CONTAINER_NAME} 启动完成！"
echo ""
echo "📊 访问地址: http://localhost:8080"
echo "📊 查看完整日志: docker logs -f $CONTAINER_NAME"