// js_mcp_server.js

const { Server, defineTool } = require('@google/gemini-cli/mcp');
const express = require('express');

// 1. Định nghĩa Công cụ (Tool)
const echoTool = defineTool({
  name: 'echoMessage',
  description: 'Lặp lại một tin nhắn được cung cấp. Tuyệt vời để kiểm tra kết nối.',
  parameters: {
    message: {
      type: 'string',
      description: 'Tin nhắn cần lặp lại.',
    },
  },
  async execute({ message }) {
    return {
      success: true,
      result: `JS Server (HTTP) đã nhận và lặp lại: ${message}`,
    };
  },
});

// 2. Định nghĩa Server MCP
const mcpServer = new Server({
  name: 'JSToolsHTTP',
  version: '1.0',
  description: 'Các công cụ Node.js cục bộ qua HTTP.',
  tools: [echoTool],
});

// 3. Khởi tạo và Chạy Server HTTP
const app = express();
const port = 8080;

// Sử dụng middleware của Gemini CLI để xử lý các yêu cầu MCP
app.use(mcpServer.middleware());

app.listen(port, () => {
  console.log(`✅ JSTools HTTP Server đang chạy tại http://localhost:${port}`);
  console.log('Giữ terminal này mở để duy trì kết nối.');
});