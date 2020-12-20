--
-- Created by IntelliJ IDEA.
-- User: gacnio
-- Date: 2020/12/10
-- Time: 11:20
-- To change this template use File | Settings | File Templates.
--


local redis_host    = "10.10.8.150"
local redis_port    = 6379

-- connection timeout for redis in ms. don't set this too high!
local redis_timeout = 200

-- check a set with this key for blacklist entries
local redis_key     = "ip_blacklist"

-- cache lookups for this many seconds
local cache_ttl     = 1

-- end configuration

local ip                 = ngx.var.remote_addr
local ip_blacklist_cache = ngx.shared.ip_blacklist_cache

-- setup a local cache
if cache_ttl > 0 then
    -- lookup the value in the cache
    local cache_result = ip_blacklist_cache:get(ip)
    if cache_result then
        ngx.log(ngx.DEBUG, "ip_blacklist: found result in cache for "..ip.." -> "..cache_result)

        if cache_result == 0 then
            ngx.log(ngx.DEBUG, "ip_blacklist: (cache) no result found for "..ip)
            return
        end

        ngx.log(ngx.INFO, "ip_blacklist: (cache) "..ip.." is blacklisted")
        return ngx.exit(ngx.HTTP_FORBIDDEN)
    end
end

-- lookup against redis
local resty = require "resty.redis"
local redis = resty:new()

redis:set_timeout(redis_timeout)

local connected, err = redis:connect(redis_host, redis_port)
if not connected then
    ngx.log(ngx.ERR, "ip_blacklist: could not connect to redis @"..redis_host..": "..err)
    return
end

local result, err = redis:sismember("ip_blacklist", ip)
if not result then
    ngx.log(ngx.ERR, "ip_blacklist: lookup failed for "..ip..":"..err)
    return
end

-- cache the result from redis
if cache_ttl > 0 then
    ip_blacklist_cache:set(ip, result, cache_ttl)
end

redis:set_keepalive(10000, 2)
if result == 0 then
    ngx.log(ngx.INFO, "ip_blacklist: no result found for "..ip)
    return
end

ngx.log(ngx.INFO, "ip_blacklist: "..ip.." is blacklisted")
return ngx.exit(ngx.HTTP_FORBIDDEN)