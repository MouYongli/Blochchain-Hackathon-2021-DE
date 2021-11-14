# Backend for the Application

## Deploy
1. make sure you have docker in your system
2. adjust configuration for yourself, the configuration file is on the directory `app/config/config.json`, server is for this app backend, chain is for jdchain
3. then build the image, go to this directory, and run `docker build --force-rm -t ${app_name} .`, replace `${app_name}` with any name you like
4. after that, then run the image, run `docker run -p ${port}:8090 -d ${app_name}`, replace `${port}` with any port you like and `${app_name}` which you have replace last step
5. enjoy the backend
