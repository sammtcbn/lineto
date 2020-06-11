#coding:UTF-8
import requests

def main():
    url = "https://notify-api.line.me/api/notify"
    token = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    headers = {"Authorization" : "Bearer "+ token}

    message =  'hello\r\nsam'
    payload = {"message" :  message}
    r = requests.post(url ,headers = headers ,params=payload)

#   files = {"imageFile": open("test.jpg", "rb")}
#   r = requests.post(url ,headers = headers ,params=payload, files=files)

if __name__ == '__main__':
    main()
