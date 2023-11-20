# 使用 Node.js 镜像
FROM node:14

# 设置工作目录
WORKDIR /app

# 复制 package.json 和 package-lock.json (如果有的话)
COPY package*.json ./

# 安装项目依赖
RUN npm install

# 复制项目文件
COPY . .

# 构建应用
RUN npm run build

# 暴露端口
EXPOSE 3030

# 启动服务
CMD ["npm", "run", "serve"]
