# gowfa⛈
🚨This simple program is written with GO to forecast weather

## You can use with 2 methods : 
1. Api version
2. Cli version

## How to use:

First of all you need to create an .env file and add "". <br/>
The key should be api key from https://weatherapi.com/ . <br/>

## Run:

**Cli mode :**
```bash
  go run ./cmd/app/main.go -c="name of city"
```
![image](https://github.com/pooulad/gowfa/assets/86445458/ee49e779-74ba-44a2-b50e-6987b31be8fd)


**Api mode :**

```bash
  go run ./cmd/app/main.go -c="name of city" -a
```
Test api with curl

```bash
  curl localhost:3000/forecast?city=Tehran
```

![image](https://github.com/pooulad/gowfa/assets/86445458/58ba715c-189d-44e0-9cf8-17dae12de614)


The response is large json data about weather of city
