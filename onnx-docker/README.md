# onnx-docker

## Run
1. make sure that you have installed docker on your computer
2. run command to build image, `docker build --force-rm -t ${image_name} .`, replace `${image_name}` with any name you like
3. after building image, run command, `docker run --name=${container_name} -v ${uploads_name}:/model/uploads:rw -p ${port}:5000 -d ${image_name}`, replace three names for this command