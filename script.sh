# FILES=pb/*.proto
# for f in $FILES
# do 
# 	echo $f
#   #PB=$(echo $f)
  
#   protoc $f --cpp_out=../cpp
#   protoc $f --java_out=../java
# done


protoc -I . todo.proto --go_out=plugins=grpc:.
#protoc --proto_path=proto  --go_out=plugins=grpc:proto service.proto