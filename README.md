# autocomplete-with-redis-1

A simple scalabale implimentation of search autocomplete using Redis based on the [blog of Salvatore Sanfilippo (Antirez)](http://oldblog.antirez.com/post/autocomplete-with-redis.html) , the creator of Redis . On the backend I used Golang. 

# How To Run 
Using Docker - 
```
git clone https://github.com/mdnurahmed/autocomplete-with-redis-1
cd autocomplete-with-redis-1
docker-compose up --build
```

Then go to localhost:3000 in the browser . I have included redisinsight with the docker-compose file. So if you wanna see how the search strings are stored in redis go to localhost:8001 and connect to the Redis inside the docker network with the following credentials -

```
Host : redis 
Port : 6379 
Name: redis 
Username :
Password :
```
