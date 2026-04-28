# install gemini cli
```
npm install -g @google/gemini-cli
```


# tạo môi trường ảo
```
python3 -m venv venv
```

# kích hoạt môi trường ảo
```
source venv/bin/activate
```

# cài đặt FastMCP để test 
```
(venv) $ pip install fastmcp
```
# lệnh thiết lập mcp connect đến gemini cli
```
(venv) $ fastmcp install gemini-cli basic_mcp_server.py
```
# mở tệp cấu hình gemini cli tại : ~/.gemini/settings.json
```
vim ~/.gemini/settings.json
```
sửa thành
```
"mcp": {
  "servers": [
    {
      "name": "basic_mcp_server",
      "command": [
        "/home/gmo/Tran_Vu_Thanh_Lam/AI_gemini_CLI_call_MCP/venv/bin/fastmcp", // HOẶC .\venv\Scripts\fastmcp.exe trên Windows
        "run",
        "basic_mcp_server.py"
      ],
      "cwd": "/home/gmo/Tran_Vu_Thanh_Lam/AI_gemini_CLI_call_MCP",
      "transport": "STDIO"
    }
  ]
}
```

# thoát môi trường ảo
```
deactivate
```







vim ~/.gemini/settings.json




gcloud projects create lamtvt-dev-123 --name="lamtvt"



gcloud projects list



{
  "mcp": {
    "inputs": [
      {
        "type": "promptString",
        "id": "apiKey",
        "description": "Firecrawl API Key",
        "password": true
      }
    ],
    "servers": {
      "firecrawl": {
        "command": "npx",
        "args": ["-y", "firecrawl-mcp"],
        "env": {
          "FIRECRAWL_API_KEY": "fc-6e16ee27ac75433089068e591309e213"
        }
      }
    }
  }
}

API KEY: fc-6e16ee27ac75433089068e591309e213

Remote hosted URL
https://mcp.firecrawl.dev/fc-05b463f55c2f4793bb2eef75a70916e0/v2/mcp

# settings gemini cli clound
```
{
  "selectedAuthType": "oauth-personal",
  "mcpServers": {
    "firecrawl": {
      "httpUrl": "https://mcp.firecrawl.dev/fc-05b463f55c2f4793bb2eef75a70916e0/v2/mcp",
      "trust": true
    }
  }
}
```




{
    "selectedAuthType": "oauth-personal",
    "mcpServers": {
        "testserver": {
            "httpUrl": "https://issue-18-dot-api-ai-mcp-dot-brantect-sl5lqzeyoa-an.a.run.app/uapi/ai/mcp",
            "headers": {
                "Accept": "application/json, text/event-stream",
                "Authorization": "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTc0MDk4NDYsImlhdCI6MTc1NzMyMzQ0NiwibG9nX29uX25tIjoiIiwibmJmIjoxNzU3MzIzNDQ2LCJzZXNzaW9uX2lkIjoiNGUyMmFjMTgtMjQ2YS00NTUyLWE0ODctMGZkZGU5NjcwNjMxIn0.KDq37oFyfl-ovdyFwJBYgAPdRBo25XihsytdevhrxursmZLDf2HvED62nqWa7w2UOHSlFTkZhYBukulIQ89t8IMcbwqWLlkvGmMsQktkqVId_SDgJ7DxI4Ar2pKZwrrBJtZ47vz_74kkUTOumDoyiI1lZYsnTHBgotSAx5WLeLk4CggZ2Dm1dVuhZMWDoDg49tmKwTPrWmMbdKNvvZBSaA2XvXjYNUij2tya9bYWz7Cgqwu4MIak9nzxmHuZ1Vabj1Ri4fGM_SsgCDjd8Y5Ic3GJk8xlf9vihiIoDQgqESm5pZHGNrdSnn_djq_kYDsTK_TkpDN75dHzYr_6j3ZkGQ"
            }
        }
    }
}