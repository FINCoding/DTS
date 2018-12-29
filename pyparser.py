import requests
import os
import codecs
from bs4 import BeautifulSoup

directory = 'parsing_files'
path_to_dir = os.path.join(os.getcwd(), directory)

def get_http(url,upd=False):
    file = get_file_name(url)
    if need_send_request(path_to_dir, file+'.html') or upd:
        headers = {'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36'}
        response = requests.get(url, headers=headers)
        new_file = os.path.join(path_to_dir, file+'.html')
        with open(new_file, 'wb') as file:
            file.write(response.text.encode("utf-8"))

def need_send_request(path, file):
    lst = os.listdir(path=path)
    if file not in lst: return True
    return False

def get_file_name(url):
    lst = url.split("/")
    return lst[len(lst)-2]

def get_page_data(html):
    soup = BeautifulSoup(html, 'lxml')
    for string in soup.strings:
        if string == 'Стоимость обычного ТО, руб.*': break
    print('TEST BEGIN', string.parent, "TEST END")
    tds = soup.find_all('table', {'class': 'responsive-table'})
    print(tds)

def parsing():
    for f in os.listdir(path=path_to_dir):
        fileObj = codecs.open(os.path.join(path_to_dir,f), "r", "utf_8_sig")
        get_page_data(fileObj.read())

if __name__ == '__main__':
    get_http('https://www.tts.ru/service/lada/' )
    parsing()
