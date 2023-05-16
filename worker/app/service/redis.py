from setting import config

if config.REDIS_URI==None:
    from redislite import Redis

class RedisClient:
    def __init__(self, host='redis', port=6379, db=0):
        self.redis = Redis(host=host, port=port, db=db)

    def set(self, key, value):
        self.redis.set(key, value)

    def get(self, key):
        return self.redis.get(key)