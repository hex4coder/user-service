FROM scratch

# set working dir
WORKDIR /app

# copy file main
COPY  .env ./
COPY main ./

# port to use
EXPOSE 9000

# run program
CMD [ "./main" ]