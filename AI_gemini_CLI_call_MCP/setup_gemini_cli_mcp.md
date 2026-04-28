# MCP Setup in Gemini CLI Cloud

## STEP 1: Open Cloud Shell
Go to:  
https://console.cloud.google.com/welcome?cloudshell=true

---

## STEP 2: Open settings.json file
Cases:  

1. First time (Gemini not installed yet): just run
```
vim ~/.gemini/settings.json
```

2. Gemini is already installed and you are inside gemini-cli:  
   Type:
```
/quit
```
   to exit CLI, then run:
```
vim ~/.gemini/settings.json
```
---

## STEP 3: Add JSON Configuration
Insert the following JSON snippet into settings.json and save:
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
```
---

## STEP 4: Start Gemini CLI Cloud
Run:
```
gemini
```
---

## STEP 5: List MCP Servers
Inside gemini-cli, type:
```
/mcp list
```

You should see the newly added server.
