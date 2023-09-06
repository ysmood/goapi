echo '== login =='
curl localhost:3000/login -id '{"username": "admin", "password": "123456"}'
# Set-Cookie: token=123456

echo
echo
echo '== get posts =='
curl 'localhost:3000/users/3/posts?keyword=sky' -H 'Cookie: token=123456'
# {"data":["post1","post2"],"meta":"User 3 using keyword: sky"}

echo
echo
echo '== print openapi doc =='
curl localhost:3000/openapi.json
# output the openapi json doc