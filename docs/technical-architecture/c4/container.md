# C4: Containers

```mermaid
flowchart TB
    subgraph vkc["vkc (single Go binary)"]
        resolver["command resolver"]
        transport["HTTP transport"]
        cache["spec cache\n~/.cache/vkc"]
        tokens["token store\nKeychain / ~/.config/vkc"]
        output["output renderer"]
    end
    vikunja["Vikunja API"] 
    
    resolver --> transport
    transport --> cache
    transport --> tokens
    transport -->|"HTTPS"| vikunja
    transport --> output
```

## Runtime containers

There is exactly one process: the `vkc` binary. Two on-disk stores (spec cache, token store) are containers in C4 terms because they outlive the process and hold state users can inspect or break. The cache is disposable (safe to delete); the token store is sensitive (chmod 600 file fallback, keychain preferred).
