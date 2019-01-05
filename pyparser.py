import requests
import os
import codecs
from transliterate import translit
from bs4 import BeautifulSoup

directory = 'parsing_files'
path_to_dir = os.path.join(os.getcwd(), directory)
directory_data = 'data_from_parsing_files'
path_dir_data = os.path.join(os.getcwd(), directory_data)

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
    tag_lst = get_td_lst(html)
    write_auto_file(tag_lst)

def write_auto_file(tag_lst):
    file = os.path.join(path_dir_data, 'lada.txt')
    written = []
    with open(file, 'wb') as file:
        for t in tag_lst:
            text = t.find_previous_sibling("th").string.encode("utf-8")
            text = text.strip()
            text = text.split()[0]
            if text.decode("utf-8").isalpha():
                text = translit(text.decode("utf-8"), reversed=True).encode("utf-8")
            if text not in written:
                written.append(text)
                file.write(text+'\r'.encode("utf-8"))

def get_td_lst(html):
    soup = BeautifulSoup(html, 'lxml')
    tag_lst = soup.find_all("td", attrs={"data-title": "Стоимость обычного ТО, руб.*"})
    return tag_lst

def parsing():
    for f in os.listdir(path=path_to_dir):
        fileObj = codecs.open(os.path.join(path_to_dir,f), "r", "utf_8_sig")
        get_page_data(fileObj.read())

if __name__ == '__main__':
    get_http('https://www.tts.ru/service/lada/' )
    parsing()

