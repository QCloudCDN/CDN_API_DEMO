## description

This is the **Lua** version of the open api demo, which implements the *refresh_cdn_url* api based on **OpenResty**.

## third-part package

```bash
opm get ledgetech/lua-resty-http
opm get jkeys089/lua-resty-hmac
```

## usage

Add and modify your code

```lua
# nginx.conf:

location = /refresh_cdn_url {
    content_by_lua_file conf/lua/refresh_cdn_url.lua;
}

- refresh_cdn_url:
local qcloud_secret_id = ''     -- pls use your secret id
local qcloud_secret_key = ''    -- pls use your secret key
```

post demo

```json
# POST ****/refresh_cdn_url

{
    "urls": [
        "https://xxxx/index.html"
         "https://xxxx/yyyy.png",
    ]
}
```

